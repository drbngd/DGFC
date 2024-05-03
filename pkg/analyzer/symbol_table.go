package analyzer

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

type TokenType string
type VariableType string

type Symbol struct {
	Name          string
	NodeType      string
	ReturnType    string // for variables it acts as the data type
	Scope         string
	IsProcedure   bool
	ParamTypeList []string
	IsArray       bool
	ArraySize     string
	Index         int
	//IsProcedure bool
	//ParameterList []string
	//TokenType     string
	//TokenValue    string
	//IsKeyword     bool
}

type SymbolTable struct {
	table        map[string]Symbol
	index        int
	IfElseCount  int
	ForLoopCount int
	errors       []string
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		table:        make(map[string]Symbol),
		index:        0,
		IfElseCount:  0,
		ForLoopCount: 0,
	}
}

func (st *SymbolTable) NewSymbol(name string, nodeType string, returnType string, scope string, isProcedure bool, paramList []string, isArray bool, arraySize string) Symbol {
	return Symbol{
		Name:          name,
		NodeType:      nodeType,
		ReturnType:    returnType,
		Scope:         scope,
		IsProcedure:   isProcedure,
		ParamTypeList: paramList,
		IsArray:       isArray,
		ArraySize:     arraySize,
		Index:         st.GetIndex(),
	}
}

func (st *SymbolTable) AddSymbol(s Symbol) (bool, Symbol) {
	// CASE: if symbol is a procedure, use a different key: s.Name + ".PROC"
	st.IncrementIndex()
	if s.IsProcedure {
		st.table[s.Name+".PROC"] = s
		return true, s
		// CASE: normal non procedure symbol
	} else {
		st.table[s.Name] = s
		return true, s
	}
}

func (st *SymbolTable) IncrementIndex() {
	st.index++
}

func (st *SymbolTable) GetIndex() int {
	return st.index

}

func (st *SymbolTable) IfElseEncountered() {
	st.IfElseCount++
}

func (st *SymbolTable) ForLoopEncountered() {
	st.ForLoopCount++
}

func (st *SymbolTable) PrintSymbolTable() {
	// Convert the map to a slice of keys for sorting
	keys := make([]string, 0, len(st.table))
	for key := range st.table {
		keys = append(keys, key)
	}

	// Sort the keys by the index of their corresponding symbols
	sort.Slice(keys, func(i, j int) bool {
		return st.table[keys[i]].Index < st.table[keys[j]].Index
	})

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Index\tKey\tName\tNodeType\tReturnType\tScope\tIsProcedure\tParamTypeList\tIsArray\tArraySize\t")
	for _, key := range keys {
		value := st.table[key]
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\t%v\t%v\t%v\t%s\t\n", value.Index, key, value.Name, value.NodeType, value.ReturnType, value.Scope, value.IsProcedure, value.ParamTypeList, value.IsArray, value.ArraySize)
	}
	w.Flush()
}
