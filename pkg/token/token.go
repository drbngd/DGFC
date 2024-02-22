package token

// convert to lower - DONE
// capture identifiers properly - DONE
// should the lexer decide the type or the parser -
// arrays capture - no
// string capture - yes
// number with decimal upto one point - yes
// comment verify - lexer

type TokenType string

// struct to represent a token
type Token struct {
	Type  TokenType
	Value string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "END OF FILE"

	// Identifiers + literals
	IDENTIFIER = "IDENTIFIER"
	NUMBER     = "NUMBER"
	STRING     = "STRING"

	// Operators
	ASSIGN = ":="

	// Arithmetic
	ADD   = "+"
	TIMES = "*"
	DIV   = "/"
	SUB   = "-"
	EXCL  = "!"

	// Relational
	EQ     = "=="
	NOT_EQ = "!="
	LT     = "<"
	LTE    = "<="
	GT     = ">"
	GTE    = ">="

	// Boolean
	AND = "&"
	OR  = "|"
	// 'not' has been moved to keywords and reserved words

	// Delimiters
	SEMICOLON = ";"
	COLON     = ":"
	COMMA     = ","
	PERIOD    = "."

	// Brackets
	LPAREN  = "("
	RPAREN  = ")"
	LCURLY  = "{"
	RCURLY  = "}"
	LSQUARE = "["
	RSQUARE = "]"

	// Reserved words & Keywords
	PROGRAM   = "program"
	IS        = "is"
	BEGIN     = "begin"
	GLOBAL    = "global"
	PROCEDURE = "procedure"
	VARIABLE  = "variable"
	INTEGER   = "integer"
	FLOAT     = "float"
	STR       = "string"
	BOOLEAN   = "bool"
	TRUE      = "true"
	FALSE     = "false"
	NOT       = "not"
	IF        = "if"
	THEN      = "then"
	ELSE      = "else"
	FOR       = "for"
	END       = "end"
	RETURN    = "return"
)

// a hash map of keywords (symbol table)
var keywordsTB = map[string]TokenType{
	"program":   PROGRAM,
	"is":        IS,
	"begin":     BEGIN,
	"global":    GLOBAL,
	"procedure": PROCEDURE,
	"variable":  VARIABLE,
	"integer":   INTEGER,
	"float":     FLOAT,
	"string":    STR,
	"bool":      BOOLEAN,
	"true":      TRUE,
	"false":     FALSE,
	"not":       NOT,
	"if":        IF,
	"then":      THEN,
	"else":      ELSE,
	"for":       FOR,
	"end":       END,
	"return":    RETURN,
}

// function to check if a string belongs to the keyword hash map
func LookUp(tokenVal string) TokenType {
	tk, ok := keywordsTB[tokenVal]
	if ok {
		return tk
	}
	return IDENTIFIER // if a word isn't a keyword then it has to be an identifier
}
