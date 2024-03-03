// Code generated from example.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // example

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type exampleParser struct {
	*antlr.BaseParser
}

var ExampleParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func exampleParserInit() {
	staticData := &ExampleParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "'if'", "'then'", "'else'", "'while'", "'do'", "'end'",
		"':='", "'['", "']'", "'('", "')'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "ADD_OP", "MULT_OP", "IF_KW", "THEN_KW", "ELSE_KW", "WHILE_KW",
		"DO_KW", "END_KW", "COLON_EQUAL", "L_BRAKET", "R_BRAKET", "L_PAREN",
		"R_PAREN", "SEMI_COLON", "ID", "NUM", "WS",
	}
	staticData.RuleNames = []string{
		"set_of_stmts", "stmt", "if_stmt", "while_stmt", "assignment_stmt",
		"cond_expr", "expr", "term", "factor",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 17, 95, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 0, 5,
		0, 22, 8, 0, 10, 0, 12, 0, 25, 9, 0, 1, 1, 1, 1, 1, 1, 3, 1, 30, 8, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 38, 8, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 52, 8, 4, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 65, 8,
		6, 10, 6, 12, 6, 68, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 76,
		8, 7, 10, 7, 12, 7, 79, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 3, 8, 90, 8, 8, 1, 8, 3, 8, 93, 8, 8, 1, 8, 0, 2, 12, 14, 9,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 0, 95, 0, 23, 1, 0, 0, 0, 2, 29, 1, 0,
		0, 0, 4, 31, 1, 0, 0, 0, 6, 39, 1, 0, 0, 0, 8, 46, 1, 0, 0, 0, 10, 56,
		1, 0, 0, 0, 12, 58, 1, 0, 0, 0, 14, 69, 1, 0, 0, 0, 16, 92, 1, 0, 0, 0,
		18, 19, 3, 2, 1, 0, 19, 20, 5, 14, 0, 0, 20, 22, 1, 0, 0, 0, 21, 18, 1,
		0, 0, 0, 22, 25, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24,
		1, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 26, 30, 3, 4, 2, 0, 27, 30, 3, 6, 3,
		0, 28, 30, 3, 8, 4, 0, 29, 26, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 28,
		1, 0, 0, 0, 30, 3, 1, 0, 0, 0, 31, 32, 5, 3, 0, 0, 32, 33, 3, 10, 5, 0,
		33, 34, 5, 4, 0, 0, 34, 37, 3, 0, 0, 0, 35, 36, 5, 5, 0, 0, 36, 38, 3,
		0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 5, 1, 0, 0, 0, 39,
		40, 5, 6, 0, 0, 40, 41, 3, 10, 5, 0, 41, 42, 5, 7, 0, 0, 42, 43, 3, 0,
		0, 0, 43, 44, 5, 8, 0, 0, 44, 45, 5, 6, 0, 0, 45, 7, 1, 0, 0, 0, 46, 51,
		5, 15, 0, 0, 47, 48, 5, 10, 0, 0, 48, 49, 3, 12, 6, 0, 49, 50, 5, 11, 0,
		0, 50, 52, 1, 0, 0, 0, 51, 47, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53,
		1, 0, 0, 0, 53, 54, 5, 9, 0, 0, 54, 55, 3, 12, 6, 0, 55, 9, 1, 0, 0, 0,
		56, 57, 3, 12, 6, 0, 57, 11, 1, 0, 0, 0, 58, 59, 6, 6, -1, 0, 59, 60, 3,
		14, 7, 0, 60, 66, 1, 0, 0, 0, 61, 62, 10, 2, 0, 0, 62, 63, 5, 1, 0, 0,
		63, 65, 3, 14, 7, 0, 64, 61, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64, 1,
		0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 13, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69,
		70, 6, 7, -1, 0, 70, 71, 3, 16, 8, 0, 71, 77, 1, 0, 0, 0, 72, 73, 10, 2,
		0, 0, 73, 74, 5, 2, 0, 0, 74, 76, 3, 16, 8, 0, 75, 72, 1, 0, 0, 0, 76,
		79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 15, 1, 0, 0,
		0, 79, 77, 1, 0, 0, 0, 80, 81, 5, 12, 0, 0, 81, 82, 3, 12, 6, 0, 82, 83,
		5, 13, 0, 0, 83, 93, 1, 0, 0, 0, 84, 89, 5, 15, 0, 0, 85, 86, 5, 10, 0,
		0, 86, 87, 3, 12, 6, 0, 87, 88, 5, 11, 0, 0, 88, 90, 1, 0, 0, 0, 89, 85,
		1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 93, 5, 16, 0, 0,
		92, 80, 1, 0, 0, 0, 92, 84, 1, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 17, 1,
		0, 0, 0, 8, 23, 29, 37, 51, 66, 77, 89, 92,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// exampleParserInit initializes any static state used to implement exampleParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewexampleParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExampleParserInit() {
	staticData := &ExampleParserStaticData
	staticData.once.Do(exampleParserInit)
}

// NewexampleParser produces a new parser instance for the optional input antlr.TokenStream.
func NewexampleParser(input antlr.TokenStream) *exampleParser {
	ExampleParserInit()
	this := new(exampleParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ExampleParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "example.g4"

	return this
}

// exampleParser tokens.
const (
	exampleParserEOF         = antlr.TokenEOF
	exampleParserADD_OP      = 1
	exampleParserMULT_OP     = 2
	exampleParserIF_KW       = 3
	exampleParserTHEN_KW     = 4
	exampleParserELSE_KW     = 5
	exampleParserWHILE_KW    = 6
	exampleParserDO_KW       = 7
	exampleParserEND_KW      = 8
	exampleParserCOLON_EQUAL = 9
	exampleParserL_BRAKET    = 10
	exampleParserR_BRAKET    = 11
	exampleParserL_PAREN     = 12
	exampleParserR_PAREN     = 13
	exampleParserSEMI_COLON  = 14
	exampleParserID          = 15
	exampleParserNUM         = 16
	exampleParserWS          = 17
)

// exampleParser rules.
const (
	exampleParserRULE_set_of_stmts    = 0
	exampleParserRULE_stmt            = 1
	exampleParserRULE_if_stmt         = 2
	exampleParserRULE_while_stmt      = 3
	exampleParserRULE_assignment_stmt = 4
	exampleParserRULE_cond_expr       = 5
	exampleParserRULE_expr            = 6
	exampleParserRULE_term            = 7
	exampleParserRULE_factor          = 8
)

// ISet_of_stmtsContext is an interface to support dynamic dispatch.
type ISet_of_stmtsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	AllSEMI_COLON() []antlr.TerminalNode
	SEMI_COLON(i int) antlr.TerminalNode

	// IsSet_of_stmtsContext differentiates from other interfaces.
	IsSet_of_stmtsContext()
}

type Set_of_stmtsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_of_stmtsContext() *Set_of_stmtsContext {
	var p = new(Set_of_stmtsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_set_of_stmts
	return p
}

func InitEmptySet_of_stmtsContext(p *Set_of_stmtsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_set_of_stmts
}

func (*Set_of_stmtsContext) IsSet_of_stmtsContext() {}

func NewSet_of_stmtsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_of_stmtsContext {
	var p = new(Set_of_stmtsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = exampleParserRULE_set_of_stmts

	return p
}

func (s *Set_of_stmtsContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_of_stmtsContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *Set_of_stmtsContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *Set_of_stmtsContext) AllSEMI_COLON() []antlr.TerminalNode {
	return s.GetTokens(exampleParserSEMI_COLON)
}

func (s *Set_of_stmtsContext) SEMI_COLON(i int) antlr.TerminalNode {
	return s.GetToken(exampleParserSEMI_COLON, i)
}

func (s *Set_of_stmtsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_of_stmtsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_of_stmtsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.EnterSet_of_stmts(s)
	}
}

func (s *Set_of_stmtsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.ExitSet_of_stmts(s)
	}
}

func (p *exampleParser) Set_of_stmts() (localctx ISet_of_stmtsContext) {
	localctx = NewSet_of_stmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, exampleParserRULE_set_of_stmts)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32840) != 0 {
		{
			p.SetState(18)
			p.Stmt()
		}
		{
			p.SetState(19)
			p.Match(exampleParserSEMI_COLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If_stmt() IIf_stmtContext
	While_stmt() IWhile_stmtContext
	Assignment_stmt() IAssignment_stmtContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = exampleParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) If_stmt() IIf_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *StmtContext) While_stmt() IWhile_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_stmtContext)
}

func (s *StmtContext) Assignment_stmt() IAssignment_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignment_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignment_stmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *exampleParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, exampleParserRULE_stmt)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case exampleParserIF_KW:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(26)
			p.If_stmt()
		}

	case exampleParserWHILE_KW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(27)
			p.While_stmt()
		}

	case exampleParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(28)
			p.Assignment_stmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF_KW() antlr.TerminalNode
	Cond_expr() ICond_exprContext
	THEN_KW() antlr.TerminalNode
	AllSet_of_stmts() []ISet_of_stmtsContext
	Set_of_stmts(i int) ISet_of_stmtsContext
	ELSE_KW() antlr.TerminalNode

	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_if_stmt
	return p
}

func InitEmptyIf_stmtContext(p *If_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_if_stmt
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = exampleParserRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) IF_KW() antlr.TerminalNode {
	return s.GetToken(exampleParserIF_KW, 0)
}

func (s *If_stmtContext) Cond_expr() ICond_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICond_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICond_exprContext)
}

func (s *If_stmtContext) THEN_KW() antlr.TerminalNode {
	return s.GetToken(exampleParserTHEN_KW, 0)
}

func (s *If_stmtContext) AllSet_of_stmts() []ISet_of_stmtsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISet_of_stmtsContext); ok {
			len++
		}
	}

	tst := make([]ISet_of_stmtsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISet_of_stmtsContext); ok {
			tst[i] = t.(ISet_of_stmtsContext)
			i++
		}
	}

	return tst
}

func (s *If_stmtContext) Set_of_stmts(i int) ISet_of_stmtsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISet_of_stmtsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISet_of_stmtsContext)
}

func (s *If_stmtContext) ELSE_KW() antlr.TerminalNode {
	return s.GetToken(exampleParserELSE_KW, 0)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.EnterIf_stmt(s)
	}
}

func (s *If_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.ExitIf_stmt(s)
	}
}

func (p *exampleParser) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, exampleParserRULE_if_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(exampleParserIF_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.Cond_expr()
	}
	{
		p.SetState(33)
		p.Match(exampleParserTHEN_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(34)
		p.Set_of_stmts()
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == exampleParserELSE_KW {
		{
			p.SetState(35)
			p.Match(exampleParserELSE_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.Set_of_stmts()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhile_stmtContext is an interface to support dynamic dispatch.
type IWhile_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWHILE_KW() []antlr.TerminalNode
	WHILE_KW(i int) antlr.TerminalNode
	Cond_expr() ICond_exprContext
	DO_KW() antlr.TerminalNode
	Set_of_stmts() ISet_of_stmtsContext
	END_KW() antlr.TerminalNode

	// IsWhile_stmtContext differentiates from other interfaces.
	IsWhile_stmtContext()
}

type While_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_stmtContext() *While_stmtContext {
	var p = new(While_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_while_stmt
	return p
}

func InitEmptyWhile_stmtContext(p *While_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_while_stmt
}

func (*While_stmtContext) IsWhile_stmtContext() {}

func NewWhile_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_stmtContext {
	var p = new(While_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = exampleParserRULE_while_stmt

	return p
}

func (s *While_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *While_stmtContext) AllWHILE_KW() []antlr.TerminalNode {
	return s.GetTokens(exampleParserWHILE_KW)
}

func (s *While_stmtContext) WHILE_KW(i int) antlr.TerminalNode {
	return s.GetToken(exampleParserWHILE_KW, i)
}

func (s *While_stmtContext) Cond_expr() ICond_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICond_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICond_exprContext)
}

func (s *While_stmtContext) DO_KW() antlr.TerminalNode {
	return s.GetToken(exampleParserDO_KW, 0)
}

func (s *While_stmtContext) Set_of_stmts() ISet_of_stmtsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISet_of_stmtsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISet_of_stmtsContext)
}

func (s *While_stmtContext) END_KW() antlr.TerminalNode {
	return s.GetToken(exampleParserEND_KW, 0)
}

func (s *While_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.EnterWhile_stmt(s)
	}
}

func (s *While_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.ExitWhile_stmt(s)
	}
}

func (p *exampleParser) While_stmt() (localctx IWhile_stmtContext) {
	localctx = NewWhile_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, exampleParserRULE_while_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(exampleParserWHILE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Cond_expr()
	}
	{
		p.SetState(41)
		p.Match(exampleParserDO_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Set_of_stmts()
	}
	{
		p.SetState(43)
		p.Match(exampleParserEND_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(exampleParserWHILE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignment_stmtContext is an interface to support dynamic dispatch.
type IAssignment_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON_EQUAL() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	L_BRAKET() antlr.TerminalNode
	R_BRAKET() antlr.TerminalNode

	// IsAssignment_stmtContext differentiates from other interfaces.
	IsAssignment_stmtContext()
}

type Assignment_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignment_stmtContext() *Assignment_stmtContext {
	var p = new(Assignment_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_assignment_stmt
	return p
}

func InitEmptyAssignment_stmtContext(p *Assignment_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_assignment_stmt
}

func (*Assignment_stmtContext) IsAssignment_stmtContext() {}

func NewAssignment_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_stmtContext {
	var p = new(Assignment_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = exampleParserRULE_assignment_stmt

	return p
}

func (s *Assignment_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Assignment_stmtContext) ID() antlr.TerminalNode {
	return s.GetToken(exampleParserID, 0)
}

func (s *Assignment_stmtContext) COLON_EQUAL() antlr.TerminalNode {
	return s.GetToken(exampleParserCOLON_EQUAL, 0)
}

func (s *Assignment_stmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Assignment_stmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Assignment_stmtContext) L_BRAKET() antlr.TerminalNode {
	return s.GetToken(exampleParserL_BRAKET, 0)
}

func (s *Assignment_stmtContext) R_BRAKET() antlr.TerminalNode {
	return s.GetToken(exampleParserR_BRAKET, 0)
}

func (s *Assignment_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assignment_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assignment_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.EnterAssignment_stmt(s)
	}
}

func (s *Assignment_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.ExitAssignment_stmt(s)
	}
}

func (p *exampleParser) Assignment_stmt() (localctx IAssignment_stmtContext) {
	localctx = NewAssignment_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, exampleParserRULE_assignment_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(exampleParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == exampleParserL_BRAKET {
		{
			p.SetState(47)
			p.Match(exampleParserL_BRAKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.expr(0)
		}
		{
			p.SetState(49)
			p.Match(exampleParserR_BRAKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(53)
		p.Match(exampleParserCOLON_EQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICond_exprContext is an interface to support dynamic dispatch.
type ICond_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsCond_exprContext differentiates from other interfaces.
	IsCond_exprContext()
}

type Cond_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCond_exprContext() *Cond_exprContext {
	var p = new(Cond_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_cond_expr
	return p
}

func InitEmptyCond_exprContext(p *Cond_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_cond_expr
}

func (*Cond_exprContext) IsCond_exprContext() {}

func NewCond_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cond_exprContext {
	var p = new(Cond_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = exampleParserRULE_cond_expr

	return p
}

func (s *Cond_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Cond_exprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Cond_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cond_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cond_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.EnterCond_expr(s)
	}
}

func (s *Cond_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.ExitCond_expr(s)
	}
}

func (p *exampleParser) Cond_expr() (localctx ICond_exprContext) {
	localctx = NewCond_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, exampleParserRULE_cond_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Term() ITermContext
	Expr() IExprContext
	ADD_OP() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = exampleParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) ADD_OP() antlr.TerminalNode {
	return s.GetToken(exampleParserADD_OP, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *exampleParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *exampleParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, exampleParserRULE_expr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.term(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, exampleParserRULE_expr)
			p.SetState(61)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(62)
				p.Match(exampleParserADD_OP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(63)
				p.term(0)
			}

		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Factor() IFactorContext
	Term() ITermContext
	MULT_OP() antlr.TerminalNode

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = exampleParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermContext) MULT_OP() antlr.TerminalNode {
	return s.GetToken(exampleParserMULT_OP, 0)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *exampleParser) Term() (localctx ITermContext) {
	return p.term(0)
}

func (p *exampleParser) term(_p int) (localctx ITermContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewTermContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITermContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, exampleParserRULE_term, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Factor()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTermContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, exampleParserRULE_term)
			p.SetState(72)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(73)
				p.Match(exampleParserMULT_OP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(74)
				p.Factor()
			}

		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	Expr() IExprContext
	R_PAREN() antlr.TerminalNode
	ID() antlr.TerminalNode
	L_BRAKET() antlr.TerminalNode
	R_BRAKET() antlr.TerminalNode
	NUM() antlr.TerminalNode

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = exampleParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = exampleParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(exampleParserL_PAREN, 0)
}

func (s *FactorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FactorContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(exampleParserR_PAREN, 0)
}

func (s *FactorContext) ID() antlr.TerminalNode {
	return s.GetToken(exampleParserID, 0)
}

func (s *FactorContext) L_BRAKET() antlr.TerminalNode {
	return s.GetToken(exampleParserL_BRAKET, 0)
}

func (s *FactorContext) R_BRAKET() antlr.TerminalNode {
	return s.GetToken(exampleParserR_BRAKET, 0)
}

func (s *FactorContext) NUM() antlr.TerminalNode {
	return s.GetToken(exampleParserNUM, 0)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(exampleListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *exampleParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, exampleParserRULE_factor)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case exampleParserL_PAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Match(exampleParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.expr(0)
		}
		{
			p.SetState(82)
			p.Match(exampleParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case exampleParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.Match(exampleParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(85)
				p.Match(exampleParserL_BRAKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(86)
				p.expr(0)
			}
			{
				p.SetState(87)
				p.Match(exampleParserR_BRAKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case exampleParserNUM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(91)
			p.Match(exampleParserNUM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *exampleParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 7:
		var t *TermContext = nil
		if localctx != nil {
			t = localctx.(*TermContext)
		}
		return p.Term_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *exampleParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *exampleParser) Term_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
