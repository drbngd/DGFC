package ast

type Node interface {
	NodeType() string
	ToString() string
}

type Program struct {
	Header *ProgramHeader
	Body   *ProgramBody
}

func (p *Program) NodeType() string { return "Program" }
func (p *Program) ToString() string {
	return p.Header.ToString() + "\n" + p.Body.ToString() + "."
}

type ProgramHeader struct {
	Identifier *Identifier
}

func (ph *ProgramHeader) NodeType() string { return "ProgramHeader" }
func (ph *ProgramHeader) ToString() string { return "program " + ph.Identifier.ToString() + "is" }

type ProgramBody struct {
	Declarations []Declaration
	Statements   []Statement
}

func (pb *ProgramBody) NodeType() string { return "ProgramBody" }
func (pb *ProgramBody) ToString() string {
	returnString := ""
	for _, d := range pb.Declarations {
		returnString += d.ToString() + "\n"
	}
	returnString += "begin\n"
	for _, s := range pb.Statements {
		returnString += s.ToString() + "\n"
	}
	returnString += "end program"
	return returnString
}

type Declaration interface {
	Node
	declarationNode()
}

type ProcedureDeclaration struct {
	IsGlobal bool
	Header   *ProcedureHeader
	Body     *ProcedureBody
}

func (pd *ProcedureDeclaration) declarationNode() {}
func (pd *ProcedureDeclaration) NodeType() string { return "ProcedureDeclaration" }
func (pd *ProcedureDeclaration) ToString() string {
	returnString := ""
	if pd.IsGlobal {
		returnString += "global "
	}
	returnString += pd.Header.ToString() + "\n" + pd.Body.ToString()
	return returnString
}

type ProcedureHeader struct {
	Identifier    *Identifier
	Type          *TypeMark
	ParameterList *ParameterList
}

func (ph *ProcedureHeader) NodeType() string { return "ProcedureHeader" }
func (ph *ProcedureHeader) ToString() string {
	returnString := "procedure " + ph.Identifier.ToString() + " : " + ph.Type.ToString() + " ("
	if ph.ParameterList != nil {
		returnString += ph.ParameterList.ToString()
	}
	returnString += ")"
	return returnString
}

type ParameterList struct {
	Parameters []Parameter
}

func (pl *ParameterList) NodeType() string { return "ParameterList" }
func (pl *ParameterList) ToString() string {
	returnString := ""
	for _, p := range pl.Parameters {
		returnString += p.ToString() + ", "
	}
	return returnString
}

type Parameter struct {
	VariableDeclaration *VariableDeclaration
}

func (p *Parameter) NodeType() string { return "Parameter" }
func (p *Parameter) ToString() string {
	return p.VariableDeclaration.ToString()
}

type ProcedureBody struct {
	Declarations []Declaration
	Statements   []Statement
}

func (pb *ProcedureBody) NodeType() string { return "ProcedureBody" }
func (pb *ProcedureBody) ToString() string {
	returnString := ""
	for _, d := range pb.Declarations {
		returnString += d.ToString() + "\n"
	}
	returnString += "begin\n"
	for _, s := range pb.Statements {
		returnString += s.ToString() + "\n"
	}
	returnString += "end procedure"
	return returnString
}

type VariableDeclaration struct {
	IsGlobal   bool
	IsArray    bool
	Identifier *Identifier
	Type       *TypeMark
	Bound      *Bound
}

func (vd *VariableDeclaration) declarationNode() {}
func (vd *VariableDeclaration) NodeType() string { return "VariableDeclaration" }
func (vd *VariableDeclaration) ToString() string {
	returnString := ""
	if vd.IsGlobal {
		returnString += "global "
	}
	returnString += "variable " + vd.Identifier.ToString() + " : " + vd.Type.ToString()
	if vd.IsArray {
		returnString += "[" + vd.Bound.ToString() + "]"
	}
	return returnString
}

type TypeMark struct {
	Name string
}

func (tm *TypeMark) NodeType() string { return "TypeMark" }
func (tm *TypeMark) ToString() string { return tm.Name }

type Bound struct {
	Value *Number
}

func (b *Bound) NodeType() string { return "Bound" }
func (b *Bound) ToString() string { return b.Value.ToString() }

type ProcedureCall struct {
	Identifier    *Identifier
	ArguementList *ArgumentList
}

func (pc *ProcedureCall) NodeType() string { return "ProcedureCall" }
func (pc *ProcedureCall) ToString() string {
	return pc.Identifier.ToString() + " (" + pc.ArguementList.ToString() + ")"
}

type Statement interface {
	Node
	statementNode()
}

type AssignmentStatement struct {
	Destination *Destination
	Expression  *Expression
}

func (as *AssignmentStatement) statementNode()   {}
func (as *AssignmentStatement) NodeType() string { return "AssignmentStatement" }
func (as *AssignmentStatement) ToString() string {
	return as.Destination.ToString() + " := " + as.Expression.ToString()
}

type Destination struct {
	IsArray    bool
	Identifier *Identifier
	Expression *Expression
}

func (ds *Destination) NodeType() string { return "Destination" }
func (ds *Destination) ToString() string {
	if ds.IsArray {
		return ds.Identifier.ToString() + " " + "[" + ds.Expression.ToString() + "]"
	}
	return ds.Identifier.ToString()
}

type IfStatement struct {
	Condition *Expression
	ThenBlock []Statement
	ElseBlock []Statement
}

func (is *IfStatement) statementNode()   {}
func (is *IfStatement) NodeType() string { return "IfStatement" }
func (is *IfStatement) ToString() string {
	returnString := "if ( " + is.Condition.ToString() + " ) then\n"
	for _, s := range is.ThenBlock {
		returnString += s.ToString() + "\n"
	}
	if is.ElseBlock != nil {
		returnString += "else\n"
		for _, s := range is.ElseBlock {
			returnString += s.ToString() + "\n"
		}
	}
	return returnString + "end if"
}

type LoopStatement struct {
	Initialization *AssignmentStatement
	Condition      *Expression
	Body           []Statement
}

func (ls *LoopStatement) statementNode()   {}
func (ls *LoopStatement) NodeType() string { return "LoopStatement" }
func (ls *LoopStatement) ToString() string {
	returnString := "for ("
	if ls.Initialization != nil {
		returnString += ls.Initialization.ToString() + " ;"
	} else {
		returnString += " ;"
	}

	if ls.Condition != nil {
		returnString += ls.Condition.ToString() + " )"
	} else {
		returnString += " )"
	}

	for _, s := range ls.Body {
		returnString += s.ToString() + "\n"
	}
	return returnString + "end for"
}

type ReturnStatement struct {
	Expression *Expression
}

func (rs *ReturnStatement) statementNode()   {}
func (rs *ReturnStatement) NodeType() string { return "ReturnStatement" }
func (rs *ReturnStatement) ToString() string {
	if rs.Expression != nil {
		return "return " + rs.Expression.ToString()
	}
	return "return"
}

type Identifier struct {
	Name string // implement rules for valid identifier names
}

func (i *Identifier) NodeType() string { return "Identifier" }
func (i *Identifier) ToString() string { return i.Name }

type Expression struct {
	IsNot     bool
	ArithOp   *ArithmeticOperation
	AndOrList []AndOrExpression
}

func (e *Expression) NodeType() string { return "Expression" }
func (e *Expression) ToString() string {
	returnString := ""
	if e.IsNot {
		returnString += "not "
	}
	if e.ArithOp != nil {
		returnString += e.ArithOp.ToString()
	}
	if e.AndOrList != nil {
		for _, aoe := range e.AndOrList {
			returnString += " " + aoe.ToString()
		}
	}
	return returnString
}

type AndOrExpression struct {
	Operator   string
	Expression *Expression
}

func (ae *AndOrExpression) NodeType() string { return "AndOrExpression" }
func (ae *AndOrExpression) ToString() string {
	return ae.Operator + " " + ae.Expression.ToString()
}

type ArithmeticOperation struct {
	Relation   *Relation
	AddSubList []AddSubExpression
}

func (ao *ArithmeticOperation) NodeType() string { return "ArithmeticOperation" }
func (ao *ArithmeticOperation) ToString() string {
	returnString := ao.Relation.ToString()
	if ao.AddSubList != nil {
		for _, ase := range ao.AddSubList {
			returnString += " " + ase.ToString()
		}
	}
	return returnString
}

type AddSubExpression struct {
	Operator            string
	ArithmeticOperation *ArithmeticOperation
}

func (ae *AddSubExpression) NodeType() string { return "AddSubExpression" }
func (ae *AddSubExpression) ToString() string {
	return ae.Operator + " " + ae.ArithmeticOperation.ToString()
}

type Relation struct {
	Term                    *Term
	RelationalOperationList []RelationalExpression
}

func (r *Relation) NodeType() string { return "Relation" }
func (r *Relation) ToString() string {
	returnString := r.Term.ToString()
	if r.RelationalOperationList != nil {
		for _, roe := range r.RelationalOperationList {
			returnString += " " + roe.ToString()
		}
	}
	return returnString
}

type RelationalExpression struct {
	Operator string
	Term     *Term
}

func (roe *RelationalExpression) NodeType() string { return "RelationalExpression" }
func (roe *RelationalExpression) ToString() string {
	return roe.Operator + " " + roe.Term.ToString()
}

type Term struct {
	Factor      *Factor
	MultDivList []MultDivExpression
}

func (t *Term) NodeType() string { return "Term" }
func (t *Term) ToString() string {
	returnString := t.Factor.ToString()
	if t.MultDivList != nil {
		for _, mde := range t.MultDivList {
			returnString += " " + mde.ToString()
		}
	}
	return returnString
}

type MultDivExpression struct {
	Operator string
	Factor   *Factor
}

func (mde *MultDivExpression) NodeType() string { return "MultDivExpression" }
func (mde *MultDivExpression) ToString() string {
	return mde.Operator + " " + mde.Factor.ToString()
}

type Factor struct {
	IsExpression    bool
	IsProcedureCall bool
	IsNegative      bool
	IsName          bool
	IsNumber        bool
	IsString        bool
	IsBool          bool

	Expression    *Expression
	ProcedureCall *ProcedureCall
	Name          *Name
	Number        *Number
	String        *String
	BoolValue     string
}

func (f *Factor) NodeType() string { return "Factor" }
func (f *Factor) ToString() string {
	if f.IsExpression {
		return f.Expression.ToString()
	}
	if f.IsProcedureCall {
		return f.ProcedureCall.ToString()
	}

	prefixStr := ""
	if f.IsNegative {
		prefixStr = "-"
	}

	if f.IsName {
		return prefixStr + f.Name.ToString()
	}
	if f.IsNumber {
		return prefixStr + f.Number.ToString()
	}
	if f.IsString {
		return prefixStr + f.String.ToString()
	}
	if f.IsBool {
		return prefixStr + f.BoolValue
	}

	return ""
}

type Name struct {
	IsArray    bool
	Identifier *Identifier
	Expression *Expression
}

func (n *Name) NodeType() string { return "Identifier" }
func (n *Name) ToString() string {
	if n.IsArray {
		return n.Identifier.ToString() + " " + "[" + n.Expression.ToString() + "]"
	}
	return n.Identifier.ToString()

}

type ArgumentList struct {
	Arguments []Expression
}

func (al *ArgumentList) NodeType() string { return "ArgumentList" }
func (al *ArgumentList) ToString() string {
	returnString := ""
	for _, a := range al.Arguments {
		returnString += a.ToString() + ", "
	}
	return returnString
}

type Number struct {
	Value string // implement rules for valid number values
}

func (n *Number) NodeType() string { return "Number" }
func (n *Number) ToString() string { return n.Value }

type String struct {
	Value string // implement rules for valid string values
}

func (s *String) NodeType() string { return "String" }
func (s *String) ToString() string { return s.Value }

//func TreePrint(node Node, indent string) {
//	switch n := node.(type) {
//	case *Program:
//		fmt.Println(indent + "📦 PROGRAM")
//		TreePrint(n.Header, indent+"  ")
//		TreePrint(n.Body, indent+"  ")
//
//	case *ProgramHeader:
//		fmt.Println(indent + "├─ PROGRAM HEADER " + n.Identifier.Name)
//
//	case *ProgramBody:
//		fmt.Println(indent + "├─ PROGRAM BODY")
//		for _, d := range n.Declarations {
//			TreePrint(d, indent+"  ")
//		}
//		for _, s := range n.Statements {
//			TreePrint(s, indent+"  ")
//		}
//
//	case *ProcedureDeclaration:
//		fmt.Print(indent + "├─ PROCEDURE ")
//		if n.IsGlobal {
//			fmt.Print("(global) ")
//		}
//		fmt.Println(n.Header.Identifier.Name)
//		TreePrint(n.Header, indent+"  ")
//		TreePrint(n.Body, indent+" │ ")
//
//	case *ProcedureHeader:
//		fmt.Print(indent + "├─ PARAMETERS ")
//		TreePrint(n.ParameterList, indent+"  ")
//		fmt.Println(indent + "└─ RETURN TYPE " + n.Type.Name)
//
//	case *VariableDeclaration:
//		fmt.Print(indent + "├─ VARIABLE ")
//		if n.IsGlobal {
//			fmt.Print("(global) ")
//		}
//		fmt.Print(n.Identifier.Name, ":", n.Type.Name)
//		if n.IsArray {
//			fmt.Print("[", n.Bound, "]")
//		}
//		fmt.Println()
//
//	case *ParameterList:
//		for _, p := range n.Parameters {
//			TreePrint(p, indent+"  ")
//		}
//
//	case *Parameter:
//		TreePrint(n.VariableDeclaration, indent)
//
//	case *ProcedureBody:
//		fmt.Println(indent + "├─ PROCEDURE BODY")
//		for _, d := range n.Declarations {
//			TreePrint(d, indent+"  ")
//		}
//		for _, s := range n.Statements {
//			TreePrint(s, indent+"  ")
//		}
//
//	case *AssignmentStatement:
//		fmt.Println(indent + "├─ ASSIGN")
//		TreePrint(n.Destination, indent+" │ ")
//		TreePrint(n.Expression, indent+" │ ")
//
//	case *IfStatement:
//		fmt.Println(indent + "├─ IF")
//		TreePrint(n.Condition, indent+" │ ")
//		fmt.Println(indent + "├─ THEN")
//		for _, s := range n.ThenBlock {
//			TreePrint(s, indent+" │ ")
//		}
//		if n.ElseBlock != nil {
//			fmt.Println(indent + "├─ ELSE")
//			for _, s := range n.ElseBlock {
//				TreePrint(s, indent+" │ ")
//			}
//		}
//
//	case *LoopStatement:
//		fmt.Println(indent + "├─ LOOP")
//		if n.Initialization != nil {
//			TreePrint(n.Initialization, indent+" │ ")
//		}
//		TreePrint(n.Condition, indent+" │ ")
//		fmt.Println(indent + "├─ BODY")
//		for _, s := range n.Body {
//			TreePrint(s, indent+" │ ")
//		}
//
//	case *ReturnStatement:
//		fmt.Println(indent + "├─ RETURN")
//		if n.Expression != nil {
//			TreePrint(n.Expression, indent+" │ ")
//		}
//
//	case *ProcedureCall:
//		fmt.Println(indent + "├─ PROCEDURE CALL " + n.Identifier.Name)
//		TreePrint(n.ArguementList, indent+" │ ")
//
//	case *Identifier:
//		fmt.Println(indent+"└─ ", n.Name)
//
//	case *Expression:
//		fmt.Println(indent + "├─ EXPRESSION")
//		if n.IsNot {
//			fmt.Println(indent + " │ └─ NOT")
//		}
//		if n.ArithOp != nil {
//			TreePrint(n.ArithOp, indent+" │ ")
//		}
//		for _, aoe := range n.AndOrList {
//			TreePrint(aoe, indent+" │ ")
//		}
//
//	case *AndOrExpression:
//		fmt.Println(indent + " │ └─ " + n.Operator)
//		TreePrint(n.Expression, indent+"   ")
//
//	case *ArithmeticOperation:
//		fmt.Println(indent + " │ └─ ARITHMETIC OPERATION")
//		TreePrint(n.Relation, indent+"   ")
//		for _, ase := range n.AddSubList {
//			TreePrint(ase, indent+"   ")
//		}
//
//	case *AddSubExpression:
//		fmt.Println(indent + "   └─ " + n.Operator)
//		TreePrint(n.ArithmeticOperation, indent+"     ")
//
//	case *Relation:
//		fmt.Println(indent + "   └─ RELATION")
//		TreePrint(n.Term, indent+"     ")
//		for _, roe := range n.RelationalOperationList {
//			TreePrint(roe, indent+"     ")
//		}
//
//	case *RelationalExpression:
//		fmt.Println(indent + "     └─ " + n.Operator)
//		TreePrint(n.Term, indent+"       ")
//
//	case *Term:
//		fmt.Println(indent + "     └─ TERM")
//		TreePrint(n.Factor, indent+"       ")
//		for _, mde := range n.MultDivList {
//			TreePrint(mde, indent+"       ")
//		}
//
//	case *MultDivExpression:
//		fmt.Println(indent + "       └─ " + n.Operator)
//		TreePrint(n.Factor, indent+"         ")
//
//	case *Factor:
//		fmt.Print(indent + "       └─ ")
//		if n.IsNegative {
//			fmt.Print("-")
//		}
//		if n.IsName {
//			TreePrint(n.Name, indent+"         ")
//		} else if n.IsNumber {
//			fmt.Println("NUMBER " + n.Number.Value)
//		} else if n.IsString {
//			fmt.Println("STRING " + n.String.Value)
//		} else if n.IsBool {
//			fmt.Println("BOOLEAN " + n.BoolValue)
//		} else if n.IsExpression {
//			TreePrint(n.Expression, indent+"         ")
//		} else if n.IsProcedureCall {
//			TreePrint(n.ProcedureCall, indent+"         ")
//		}
//
//	case *Name:
//		fmt.Print(indent + "         └─ ") // Adjusted for alignment
//		fmt.Println(n.Identifier.Name)
//		if n.IsArray {
//			fmt.Println(indent + "           └─ ARRAY INDEX")
//			TreePrint(n.Expression, indent+"             ")
//		}
//
//	case *ArgumentList:
//		fmt.Println(indent + " │ └─ ARGUMENTS")
//		for _, exp := range n.Arguments {
//			TreePrint(exp, indent+"   ")
//		}
//
//	case *Number:
//		fmt.Println(indent + "└─ NUMBER " + n.Value)
//
//	case *String:
//		fmt.Println(indent + "└─ STRING " + n.Value)
//
//	case *TypeMark:
//		fmt.Println(indent+"└─ TYPE", n.Name)
//
//	case *Bound:
//		fmt.Println(indent+"└─ BOUND", n.Value.Value)
//
//	default:
//		fmt.Printf(indent+"└─ %s (unhandled type)\n", n.NodeType())
//	}
//}
