package sym

import "fmt"

type Resolver struct {
	interpreter *Interpreter
	scopes      []map[string]bool
}

func NewResolver(interpreter *Interpreter) *Resolver {
	return &Resolver{
		interpreter: interpreter,
		scopes:      make([]map[string]bool, 0),
	}
}

func (r *Resolver) resolveStatements(statements []Stmt) {
	defer func() {
		err := recover()
		if err != nil {
			panic(err)
		}
	}()
	for _, statement := range statements {
		r.resolveStatement(statement)
	}
}

func (r *Resolver) resolveStatement(statement Stmt) {
	statement.Accept(r)
}

func (r *Resolver) resolveExpression(expression Expr) {
	expression.Accept(r)
}

func (r *Resolver) resolveLocal(expression Expr, name Token) {
	for i := len(r.scopes) - 1; i >= 0; i-- {
		_, ok := r.scopes[i][name.Lexeme]
		if ok {
			r.interpreter.resolve(expression, len(r.scopes)-1-i)
			return
		}
	}
}

func (r *Resolver) beginScope() {
	r.scopes = append(r.scopes, make(map[string]bool))
}

func (r *Resolver) endScope() {
	r.scopes = r.scopes[:len(r.scopes)-1]
}

func (r *Resolver) declare(name Token) {
	if len(r.scopes) == 0 {
		return
	}
	_, ok := r.scopes[len(r.scopes)-1][name.Lexeme]
	if ok {
		panic(fmt.Sprintf("Variable '%s' already declared in this scope.", name.Lexeme))
	}
	r.scopes[len(r.scopes)-1][name.Lexeme] = false
}

func (r *Resolver) define(name Token) {
	if len(r.scopes) == 0 {
		return
	}
	r.scopes[len(r.scopes)-1][name.Lexeme] = true
}

func (r *Resolver) visitAssignExpr(expression *AssignExpr) interface{} {
	r.resolveExpression(expression.Value)
	r.resolveLocal(expression, expression.Name)
	return nil
}

func (r *Resolver) visitBinaryExpr(expression *BinaryExpr) interface{} {
	r.resolveExpression(expression.Left)
	r.resolveExpression(expression.Right)
	return nil
}

func (r *Resolver) visitLiteralExpr(expression *LiteralExpr) interface{} {
	return nil
}

func (r *Resolver) visitLogicalExpr(expression *LogicalExpr) interface{} {
	r.resolveExpression(expression.Left)
	r.resolveExpression(expression.Right)
	return nil
}

func (r *Resolver) visitUnaryExpr(expression *UnaryExpr) interface{} {
	r.resolveExpression(expression.Right)
	return nil
}

func (r *Resolver) visitVarExpr(expression *VarExpr) interface{} {
	if len(r.scopes) > 0 {
		value, ok := r.scopes[len(r.scopes)-1][expression.Name.Lexeme]
		if ok && value == false {
			panic("Can't read local variable in its own initializer.")
		}
	}
	r.resolveLocal(expression, expression.Name)
	return nil
}

func (r *Resolver) visitBlockStmt(statement *BlockStmt) interface{} {
	r.beginScope()
	r.resolveStatements(statement.Statements)
	r.endScope()
	return nil
}

func (r *Resolver) visitBreakStmt(statement *BreakStmt) interface{} {
	return nil
}

func (r *Resolver) visitExpressionStmt(statement *ExpressionStmt) interface{} {
	r.resolveExpression(statement.Expression)
	return nil
}

func (r *Resolver) visitIfStmt(statement *IfStmt) interface{} {
	r.resolveExpression(statement.Condition)
	r.resolveStatement(statement.Then)
	return nil
}

func (r *Resolver) visitLoopStmt(statement *LoopStmt) interface{} {
	r.resolveStatement(statement.Body)
	return nil
}

func (r *Resolver) visitPrintStmt(statement *PrintStmt) interface{} {
	r.resolveExpression(statement.Expression)
	return nil
}

func (r *Resolver) visitVarStmt(statement *VarStmt) interface{} {
	r.declare(statement.Name)
	if statement.Initializer != nil {
		r.resolveExpression(statement.Initializer)
	}
	r.define(statement.Name)
	return nil
}
