package sym

import "fmt"

type Visitor interface {
	visitAssignExpr(expr *AssignExpr) interface{}
	visitBinaryExpr(expr *BinaryExpr) interface{}
	visitLiteralExpr(expr *LiteralExpr) interface{}
	visitUnaryExpr(expr *UnaryExpr) interface{}
	visitVarExpr(expr *VarExpr) interface{}

	visitBlockStmt(stmt *BlockStmt) interface{}
	visitBreakStmt(stmt *BreakStmt) interface{}
	visitExpressionStmt(stmt *ExpressionStmt) interface{}
	visitIfStmt(stmt *IfStmt) interface{}
	visitLoopStmt(stmt *LoopStmt) interface{}
	visitPrintStmt(stmt *PrintStmt) interface{}
	visitVarStmt(stmt *VarStmt) interface{}
}

/*
/ Expressions
*/
type Expr interface {
	Accept(visitor Visitor) interface{}
}

type AssignExpr struct {
	Name  Token
	Value Expr
}

func NewAssignExpr(name Token, value Expr) *AssignExpr {
	return &AssignExpr{
		Name:  name,
		Value: value,
	}
}

func (ae *AssignExpr) Accept(visitor Visitor) interface{} {
	return visitor.visitAssignExpr(ae)
}

func (ae *AssignExpr) String() string {
	return fmt.Sprintf("AssignExpr {Name %v,Value: %v}", ae.Name, ae.Value)
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

type VarExpr struct {
	Name Token
}

func NewVarExpr(name Token) *VarExpr {
	return &VarExpr{
		Name: name,
	}
}

func (ve *VarExpr) Accept(visitor Visitor) interface{} {
	return visitor.visitVarExpr(ve)
}

func (ve *VarExpr) String() string {
	return fmt.Sprintf("VarExpr {Name: %v}", ve.Name)
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

type IfStmt struct {
	Condition Expr
	Then      Stmt
}

func NewIfStmt(condition Expr, then Stmt) *IfStmt {
	return &IfStmt{
		Condition: condition,
		Then:      then,
	}
}

func (is *IfStmt) Accept(visitor Visitor) interface{} {
	return visitor.visitIfStmt(is)
}

func (is *IfStmt) String() string {
	return fmt.Sprintf("IfStmt {Condition: %v,Then: %v}", is.Condition, is.Then)
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

type VarStmt struct {
	Name        Token
	Initializer Expr
}

func NewVarStmt(name Token, initializer Expr) *VarStmt {
	return &VarStmt{
		Name:        name,
		Initializer: initializer,
	}
}

func (vs *VarStmt) Accept(visitor Visitor) interface{} {
	return visitor.visitVarStmt(vs)
}

func (vs *VarStmt) String() string {
	return fmt.Sprintf("VarStmt {Name: %v,Initializer: %v}", vs.Name, vs.Initializer)
}
