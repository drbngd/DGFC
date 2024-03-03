// Code generated from example.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type exampleLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ExampleLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func examplelexerLexerInit() {
	staticData := &ExampleLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"ADD_OP", "MULT_OP", "IF_KW", "THEN_KW", "ELSE_KW", "WHILE_KW", "DO_KW",
		"END_KW", "COLON_EQUAL", "L_BRAKET", "R_BRAKET", "L_PAREN", "R_PAREN",
		"SEMI_COLON", "ID", "NUM", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 95, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 14, 4, 14, 80, 8, 14, 11, 14, 12, 14, 81, 1, 15, 4, 15, 85, 8, 15, 11,
		15, 12, 15, 86, 1, 16, 4, 16, 90, 8, 16, 11, 16, 12, 16, 91, 1, 16, 1,
		16, 0, 0, 17, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 1, 0, 5,
		2, 0, 43, 43, 45, 45, 2, 0, 42, 42, 47, 47, 1, 0, 97, 122, 1, 0, 48, 57,
		3, 0, 9, 10, 13, 13, 32, 32, 97, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 1, 35, 1, 0,
		0, 0, 3, 37, 1, 0, 0, 0, 5, 39, 1, 0, 0, 0, 7, 42, 1, 0, 0, 0, 9, 47, 1,
		0, 0, 0, 11, 52, 1, 0, 0, 0, 13, 58, 1, 0, 0, 0, 15, 61, 1, 0, 0, 0, 17,
		65, 1, 0, 0, 0, 19, 68, 1, 0, 0, 0, 21, 70, 1, 0, 0, 0, 23, 72, 1, 0, 0,
		0, 25, 74, 1, 0, 0, 0, 27, 76, 1, 0, 0, 0, 29, 79, 1, 0, 0, 0, 31, 84,
		1, 0, 0, 0, 33, 89, 1, 0, 0, 0, 35, 36, 7, 0, 0, 0, 36, 2, 1, 0, 0, 0,
		37, 38, 7, 1, 0, 0, 38, 4, 1, 0, 0, 0, 39, 40, 5, 105, 0, 0, 40, 41, 5,
		102, 0, 0, 41, 6, 1, 0, 0, 0, 42, 43, 5, 116, 0, 0, 43, 44, 5, 104, 0,
		0, 44, 45, 5, 101, 0, 0, 45, 46, 5, 110, 0, 0, 46, 8, 1, 0, 0, 0, 47, 48,
		5, 101, 0, 0, 48, 49, 5, 108, 0, 0, 49, 50, 5, 115, 0, 0, 50, 51, 5, 101,
		0, 0, 51, 10, 1, 0, 0, 0, 52, 53, 5, 119, 0, 0, 53, 54, 5, 104, 0, 0, 54,
		55, 5, 105, 0, 0, 55, 56, 5, 108, 0, 0, 56, 57, 5, 101, 0, 0, 57, 12, 1,
		0, 0, 0, 58, 59, 5, 100, 0, 0, 59, 60, 5, 111, 0, 0, 60, 14, 1, 0, 0, 0,
		61, 62, 5, 101, 0, 0, 62, 63, 5, 110, 0, 0, 63, 64, 5, 100, 0, 0, 64, 16,
		1, 0, 0, 0, 65, 66, 5, 58, 0, 0, 66, 67, 5, 61, 0, 0, 67, 18, 1, 0, 0,
		0, 68, 69, 5, 91, 0, 0, 69, 20, 1, 0, 0, 0, 70, 71, 5, 93, 0, 0, 71, 22,
		1, 0, 0, 0, 72, 73, 5, 40, 0, 0, 73, 24, 1, 0, 0, 0, 74, 75, 5, 41, 0,
		0, 75, 26, 1, 0, 0, 0, 76, 77, 5, 59, 0, 0, 77, 28, 1, 0, 0, 0, 78, 80,
		7, 2, 0, 0, 79, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0,
		81, 82, 1, 0, 0, 0, 82, 30, 1, 0, 0, 0, 83, 85, 7, 3, 0, 0, 84, 83, 1,
		0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87,
		32, 1, 0, 0, 0, 88, 90, 7, 4, 0, 0, 89, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0,
		0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94,
		6, 16, 0, 0, 94, 34, 1, 0, 0, 0, 4, 0, 81, 86, 91, 1, 6, 0, 0,
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

// exampleLexerInit initializes any static state used to implement exampleLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewexampleLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExampleLexerInit() {
	staticData := &ExampleLexerLexerStaticData
	staticData.once.Do(examplelexerLexerInit)
}

// NewexampleLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewexampleLexer(input antlr.CharStream) *exampleLexer {
	ExampleLexerInit()
	l := new(exampleLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ExampleLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "example.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// exampleLexer tokens.
const (
	exampleLexerADD_OP      = 1
	exampleLexerMULT_OP     = 2
	exampleLexerIF_KW       = 3
	exampleLexerTHEN_KW     = 4
	exampleLexerELSE_KW     = 5
	exampleLexerWHILE_KW    = 6
	exampleLexerDO_KW       = 7
	exampleLexerEND_KW      = 8
	exampleLexerCOLON_EQUAL = 9
	exampleLexerL_BRAKET    = 10
	exampleLexerR_BRAKET    = 11
	exampleLexerL_PAREN     = 12
	exampleLexerR_PAREN     = 13
	exampleLexerSEMI_COLON  = 14
	exampleLexerID          = 15
	exampleLexerNUM         = 16
	exampleLexerWS          = 17
)
