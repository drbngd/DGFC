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
	Declarations *[]Declaration
	Statements   *[]Statement
}

func (pb *ProgramBody) NodeType() string { return "ProgramBody" }
func (pb *ProgramBody) ToString() string {
	returnString := ""
	for _, d := range *pb.Declarations {
		returnString += d.ToString() + "\n"
	}
	returnString += "begin\n"
	for _, s := range *pb.Statements {
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
	Identifier *Identifier
	Type       *TypeMark
	Parameters *ParameterList
}

func (ph *ProcedureHeader) NodeType() string { return "ProcedureHeader" }
func (ph *ProcedureHeader) ToString() string {
	returnString := "procedure " + ph.Identifier.ToString() + " : " + ph.Type.ToString() + " ("
	if ph.Parameters != nil {
		returnString += ph.Parameters.ToString()
	}
	returnString += ")"
	return returnString
}

type ParameterList struct {
	Parameters *[]Parameter
}

func (pl *ParameterList) NodeType() string { return "ParameterList" }
func (pl *ParameterList) ToString() string {
	returnString := ""
	for _, p := range *pl.Parameters {
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
	Declarations *[]Declaration
	Statements   *[]Statement
}

func (pb *ProcedureBody) NodeType() string { return "ProcedureBody" }
func (pb *ProcedureBody) ToString() string {
	returnString := ""
	for _, d := range *pb.Declarations {
		returnString += d.ToString() + "\n"
	}
	returnString += "begin\n"
	for _, s := range *pb.Statements {
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
	ProcedureName *Identifier
	ArguementList *ArguementList
}

func (pc *ProcedureCall) NodeType() string { return "ProcedureCall" }
func (pc *ProcedureCall) ToString() string {
	return pc.ProcedureName.ToString() + " (" + pc.ArguementList.ToString() + ")"
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
	IsArray      bool
	VariableName *Identifier
	Index        *Expression
}

func (ds *Destination) NodeType() string { return "Destination" }
func (ds *Destination) ToString() string {
	if ds.IsArray {
		return ds.VariableName.ToString() + " " + "[" + ds.Index.ToString() + "]"
	}
	return ds.VariableName.ToString()
}

type IfStatement struct {
	Condition *Expression
	ThenBlock *[]Statement
	ElseBlock *[]Statement
}

func (is *IfStatement) statementNode()   {}
func (is *IfStatement) NodeType() string { return "IfStatement" }
func (is *IfStatement) ToString() string {
	returnString := "if ( " + is.Condition.ToString() + " ) then\n"
	for _, s := range *is.ThenBlock {
		returnString += s.ToString() + "\n"
	}
	if is.ElseBlock != nil {
		returnString += "else\n"
		for _, s := range *is.ElseBlock {
			returnString += s.ToString() + "\n"
		}
	}
	return returnString + "end if"
}

type LoopStatement struct {
	Initialization *AssignmentStatement
	Condition      *Expression
	Body           *[]Statement
}

func (ls *LoopStatement) getStatementToString() string { return ls.ToString() }
func (ls *LoopStatement) NodeType() string             { return "LoopStatement" }
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

	for _, s := range *ls.Body {
		returnString += s.ToString() + "\n"
	}
	return returnString + "end for"
}

type ReturnStatement struct {
	Expression *Expression
}

func (rs *ReturnStatement) getStatementToString() string { return rs.ToString() }
func (rs *ReturnStatement) NodeType() string             { return "ReturnStatement" }
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
	HasNotOperator bool
	ArithOp        *ArithOp
	AndOrList      *[]AndOrExpression
}

func (e *Expression) NodeType() string { return "Expression" }
func (e *Expression) ToString() string {
	returnString := ""
	if e.HasNotOperator {
		returnString += "not "
	}
	if e.ArithOp != nil {
		returnString += e.ArithOp.ToString()
	}
	if e.AndOrList != nil {
		for _, aoe := range *e.AndOrList {
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

type ArithOp struct {
	Relation   *Relation
	AddSubList *[]AddSubExpression
}

func (ao *ArithOp) NodeType() string { return "ArithmeticOperation" }
func (ao *ArithOp) ToString() string {
	returnString := ao.Relation.ToString()
	if ao.AddSubList != nil {
		for _, ase := range *ao.AddSubList {
			returnString += " " + ase.ToString()
		}
	}
	return returnString
}

type AddSubExpression struct {
	Operator           string
	ArithmeticOperator *ArithOp
}

func (ae *AddSubExpression) NodeType() string { return "AddSubExpression" }
func (ae *AddSubExpression) ToString() string {
	return ae.Operator + " " + ae.ArithmeticOperator.ToString()
}

type Relation struct {
	Term                   *Term
	RelationalOperatorList *[]RelationalOperatorExpression
}

func (r *Relation) NodeType() string { return "Relation" }
func (r *Relation) ToString() string {
	returnString := r.Term.ToString()
	if r.RelationalOperatorList != nil {
		for _, roe := range *r.RelationalOperatorList {
			returnString += " " + roe.ToString()
		}
	}
	return returnString
}

type RelationalOperatorExpression struct {
	Operator string
	Relation *Relation
}

func (roe *RelationalOperatorExpression) NodeType() string { return "RelationalOperatorExpression" }
func (roe *RelationalOperatorExpression) ToString() string {
	return roe.Operator + " " + roe.Relation.ToString()
}

type Term struct {
	Factor      *Factor
	MultDivList *[]MultDivExpression
}

func (t *Term) NodeType() string { return "Term" }
func (t *Term) ToString() string {
	returnString := t.Factor.ToString()
	if t.MultDivList != nil {
		for _, mde := range *t.MultDivList {
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
	if f.IsName {
		return f.Name.ToString()
	}
	if f.IsNumber {
		return f.Number.ToString()
	}
	if f.IsString {
		return f.String.ToString()
	}
	if f.IsBool {
		return f.BoolValue
	}
	return ""
}

type Name struct {
	IsArray bool
	Name    *Identifier
	Index   *Expression
}

func (n *Name) NodeType() string { return "Identifier" }
func (n *Name) ToString() string {
	if n.IsArray {
		return n.Name.ToString() + " " + "[" + n.Index.ToString() + "]"
	}
	return n.Name.ToString()

}

type ArguementList struct {
	Arguements *[]Expression
}

func (al *ArguementList) NodeType() string { return "ArguementList" }
func (al *ArguementList) ToString() string {
	returnString := ""
	for _, a := range *al.Arguements {
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
