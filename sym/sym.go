package sym

import (
	"fmt"
	"os"
)

type Runtime struct {
	interpreter *Interpreter
}

func NewRuntime() Runtime {
	interpreter := NewInterpreter()
	return Runtime{
		interpreter: interpreter,
	}
}

func (r *Runtime) ExecFile(path string) {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	r.exec(string(input))
}

func (r *Runtime) exec(source string) {
	lexer := NewLexer(source)
	tokens := lexer.scanTokens()
	r.debugTokens(tokens)
	parser := NewParser(tokens)
	statements := parser.parse()
	r.debugStatements(statements)
	resolver := NewResolver(r.interpreter)
	resolver.resolveStatements(statements)
	result := r.interpreter.interpret(statements)
	fmt.Printf("Result: %v\n", result)
}

func (r *Runtime) debugTokens(tokens []Token) {
	for _, token := range tokens {
		fmt.Printf("%#v\n", token)
	}
}

func (r *Runtime) debugStatements(statements []Stmt) {
	for _, statement := range statements {
		r.debugStatement(fmt.Sprintf("%v", statement))
	}
}

func (r *Runtime) debugStatement(output string) {
	indentation := 0
	for i, c := range output {
		if c == '{' || c == '[' {
			fmt.Print(string(c))
			fmt.Println()
			indentation += 2
			r.indent(indentation)
		} else if output[i] == '}' || output[i] == ']' {
			fmt.Println()
			indentation -= 2
			r.indent(indentation)
			fmt.Print(string(c))
		} else if c == ',' {
			fmt.Print(string(c))
			fmt.Println()
			r.indent(indentation)
		} else {
			fmt.Print(string(c))
		}
	}
	fmt.Println()
}

func (r *Runtime) indent(indentation int) {
	for range indentation {
		fmt.Print(" ")
	}
}
