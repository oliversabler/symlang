package sym

import "fmt"

type Visitor interface {
	visitBinaryExpr(expr *BinaryExpr) interface{}
	visitLiteralExpr(expr *LiteralExpr) interface{}
	visitUnaryExpr(expr *UnaryExpr) interface{}

	visitExpressionStmt(stmt *ExpressionStmt) interface{}
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
