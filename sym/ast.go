package sym

import "fmt"

type Visitor interface {
	visitBinaryExpr(expr *BinaryExpr) interface{}
	visitLiteralExpr(expr *LiteralExpr) interface{}
	visitUnaryExpr(expr *UnaryExpr) interface{}

	visitBlockStmt(stmt *BlockStmt) interface{}
	visitBreakStmt(stmt *BreakStmt) interface{}
	visitExpressionStmt(stmt *ExpressionStmt) interface{}
	visitLoopStmt(stmt *LoopStmt) interface{}
	visitPrintStmt(stmt *PrintStmt) interface{}
}

/*
/ Expressions
*/
type Expr interface {
	Accept(visitor Visitor) interface{}
}

type BinaryExpr struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func NewBinaryExpr(left Expr, operator Token, right Expr) *BinaryExpr {
	return &BinaryExpr{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func (be *BinaryExpr) Accept(visitor Visitor) interface{} {
	return visitor.visitBinaryExpr(be)
}

func (be *BinaryExpr) String() string {
	return fmt.Sprintf("BinaryExpr {Left: %v,Operator: %v,Right: %v}",
		be.Left, be.Operator, be.Right)
}

type LiteralExpr struct {
	Value interface{}
}

func NewLiteralExpr(value interface{}) *LiteralExpr {
	return &LiteralExpr{
		Value: value,
	}
}

func (le *LiteralExpr) Accept(visitor Visitor) interface{} {
	return visitor.visitLiteralExpr(le)
}

func (le *LiteralExpr) String() string {
	return fmt.Sprintf("LiteralExpr {Value: %v}", le.Value)
}

type UnaryExpr struct {
	Operator Token
	Right    Expr
}

func NewUnaryExpr(operator Token, right Expr) *UnaryExpr {
	return &UnaryExpr{
		Operator: operator,
		Right:    right,
	}
}

func (ue *UnaryExpr) Accept(visitor Visitor) interface{} {
	return visitor.visitUnaryExpr(ue)
}

func (ue *UnaryExpr) String() string {
	return fmt.Sprintf("UnaryExpr {Operator: %v,Right: %v}", ue.Operator, ue.Right)
}

/*
/ Statements
*/
type Stmt interface {
	Accept(visitor Visitor) interface{}
}

type BlockStmt struct {
	Statements []Stmt
}

func NewBlockStmt(statements []Stmt) *BlockStmt {
	return &BlockStmt{
		Statements: statements,
	}
}

func (bs *BlockStmt) Accept(visitor Visitor) interface{} {
	return visitor.visitBlockStmt(bs)
}

func (bs *BlockStmt) String() string {
	return fmt.Sprintf("BlockStmt {Statements: %v}", bs.Statements)
}

type BreakStmt struct {
	Token Token
}

func NewBreakStmt(token Token) *BreakStmt {
	return &BreakStmt{}
}

func (bs *BreakStmt) Accept(visitor Visitor) interface{} {
	return visitor.visitBreakStmt(bs)
}

func (bs *BreakStmt) String() string {
	return fmt.Sprintf("BreakStmt {Token: %v}", bs.Token)
}

type ExpressionStmt struct {
	Expression Expr
}

func NewExpressionStmt(expression Expr) *ExpressionStmt {
	return &ExpressionStmt{
		Expression: expression,
	}
}

func (es *ExpressionStmt) Accept(visitor Visitor) interface{} {
	return visitor.visitExpressionStmt(es)
}

func (es *ExpressionStmt) String() string {
	return fmt.Sprintf("ExpressionStmt {Expression: %v}", es.Expression)
}

type LoopStmt struct {
	Body Stmt
}

func NewLoopStmt(body Stmt) *LoopStmt {
	return &LoopStmt{
		Body: body,
	}
}

func (ls *LoopStmt) Accept(visitor Visitor) interface{} {
	return visitor.visitLoopStmt(ls)
}

func (ls *LoopStmt) String() string {
	return fmt.Sprintf("LoopStmt {Body: %v}", ls.Body)
}

type PrintStmt struct {
	Expression Expr
}

func NewPrintStmt(expression Expr) *PrintStmt {
	return &PrintStmt{
		Expression: expression,
	}
}

func (ps *PrintStmt) Accept(visitor Visitor) interface{} {
	return visitor.visitPrintStmt(ps)
}

func (ps *PrintStmt) String() string {
	return fmt.Sprintf("PrintStmt {Expression: %v}", ps.Expression)
}
