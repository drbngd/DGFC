// Code generated from example.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // example

import "github.com/antlr4-go/antlr/v4"

// exampleListener is a complete listener for a parse tree produced by exampleParser.
type exampleListener interface {
	antlr.ParseTreeListener

	// EnterSet_of_stmts is called when entering the set_of_stmts production.
	EnterSet_of_stmts(c *Set_of_stmtsContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterIf_stmt is called when entering the if_stmt production.
	EnterIf_stmt(c *If_stmtContext)

	// EnterWhile_stmt is called when entering the while_stmt production.
	EnterWhile_stmt(c *While_stmtContext)

	// EnterAssignment_stmt is called when entering the assignment_stmt production.
	EnterAssignment_stmt(c *Assignment_stmtContext)

	// EnterCond_expr is called when entering the cond_expr production.
	EnterCond_expr(c *Cond_exprContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// ExitSet_of_stmts is called when exiting the set_of_stmts production.
	ExitSet_of_stmts(c *Set_of_stmtsContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitIf_stmt is called when exiting the if_stmt production.
	ExitIf_stmt(c *If_stmtContext)

	// ExitWhile_stmt is called when exiting the while_stmt production.
	ExitWhile_stmt(c *While_stmtContext)

	// ExitAssignment_stmt is called when exiting the assignment_stmt production.
	ExitAssignment_stmt(c *Assignment_stmtContext)

	// ExitCond_expr is called when exiting the cond_expr production.
	ExitCond_expr(c *Cond_exprContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)
}
