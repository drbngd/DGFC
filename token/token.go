package token

type TokenType string

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
	END_PROG  = "end program"
	GLOBAL    = "global"
	PROCEDURE = "procedure"
	END_PROC  = "end procedure"
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
	END_IF    = "end if"
	FOR       = "for"
	END_FOR   = "end for"
	END       = "end"
	RETURN    = "return"
)

// a hash map of keywords (symbol table)
var symbolTb = map[string]TokenType{
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
	tk, ok := symbolTb[tokenVal]
	if ok {
		return tk
	}
	return IDENTIFIER // if a word isn't a keyword then it has to be an identifier
}
