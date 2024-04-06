package lexer

import (
	"DGFC/pkg/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
    /* ok so it's also time to test
    /* nested comments; /*can you do it??*/
    */
    */
    PROGRAM grant IS
    global variable JAKE : integer;
    global variable ryan : integer[3];
    global variable zach : integer;
    variable tmp : integer;

    procedure if_proc : integer()
        variable i : integer;
        procedure dummy: float()
        // this should just hide the i of the outter environment
        variable i: float;
        variable tst : bool;
        begin
            i := 4.5;
            tst := putString("passed");
            return (0);
        end procedure;
        begin
            if(true) then jake := jake + 1;
            end if;
            return (0);
    end procedure;

    procedure for_proc : integer()
                variable i : integer;
        begin
            for(i := 0; i < zach)
            end for;
            return 1;
        end procedure;

    begin
        tmp := if_proc();
        tmp := for_proc();
    end program
    `

	tests := []struct {
		expectedType token.TokenType
		expectedVal  string
	}{
		{token.PROGRAM, "program"},
		{token.IDENTIFIER, "grant"},
		{token.IS, "is"},
		{token.GLOBAL, "global"},
		{token.VARIABLE, "variable"},
		{token.IDENTIFIER, "jake"},
		{token.COLON, ":"},
		{token.INTEGER, "integer"},
		{token.SEMICOLON, ";"},
		{token.GLOBAL, "global"},
		{token.VARIABLE, "variable"},
		{token.IDENTIFIER, "ryan"},
		{token.COLON, ":"},
		{token.INTEGER, "integer"},
		{token.LSQUARE, "["},
		{token.NUMBER, "3"},
		{token.RSQUARE, "]"},
		{token.SEMICOLON, ";"},
		{token.GLOBAL, "global"},
		{token.VARIABLE, "variable"},
		{token.IDENTIFIER, "zach"},
		{token.COLON, ":"},
		{token.INTEGER, "integer"},
		{token.SEMICOLON, ";"},
		{token.VARIABLE, "variable"},
		{token.IDENTIFIER, "tmp"},
		{token.COLON, ":"},
		{token.INTEGER, "integer"},
		{token.SEMICOLON, ";"},
		{token.PROCEDURE, "procedure"},
		{token.IDENTIFIER, "if_proc"},
		{token.COLON, ":"},
		{token.INTEGER, "integer"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.VARIABLE, "variable"},
		{token.IDENTIFIER, "i"},
		{token.COLON, ":"},
		{token.INTEGER, "integer"},
		{token.SEMICOLON, ";"},
		{token.PROCEDURE, "procedure"},
		{token.IDENTIFIER, "dummy"},
		{token.COLON, ":"},
		{token.FLOAT, "float"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.VARIABLE, "variable"},
		{token.IDENTIFIER, "i"},
		{token.COLON, ":"},
		{token.FLOAT, "float"},
		{token.SEMICOLON, ";"},
		{token.VARIABLE, "variable"},
		{token.IDENTIFIER, "tst"},
		{token.COLON, ":"},
		{token.BOOLEAN, "bool"},
		{token.SEMICOLON, ";"},
		{token.BEGIN, "begin"},
		{token.IDENTIFIER, "i"},
		{token.ASSIGN, ":="},
		{token.NUMBER, "4.5"},
		{token.SEMICOLON, ";"},
		{token.IDENTIFIER, "tst"},
		{token.ASSIGN, ":="},
		{token.IDENTIFIER, "putstring"},
		{token.LPAREN, "("},
		{token.STRING, "\"passed\""},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RETURN, "return"},
		{token.LPAREN, "("},
		{token.NUMBER, "0"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.END, "end"},
		{token.PROCEDURE, "procedure"},
		{token.SEMICOLON, ";"},
		{token.BEGIN, "begin"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.TRUE, "true"},
		{token.RPAREN, ")"},
		{token.THEN, "then"},
		{token.IDENTIFIER, "jake"},
		{token.ASSIGN, ":="},
		{token.IDENTIFIER, "jake"},
		{token.ADD, "+"},
		{token.NUMBER, "1"},
		{token.SEMICOLON, ";"},
		{token.END, "end"},
		{token.IF, "if"},
		{token.SEMICOLON, ";"},
		{token.RETURN, "return"},
		{token.LPAREN, "("},
		{token.NUMBER, "0"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.END, "end"},
		{token.PROCEDURE, "procedure"},
		{token.SEMICOLON, ";"},
		{token.PROCEDURE, "procedure"},
		{token.IDENTIFIER, "for_proc"},
		{token.COLON, ":"},
		{token.INTEGER, "integer"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.VARIABLE, "variable"},
		{token.IDENTIFIER, "i"},
		{token.COLON, ":"},
		{token.INTEGER, "integer"},
		{token.SEMICOLON, ";"},
		{token.BEGIN, "begin"},
		{token.FOR, "for"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "i"},
		{token.ASSIGN, ":="},
		{token.NUMBER, "0"},
		{token.SEMICOLON, ";"},
		{token.IDENTIFIER, "i"},
		{token.LT, "<"},
		{token.IDENTIFIER, "zach"},
		{token.RPAREN, ")"},
		{token.END, "end"},
		{token.FOR, "for"},
		{token.SEMICOLON, ";"},
		{token.RETURN, "return"},
		{token.NUMBER, "1"},
		{token.SEMICOLON, ";"},
		{token.END, "end"},
		{token.PROCEDURE, "procedure"},
		{token.SEMICOLON, ";"},
		{token.BEGIN, "begin"},
		{token.IDENTIFIER, "tmp"},
		{token.ASSIGN, ":="},
		{token.IDENTIFIER, "if_proc"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.IDENTIFIER, "tmp"},
		{token.ASSIGN, ":="},
		{token.IDENTIFIER, "for_proc"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.END, "end"},
		{token.PROGRAM, "program"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tk := l.NextToken()

		if tk.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q",
				i, tt.expectedType, tk.Type)
		}

		if tk.Value != tt.expectedVal {
			t.Fatalf("tests[%d] - value wrong, expected=%q, got=%q",
				i, tt.expectedVal, tk.Value)
		}
	}
}
