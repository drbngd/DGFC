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
		print("Analyzing Program\n")
		st.Analyze(node.Header, "GLOBAL")
		st.Analyze(node.Body, "GLOBAL")

		return "", nil

	case *ast.ProgramHeader:
		print("Analyzing Program Header\n")
		st.Analyze(node.Identifier, "GLOBAL")

		return "", nil

	case *ast.ProgramBody:
		print("Analyzing Program Body\n")
		// all declarations in Progam Body are global
		for _, d := range node.Declarations {
			st.Analyze(d, "GLOBAL")
		}
		for _, s := range node.Statements {
			st.Analyze(s, "GLOBAL")
		}

		return "", nil

	case *ast.ProcedureDeclaration:
		print("Analyzing Procedure Declaration\n")
		// 0: get procedure name, return type, scope, and param list
		procName := node.Header.Identifier.Name
		procReturnType := node.Header.Type.Name

		newScope := scope + "." + procName
		if node.IsGlobal {
			newScope = "GLOBAL" + "." + procName
			print(newScope)
			print("\n")
		}
		print("\n")
		print(newScope)
		print("\n")
		print(node.IsGlobal)
		print("\n")

		paramList, paramTypeList := GetProcedureParams(*node)

		// 1: check if procedure is already declared
		_, ok := st.ProcedureExists(procName, newScope)
		if ok {
			return "", fmt.Errorf("Procedure Declaration: Procedure '%s' already declared in the current scope", procName)
		}
		hashKey := st.GetProcedureKey(procName, newScope)
		//if ok && SameIndex(symObj.Index, st.GetIndex()) && InSameScope(symObj.Scope, scope) {
		//	return "", fmt.Errorf("Procedure Declaration: Procedure '%s' already declared in the same scope", procName)
		//}

		// 2: add procedure to symbol table
		sym := st.NewSymbol(procName, "Procedure", procReturnType, scope, true, paramTypeList, false, "0")
		st.table[procName+hashKey+".PROC"] = sym

		// 3: add procedure's params to symbol table
		for _, p := range paramList {
			paramName := p.Identifier.Name
			paramType := p.Type.Name
			isArray := p.IsArray
			if p.IsGlobal {
				return "", fmt.Errorf("Procedure Declaration: Parameters cannot be global")
			}
			sym := st.NewSymbol(paramName, "Procedure Param", paramType, newScope, false, nil, isArray, "0")
			st.AddSymbol(sym)
			st.IncrementIndex()
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
		print("Analyzing Procedure Header\n")
		st.Analyze(node.Identifier, scope)

		return "", nil

	case *ast.ParameterList:
		print("Analyzing Parameter List\n")
		for _, p := range node.Parameters {
			st.Analyze(&p, scope)
		}

		return "", nil

	case *ast.Parameter:
		print("Analyzing Parameter\n")
		st.Analyze(node.VariableDeclaration, scope)

		return "", nil

	case *ast.ProcedureBody:
		print("Analyzing Procedure Body\n")
		for _, d := range node.Declarations {
			st.Analyze(d, scope)
		}
		for _, s := range node.Statements {
			st.Analyze(s, scope)
		}

		return "", nil

	case *ast.VariableDeclaration:
		print("Analyzing Variable Declaration\n")

		print("on step 0\n")
		// 0: get variable name, type, and scope
		name := &node.Identifier.Name
		print(*name)
		print("\non step 0.1\n")
		varType := &node.Type.Name
		print("on step 0.2\n")
		arrBound := "0"
		if node.IsArray {
			arrBound = node.Bound.Value.Value
			if GetNumberType(arrBound) != token.INTEGER {
				return "", fmt.Errorf("Variable Declaration: Array bound must be an integer")
			}
		}
		print("on step 0.3\n")
		newScope := scope
		if node.IsGlobal {
			newScope = "GLOBAL"
		}

		print("on step 1\n")
		// 1: check if already exists in the symbol table with the same scope
		varSym, ok := st.table[*name]
		if ok && ScopeCompatible(newScope, varSym.Scope) && SameIndex(varSym.Index, st.GetIndex()) {
			return "", fmt.Errorf("Variable Declaration: Variable '%s' already declared in the same scope", name)
		}

		print("on step 2\n")
		// 2: add to symbol table
		sym := st.NewSymbol(*name, "Variable", *varType, newScope, false, nil, node.IsArray, arrBound)
		st.AddSymbol(sym)

		print("on step 3\n")
		// 3: return the type of the variable
		return ReturnType(*varType), nil

	case *ast.TypeMark:
		print("Analyzing TypeMark\n")
		// 0: get type name
		typeName := node.Name

		// 1: check if type is valid
		if typeName != token.INTEGER && typeName != token.FLOAT && typeName != token.BOOLEAN && typeName != token.STRING {
			return "", fmt.Errorf("TypeMark: Invalid type '%s'", typeName)
		}

		// 2: return the type
		return ReturnType(typeName), nil

	case *ast.Bound:
		print("Analyzing Bound\n")
		_, err := st.Analyze(node.Value, scope)
		if err != nil {
			return "", err
		}
		if GetNumberType(node.Value.Value) != token.INTEGER {
			return "", fmt.Errorf("Bound: Bound must be an integer")
		}

		return "", nil

	case *ast.ProcedureCall:
		print("Analyzing Procedure Call\n")
		// 0: get procedure name
		procName := node.Identifier.Name

		// 1: check if procedure is already declared in the current scope, get return type and param list
		symObj, ok := st.table[procName+".PROC"]
		if !ok {
			return "", fmt.Errorf("Procedure Call: Procedure '%s' not found in symbol table", procName)
		} else {
			if !ScopeCompatible(scope, symObj.Scope) {
				return "", fmt.Errorf("Procedure Call: Procedure '%s' not in scope", procName)
			}
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
		print("Analyzing Assignment Statement\n")
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
		print("Analyzing Destination\n")
		// 0: check if present in symbol table & in current scope, get type
		destName := node.Identifier.Name
		symObj, ok := st.table[destName]
		if !ok {
			return "", fmt.Errorf("Destination: '%s' not found in symbol table", destName)
		} else {
			if !ScopeCompatible(scope, symObj.Scope) {
				return "", fmt.Errorf("Destination: '%s' not in scope", destName)
			}
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
		print("Analyzing If Statement\n")
		// 0: generate new scope string & increment if/else count
		newScope := scope + ".IF."
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
		print("Analyzing Loop Statement\n")
		// 0: generate new scope string & increment for loop count
		newScope := scope + ".LOOP."
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
		print("Analyzing Return Statement\n")
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
		print("Analyzing Identifier\n")
		// maybe there is more to do here
		return "", nil

	case *ast.Expression:
		print("Analyzing Expression\n")
		print("on step 0\n")
		if node == nil {
			print("Expression: nil\n")
			return "", nil
		}

		if node.ArithOp == nil {
			print("ArithOp: nil\n")
			//return "", nil
		}
		// 0: analyze arithmetic operation & type check
		arithopType, _ := st.Analyze(node.ArithOp, scope)
		// if not isn't present and no list of bitwise and/or expressions is present, then type doesn't matter
		print("on step 0.1\n")
		if !node.IsNot && len(node.AndOrList) == 0 {
			return arithopType, nil
		}

		print("on step 1\n")
		// 1: if an bitwise and/or/not list is present, then arithmetic operation must be of type int
		if arithopType != token.INTEGER {
			return "", fmt.Errorf("Expression: Bitwise operations must be of type int")
		}
		// TODO(codegen): if type is float, we will truncate it to int in code generation

		print("on step 2\n")
		// 2: analyze AndOr list expression & type check for int
		for _, e := range node.AndOrList {
			st.Analyze(&e, scope)
		}

		return ReturnType(token.INTEGER), nil

	case *ast.AndOrExpression:
		print("Analyzing AndOr Expression\n")
		exprType, _ := st.Analyze(node.Expression, scope)
		if exprType != token.INTEGER {
			return "", fmt.Errorf("Bitwise Expression: Expression must be of type int")
		}
		return token.INTEGER, nil

	case *ast.ArithmeticOperation:
		print("Analyzing Arithmetic Operation\n")
		// 0: analyze relation & type check
		relType, _ := st.Analyze(node.Relation, scope)
		arithopReturnType := relType
		// if no list of +, - expressions is present, then type doesn't matter
		if len(node.AddSubList) == 0 {
			return arithopReturnType, nil
		}

		// 1: if a list of +, - expressions is present, then relation must be of type int or float
		if relType != token.INTEGER && relType != token.FLOAT {
			return "", fmt.Errorf("Arithmetic Operation: Relation must be of type int or float")
		}

		// 2: analyze AddSub list expression & type check for int or float
		for _, e := range node.AddSubList {
			exprType, _ := st.Analyze(&e, scope)
			if exprType == token.FLOAT { // if float is present, then int is cast to float
				arithopReturnType = token.FLOAT
			}
		}

		// 3: return the type of the arithmetic operation
		return arithopReturnType, nil

	case *ast.AddSubExpression:
		print("Analyzing AddSub Expression\n")
		exprType, _ := st.Analyze(node.ArithmeticOperation, scope)
		if exprType != token.INTEGER && exprType != token.FLOAT {
			return "", fmt.Errorf("AddSub Expression: Arithmetic Operation must be of type int or float")
		}
		return exprType, nil

	case *ast.Relation:
		print("Analyzing Relation\n")
		// float - float, int | bool with int | bool, string with string
		// 0: analyze the term
		termType, _ := st.Analyze(node.Term, scope)
		// if no list of relational expressions is present, then type doesn't matter
		if len(node.RelationalOperationList) == 0 {
			return termType, nil
		}

		// 1: if a list of relational expressions is present, then we must to proper type checking
		leftTermType := string(termType)
		for _, e := range node.RelationalOperationList {
			relOpType, _ := st.Analyze(&e, scope)
			rightTermType := string(relOpType)
			if !RelationTypeCompatible(e.Operator, leftTermType, rightTermType) {
				return "", fmt.Errorf("Relation: Type mismatch")
			}
			leftTermType = rightTermType
		}

		// 2: return the type of the relation
		return token.BOOLEAN, nil

	case *ast.RelationalExpression:
		print("Analyzing Relational Expression\n")
		// 0: analyze the relation
		relType, _ := st.Analyze(node.Term, scope)
		// strings can only be used with == and !=
		if relType == token.STR && (node.Operator != "==" && node.Operator != "!=") {
			return "", fmt.Errorf("Relational Expression: Strings can only be used with == and !=")
		}
		return relType, nil

	case *ast.Term:
		print("Analyzing Term\n")
		// 0: analyze the factor
		factorType, _ := st.Analyze(node.Factor, scope)
		termReturnType := factorType
		if len(node.MultDivList) == 0 {
			return factorType, nil
		}

		// 1: if a list of *, / expressions is present, then factor must be of type int or float
		if factorType != token.INTEGER && factorType != token.FLOAT {
			return "", fmt.Errorf("Term: Factor must be of type int or float")
		}

		// 2: analyze the mult div list
		for _, e := range node.MultDivList {
			multdivType, _ := st.Analyze(&e, scope)
			if multdivType == token.FLOAT {
				// if float is present, then int is cast to float
				termReturnType = multdivType
			}
		}

		// 3: return the type of the term
		return termReturnType, nil

	case *ast.MultDivExpression:
		print("Analyzing MultDiv Expression\n")
		// 0: analuze mult, div term
		termType, _ := st.Analyze(node.Factor, scope)
		if termType != token.INTEGER && termType != token.FLOAT {
			return "", fmt.Errorf("MultDiv Expression: Factor must be of type int or float")
		}
		return termType, nil

	case *ast.Factor:
		print("Analyzing Factor\n")
		if node.IsExpression {
			exprType, _ := st.Analyze(node.Expression, scope)
			return exprType, nil
		}

		if node.IsProcedureCall {
			procType, _ := st.Analyze(node.ProcedureCall, scope)
			return procType, nil
		}

		if node.IsName {
			nameType, _ := st.Analyze(node.Name, scope)
			return nameType, nil
		}

		if node.IsNumber {
			numType, _ := st.Analyze(node.Number, scope)
			return numType, nil
		}

		if node.IsString {
			st.Analyze(node.String, scope)
		}

		if node.IsBool {
			if node.BoolValue == "true" || node.BoolValue == "false" {
				return token.BOOLEAN, nil
			}
			return "", fmt.Errorf("Factor: Invalid boolean value")
		}

		return "", nil

	case *ast.Name:
		print("Analyzing Name\n")
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
		print("Analyzing Argument List\n")
		for _, a := range node.Arguments {
			st.Analyze(&a, scope)
		}

		return "", nil

	case *ast.Number:
		print("Analyzing Number\n")
		num := node.Value
		return ReturnType(GetNumberType(num)), nil

	case *ast.String:
		print("Analyzing String\n")
		return token.STR, nil

	default:
		print("Analyzing Default\n")
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

func SameIndex(ind1, ind2 int) bool {
	return ind1 == ind2
}

func InSameScope(scope1, scope2 string) bool {
	return scope1 == scope2
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

func (st *SymbolTable) GetProcedureKey(name, scope string) string {
	symObj, ok := st.table[name+".PROC"]
	if ok && ScopeCompatible(symObj.Scope, scope) {
		count := CountOccurrences(scope, name)
		hashKey := strings.Repeat(".PROC", count)
		return hashKey
	} else {
		return ""
	}
}

func CountOccurrences(scope, name string) int {
	// Split the scope string by the period character
	elements := strings.Split(scope, ".")

	// Initialize a counter
	count := 0

	// Iterate over the elements
	for _, element := range elements {
		// If the element matches the name, increment the counter
		if element == name {
			count++
		}
	}

	// Return the count
	return count
}

func (st *SymbolTable) ProcedureExists(name, scope string) (string, bool) {
	// get count
	count := CountOccurrences(scope, name)
	key := strings.Repeat(".PROC", count)
	_, ok := st.table[name+key+".PROC"]
	if ok {
		return key, true
	}
	return key, false
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
