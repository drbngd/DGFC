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
		// nothing to do in the header

	case *ast.ProgramBody:
		// all declarations in Progam Body are global
		for _, d := range node.Declarations {
			nodeType, err := st.Analyze(d, "GLOBAL")
			if err != nil {
				return nodeType, err
			}
		}

		for _, s := range node.Statements {
			//exprType, err := st.Analyze(s, "GLOBAL")
			nodeType, err := st.Analyze(s, "")
			if err != nil {
				return nodeType, err
			}
		}

	case *ast.ProcedureDeclaration:
		// get procedure name for scope
		procName := node.Header.Identifier.Name
		currentScope := scope + "." + procName
		if node.IsGlobal {
			currentScope = "GLOBAL"
		}

		// get param list & return type
		paramList := GetProcedureParams(*node)
		returnType := node.Header.Type.Name

		// first analyze the header & body for any errors
		nodeType, err := st.Analyze(node.Header, currentScope)
		if err != nil {
			return nodeType, err
		}
		nodeType, err = st.Analyze(node.Body, currentScope)
		if err != nil {
			return nodeType, err
		}

		// no errors found -> add to symbol table
		sym := NewSymbol(procName, "Procedure", returnType, scope, true, paramList, false, 0)
		st.AddSymbol(sym)

	case *ast.ProcedureHeader:
		// nothing to do in the header

	case *ast.ParameterList:
		// nothing to do, handled in ProcedureDeclaration

	case *ast.Parameter:
		// nothing to do, handled in ProcedureDeclaration

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
		sym := NewSymbol(name, "Variable", varType, scope, false, nil, node.IsArray, arrBound)
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
		relType, err := st.Analyze(node.Relation, scope)
		if err != nil {
			return "", err
		}
		return relType, nil

	case *ast.Term:
		// analyze the factor
		factorType, err := st.Analyze(node.Factor, scope)
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
		// Handle Factor node
		// ...

	case *ast.Name:
		// Handle Name node
		// ...

	case *ast.ArgumentList:
		// Handle ArgumentList node
		// ...

	case *ast.Number:
		// Handle Number node
		// ...

	case *ast.String:
		// Handle String node
		// ...

	default:
		// Handle unhandled types
		// ...
	}

	return "", nil
}

func GetProcedureParams(pd ast.ProcedureDeclaration) []string {
	params := []string{}
	for _, p := range pd.Header.ParameterList.Parameters {
		params = append(params, p.VariableDeclaration.Type.Name)
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
