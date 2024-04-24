package analyzer

type TokenType string
type VariableType string

type Symbol struct {
	Name        string
	NodeType    string
	ReturnType  string // for variables it acts as the data type
	Scope       string
	IsProcedure bool
	ParamList   []string
	IsArray     bool
	ArraySize   string
	//IsProcedure bool
	//ParameterList []string
	//TokenType     string
	//TokenValue    string
	//IsKeyword     bool
}

type SymbolTable struct {
	table map[string]Symbol
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		table: make(map[string]Symbol),
	}
}

func NewSymbol(name string, nodeType string, returnType string, scope string, isProcedure bool, paramList []string, isArray bool, arraySize string) Symbol {
	return Symbol{
		Name:        name,
		NodeType:    nodeType,
		ReturnType:  returnType,
		Scope:       scope,
		IsProcedure: isProcedure,
		ParamList:   paramList,
		IsArray:     isArray,
		ArraySize:   arraySize,
	}
}

func (st *SymbolTable) AddSymbol(s Symbol) (bool, Symbol) {
	// CASE: if symbol is a procedure, use a different key: s.Name + ".PROC"
	if s.IsProcedure {
		if _, ok := st.table[s.Name+".PROC"]; ok {
			return ok, s
		} else {
			st.table[s.Name+".PROC"] = s
			return true, s
		}
		// CASE: normal non procedure symbol
	} else {
		if _, ok := st.table[s.Name]; ok {
			return false, s
		} else {
			st.table[s.Name] = s
			return true, s
		}
	}
}

//func (st *SymbolTable) PrintTable() {
//	fmt.Println("Symbol Table:")
//	fmt.Println("-----------------------------------------------------------------------------")
//	fmt.Printf("%-15s %-15s %-15s %-15s %-15s %-10s %-10s %-10s %-10s\n", "Name", "NodeType", "DataType", "ReturnType", "Scope", "IsProcedure", "ParamList", "IsArray", "ArraySize")
//	fmt.Println("-----------------------------------------------------------------------------")
//	for _, symbol := range st.table {
//		fmt.Printf("%-15s %-15v %-15s %-15s %-15s %-10v %-10v %-10v %-10d\n", symbol.Name, symbol.NodeType, symbol.DataType, symbol.ReturnType, symbol.Scope, symbol.IsProcedure, symbol.ParamList, symbol.IsArray, symbol.ArraySize)
//	}
//	fmt.Println("-----------------------------------------------------------------------------")
//}
