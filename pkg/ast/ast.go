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

func (p *Program) Type() string              {}
func (p *Program) Value() string             {}
func (p *Program) Children() []Node          {}
func (p *Program) Position() (line, col int) {}

// program header node
type ProgramHeader struct {
	Name *Identifier
}

func (ph *ProgramHeader) Type() string              {}
func (ph *ProgramHeader) Value() string             {}
func (ph *ProgramHeader) Children() []Node          {}
func (ph *ProgramHeader) Position() (line, col int) {}

// program body node
type ProgramBody struct {
	Declaration []Declaration // TODO - a little iffy on this
	Statement   []Statement   // TODO - a little iffy on this as well
}

func (pb *ProgramBody) Type() string              {}
func (pb *ProgramBody) Value() string             {}
func (pb *ProgramBody) Children() []Node          {}
func (pb *ProgramBody) Position() (line, col int) {}

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

func (pd *ProcedureDeclaration) Type() string              {}
func (pd *ProcedureDeclaration) Value() string             {}
func (pd *ProcedureDeclaration) Children() []Node          {}
func (pd *ProcedureDeclaration) Position() (line, col int) {}

// procedure header node
type ProcedureHeader struct {
	Name          *Identifier
	ReturnType    *TypeMark
	ParameterList []*VariableDeclaration
}

func (ph *ProcedureHeader) Type() string              {}
func (ph *ProcedureHeader) Value() string             {}
func (ph *ProcedureHeader) Children() []Node          {}
func (ph *ProcedureHeader) Position() (line, col int) {}

// parameter list <=> parameter <=> variable declaration node
// TODO - should this be fine to implement a list as well
// TODO - might need to remove left recursion
type VariableDeclaration struct {
	Name      *Identifier
	DataType  *TypeMark // given grammar calls it Type_Mark
	Array     bool
	ArraySize *Bound // in the grammar it is bound
}

func (vd *VariableDeclaration) Type() string              {}
func (vd *VariableDeclaration) Value() string             {}
func (vd *VariableDeclaration) Children() []Node          {}
func (vd *VariableDeclaration) Position() (line, col int) {}

// type mark node
type TypeMark struct {
	Name      *Identifier
	IfInteger bool
	IfFloat   bool
	IfString  bool
	IfBoolean bool
}

func (tm *TypeMark) Type() string              {}
func (tm *TypeMark) Value() string             {}
func (tm *TypeMark) Children() []Node          {}
func (tm *TypeMark) Position() (line, col int) {}

// bound node
type Bound Number

// procedure body node
type ProcedureBody struct {
	Declaration []Declaration // TODO - a little iffy on this
	Statement   []Statement   // TODO - a little iffy on this as well
}

func (pb *ProcedureBody) Type() string              {}
func (pb *ProcedureBody) Value() string             {}
func (pb *ProcedureBody) Children() []Node          {}
func (pb *ProcedureBody) Position() (line, col int) {}

// procedure call statement node
type ProcedureCallStatement struct {
	Name         *Identifier
	ArgumentList []*ArgumentList
}

func (pcs *ProcedureCallStatement) Type() string              {}
func (pcs *ProcedureCallStatement) Value() string             {}
func (pcs *ProcedureCallStatement) Children() []Node          {}
func (pcs *ProcedureCallStatement) Position() (line, col int) {}

// argument list node
type ArgumentList struct {
	ExpressionList       []*Expression
	IfMultipleExpression bool
}

func (al *ArgumentList) Type() string              {}
func (al *ArgumentList) Value() string             {}
func (al *ArgumentList) Children() []Node          {}
func (al *ArgumentList) Position() (line, col int) {}

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

func (as *AssignmentStatement) Type() string              {}
func (as *AssignmentStatement) Value() string             {}
func (as *AssignmentStatement) Children() []Node          {}
func (as *AssignmentStatement) Position() (line, col int) {}

// destination node
type Destination struct {
	Name  *Identifier
	Array bool
	Index *Expression
}

func (d *Destination) Type() string              {}
func (d *Destination) Value() string             {}
func (d *Destination) Children() []Node          {}
func (d *Destination) Position() (line, col int) {}

// if statement node
type IfStatement struct {
	IfExpression *Expression
	ThenBody     []Statement // since there is a *
	ElseBody     []Statement // since there is a *
}

func (is *IfStatement) Type() string              {}
func (is *IfStatement) Value() string             {}
func (is *IfStatement) Children() []Node          {}
func (is *IfStatement) Position() (line, col int) {}

// loop statement node
type LoopStatement struct {
	Initialization *AssignmentStatement
	Condition      *Expression
	Body           []Statement
}

func (ls *LoopStatement) Type() string              {}
func (ls *LoopStatement) Value() string             {}
func (ls *LoopStatement) Children() []Node          {}
func (ls *LoopStatement) Position() (line, col int) {}

// return statement node
type ReturnStatement struct {
	ReturnValue *Expression
}

func (rs *ReturnStatement) Type() string              {}
func (rs *ReturnStatement) Value() string             {}
func (rs *ReturnStatement) Children() []Node          {}
func (rs *ReturnStatement) Position() (line, col int) {}

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

func (e *Expression) Type() string              {}
func (e *Expression) Value() string             {}
func (e *Expression) Children() []Node          {}
func (e *Expression) Position() (line, col int) {}

// arithmetic operand node
type ArithmeticOperand struct {
	IfPrecedingNOT bool
	Relation       *Relation
	OperandList    []*ArithmeticOperand
}

func (ao *ArithmeticOperand) Type() string              {}
func (ao *ArithmeticOperand) Value() string             {}
func (ao *ArithmeticOperand) Children() []Node          {}
func (ao *ArithmeticOperand) Position() (line, col int) {}

// relation node
type Relation struct {
	FirstTerm         *Term
	FollowingTermList []*Term
}

func (r *Relation) Type() string              {}
func (r *Relation) Value() string             {}
func (r *Relation) Children() []Node          {}
func (r *Relation) Position() (line, col int) {}

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

func (t *Term) Type() string              {}
func (t *Term) Value() string             {}
func (t *Term) Children() []Node          {}
func (t *Term) Position() (line, col int) {}

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

func (f *Factor) Type() string              {}
func (f *Factor) Value() string             {}
func (f *Factor) Children() []Node          {}
func (f *Factor) Position() (line, col int) {}

// name node
type Name struct {
	IfPrecedingUnaryMINUS bool
	Identifier            *Identifier
	IfArray               bool
	Expression            *Expression
}

func (n *Name) Type() string              {}
func (n *Name) Value() string             {}
func (n *Name) Children() []Node          {}
func (n *Name) Position() (line, col int) {}

// identifier node
type Identifier struct {
	Name                   string
	IfStartsWithUNDERSCORE bool
}

func (i *Identifier) Type() string              {}
func (i *Identifier) Value() string             {}
func (i *Identifier) Children() []Node          {}
func (i *Identifier) Position() (line, col int) {}

// number node
type Number struct {
	Content                string
	IfStartsWithUNDERSCORE bool
	IfContainsPERIOD       bool
}

func (n *Number) Type() string              {}
func (n *Number) Value() string             {}
func (n *Number) Children() []Node          {}
func (n *Number) Position() (line, col int) {}

// string node
type String struct {
	Content               string
	IfContainsDoubleQuote bool
}

func (s *String) Type() string              {}
func (s *String) Value() string             {}
func (s *String) Children() []Node          {}
func (s *String) Position() (line, col int) {}
