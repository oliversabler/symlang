package sym

import "fmt"

type Interpreter struct {
	currentValue interface{}
	environment  *Environment
	globals      *Environment
	locals       map[Expr]int
}

func NewInterpreter() *Interpreter {
	globals := NewEnvironment()
	return &Interpreter{
		currentValue: nil,
		environment:  globals,
		globals:      globals,
		locals:       make(map[Expr]int),
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

func (i Interpreter) resolve(expression Expr, depth int) {
	i.locals[expression] = depth
}

func (i *Interpreter) executeBlock(statements []Stmt, environment *Environment) {
	previous := i.environment
	i.environment = environment
	for _, statement := range statements {
		i.execute(statement)
	}
	i.environment = previous
}

func (i *Interpreter) visitAssignExpr(expression *AssignExpr) interface{} {
	value := i.evaluate(expression.Value)
	distance, ok := i.locals[expression]
	if ok {
		i.environment.assignAt(distance, expression.Name.Lexeme, value)
	} else {
		i.globals.assign(expression.Name.Lexeme, value)
	}
	return value
}

func (i *Interpreter) visitBinaryExpr(expression *BinaryExpr) interface{} {
	left := i.evaluate(expression.Left)
	right := i.evaluate(expression.Right)
	switch expression.Operator.TokenType {
	case NOTEQUAL:
		return !i.isEqual(left, right)
	case EQUAL:
		return i.isEqual(left, right)
	case PLUS:
		switch t := left.(type) {
		case float64:
			rightValue, rightOk := right.(float64)
			if rightOk {
				return t + rightValue
			}
		case string:
			rightValue, rightOk := right.(string)
			if rightOk {
				return t + rightValue
			}
		default:
			panic(fmt.Sprintf("Operands must be two numbers or two strings, got %v and %v.", left, right))
		}
	default:
		leftValue, leftOk := left.(float64)
		rightValue, rightOk := right.(float64)
		if !leftOk || !rightOk {
			panic(fmt.Sprintf("Operands must be two numbers, got %v and %v.", left, right))
		}
		switch expression.Operator.TokenType {
		case MINUS:
			return leftValue - rightValue
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

func (i *Interpreter) visitCallExpr(expression *CallExpr) interface{} {
	callee := i.evaluate(expression.Callee)
	var arguments []interface{}
	for _, argument := range expression.Arguments {
		argument := i.evaluate(argument)
		arguments = append(arguments, argument)
	}
	function, ok := callee.(SymCallable)
	if !ok {
		panic("Can only call functions.")
	}
	if len(arguments) != function.Arity() {
		panic(fmt.Sprintf("Expected %d arguments but got %d.", function.Arity(), len(arguments)))
	}
	value := function.Call(i, arguments)
	return value
}

func (i *Interpreter) visitLiteralExpr(expression *LiteralExpr) interface{} {
	return expression.Value
}

func (i *Interpreter) visitLogicalExpr(expression *LogicalExpr) interface{} {
	left := i.evaluate(expression.Left)
	if expression.Operator.TokenType == OR {
		if i.isTruthy(left) {
			return left
		}
	} else {
		if !i.isTruthy(left) {
			return left
		}
	}
	return i.evaluate(expression.Right)
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

func (i *Interpreter) visitVarExpr(expression *VarExpr) interface{} {
	return i.variableLookup(expression.Name, expression)
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

func (i *Interpreter) visitFunctionStmt(statement *FunctionStmt) interface{} {
	function := NewSymFunction(statement, i.environment)
	i.environment.define(statement.Name.Lexeme, function)
	return nil
}

func (i *Interpreter) visitIfStmt(statement *IfStmt) interface{} {
	condition := i.evaluate(statement.Condition)
	if i.isTruthy(condition) {
		return i.execute(statement.Then)
	}
	return nil
}

func (i *Interpreter) visitLoopStmt(statement *LoopStmt) interface{} {
	for {
		if i.loop(statement.Body) == "BREAK" {
			break
		}
	}
	return nil
}

type ActionType = string

type LoopAction struct {
	actionType ActionType
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

func (i *Interpreter) visitPrintStmt(statement *PrintStmt) interface{} {
	value := i.evaluate(statement.Expression)
	fmt.Printf("%v\n", value)
	i.currentValue = value
	return value
}

func (i *Interpreter) visitReturnStmt(statement *ReturnStmt) interface{} {
	value := i.evaluate(statement.Value)
	panic(NewSymReturn(value))
}

func (i *Interpreter) visitVarStmt(statement *VarStmt) interface{} {
	var value interface{}
	if statement.Initializer != nil {
		value = i.evaluate(statement.Initializer)
	}
	i.environment.define(statement.Name.Lexeme, value)
	return value
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

func (i *Interpreter) variableLookup(name Token, expression Expr) interface{} {
	distance, ok := i.locals[expression]
	if ok {
		return i.environment.getAt(distance, name.Lexeme)
	} else {
		value, ok := i.globals.get(name.Lexeme)
		if !ok {
			panic(fmt.Sprintf("Undefined variable '%s'.", name.Lexeme))
		}
		return value
	}
}
