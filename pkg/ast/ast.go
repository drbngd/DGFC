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
	Name string
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
	Name          string
	ReturnType    string
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
	Name      string
	DataType  string // given grammar calls it Type_Mark
	Array     bool
	ArraySize int
}

func (vd *VariableDeclaration) Type() string              {}
func (vd *VariableDeclaration) Value() string             {}
func (vd *VariableDeclaration) Children() []Node          {}
func (vd *VariableDeclaration) Position() (line, col int) {}

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
// TODO - will probably require elimination of left recursion
type ProcedureCallStatement struct{}

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
	Name  string
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
type Expression interface{}
