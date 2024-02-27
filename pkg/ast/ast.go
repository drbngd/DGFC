package ast

// fundamental node interface
type Node interface {
	Type() string
	Value() string
	Children() []Node
	Position() (line, col int)
}

// program node
type Program struct {
	Header *ProgramHeader
	Body   *ProgramBody
}

//func (p *Program) Children() []Node          {}
//func (p *Program) Position() (line, col int) {} nn

// program header node
type ProgramHeader struct {
	Name *Identifier
}

// program body node
type ProgramBody struct {
	Declaration []Declaration // TODO - a little iffy on this
	Statement   []Statement   // TODO - a little iffy on this as well
}

// declaration interface -  as there are two types of declarations
type Declaration interface {
	Type() string
	Node
}

// procedure declaration node
type ProcedureDeclaration struct {
	Header *ProcedureHeader
	Body   *ProcedureBody
}

// procedure header node
type ProcedureHeader struct {
	Name          *Identifier
	ReturnType    *TypeMark
	ParameterList []*VariableDeclaration
}

// parameter list <=> parameter <=> variable declaration node
// TODO - should this be fine to implement a list as well
// TODO - might need to remove left recursion
type VariableDeclaration struct {
	Name      *Identifier
	DataType  *TypeMark // given grammar calls it Type_Mark
	Array     bool
	ArraySize *Bound // in the grammar it is bound
}

// type mark node
type TypeMark struct {
	Name      *Identifier
	IfInteger bool
	IfFloat   bool
	IfString  bool
	IfBoolean bool
}

// bound node
type Bound Number

// procedure body node
type ProcedureBody struct {
	Declaration []Declaration // TODO - a little iffy on this
	Statement   []Statement   // TODO - a little iffy on this as well
}

// procedure call statement node
type ProcedureCallStatement struct {
	Name         *Identifier
	ArgumentList []*ArgumentList
}

// argument list node
type ArgumentList struct {
	ExpressionList       []*Expression
	IfMultipleExpression bool
}

// statement interface
type Statement interface {
	Type() string
	Node
}

// assignment statement node
type AssignmentStatement struct {
	Destination *Destination
	Source      *Expression // TODO - a little iffy on this
}

// destination node
type Destination struct {
	Name  *Identifier
	Array bool
	Index *Expression
}

// if statement node
type IfStatement struct {
	IfExpression *Expression
	ThenBody     []Statement // since there is a *
	ElseBody     []Statement // since there is a *
}

// loop statement node
type LoopStatement struct {
	Initialization *AssignmentStatement
	Condition      *Expression
	Body           []Statement
}

// return statement node
type ReturnStatement struct {
	ReturnValue *Expression
}

// identifier node TODO - do I need this?

// expression interface
// TODO - eliminate left recursion
type Expression struct {
	IfPrecedingAND    bool
	IfPrecedingOR     bool
	IfPrecedingCOMMA  bool
	ArithmeticOperand *ArithmeticOperand
	ExpressionList    []*Expression
}

// arithmetic operand node
type ArithmeticOperand struct {
	IfPrecedingNOT bool
	Relation       *Relation
	OperandList    []*ArithmeticOperand
}

// relation node
type Relation struct {
	FirstTerm         *Term
	FollowingTermList []*Term
}

// term node

type Term struct {
	IfPrecedingLT       bool
	IfPrecedingGT       bool
	IfPrecedingLTEQ     bool
	IfPrecedingGTEQ     bool
	IfPrecedingEQ       bool
	IfPrecedingNOTEQ    bool
	FirstFactor         *Factor
	FollowingFactorList []*Factor
}

// factor node
type Factor struct {
	IfPrecedingMULTIPLY bool
	IfPrecedingDIVIDE   bool
	Expression          *Expression
	ProcedureCall       *ProcedureCallStatement
	IfUnaryMinus        bool
	Name                *Name
	Number              *Number
	String              *String
	IfTrue              bool
	IfFalse             bool
}

// name node
type Name struct {
	IfPrecedingUnaryMINUS bool
	Identifier            *Identifier
	IfArray               bool
	Expression            *Expression
}

// identifier node
type Identifier struct {
	Content              string
	StartsWithUNDERSCORE bool
}

// number node
type Number struct {
	Content              string
	StartsWithUNDERSCORE bool
	ContainsPERIOD       bool
}

func (n *Number) Type() string {
	if n.ContainsPERIOD {
		return "float"
	} else {
		return "integer"
	}
}

// string node
type String struct {
	Content               string
	IfContainsDoubleQuote bool
}
