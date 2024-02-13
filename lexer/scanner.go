package lexer

import (
	"DGFC/token"
	"strings"
)

// struct to represent a pointer to scan the input string
type ScanPointer struct {
	input    string
	currPos  int  // current pointer position
	nextPos  int  // next pointer position
	currChar byte // currChar at current position
}

// function to create a new ScanPointer
func New(input string) *ScanPointer {
	// first convert the input string to lower case
	// TODO - add row and col for the current position
	input = strings.ToLower(input)
	sp := &ScanPointer{input: input}
	sp.ReadNextChar()
	return sp
}

// method to progress through input string char-by-char
func (sp *ScanPointer) ReadNextChar() {
	if sp.nextPos >= len(sp.input) {
		sp.currChar = 0 // ASCII(0) is NUL
	} else {
		sp.currChar = sp.input[sp.nextPos]
	}
	sp.currPos = sp.nextPos
	sp.nextPos += 1
}

// method to scan the next token
func (sp *ScanPointer) NextToken() token.Token {
	var tk token.Token

	sp.EatWhitepace()
	sp.SkipComments()

	switch sp.currChar {
	case ';':
		tk = NewToken(token.SEMICOLON, ";")
	case ':':
		if sp.PeekNext() == '=' {
			sp.ReadNextChar()
			tk = NewToken(token.ASSIGN, ":=")
		} else {
			tk = NewToken(token.COLON, ":")
		}
	case ',':
		tk = NewToken(token.COMMA, ",")
	case '.':
		if sp.PeekNext() == 0 {
			tk.Value = ""
			tk.Type = token.EOF
		} else {
			tk = NewToken(token.PERIOD, ".")
		}
	case '(':
		tk = NewToken(token.LPAREN, "(")
	case ')':
		tk = NewToken(token.RPAREN, ")")
	case '{':
		tk = NewToken(token.LCURLY, "{")
	case '}':
		tk = NewToken(token.RCURLY, "}")
	case '[':
		tk = NewToken(token.LSQUARE, "[")
	case ']':
		tk = NewToken(token.RSQUARE, "]")
	case '+':
		tk = NewToken(token.ADD, "+")
	case '*':
		tk = NewToken(token.TIMES, "*")
	case '/':
		tk = NewToken(token.DIV, "/")
	case '-':
		tk = NewToken(token.SUB, "-")
	case '!':
		if sp.PeekNext() == '=' {
			sp.ReadNextChar()
			tk = NewToken(token.NOT_EQ, "!=")
		} else {
			tk = NewToken(token.EXCL, "!")
		}
	case '&':
		tk = NewToken(token.AND, "&")
	case '|':
		tk = NewToken(token.OR, "|")
	case '<':
		if sp.PeekNext() == '=' {
			sp.ReadNextChar()
			tk = NewToken(token.LTE, "<=")
		} else {
			tk = NewToken(token.LT, "<")
		}
	case '>':
		if sp.PeekNext() == '=' {
			sp.ReadNextChar()
			tk = NewToken(token.GTE, ">=")
		} else {
			tk = NewToken(token.GT, ">")
		}
	case '=':
		if sp.PeekNext() == '=' {
			sp.ReadNextChar()
			tk = NewToken(token.EQ, "==")
		}
	case '"':
		if sp.PeekNext() == '"' {
			sp.ReadNextChar()
			tk = NewToken(token.STRING, "\"\"")
		} else {
			sp.ReadNextChar()
			temp_val := sp.ReadString()
			//sp.ReadNextChar()
			tk = NewToken(token.STRING, "\""+temp_val+"\"")
		}
	case 0:
		tk.Value = ""
		tk.Type = token.EOF
	default:
		//check for identifiers, numbers, keywords
		if IsLetter(sp.currChar) {
			tk.Value = sp.ReadIdentifier()
			tk.Type = token.LookUp(tk.Value)
			return tk
		} else if IsDigit(sp.currChar) {
			tk.Value = sp.ReadNumber()
			tk.Type = token.NUMBER
			return tk
		} else {
			tk = NewToken(token.ILLEGAL, string(sp.currChar))
		}
	}

	sp.ReadNextChar()
	return tk
}

// function to generate a new Token variable
func NewToken(tokenType token.TokenType, tokenValue string) token.Token {
	return token.Token{Type: tokenType, Value: tokenValue}
}

// method to peek forward
func (sp *ScanPointer) PeekNext() byte {
	if sp.nextPos >= len(sp.input) {
		return 0
	} else {
		return sp.input[sp.nextPos]
	}
}

// function to check if character is a letter or underscore
func IsLetter(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

// function to check if character is a digit
func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// method to detect & read strings
func (sp *ScanPointer) ReadString() string {
	startPos := sp.currPos
	for sp.PeekNext() != '"' {
		sp.ReadNextChar()
	}
	sp.ReadNextChar()
	return sp.input[startPos:sp.currPos]
}

// method to detect & read identifiers
func (sp *ScanPointer) ReadIdentifier() string {
	startPos := sp.currPos
	if IsLetter(sp.currChar) {
		sp.ReadNextChar()
		for IsLetter(sp.currChar) || IsDigit(sp.currChar) || sp.currChar == '_' {
			sp.ReadNextChar()
		}
	}
	return sp.input[startPos:sp.currPos]
}

// method to detect & read a number -  integer and floating point
func (sp *ScanPointer) ReadNumber() string {
	startPos := sp.currPos
	decimal_point_count := 0
	for IsDigit(sp.currChar) {
		sp.ReadNextChar()
		if sp.currChar == '.' && IsDigit(sp.PeekNext()) && decimal_point_count <= 1 {
			decimal_point_count++
			sp.ReadNextChar()
		}
	}
	return sp.input[startPos:sp.currPos]
}

// method to detect & skip whitespaces - new line, tab, space, carriage return
func (sp *ScanPointer) EatWhitepace() {
	for sp.currChar == ' ' || sp.currChar == '\n' || sp.currChar == '\t' || sp.currChar == '\r' {
		sp.ReadNextChar()
	}
}

// method to detect & skip comments - single and multi-line
func (sp *ScanPointer) SkipComments() {
	for {
		sp.EatWhitepace()
		if sp.currChar == '/' {
			if sp.PeekNext() == '/' { // Single-line comment
				sp.ReadNextChar() // Consume '/'
				for sp.PeekNext() != '\n' && sp.PeekNext() != 0 {
					sp.ReadNextChar()
				}
				sp.ReadNextChar() // Move pointer to next line
			} else if sp.PeekNext() == '*' { // Multi-line comment
				sp.ReadNextChar() // Consume '*'
				multiCommentDepth := 1
				for multiCommentDepth > 0 {
					sp.ReadNextChar()
					if sp.currChar == '*' && sp.PeekNext() == '/' {
						sp.ReadNextChar() // Consume '/'
						multiCommentDepth--
					} else if sp.currChar == '/' && sp.PeekNext() == '*' {
						sp.ReadNextChar() // Consume '*'
						multiCommentDepth++
					}
				}
				sp.ReadNextChar() // Move pointer to next line
			} else {
				sp.ReadNextChar()
				return // Not a comment, return
			}
		} else {
			return // Not a comment, return
		}
	}
}
