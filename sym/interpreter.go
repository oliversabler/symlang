package sym

import "fmt"

type ActionType = string

type LoopAction struct {
	actionType ActionType
}

type Interpreter struct {
	environment  *Environment
	currentValue interface{}
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		currentValue: nil,
	}
}

func (i *Interpreter) interpret(statements []Stmt) interface{} {
	defer func() {
		err := recover()
		if err != nil {
			panic(err)
		}
	}()

	for _, statement := range statements {
		i.execute(statement)
	}

	return i.currentValue
}

func (i *Interpreter) evaluate(expression Expr) interface{} {
	return expression.Accept(i)
}

func (i *Interpreter) execute(statement Stmt) interface{} {
	return statement.Accept(i)
}

func (i *Interpreter) executeBlock(statements []Stmt, environment *Environment) {
	previous := i.environment
	i.environment = environment
	for _, statement := range statements {
		i.execute(statement)
	}
	i.environment = previous
}

func (i *Interpreter) visitBinaryExpr(expression *BinaryExpr) interface{} {
	left := i.evaluate(expression.Left)
	right := i.evaluate(expression.Right)

	switch expression.Operator.TokenType {
	case NOTEQUAL:
		return !i.isEqual(left, right)
	case EQUAL:
		return i.isEqual(left, right)
	default:
		leftValue, leftOk := left.(float64)
		rightValue, rightOk := right.(float64)

		if !leftOk || !rightOk {
			panic(fmt.Sprintf("Operands must be two numbers, got %v and %v.", left, right))
		}

		switch expression.Operator.TokenType {
		case MINUS:
			return leftValue - rightValue
		case PLUS:
			return leftValue + rightValue
		case DIVIDE:
			return leftValue / rightValue
		case MULTIPLY:
			return leftValue * rightValue
		case GREATER:
			return leftValue > rightValue
		case GREATEREQUAL:
			return leftValue >= rightValue
		case LESS:
			return leftValue < rightValue
		case LESSEQUAL:
			return leftValue <= rightValue
		}
	}

	panic("You done messed up.")
}

func (i *Interpreter) visitLiteralExpr(expression *LiteralExpr) interface{} {
	return expression.Value
}

func (i *Interpreter) visitUnaryExpr(expression *UnaryExpr) interface{} {
	right := i.evaluate(expression.Right)

	switch expression.Operator.TokenType {
	case BANG:
		return !i.isTruthy(right)
	case MINUS:
		value, ok := right.(float64)
		if ok {
			return -value
		}

		panic(fmt.Sprintf("Operand must be a number, got %v.", expression.Right))
	default:
		panic("You done messed up.")
	}
}

func (i *Interpreter) visitBlockStmt(statement *BlockStmt) interface{} {
	blockEnvironment := NewEnvironmentWithEnclosing(i.environment)
	i.executeBlock(statement.Statements, blockEnvironment)
	return nil
}

func (i *Interpreter) visitBreakStmt(statement *BreakStmt) interface{} {
	panic(LoopAction{"BREAK"})
}

func (i *Interpreter) visitExpressionStmt(statement *ExpressionStmt) interface{} {
	value := i.evaluate(statement.Expression)

	i.currentValue = value

	return value
}

func (i *Interpreter) visitLoopStmt(statement *LoopStmt) interface{} {
	for {
		if i.loop(statement.Body) == "BREAK" {
			break
		}
	}
	return nil
}

func (i *Interpreter) loop(body Stmt) (actionType ActionType) {
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case LoopAction:
				actionType = r.actionType
				break
			default:
				panic(r)
			}
		}
	}()

	i.execute(body)
	return ""
}

func (i *Interpreter) isEqual(left interface{}, right interface{}) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil {
		return false
	}

	return left == right
}

func (i *Interpreter) isTruthy(object interface{}) bool {
	if object == nil {
		return false
	}

	isBool, ok := object.(bool)
	if ok {
		return isBool
	}

	return true
}
