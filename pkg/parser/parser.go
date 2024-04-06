package parser

import (
	"DGFC/pkg/ast"
	"DGFC/pkg/lexer"
	"DGFC/pkg/token"
	"fmt"
)

type ParserPointer struct {
	sp           *lexer.ScanPointer
	currentToken token.Token
	peekToken    token.Token
	errors       []string
}

func New(sp *lexer.ScanPointer) *ParserPointer {
	pp := &ParserPointer{sp: sp, errors: []string{}}

	// Read two tokens, so currentToken and peekToken are both set
	pp.NextToken()
	pp.NextToken() // only once to make sure PROGRAM is the first token
	//fmt.Printf("Current Token is: %+v\n", pp.currentToken.Value)
	return pp
}

// for testing
func (pp *ParserPointer) GetCurrentToken() token.Token {
	return pp.currentToken
}

func (pp *ParserPointer) NextToken() {
	pp.currentToken = pp.peekToken
	pp.peekToken = pp.sp.NextToken()
	fmt.Printf("Current Token is: %+v\n", pp.currentToken.Value)
}

func (pp *ParserPointer) CurrentTokenIs(t token.TokenType) bool {
	return pp.currentToken.Type == t
}

func (pp *ParserPointer) NextTokenIs(t token.TokenType) bool {
	return pp.peekToken.Type == t
}

func (pp *ParserPointer) ReportError(msg string) {
	line, col := pp.sp.GetPosition()
	errorMessage := fmt.Sprintf("Error at line %d, column %d: %s. Received token: %s", line, col, msg, pp.currentToken.Value)

	pp.errors = append(pp.errors, errorMessage)
	//print(errorMessage)

}

func (pp *ParserPointer) GetErrors() []string {
	return pp.errors
}

/*---Parse Functions for Individual Grammar Rules---*/
// convention followed: each parse function ends with a NextToken() call or a return statement
// exception: ParseStatements() function, which is called by multiple grammar rules

func (pp *ParserPointer) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Header = pp.ParseProgramHeader()
	program.Body = pp.ParseProgramBody()

	// check for END PROGRAM
	if !pp.CurrentTokenIs(token.PERIOD) {
		pp.ReportError("[Program] expected '.' at the end of the program")
		return nil
	}

	return program
}

func (pp *ParserPointer) ParseProgramHeader() *ast.ProgramHeader {
	programHeader := &ast.ProgramHeader{Identifier: &ast.Identifier{}}

	if !pp.CurrentTokenIs(token.PROGRAM) {
		pp.ReportError("[Program Header] expected PROGRAM keyword")
		return nil
	}
	pp.NextToken()

	if !pp.CurrentTokenIs(token.IDENTIFIER) {
		pp.ReportError("[Program Header] expected IDENTIFIER after PROGRAM keyword")
		return nil
	}

	fmt.Printf("Current Token is: %+v\n", pp.currentToken.Value)

	programHeader.Identifier.Name = pp.currentToken.Value
	pp.NextToken()

	if !pp.CurrentTokenIs(token.IS) {
		pp.ReportError("[Program Header] expected IS keyword")
		return nil
	}
	pp.NextToken()

	return programHeader

}

func (pp *ParserPointer) ParseProgramBody() *ast.ProgramBody {
	// TODO - check for semicolon at the end of each statement & declaration
	// TODO - check for being and end program

	programBody := &ast.ProgramBody{}
	programBody.Declarations = pp.ParseDeclarations() // TODO: make all var/proc declarations global
	print("back to parse program body \n")

	programBody.Statements = pp.ParseStatements()
	print("back to parse program body \n")

	if !(pp.CurrentTokenIs(token.END) && pp.NextTokenIs(token.PROGRAM)) {
		pp.ReportError("[Program Body] expected END PROGRAM after declarations and statements")
		return nil
	}

	pp.NextToken() // consume the END token
	pp.NextToken() // consume the PROGRAM token

	// semicolon check in ParseDeclaration and ParseStatement

	return programBody
}

func (pp *ParserPointer) ParseDeclarations() []ast.Declaration {
	declarations := []ast.Declaration{}
	isGlobal := false

	// keep appending to declarations until we reach the BEGIN token
	// as acc to our grammar rules, all declarations come before the BEGIN token
	for !pp.CurrentTokenIs(token.BEGIN) {
		// check for global
		if pp.CurrentTokenIs(token.GLOBAL) {
			isGlobal = true
			pp.NextToken()
		}
		// parse either
		switch pp.currentToken.Type {
		case token.PROCEDURE:
			procedure := pp.ParseProcedureDeclaration(isGlobal)
			declarations = append(declarations, &*procedure)
		case token.VARIABLE:
			variable := pp.ParseVariableDeclaration(isGlobal)
			declarations = append(declarations, &*variable)
			fmt.Printf("variable declaration added to declarations\n")
		default:
			pp.ReportError("[Declaration] unexpected TOKEN encountered, expected either PROCEDURE or VARIABLE keyword")
			return nil
		}

		// check for semicolon
		if !pp.CurrentTokenIs(token.SEMICOLON) {
			pp.ReportError("[Declaration] expected SEMICOLON after declaration")
			return nil
		} else {
			pp.NextToken()
		}
	}

	pp.NextToken() // consume the BEGIN token

	return declarations
}

func (pp *ParserPointer) ParseProcedureDeclaration(isGlobal bool) *ast.ProcedureDeclaration {
	procedure := &ast.ProcedureDeclaration{}

	procedure.IsGlobal = isGlobal
	procedure.Header = pp.ParserProcedureHeader()
	procedure.Body = pp.ParseProcedureBody()

	return procedure

}

func (pp *ParserPointer) ParserProcedureHeader() *ast.ProcedureHeader {
	procedureHeader := &ast.ProcedureHeader{Identifier: &ast.Identifier{}}

	if !pp.CurrentTokenIs(token.PROCEDURE) {
		pp.ReportError("[Procedure Header] expected PROCEDURE keyword")
		return nil
	}
	pp.NextToken()
	fmt.Printf("onto identifer \n")

	if !pp.CurrentTokenIs(token.IDENTIFIER) {
		pp.ReportError("[Procedure Header] expected IDENTIFIER after PROCEDURE keyword")
		return nil
	}
	procedureHeader.Identifier.Name = pp.currentToken.Value
	pp.NextToken()

	fmt.Printf("onto colon \n")
	if !pp.CurrentTokenIs(token.COLON) {
		pp.ReportError("[Procedure Header] expected COLON after IDENTIFIER in PROCEDURE header")
		return nil
	}
	pp.NextToken()

	fmt.Printf("onto type \n")
	procedureHeader.Type = pp.ParseTypeMark()

	if !pp.CurrentTokenIs(token.LPAREN) {
		pp.ReportError("[Procedure Header] expected LPAREN after TYPE in PROCEDURE header")
		return nil
	}
	pp.NextToken()
	procedureHeader.ParameterList = pp.ParseParameterList()

	if !pp.CurrentTokenIs(token.RPAREN) {
		pp.ReportError("[Procedure Header] expected RPAREN after PARAMETER LIST in PROCEDURE header")
		return nil
	}
	pp.NextToken()

	return procedureHeader

}

func (pp *ParserPointer) ParseParameterList() *ast.ParameterList {
	parameterList := &ast.ParameterList{Parameters: []ast.Parameter{}}

	if pp.CurrentTokenIs(token.RPAREN) {
		return parameterList
	}

	// parameter list would never be global
	firstParameter := pp.ParseParameter(false)
	parameterList.Parameters = append(parameterList.Parameters, *firstParameter)

	for pp.CurrentTokenIs(token.COMMA) {
		pp.NextToken()
		// check for multiple commas
		if pp.CurrentTokenIs(token.COMMA) {
			pp.ReportError("[Parameter List] expected PARAMETER after COMMA, found another COMMA")
			return nil
		}

		parameter := pp.ParseParameter(false)
		parameterList.Parameters = append(parameterList.Parameters, *parameter)

	}

	// todo - add case where there are multiple commas but no parameters

	return parameterList
}

func (pp *ParserPointer) ParseParameter(isGlobal bool) *ast.Parameter {
	parameter := &ast.Parameter{}
	parameter.VariableDeclaration = pp.ParseVariableDeclaration(isGlobal)

	return parameter
}

func (pp *ParserPointer) ParseProcedureBody() *ast.ProcedureBody {
	procedureBody := &ast.ProcedureBody{}
	procedureBody.Declarations = pp.ParseDeclarations()
	print("back to parse procedure body \n")

	print("onto procedure body statements \n")

	procedureBody.Statements = pp.ParseStatements()
	print("back to parse procedure body \n")
	print("current token is: " + pp.currentToken.Value + "\n")

	print("onto end procedure \n")
	if !(pp.CurrentTokenIs(token.END) && pp.NextTokenIs(token.PROCEDURE)) {
		print("error in end procedure \n")
		pp.ReportError("[Procedure Body] expected END PROCEDURE after declarations and statements")
		return nil
	}

	print("end procedure detected \n")
	pp.NextToken() // consume the END token
	pp.NextToken() // consume the PROCEDURE token

	return procedureBody
}

func (pp *ParserPointer) ParseVariableDeclaration(isGlobal bool) *ast.VariableDeclaration {
	variableDeclaration := &ast.VariableDeclaration{IsGlobal: isGlobal, Identifier: &ast.Identifier{}}

	if !pp.CurrentTokenIs(token.VARIABLE) {
		pp.ReportError("[Variable Declaration] expected VARIABLE keyword")
		return nil
	}
	pp.NextToken()

	fmt.Printf("onto identifer \n")
	if !pp.CurrentTokenIs(token.IDENTIFIER) {
		pp.ReportError("[Variable Declaration] expected IDENTIFIER after VARIABLE keyword")
		return nil
	}
	variableDeclaration.Identifier.Name = pp.currentToken.Value
	pp.NextToken()

	fmt.Printf("onto colon \n")
	if !pp.CurrentTokenIs(token.COLON) {
		pp.ReportError("[Variable Declaration] expected COLON after IDENTIFIER")
		return nil
	}
	pp.NextToken()

	variableDeclaration.Type = pp.ParseTypeMark()

	fmt.Printf("onto bound \n")
	if pp.CurrentTokenIs(token.LSQUARE) {
		pp.NextToken()
		if !pp.CurrentTokenIs(token.NUMBER) {
			pp.ReportError("[Variable Declaration] expected NUMBER after LEFT SQUARE BRACE in array declaration")
			return nil
		}
		variableDeclaration.Bound = pp.ParseBound()

		if !pp.CurrentTokenIs(token.RSQUARE) {
			pp.ReportError("[Variable Declaration] expected RIGHT SQUARE BRACE in array declaration")
			return nil
		}
		variableDeclaration.IsArray = true
		pp.NextToken()
		fmt.Printf("bound present \n")
	} else {
		variableDeclaration.IsArray = false
	}

	fmt.Printf("back to declaration function \n")
	return variableDeclaration
}

func (pp *ParserPointer) ParseTypeMark() *ast.TypeMark {
	typeMark := &ast.TypeMark{}

	switch pp.currentToken.Type {
	case token.INTEGER:
		typeMark.Name = "integer"
		pp.NextToken()
	case token.FLOAT:
		typeMark.Name = "float"
		pp.NextToken()
	case token.STR:
		typeMark.Name = "string"
		pp.NextToken()
	case token.BOOLEAN:
		typeMark.Name = "bool"
		pp.NextToken()
	default:
		pp.ReportError("[Type Mark] expected INTEGER, FLOAT or BOOL")
		return nil
	}

	return typeMark
}

func (pp *ParserPointer) ParseBound() *ast.Bound {
	bound := &ast.Bound{}
	bound.Value = pp.ParseNumber()

	return bound
}

// (exception) ParseStatement doesn't consume the END or ELSE token
func (pp *ParserPointer) ParseStatements() []ast.Statement {
	statements := []ast.Statement{}

	// TODO - should the for loop look for a semicolon?
	for !(pp.CurrentTokenIs(token.END) || pp.CurrentTokenIs(token.ELSE)) {
		switch pp.currentToken.Type {
		case token.IF:
			ifStatement := pp.ParseIfStatement()
			statements = append(statements, ifStatement)
		case token.FOR:
			print("parsing for statement \n")
			forStatement := pp.ParseLoopStatement()
			print("for statement parsed \n")
			print(forStatement.ToString())
			statements = append(statements, forStatement)
		case token.RETURN:
			print("parsing return statement \n")
			returnStatement := pp.ParseReturnStatement()
			print("return statement parsed \n")
			statements = append(statements, returnStatement)
		case token.IDENTIFIER:
			print("parsing assignment statement \n")
			assignmentStatement := pp.ParseAssignmentStatement()
			print("assignment statement parsed \n")
			statements = append(statements, assignmentStatement)
		default:
			pp.ReportError("[Statement] expected IF, FOR, RETURN or ASSIGNMENT statement")
			return nil
		}

		// check for semicolon
		print("checking for semicolon \n")
		if !pp.CurrentTokenIs(token.SEMICOLON) {
			pp.ReportError("[Statement] expected SEMICOLON after statement")
			return nil
		} else {
			print("semicolon detected \n")
			pp.NextToken()
		}
	}

	// (Exception) we will not consume the END or ELSE token here
	// each grammar rule that calls ParseStatements will consume the END or ELSE token
	print("going back to where ParseStatements was called \n")

	return statements
}

func (pp *ParserPointer) ParseProcedureCall() *ast.ProcedureCall {
	procedureCall := &ast.ProcedureCall{Identifier: &ast.Identifier{}}

	if !pp.CurrentTokenIs(token.IDENTIFIER) {
		pp.ReportError("[Procedure Call] expected IDENTIFIER")
		return nil
	}
	print("parsing procedure call \n")

	procedureCall.Identifier.Name = pp.currentToken.Value
	pp.NextToken()

	if !pp.CurrentTokenIs(token.LPAREN) {
		pp.ReportError("[Procedure Call] expected LPAREN after IDENTIFIER")
		return nil
	}
	print("LPAREN detected \n")
	pp.NextToken()

	print("now parsing argument list \n")
	procedureCall.ArguementList = pp.ParseArgumentList()

	if !pp.CurrentTokenIs(token.RPAREN) {
		pp.ReportError("[Procedure Call] expected RPAREN after ARGUEMENT LIST")
		return nil
	}
	print("RPAREN detected \n")
	pp.NextToken()

	return procedureCall
}

func (pp *ParserPointer) ParseAssignmentStatement() *ast.AssignmentStatement {
	assignmentStatement := &ast.AssignmentStatement{Destination: &ast.Destination{Identifier: &ast.Identifier{}}, Expression: &ast.Expression{}}

	print("parsing assignment statement \n")
	assignmentStatement.Destination = pp.ParseDestination()

	print("destination parsed, onto assignment operator \n")

	if !pp.CurrentTokenIs(token.ASSIGN) {
		pp.ReportError("[Assignment Statement] expected ASSIGNMENT operator")
		return nil
	}
	pp.NextToken()

	print("assignment operator parsed, onto expression \n")

	assignmentStatement.Expression = pp.ParseExpression()

	return assignmentStatement
}

func (pp *ParserPointer) ParseDestination() *ast.Destination {
	destination := &ast.Destination{}

	destination.Identifier = pp.ParseIdentifier()

	if pp.CurrentTokenIs(token.LSQUARE) {
		pp.NextToken()

		destination.Expression = pp.ParseExpression()

		if !pp.CurrentTokenIs(token.RSQUARE) {
			pp.ReportError("[Destination] expected RSQUARE after EXPRESSION in array call")
			return nil
		}
		pp.NextToken()

	}
	return destination
}

func (pp *ParserPointer) ParseIfStatement() *ast.IfStatement {
	ifStatement := &ast.IfStatement{}

	if !pp.CurrentTokenIs(token.IF) {
		pp.ReportError("[If Statement] expected IF keyword")
		return nil
	}
	pp.NextToken()

	if !pp.CurrentTokenIs(token.LPAREN) {
		pp.ReportError("[If Statement] expected LEFT PARENTHESIS after IF keyword")
		return nil
	}
	pp.NextToken()

	ifStatement.Condition = pp.ParseExpression()
	// if expr is NIL, then throw error and return NIL
	if ifStatement.Condition == nil {
		pp.ReportError("[If Statement] expected EXPRESSION after LEFT PARENTHESIS")
		return nil
	}

	fmt.Printf("if condition parsed, onto RPAREN \n")

	if !pp.CurrentTokenIs(token.RPAREN) {
		pp.ReportError("[If Statement] expected RIGHT PARENTHESIS after IF condition")
		return nil
	}
	pp.NextToken()

	if !pp.CurrentTokenIs(token.THEN) {
		pp.ReportError("[If Statement] expected THEN keyword after IF condition")
		return nil
	}
	pp.NextToken()

	ifStatement.ThenBlock = pp.ParseStatements()
	print("back to parse if statement \n")

	if pp.CurrentTokenIs(token.ELSE) {
		pp.NextToken()
		ifStatement.ElseBlock = pp.ParseStatements()
		print("back to parse loop else block \n")
	}

	print("done with ELSE, onto END IF \n")

	if !(pp.CurrentTokenIs(token.END) && pp.NextTokenIs(token.IF)) {
		pp.ReportError("[If Statement] expected END IF after THEN and ELSE blocks")
		return nil
	}
	pp.NextToken() // consume the END token
	pp.NextToken() // consume the IF token
	// this was inccorrect as I was consuming the semicolon
	//pp.NextToken() // move to the next token after IF-ELSE block

	fmt.Printf("done with IF, back to statement function\n")
	return ifStatement
}

func (pp *ParserPointer) ParseLoopStatement() *ast.LoopStatement {
	forStatement := &ast.LoopStatement{}

	if !pp.CurrentTokenIs(token.FOR) {
		pp.ReportError("[For Statement] expected FOR keyword")
		return nil
	}
	pp.NextToken()

	if !pp.CurrentTokenIs(token.LPAREN) {
		pp.ReportError("[For Statement] expected LEFT PARENTHESIS after FOR keyword")
		return nil
	}
	pp.NextToken()

	forStatement.Initialization = pp.ParseAssignmentStatement()

	if !pp.CurrentTokenIs(token.SEMICOLON) {
		pp.ReportError("[For Statement] expected SEMICOLON after INITIALIZATION")
		return nil
	}
	pp.NextToken()

	forStatement.Condition = pp.ParseExpression()

	if !pp.CurrentTokenIs(token.RPAREN) {
		pp.ReportError("[For Statement] expected RIGHT PARENTHESIS after CONDITION")
		return nil
	}
	pp.NextToken()

	forStatement.Body = pp.ParseStatements()
	print("back to parse loop body \n")
	print("current token is: " + pp.currentToken.Value + "\n")

	if !(pp.CurrentTokenIs(token.END) && pp.NextTokenIs(token.FOR)) {
		pp.ReportError("[For Statement] expected END FOR after LOOP body")
		return nil
	}

	pp.NextToken() // consume the END token
	pp.NextToken() // consume the FOR token and move to next token

	print("done with FOR, back to statement function \n")
	return forStatement
}

func (pp *ParserPointer) ParseReturnStatement() *ast.ReturnStatement {
	returnStatement := &ast.ReturnStatement{}

	if !pp.CurrentTokenIs(token.RETURN) {
		pp.ReportError("[Return Statement] expected RETURN keyword")
		return nil
	}
	pp.NextToken()

	returnStatement.Expression = pp.ParseExpression()

	fmt.Printf("expression detected, back to statement function\n")

	return returnStatement
}

func (pp *ParserPointer) ParseIdentifier() *ast.Identifier {
	identifier := &ast.Identifier{}

	if !pp.CurrentTokenIs(token.IDENTIFIER) {
		pp.ReportError("[Identifier] expected IDENTIFIER")
		return nil
	}
	identifier.Name = pp.currentToken.Value
	pp.NextToken()

	return identifier
}

// expression: ('not'?) arithOp ( '&' | '|' expression)*;
func (pp *ParserPointer) ParseExpression() *ast.Expression {
	expression := &ast.Expression{}

	if pp.CurrentTokenIs(token.NOT) {
		expression.IsNot = true
		pp.NextToken()
	}

	expression.ArithOp = pp.ParseArithmeticOperation()
	if expression.ArithOp == nil {
		pp.ReportError("[Expression] expected ARITHMETIC OPERATION")
		return nil
	}

	expression.AndOrList = pp.ParseAndOrList()

	return expression
}

func (pp *ParserPointer) ParseAndOrList() []ast.AndOrExpression {
	andOrList := []ast.AndOrExpression{}

	// todo - check if an error can be thrown in the func
	for pp.CurrentTokenIs(token.AND) || pp.CurrentTokenIs(token.OR) {
		andOr := &ast.AndOrExpression{}

		if pp.CurrentTokenIs(token.AND) {
			andOr.Operator = "and"
		} else if pp.CurrentTokenIs(token.OR) {
			andOr.Operator = "or"
		}
		pp.NextToken()

		andOr.Expression = pp.ParseExpression()
		andOrList = append(andOrList, *andOr)
	}

	return andOrList
}

// arithOp: relation ( '+' | '-' arithOp)*;
func (pp *ParserPointer) ParseArithmeticOperation() *ast.ArithmeticOperation {
	arithmeticOperation := &ast.ArithmeticOperation{}

	arithmeticOperation.Relation = pp.ParseRelation()
	if arithmeticOperation.Relation == nil {
		pp.ReportError("[Arithmetic Operation] expected a RELATION")
		return nil
	}

	arithmeticOperation.AddSubList = pp.ParseAddSubList()

	return arithmeticOperation
}

func (pp *ParserPointer) ParseAddSubList() []ast.AddSubExpression {
	addSubList := []ast.AddSubExpression{}

	// todo - check if an error can be thrown in the func
	for pp.CurrentTokenIs(token.ADD) || pp.CurrentTokenIs(token.SUB) {
		addSub := &ast.AddSubExpression{}

		if pp.CurrentTokenIs(token.ADD) {
			addSub.Operator = "+"
		} else if pp.CurrentTokenIs(token.SUB) {
			addSub.Operator = "-"
		}
		pp.NextToken()

		addSub.ArithmeticOperation = pp.ParseArithmeticOperation()
		addSubList = append(addSubList, *addSub)
	}

	return addSubList
}

// relation: term ( '==' | '!=' | '<' | '<=' | '>' | '>=' term)*;
func (pp *ParserPointer) ParseRelation() *ast.Relation {
	relation := &ast.Relation{}

	relation.Term = pp.ParseTerm()
	if relation.Term == nil {
		pp.ReportError("[Relation] expected a TERM")
		return nil
	}

	relation.RelationalOperationList = pp.ParseRelationalOpList()

	return relation
}

func (pp *ParserPointer) ParseRelationalOpList() []ast.RelationalExpression {
	relationalOpList := []ast.RelationalExpression{}

	// todo - check if an error can be thrown in the func
	for pp.CurrentTokenIs(token.EQ) || pp.CurrentTokenIs(token.NOT_EQ) || pp.CurrentTokenIs(token.LT) ||
		pp.CurrentTokenIs(token.LTE) || pp.CurrentTokenIs(token.GT) || pp.CurrentTokenIs(token.GTE) {
		relationalOp := &ast.RelationalExpression{}

		if pp.CurrentTokenIs(token.EQ) {
			relationalOp.Operator = "=="
		} else if pp.CurrentTokenIs(token.NOT_EQ) {
			relationalOp.Operator = "!="
		} else if pp.CurrentTokenIs(token.LT) {
			relationalOp.Operator = "<"
		} else if pp.CurrentTokenIs(token.LTE) {
			relationalOp.Operator = "<="
		} else if pp.CurrentTokenIs(token.GT) {
			relationalOp.Operator = ">"
		} else if pp.CurrentTokenIs(token.GTE) {
			relationalOp.Operator = ">="
		}
		pp.NextToken()

		relationalOp.Term = pp.ParseTerm()
		relationalOpList = append(relationalOpList, *relationalOp)
	}

	return relationalOpList
}

// term: factor ( '*' | '/' factor)*;
func (pp *ParserPointer) ParseTerm() *ast.Term {
	term := &ast.Term{}

	term.Factor = pp.ParseFactor()
	if term.Factor == nil {
		pp.ReportError("[Term] expected a FACTOR")
		return nil
	}

	term.MultDivList = pp.ParseMulDivList()

	return term
}

func (pp *ParserPointer) ParseMulDivList() []ast.MultDivExpression {
	mulDivList := []ast.MultDivExpression{}

	// todo - check if an error can be thrown in the func
	for pp.CurrentTokenIs(token.TIMES) || pp.CurrentTokenIs(token.DIV) {
		mulDiv := &ast.MultDivExpression{}

		if pp.CurrentTokenIs(token.TIMES) {
			mulDiv.Operator = "*"
		} else if pp.CurrentTokenIs(token.DIV) {
			mulDiv.Operator = "/"
		}
		pp.NextToken()

		mulDiv.Factor = pp.ParseFactor()
		mulDivList = append(mulDivList, *mulDiv)
	}

	return mulDivList
}

// factor: '(' expression ')' | procedure_call | ('-')? (name | number | string | 'true' | 'false');
func (pp *ParserPointer) ParseFactor() *ast.Factor {
	factor := &ast.Factor{}

	print("parsing factor \n")

	if pp.CurrentTokenIs(token.SUB) {
		factor.IsNegative = true
		pp.NextToken()
	}

	print("going into switch \n")

	switch pp.currentToken.Type {
	case token.TRUE:
		factor.IsBool = true
		factor.BoolValue = token.TRUE
		pp.NextToken()
	case token.FALSE:
		factor.IsBool = true
		factor.BoolValue = token.FALSE
		pp.NextToken()
	case token.NUMBER:
		factor.IsNumber = true
		factor.Number = pp.ParseNumber()
	case token.STRING:
		print("string detected \n")
		factor.IsString = true
		factor.String = pp.ParseString()
	case token.IDENTIFIER:
		print("identifier detected \n")
		print(pp.currentToken.Value + "\n")
		// procedure call: next token is '('
		if pp.NextTokenIs(token.LPAREN) {
			factor.IsProcedureCall = true
			factor.ProcedureCall = pp.ParseProcedureCall()
		} else {
			factor.IsName = true
			factor.Name = pp.ParseName()
		}
	case token.LPAREN:
		print("left parenthesis detected \n")
		pp.NextToken()
		factor.Expression = pp.ParseExpression()
		if !pp.CurrentTokenIs(token.RPAREN) {
			pp.ReportError("[Factor] expected RIGHT PARENTHESIS after EXPRESSION")
			return nil
		}
		pp.NextToken()
	default:
		pp.ReportError("[Factor] expected NAME, NUMBER, STRING, BOOL, PROCEDURE CALL or EXPRESSION")
		return nil
	}

	return factor
}

func (pp *ParserPointer) ParseName() *ast.Name {
	name := &ast.Name{Identifier: &ast.Identifier{}}

	if !pp.CurrentTokenIs(token.IDENTIFIER) {
		pp.ReportError("[Identifier] expected IDENTIFIER")
		return nil
	}
	name.Identifier.Name = pp.currentToken.Value
	pp.NextToken()

	if pp.CurrentTokenIs(token.LSQUARE) {
		pp.NextToken()
		name.Expression = pp.ParseExpression()
		if !pp.CurrentTokenIs(token.RSQUARE) {
			pp.ReportError("[Identifier] expected RIGHT SQUARE BRACE after EXPRESSION in array call")
			return nil
		}
		name.IsArray = true
		pp.NextToken()
	} else {
		name.IsArray = false
	}

	return name
}

func (pp *ParserPointer) ParseArgumentList() *ast.ArgumentList {
	argumentList := &ast.ArgumentList{Arguments: []ast.Expression{}}

	if pp.CurrentTokenIs(token.RPAREN) {
		return argumentList
	}

	firstExpression := pp.ParseExpression()
	argumentList.Arguments = append(argumentList.Arguments, *firstExpression)

	// todo - check if an error can be thrown in the func
	for pp.CurrentTokenIs(token.COMMA) {
		pp.NextToken()
		// check for multiple commas
		if pp.CurrentTokenIs(token.COMMA) {
			pp.ReportError("[Argument List] expected EXPRESSION after COMMA, found another COMMA")
			return nil
		}
		argument := pp.ParseExpression()
		argumentList.Arguments = append(argumentList.Arguments, *argument)

	}
	print("done with argument list \n")
	return argumentList
}

func (pp *ParserPointer) ParseNumber() *ast.Number {
	number := &ast.Number{}

	if !pp.CurrentTokenIs(token.NUMBER) {
		pp.ReportError("[Number] expected NUMBER")
		return nil
	}
	number.Value = pp.currentToken.Value
	pp.NextToken()

	return number
}

func (pp *ParserPointer) ParseString() *ast.String {
	str := &ast.String{}

	if !pp.CurrentTokenIs(token.STRING) {
		pp.ReportError("[String] expected STRING")
		return nil
	}
	str.Value = pp.currentToken.Value
	pp.NextToken()

	return str
}
