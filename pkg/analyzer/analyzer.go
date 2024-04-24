package analyzer

import (
	"DGFC/pkg/ast"
	"fmt"
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
			//nodeType, err := st.Analyze(s, "GLOBAL")
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

		// get destination name
		destName := node.Destination.Identifier.Name

		// check if in symbol table - TODO: check if this is necessary
		_, ok := st.table[destName]
		if !ok {
			return "", fmt.Errorf("Assignment Statement: Destination '%s' not found.", destName)
		}

		// get expression type

	case *ast.Destination:
		// Handle Destination node
		// ...

	case *ast.IfStatement:
		// Handle IfStatement node
		// ...

	case *ast.LoopStatement:
		// Handle LoopStatement node
		// ...

	case *ast.ReturnStatement:
		// Handle ReturnStatement node
		// ...

	case *ast.Identifier:
		// Handle Identifier node
		// ...

	case *ast.Expression:
		// Handle Expression node
		// ...

	case *ast.AndOrExpression:
		// Handle AndOrExpression node
		// ...

	case *ast.ArithmeticOperation:
		// Handle ArithmeticOperation node
		// ...

	case *ast.AddSubExpression:
		// Handle AddSubExpression node
		// ...

	case *ast.Relation:
		// Handle Relation node
		// ...

	case *ast.RelationalExpression:
		// Handle RelationalExpression node
		// ...

	case *ast.Term:
		// Handle Term node
		// ...

	case *ast.MultDivExpression:
		// Handle MultDivExpression node
		// ...

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
