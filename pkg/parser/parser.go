package parser

import (
	"DGFC/pkg/ast"
	"DGFC/pkg/lexer"
	"DGFC/pkg/token"
)

type ParserPointer struct {
	sp           *lexer.ScanPointer
	currentToken token.Token
	peekToken    token.Token
}

func New(sp *lexer.ScanPointer) *ParserPointer {
	pp := &ParserPointer{sp: sp}

	// Read two tokens, so currentToken and peekToken are both set
	//pp.NextToken()
	pp.NextToken() // only once to make sure PROGRAM is the first token

	return pp
}

func (pp *ParserPointer) NextToken() {
	pp.currentToken = pp.peekToken
	pp.peekToken = pp.sp.NextToken()
}

func (pp *ParserPointer) CurrentTokenIs(t token.TokenType) bool {
	return pp.currentToken.Type == t
}

func (pp *ParserPointer) NextTokenIs(t token.TokenType) bool {
	return pp.peekToken.Type == t
}

func (pp *ParserPointer) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Header = pp.ParseProgramHeader()
	program.Body = pp.ParseProgramBody()

	return program
}

func (pp *ParserPointer) ParseProgramHeader() *ast.ProgramHeader {
	programHeader := &ast.ProgramHeader{}

	// ensure first token is PROGRAM, 2nd is IDENTIFIER, and 3rd is IS
	if pp.CurrentTokenIs(token.PROGRAM) {
		pp.NextToken()
		if pp.CurrentTokenIs(token.IDENTIFIER) && pp.NextTokenIs(token.IS) {
			programHeader.Identifier.Name = pp.currentToken.Value
			pp.NextToken()
			pp.NextToken()
			return programHeader
		}
	}
	// throw some error
	return nil
}

func (pp *ParserPointer) ParseProgramBody() *ast.ProgramBody {
	// TODO - check for semicolon at the end of each statement & declaration
	// TODO - check for being and end program

	programBody := &ast.ProgramBody{}
	programBody.Declarations = pp.ParseDeclarations()
	programBody.Statements = pp.ParseStatements()

	return programBody
}

func (pp *ParserPointer) ParseDeclarations() *[]ast.Declaration {
	declarations := &[]ast.Declaration{}

	// keep appending to declarations until we reach the BEGIN token
	for !pp.CurrentTokenIs(token.BEGIN) {
		if pp.CurrentTokenIs(token.PROCEDURE) || (pp.CurrentTokenIs(token.GLOBAL)) && pp.NextTokenIs(token.PROCEDURE) {
			procedure := pp.ParseProcedure()
			*declarations = append(*declarations, &*procedure)
		} else if pp.CurrentTokenIs(token.VARIABLE) || (pp.CurrentTokenIs(token.GLOBAL) && pp.NextTokenIs(token.VARIABLE)) {
			variable := pp.ParseVariable()
			*declarations = append(*declarations, &*variable)
		} else {
			// throw some error
			return nil
		}
	}

	return declarations

}

func (pp *ParserPointer) ParseProcedure() *ast.ProcedureDeclaration {
	procedure := &ast.ProcedureDeclaration{}
	if pp.CurrentTokenIs(token.GLOBAL) {
		procedure.IsGlobal = true
		pp.NextToken()
	} else {
		procedure.IsGlobal = false
	}

	procedure.Header = pp.ParseProcedureHeader()
	procedure.Body = pp.ParseProcedureBody()

	return procedure

	// throw some error
	return nil
}

func (pp *ParserPointer) ParseProcedureHeader() *ast.ProcedureHeader {
	procedureHeader := &ast.ProcedureHeader{}

	// ensure first token is PROCEDURE, 2nd is IDENTIFIER, and 3rd is LPAREN
	if pp.CurrentTokenIs(token.PROCEDURE) {
		pp.NextToken()
		if pp.CurrentTokenIs(token.IDENTIFIER) {
			procedureHeader.Identifier.Name = pp.currentToken.Value
			pp.NextToken()
			if pp.CurrentTokenIs(token.COLON) {
				pp.NextToken()
				if pp.CurrentTokenIs(token.INTEGER) || pp.CurrentTokenIs(token.FLOAT) ||
					pp.CurrentTokenIs(token.STR) || pp.CurrentTokenIs(token.BOOLEAN) {
					pp.NextToken()
					procedureHeader.Type.Name = pp.currentToken.Value
					if pp.CurrentTokenIs(token.LPAREN) {
						// TODO - add something which would be fine with no parameters
						parameterList := ast.ParameterList{}
						for !pp.CurrentTokenIs(token.RPAREN) {
							parameter := pp.ParseParameter()
							*parameterList.Parameters = append(*parameterList.Parameters, parameter)
						}
						procedureHeader.Parameters = &parameterList
						pp.NextToken()
						return procedureHeader
					}
				}
			}
		}
	}
	// throw some error
	return nil
}

func (pp *ParserPointer) ParseParameter() *ast.Parameter {
	parameter := &ast.Parameter{}
	parameter.VariableDeclaration = pp.ParseVariableDeclaration()
	// throw some error
	return parameter
}

func (pp *ParserPointer) ParseVariableDeclaration() *ast.VariableDeclaration {
	variableDeclaration := &ast.VariableDeclaration{}
	if pp.CurrentTokenIs(token.GLOBAL) {
		variableDeclaration.IsGlobal = true
		pp.NextToken()
	} else {
		variableDeclaration.IsGlobal = false
	}

	if pp.CurrentTokenIs(token.VARIABLE) {
		pp.NextToken()
		if pp.CurrentTokenIs(token.IDENTIFIER) {
			variableDeclaration.Identifier.Name = pp.currentToken.Value
			pp.NextToken()
			if pp.CurrentTokenIs(token.COLON) {
				pp.NextToken()
				if pp.CurrentTokenIs(token.INTEGER) || pp.CurrentTokenIs(token.FLOAT) ||
					pp.CurrentTokenIs(token.STR) || pp.CurrentTokenIs(token.BOOLEAN) {
					variableDeclaration.Type.Name = pp.currentToken.Value
					pp.NextToken()
					if pp.CurrentTokenIs(token.LSQUARE) {
						pp.NextToken()
						if pp.CurrentTokenIs(token.NUMBER) {
							variableDeclaration.Bound.Value.Value = pp.currentToken.Value
							pp.NextToken()
							if pp.CurrentTokenIs(token.RSQUARE) {
								variableDeclaration.IsArray = true
							} else {
								variableDeclaration.IsArray = false
							}
							pp.NextToken()
							return variableDeclaration
						}
					}
				}
			}
		}
	}
	// TODO - throw some error
	return nil
}

func (pp *ParserPointer) ParseProcedureBody() *ast.ProcedureBody {
	// TODO - check for semicolon at the end of each statement & declaration
	// TODO - check for being and end procedure

	procedureBody := &ast.ProcedureBody{}
	procedureBody.Declarations = pp.ParseDeclarations()
	procedureBody.Statements = pp.ParseStatements()

	return procedureBody
}
