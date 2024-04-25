package analyzer

import (
	"DGFC/pkg/ast"
	"DGFC/pkg/token"
	"fmt"
	"strings"
)

//const (
//	INT   = "integer"
//	STR   = "string"
//	FLOAT = "float"
//	BOOL  = "bool"
//)

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
		st.Analyze(node.Header, "GLOBAL")
		st.Analyze(node.Body, "GLOBAL")

		return "", nil

	case *ast.ProgramHeader:
		st.Analyze(node.Identifier, "GLOBAL")

		return "", nil

	case *ast.ProgramBody:
		// all declarations in Progam Body are global
		for _, d := range node.Declarations {
			st.Analyze(d, "GLOBAL")
		}
		for _, s := range node.Statements {
			st.Analyze(s, "GLOBAL")
		}

		return "", nil

	case *ast.ProcedureDeclaration:
		// 0: get procedure name, return type, scope, and param list
		procName := node.Header.Identifier.Name
		procReturnType := node.Header.Type.Name
		newScope := scope + "." + procName
		if node.IsGlobal {
			newScope = "GLOBAL"
		}
		paramList, paramTypeList := GetProcedureParams(*node)

		// 1: check if procedure is already declared
		_, ok := st.table[procName+".PROC"]
		if ok {
			return "", fmt.Errorf("Procedure Declaration: Procedure '%s' already declared", procName)
		}

		// 2: add procedure to symbol table
		sym := NewSymbol(procName, "Procedure", procReturnType, scope, true, paramTypeList, false, "0")
		st.AddSymbol(sym)

		// 3: add procedure's params to symbol table
		for _, p := range paramList {
			paramName := p.Identifier.Name
			paramType := p.Type.Name
			isArray := p.IsArray
			if p.IsGlobal {
				return "", fmt.Errorf("Procedure Declaration: Parameters cannot be global")
			}
			sym := NewSymbol(paramName, "Procedure Param", paramType, newScope, false, nil, isArray, "0")
			st.AddSymbol(sym)
		}

		// 4: analyze the header & body
		nodeType, err := st.Analyze(node.Header, newScope)
		if err != nil {
			return nodeType, err
		}
		nodeType, err = st.Analyze(node.Body, newScope)
		if err != nil {
			return nodeType, err
		}

		// 5: return the return type of the procedure
		return ReturnType(procReturnType), nil

	case *ast.ProcedureHeader:
		st.Analyze(node.Identifier, scope)

		return "", nil

	case *ast.ParameterList:
		for _, p := range node.Parameters {
			st.Analyze(&p, scope)
		}

		return "", nil

	case *ast.Parameter:
		st.Analyze(node.VariableDeclaration, scope)

		return "", nil

	case *ast.ProcedureBody:
		for _, d := range node.Declarations {
			st.Analyze(d, scope)
		}
		for _, s := range node.Statements {
			st.Analyze(s, scope)
		}

		return "", nil

	case *ast.VariableDeclaration:
		// 0: get variable name, type, and scope
		name := node.Identifier.Name
		varType := node.Type.Name
		arrBound := node.Bound.Value.Value
		if node.IsGlobal {
			scope = "GLOBAL"
		}

		// 1: check if already exists in the symbol table with the same scope
		varSym, ok := st.table[name]
		if ok && varSym.Scope == scope {
			return "", fmt.Errorf("Variable Declaration: Variable '%s' already declared in the same scope", name)
		}

		// 2: check that array bound is an integer
		if node.IsArray {
			if GetNumberType(arrBound) != token.INTEGER {
				return "", fmt.Errorf("Variable Declaration: Array bound must be an integer")
			}
		}

		// 3: add to symbol table
		sym := NewSymbol(name, "Variable", varType, scope, false, nil, node.IsArray, arrBound)
		st.AddSymbol(sym)

		// 4: return the type of the variable
		return ReturnType(varType), nil

	case *ast.TypeMark:
		// 0: get type name
		typeName := node.Name

		// 1: check if type is valid
		if typeName != token.INTEGER && typeName != token.FLOAT && typeName != token.BOOLEAN && typeName != token.STRING {
			return "", fmt.Errorf("TypeMark: Invalid type '%s'", typeName)
		}

		// 2: return the type
		return ReturnType(typeName), nil

	case *ast.Bound:
		_, err := st.Analyze(node.Value, scope)
		if err != nil {
			return "", err
		}
		if GetNumberType(node.Value.Value) != token.INTEGER {
			return "", fmt.Errorf("Bound: Bound must be an integer")
		}

		return "", nil

	case *ast.ProcedureCall:
		// 0: get procedure name
		procName := node.Identifier.Name

		// 1: check if procedure is already declared in the current scope, get return type and param list
		symObj, ok := st.table[procName+".PROC"]
		if !ok && !ScopeCompatible(scope, symObj.Scope) {
			return "", fmt.Errorf("Procedure Call: Procedure '%s' not found in the current scope", procName)
		}
		procParamList := symObj.ParamTypeList

		// 2: check if number of arguments match & if argument types match
		if len(node.ArguementList.Arguments) != len(procParamList) {
			return "", fmt.Errorf("Procedure Call: Number of arguments do not match")
		}
		for i, e := range node.ArguementList.Arguments {
			nodeType, err := st.Analyze(&e, scope)
			if err != nil {
				return nodeType, err
			} else if nodeType != ReturnType(procParamList[i]) {
				return "", fmt.Errorf("Procedure Call: Argument type mismatch")
			}
		}

		// 3: return the return type of the procedure
		return ReturnType(symObj.ReturnType), nil

	case *ast.AssignmentStatement:
		// assignment stat - type converted to that of destination
		// bool & int
		// int & float
		// 0: analyze destination & expression
		destType, err := st.Analyze(node.Destination, scope)
		if err != nil {
			return "", err
		}
		exprType, err := st.Analyze(node.Expression, scope)
		if err != nil {
			return "", err
		}

		// 1: type check destination & expression
		if !AssignmentTypeCompatible(string(destType), string(exprType)) {
			return "", fmt.Errorf("Assignment Statement: Destination & Expression type mismatch")
		}

		// 2: return "" as no type to be returned
		return "", nil

	case *ast.Destination:
		// 0: check if present in symbol table & in current scope, get type
		destName := node.Identifier.Name
		symObj, ok := st.table[destName]
		if !ok {
			if !ScopeCompatible(scope, symObj.Scope) {
				return "", fmt.Errorf("Destination: '%s' not in scope", destName)
			}
			return "", fmt.Errorf("Destination: '%s' not found in symbol table", destName)
		}
		destType := st.table[destName].ReturnType

		// 1: analyze identifier & expression
		st.Analyze(node.Identifier, scope)
		exprType, _ := st.Analyze(node.Expression, scope)

		// 2: if array: check if that expression is an integer, and index is within bounds
		if node.IsArray {
			if exprType != token.INTEGER {
				return "", fmt.Errorf("Destination: Array index must be an integer")
			}
			// TODO(codegen): Bounds checking to be done in code generation
		}
		return ReturnType(destType), nil

	case *ast.IfStatement:
		// 0: generate new scope string & increment if/else count
		newScope := scope + ".IF." + string(st.IfElseCount)
		st.IfElseEncountered()

		// 1:  analyze condition & type check
		condType, _ := st.Analyze(node.Condition, scope)
		if !(condType == token.BOOLEAN || condType == token.INTEGER) {
			return "", fmt.Errorf("If Statement: Condition must be of type bool (or int)")
		}

		// 2: analyze then block -> use new scope
		for _, s := range node.ThenBlock {
			st.Analyze(s, newScope+".THEN")
		}

		// 3: analyze else block
		for _, s := range node.ElseBlock {
			st.Analyze(s, newScope+".ELSE")
		}

		return "", nil

	case *ast.LoopStatement:
		// 0: generate new scope string & increment for loop count
		newScope := scope + ".LOOP." + string(st.ForLoopCount)
		st.ForLoopEncountered()

		// 1: analyze initialization statement
		st.Analyze(node.Initialization, scope)

		// 2: analyze condition statement & type check
		condType, _ := st.Analyze(node.Condition, scope)
		if !(condType == token.BOOLEAN || condType == token.INTEGER) {
			return "", fmt.Errorf("Loop Statement: Condition must be of type bool (or int)")
		}

		// 3: analyze body
		for _, s := range node.Body {
			st.Analyze(s, newScope)
		}

		return "", nil

	case *ast.ReturnStatement:
		// 0: analyze expression
		exprType, _ := st.Analyze(node.Expression, scope)

		// 1: type check return expression & procedure return type
		assocProcName := GetLastProcedureName(scope)
		symObj, ok := st.table[assocProcName+".PROC"]
		if !ok {
			return "", fmt.Errorf("Return Statement: Parent Procedure not found")
		}

		if string(exprType) != symObj.ReturnType {
			return "", fmt.Errorf("Return Statement: Return type does not match parent procedure return type")
		}

		// 2: return "" as no type to be returned
		return "", nil

	case *ast.Identifier:
		return "", nil

	case *ast.Expression:
		// 0: analyze arithmetic operation & type check
		arithopType, _ := st.Analyze(node.ArithOp, scope)
		// if no operation is present and no list of expressions is present, then type doesn't matter
		if !node.IsNot && len(node.AndOrList) == 0 {
			return "", nil
		}

		if arithopType != token.INTEGER || arithopType != token.FLOAT {
			return "", fmt.Errorf("Expression: Bitwise operations must be of type int (or float)")
		}
		// TODO(codegen): if type is float, we will truncate it to int in code generation

		// 1: analyze AndOr list expression & type check for int
		for _, e := range node.AndOrList {
			st.Analyze(&e, scope)
		}

		return "", nil

	case *ast.AndOrExpression:
		// 0: analyze the expression & type check
		exprType, _ := st.Analyze(node.Expression, scope)
		if exprType != token.INTEGER || exprType != token.FLOAT {
			return "", fmt.Errorf("Expression: Bitwise And/Or/Not expression must be of type int (or float)")
		}

		// TODO(codegen): if type is float, we will truncate it to int in code generation

		return "", nil

	case *ast.ArithmeticOperation:

	case *ast.AddSubExpression:

	case *ast.Relation:

	case *ast.RelationalExpression:

	case *ast.Term:

	case *ast.MultDivExpression:

	case *ast.Factor:

	case *ast.Name:
		// 0: get name, check if in symbol table, in scope, and if it is an array
		name := node.Identifier.Name
		symObj, ok := st.table[name]
		if !ok {
			return "", fmt.Errorf("Name: '%s' not found in symbol table", name)
		} else {
			if !ScopeCompatible(scope, symObj.Scope) {
				return "", fmt.Errorf("Name: '%s' not in scope", name)
			}
			if node.IsArray && !symObj.IsArray {
				return "", fmt.Errorf("Name: '%s' is not an array", name)
			}
			if !node.IsArray && symObj.IsArray {
				return "", fmt.Errorf("Name: '%s' is an array", name)
			}
		}

		// 1: analyze identifier & expression
		st.Analyze(node.Identifier, scope)
		exprType, _ := st.Analyze(node.Expression, scope)
		if exprType != token.INTEGER {
			return "", fmt.Errorf("Name: Array index must be an integer")
		}

		// 2: return the type of the name
		return ReturnType(symObj.ReturnType), nil

	case *ast.ArgumentList:
		for _, a := range node.Arguments {
			st.Analyze(&a, scope)
		}

		return "", nil

	case *ast.Number:
		num := node.Value
		return ReturnType(GetNumberType(num)), nil

	case *ast.String:
		return token.STR, nil

	default:
		// Handle unhandled types
		// ...
	}

	return "", nil
}

func GetProcedureParams(pd ast.ProcedureDeclaration) ([]ast.VariableDeclaration, []string) {
	params := []ast.VariableDeclaration{}
	paramsType := []string{}

	for _, p := range pd.Header.ParameterList.Parameters {
		params = append(params, *p.VariableDeclaration)
		paramsType = append(paramsType, p.VariableDeclaration.Type.Name)
	}
	return params, paramsType
}

// TODO - need to test is
func GetLastProcedureName(scope string) string {
	// Split the scope string by the period character
	elements := strings.Split(scope, ".")

	// todo - search for procedure with following naming convention: <name>-PROC
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

func TypeCompatible(t1, t2 string) bool {
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

func GetNumberType(num string) string {
	if strings.Contains(num, ".") {
		return token.FLOAT
	} else {
		return token.INTEGER
	}
}

func ScopeCompatible(currScope, symbolScope string) bool {
	return strings.HasPrefix(symbolScope, currScope)
}

func AssignmentTypeCompatible(destType, exprType string) bool {
	// If the types are the same, they are compatible
	if destType == exprType {
		return true
	}
	if (destType == token.BOOLEAN || destType == token.INTEGER) && (exprType == token.BOOLEAN || exprType == token.INTEGER) {
		return true
	}
	if (destType == token.FLOAT || destType == token.INTEGER) && (exprType == token.INTEGER || exprType == token.FLOAT) {
		return true
	}

	return false
}

func ArithmeticTypeCompatible(op, type1, type2 string) bool {
	// If either of types are int or float, they are compatible
	if (type1 == token.INTEGER || type1 == token.FLOAT) && (type2 == token.INTEGER || type2 == token.FLOAT) {
		return true
	}
	// If the operation is addition and both types are strings, they are compatible
	if op == "+" && type1 == token.STR && type2 == token.STR {
		return true
	}
	// In all other cases, the types are not compatible
	return false
}

func BitwiseTypeCompatible(op, type1, type2 string) bool {
	// If the operation is bitwise & | not and both types are int, they are compatible
	if (op == "&" || op == "|" || op == "not") && type1 == token.INTEGER && type2 == token.INTEGER {
		return true
	}
	// In all other cases, the types are not compatible
	return false
}

func RelationTypeCompatible(op, type1, type2 string) bool {
	// If the types are int or bool, they are compatible
	if (type1 == token.INTEGER || type1 == token.BOOLEAN) && (type2 == token.INTEGER || type2 == token.BOOLEAN) {
		return true
	}
	// If both types are float, they are compatible
	if type1 == token.FLOAT && type2 == token.FLOAT {
		return true
	}
	// If the operation is == or !=, the types are compatible
	if (op == "==" || op == "!=") && type1 == token.STR && type2 == token.STR {
		return true
	}
	// In all other cases, the types are not compatible
	return false
}

// todo - expr type - type must match
// arithmetic operation + - * / - int & float are compatible
// + - allows strings as well
// relation - int (converted to bool) & bool are compatible, float & float are compatible
// == , != - strings also allowed
// bitwise - & | not - only int type

// assignment stat - type converted to that of destination
// bool & int
// int & float

// If & For expression(condition) - bool type, int is cast to bool
