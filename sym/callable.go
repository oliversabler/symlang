package sym

type SymCallable interface {
	Arity() int
	Call(interpreter *Interpreter, arguments []interface{}) interface{}
}

type SymFunction struct {
	Declaration *FunctionStmt
	Closure     *Environment
}

func NewSymFunction(declaration *FunctionStmt, closure *Environment) *SymFunction {
	return &SymFunction{
		Declaration: declaration,
		Closure:     closure,
	}
}

func (sf SymFunction) Arity() int {
	return len(sf.Declaration.Params)
}

func (sf SymFunction) Call(interpreter *Interpreter, arguments []interface{}) (returnValue interface{}) {
	envlosingEnvironment := interpreter.environment
	environment := NewEnvironmentWithEnclosing(sf.Closure)
	defer func() {
		err := recover()
		if err != nil {
			symReturn, ok := err.(*SymReturn)
			if !ok {
				panic(err)
			}
			returnValue = symReturn.Value
			interpreter.environment = envlosingEnvironment
			return
		}
	}()
	for i, argument := range sf.Declaration.Params {
		environment.define(argument.Lexeme, arguments[i])
	}
	interpreter.executeBlock(sf.Declaration.Body, environment)
	return interpreter.currentValue
}

func (sf SymFunction) bind() *SymFunction {
	environment := NewEnvironmentWithEnclosing(sf.Closure)
	return NewSymFunction(sf.Declaration, environment)
}

type SymReturn struct {
	Value interface{}
}

func NewSymReturn(value interface{}) *SymReturn {
	return &SymReturn{Value: value}
}
