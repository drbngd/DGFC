package analyzer

import (
	"DGFC/pkg/ast"
	"DGFC/pkg/token"
	"fmt"
	"strings"
)

const (
	INT   = "integer"
	STR   = "string"
	FLOAT = "float"
	BOOL  = "bool"
)

type ReturnType string

type Analyzer struct {
	SymbolTable *SymbolTable
}

func New() *Analyzer {
	return &Analyzer{
		SymbolTable: NewSymbolTable(),
	}
}

func (st *SymbolTable) Analyze(node ast.Node, scope string) (ReturnType, error) {
	switch node := node.(type) {

	case *ast.Program:
		nodeType, err := st.Analyze(node.Header, "GLOBAL")
		if err != nil {
			return nodeType, err
		}

		nodeType, err = st.Analyze(node.Body, "GLOBAL")
		if err != nil {
			return nodeType, err
		}

	case *ast.ProgramHeader:
		_, err := st.Analyze(node.Identifier, "GLOBAL")
		if err != nil {
			return "", err
		}

	case *ast.ProgramBody:
		// all declarations in Progam Body are global
		for _, d := range node.Declarations {
			nodeType, err := st.Analyze(d, "GLOBAL")
			if err != nil {
				return nodeType, err
			}
		}

		for _, s := range node.Statements {
			nodeType, err := st.Analyze(s, "GLOBAL")
			if err != nil {
				return nodeType, err
			}
		}

	case *ast.ProcedureDeclaration:
		// get procedure name, return type, scope, and param list
		procName := node.Header.Identifier.Name
		procReturnType := node.Header.Type.Name
		newScope := scope + "." + procName
		if node.IsGlobal {
			newScope = "GLOBAL"
		}
		paramList := GetProcedureParams(*node)

		// first: add procedure to symbol table
		sym := NewSymbol(procName, "Procedure", procReturnType, scope, true, false, "0")
		st.AddSymbol(sym)

		// second: add procedure's params to symbol table
		for _, p := range paramList {
			paramName := p.Identifier.Name
			paramType := p.Type.Name
			isArray := p.IsArray
			if p.IsGlobal {
				return "", fmt.Errorf("Procedure Declaration: Parameters cannot be global")
			}
			sym := NewSymbol(paramName, "Procedure Param", paramType, newScope, false, isArray, "0")
			st.AddSymbol(sym)
		}

		// third: analyze the header & body
		nodeType, err := st.Analyze(node.Header, newScope)
		if err != nil {
			return nodeType, err
		}
		nodeType, err = st.Analyze(node.Body, newScope)
		if err != nil {
			return nodeType, err
		}

	case *ast.ProcedureHeader:
		_, err := st.Analyze(node.Identifier, scope)
		if err != nil {
			return "", err
		}

	case *ast.ParameterList:
		for _, p := range node.Parameters {
			nodeType, err := st.Analyze(&p, scope)
			if err != nil {
				return nodeType, err
			}
		}

	case *ast.Parameter:
		_, err := st.Analyze(node.VariableDeclaration, scope)
		if err != nil {
			return "", err
		}

	case *ast.ProcedureBody:
		for _, d := range node.Declarations {
			nodeType, err := st.Analyze(d, scope)
			if err != nil {
				return nodeType, err
			}
		}

		for _, s := range node.Statements {
			nodeType, err := st.Analyze(s, scope)
			if err != nil {
				return nodeType, err
			}
		}

	case *ast.VariableDeclaration:
		// get variable name, type, and scope
		if node.IsGlobal {
			scope = "GLOBAL"
		}
		name := node.Identifier.Name
		varType := node.Type.Name
		arrBound := node.Bound.Value.Value

		// add to symbol table
		sym := NewSymbol(name, "Variable", varType, scope, false, node.IsArray, arrBound)
		st.AddSymbol(sym)

	case *ast.TypeMark:
		// nothign to handle in TypeMark

	case *ast.Bound:
		// nothing to handle in Bound

	case *ast.ProcedureCall:
		// get procedure name
		procName := node.Identifier.Name

		// get return type of each expression in the procedure call
		returnTypeList := []string{}
		for _, e := range node.ArguementList.Arguments {
			nodeType, err := st.Analyze(&e, scope)
			if err != nil {
				return nodeType, err
			}
			returnTypeList = append(returnTypeList, (string(nodeType)))
		}

		// check if procedure exists in symbol table & get it's param list
		procSym, ok := st.table[procName+".PROC"]
		if !ok {
			return "", fmt.Errorf("Procedure Call: Procedure declaration not found")
		} else {
			// check if number of arguments match
			if len(returnTypeList) != len(procSym.ParamList) {
				return "", fmt.Errorf("Procedure Call: Number of arguments do not match")
			}

			// check if argument types match
			for i, t := range returnTypeList {
				if t != procSym.ParamList[i] {
					return "", fmt.Errorf("Procedure Call: One or more argument types do not match")
				}
			}

			// if both tests pass, return the return type of the procedure
			return ReturnType(procSym.ReturnType), nil
		}

	case *ast.AssignmentStatement:
		// first analyze destination
		destType, err := st.Analyze(node.Destination, scope)
		if err != nil {
			return "", err
		}

		//// get destination name
		//destName := node.Destination.Identifier.Name
		//
		//// check if in symbol table - TODO: check if this is necessary
		//_, ok := st.table[destName]
		//if !ok {
		//	return "", fmt.Errorf("Assignment Statement: Destination '%s' not found.", destName)
		//}

		// analyze expression
		exprType, err := st.Analyze(node.Expression, scope)
		if err != nil {
			return "", err
		}

		// check if type matches TODO - (int & bool are compatible) & (int & float are compatible)
		if !CheckTypeCompatibility(string(destType), string(exprType)) {
			return "", fmt.Errorf("Assignment Statement: Type mismatch")
		}

	case *ast.Destination:
		// check if present in symbol table
		destName := node.Identifier.Name
		_, ok := st.table[destName]
		if !ok {
			return "", fmt.Errorf("Destination: '%s' not found in symbol table", destName)
		}
		destType := st.table[destName].ReturnType

		// analyze expression
		exprType, err := st.Analyze(node.Expression, scope)
		if err != nil {
			return "", err
		}

		// if expression is an array, check if index is within bounds, and that expression is an integer
		if node.IsArray {
			// check if expression is of type integer
			if exprType != "integer" {
				return "", fmt.Errorf("Assingment Statement: Array index must be an integer")
			}

			//// TODO - bounds check to be done in Code Generation
			//// check if index is within bounds
			//if node.Expression.Bound.Value.Value > node.Expression.ArraySize {
			//	return "", fmt.Errorf("Destination: Array index out of bounds")
			//}
		}

		return ReturnType(destType), nil

	case *ast.IfStatement:
		newScope := scope + ".IF." + string(st.IfElseCount)
		st.IfElseEncountered()

		// analyze condition
		condType, err := st.Analyze(node.Condition, scope)
		if err != nil {
			return "", err
		}
		// check if condition is of type bool or int
		if condType != token.BOOLEAN && condType != token.INTEGER {
			return "", fmt.Errorf("If Statement: Condition must be of type bool or int")
		}

		// analyze then block -> use new scope
		for _, s := range node.ThenBlock {
			_, err := st.Analyze(s, newScope+".THEN")
			if err != nil {
				return "", err
			}
		}

		// analyze else block
		for _, s := range node.ElseBlock {
			_, err := st.Analyze(s, newScope+".ELSE")
			if err != nil {
				return "", err
			}
		}

	case *ast.LoopStatement:
		newScope := scope + ".LOOP." + string(st.ForLoopCount)
		st.ForLoopEncountered()

		// analyze initialization statement
		_, err := st.Analyze(node.Initialization, scope)
		if err != nil {
			return "", err
		}

		// analyze condition statement
		condType, err := st.Analyze(node.Condition, scope)
		if err != nil {
			return "", err
		}
		// check if condition is of type bool or int
		if condType != token.BOOLEAN && condType != token.INTEGER {
			return "", fmt.Errorf("Loop Statement: Condition must be of type bool or int")
		}

		// analyze body
		for _, s := range node.Body {
			_, err := st.Analyze(s, newScope)
			if err != nil {
				return "", err
			}
		}

	case *ast.ReturnStatement:
		// analyze expression
		exprType, err := st.Analyze(node.Expression, scope)
		if err != nil {
			return "", err
		}

		// need to check if return type matches procedure return type
		requiredReturnType := GetLastProcedureName(scope)
		if string(exprType) != requiredReturnType {
			return "", fmt.Errorf("Return Statement: Return type does not match procedure return type")
		}

	case *ast.Identifier:
		// nothing to do here

	case *ast.Expression:
		// a lot to do here
		// analyze arithmetic operation
		exprType, err := st.Analyze(node.ArithOp, scope)
		if err != nil {
			return "", err
		} else if exprType != token.INTEGER {
			return "", fmt.Errorf("Expression: Bitwise And/Or/Not operation must be of type int")
		}

		// analyze bitwise and/or expression
		for _, e := range node.AndOrList {
			andOrType, err := st.Analyze(&e, scope)
			if err != nil {
				return "", err
			} else if andOrType != token.INTEGER {
				return "", fmt.Errorf("Expression: Bitwise And/Or/Not expression must be of type int")
			}
		}

	case *ast.AndOrExpression:
		// analyze the expression
		exprType, err := st.Analyze(node.Expression, scope)
		if err != nil {
			return "", err
		} else if exprType != token.INTEGER {
			return "", fmt.Errorf("Expression: Bitwise And/Or/Not expression must be of type int")
		}

	case *ast.ArithmeticOperation:
		// analyze the relation
		relType, err := st.Analyze(node.Relation, scope)
		if err != nil {
			return "", err
		}

		// analyze the arithmetic operation
		// todo - add concatenation for string
		for _, e := range node.AddSubList {
			addSubType, err := st.Analyze(&e, scope)
			if err != nil {
				return "", err
			} else if !CheckTypeCompatibility(string(relType), string(addSubType)) {
				return "", fmt.Errorf("Artithmetic Expression: Type mismatch")
			}
		}

	case *ast.AddSubExpression:
		// analyze the arithmetic operation
		exprType, err := st.Analyze(node.ArithmeticOperation, scope)
		if err != nil {
			return "", err
		} else if exprType != token.INTEGER && exprType != token.FLOAT {
			return "", fmt.Errorf("Expression: Arithmetic operation must be of type int or float")
		}

	case *ast.Relation:
		// analyze the term
		termType, err := st.Analyze(node.Term, scope)
		if err != nil {
			return "", err

		}

		// analyze the relation
		// todo - add equality/inequality check for string
		for _, e := range node.RelationalOperationList {
			relType, err := st.Analyze(&e, scope)
			if err != nil {
				return "", err
			} else if !CheckTypeCompatibility(string(termType), string(relType)) {
				return "", fmt.Errorf("Relation Expression: Type mismatch - INTEGER or FLOAT expected")
			}

		}
	case *ast.RelationalExpression:
		// analyze the term
		relType, err := st.Analyze(node.Term, scope)
		if err != nil {
			return "", err
		}
		return relType, nil

	case *ast.Term:
		// analyze the factor
		_, err := st.Analyze(node.Factor, scope)
		if err != nil {
			return "", err
		}

		// analyze the multiplication/division expression
		for _, e := range node.MultDivList {
			multDivType, err := st.Analyze(&e, scope)
			if err != nil {
				return "", err
			} else if multDivType != token.INTEGER && multDivType != token.FLOAT {
				return "", fmt.Errorf("Term Expression: Type mismatch - INTEGER or FLOAT expected")
			}
		}

		// todo - add case for what type is returned

	case *ast.MultDivExpression:
		// analyze the factor
		exprType, err := st.Analyze(node.Factor, scope)
		if err != nil {
			return "", err
		}
		return exprType, nil

	case *ast.Factor:
		// analyze the expression
		if node.IsExpression {
			_, err := st.Analyze(node.Expression, scope)
			if err != nil {
				return "", err
			}
		} else if node.IsProcedureCall {
			_, err := st.Analyze(node.ProcedureCall, scope)
			if err != nil {
				return "", err
			}
		} else if node.IsNumber {
			_, err := st.Analyze(node.Number, scope)
			if err != nil {
				return "", err
			} else {
				// code to return type of number
			}
		} else if node.IsString {
			_, err := st.Analyze(node.String, scope)
			if err != nil {
				return "", err
			} else {
				return token.STR, nil
			}
		} else if node.IsName {
			_, err := st.Analyze(node.Name, scope)
			if err != nil {
				return "", err
			}
		}

	case *ast.Name:
		// check if present in symbol table & if it is array does the symbol has the record of being an array
		isArray := node.IsArray
		name := node.Identifier.Name

		// check if present in symbol table and if it is an array
		_, ok := st.table[name]
		if !ok {
			return "", fmt.Errorf("Name: '%s' not found in symbol table", name)
		} else if isArray && !st.table[name].IsArray {
			return "", fmt.Errorf("Name: '%s' is not an array", name)
		}

		// check if expression is an int value
		exprType, err := st.Analyze(node.Expression, scope)
		if err != nil {
			return "", err
		} else if exprType != token.INTEGER {
			return "", fmt.Errorf("Name: Array index must be an integer")
		}

		return ReturnType(st.table[name].ReturnType), nil

	case *ast.ArgumentList:
		for _, a := range node.Arguments {
			_, err := st.Analyze(&a, scope)
			if err != nil {
				return "", err
			}
		}

	case *ast.Number:
		num := node.Value
		return ReturnType(CheckNumberType(num)), nil

	case *ast.String:
		return token.STR, nil

	default:
		// Handle unhandled types
		// ...
	}

	return "", nil
}

func GetProcedureParams(pd ast.ProcedureDeclaration) []ast.VariableDeclaration {
	params := []ast.VariableDeclaration{}

	for _, p := range pd.Header.ParameterList.Parameters {
		params = append(params, *p.VariableDeclaration)
	}
	return params
}

// TODO - need to test is
func GetLastProcedureName(scope string) string {
	// Split the scope string by the period character
	elements := strings.Split(scope, ".")

	// Iterate over the elements in reverse order
	for i := len(elements) - 1; i >= 0; i-- {
		// If the element is not a keyword (GLOBAL, IF, LOOP, ELSE), return it
		if elements[i] != "ELSE" && elements[i] != "THEN" && elements[i] != "IF" && elements[i] != "LOOP" && elements[i] != "GLOBAL" {
			return elements[i]
		}
	}

	// If no procedure name is found, return an empty string
	return ""
}

func CheckTypeCompatibility(t1, t2 string) bool {
	if (t1 == token.BOOLEAN || t1 == token.INTEGER) && (t2 == token.BOOLEAN || t2 == token.INTEGER) {
		return true
	} else if (t1 == token.FLOAT || t1 == token.INTEGER) && (t2 == token.FLOAT || t2 == token.INTEGER) {
		return true
		//} else if t1 == token.STRING && t2 == token.STRING {
		//	return true
	} else {
		return false
	}
}

func CheckNumberType(num string) string {
	if strings.Contains(num, ".") {
		return token.FLOAT
	} else {
		return token.INTEGER
	}
}
