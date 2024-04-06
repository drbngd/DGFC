// Code generated from example.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // example

import "github.com/antlr4-go/antlr/v4"

// BaseexampleListener is a complete listener for a parse tree produced by exampleParser.
type BaseexampleListener struct{}

var _ exampleListener = &BaseexampleListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseexampleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseexampleListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseexampleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseexampleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSet_of_stmts is called when production set_of_stmts is entered.
func (s *BaseexampleListener) EnterSet_of_stmts(ctx *Set_of_stmtsContext) {}

// ExitSet_of_stmts is called when production set_of_stmts is exited.
func (s *BaseexampleListener) ExitSet_of_stmts(ctx *Set_of_stmtsContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseexampleListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseexampleListener) ExitStmt(ctx *StmtContext) {}

// EnterIf_stmt is called when production if_stmt is entered.
func (s *BaseexampleListener) EnterIf_stmt(ctx *If_stmtContext) {}

// ExitIf_stmt is called when production if_stmt is exited.
func (s *BaseexampleListener) ExitIf_stmt(ctx *If_stmtContext) {}

// EnterWhile_stmt is called when production while_stmt is entered.
func (s *BaseexampleListener) EnterWhile_stmt(ctx *While_stmtContext) {}

// ExitWhile_stmt is called when production while_stmt is exited.
func (s *BaseexampleListener) ExitWhile_stmt(ctx *While_stmtContext) {}

// EnterAssignment_stmt is called when production assignment_stmt is entered.
func (s *BaseexampleListener) EnterAssignment_stmt(ctx *Assignment_stmtContext) {}

// ExitAssignment_stmt is called when production assignment_stmt is exited.
func (s *BaseexampleListener) ExitAssignment_stmt(ctx *Assignment_stmtContext) {}

// EnterCond_expr is called when production cond_expr is entered.
func (s *BaseexampleListener) EnterCond_expr(ctx *Cond_exprContext) {}

// ExitCond_expr is called when production cond_expr is exited.
func (s *BaseexampleListener) ExitCond_expr(ctx *Cond_exprContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseexampleListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseexampleListener) ExitExpr(ctx *ExprContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseexampleListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseexampleListener) ExitTerm(ctx *TermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseexampleListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseexampleListener) ExitFactor(ctx *FactorContext) {}
