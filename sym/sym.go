package sym

import (
	"fmt"
	"os"
)

type Runtime struct {
}

func NewRuntime() Runtime {
	return Runtime{}
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
	//	for _, token := range tokens {
	//		fmt.Printf("%#v\n", token)
	//	}

	parser := NewParser(tokens)
	statements := parser.parse()
	for _, statement := range statements {
		r.debugStatement(fmt.Sprintf("%v", statement))
	}
}

func (r *Runtime) debugStatement(output string) {
	indentation := 0
	for i, c := range output {
		if c == '{' {
			fmt.Print(string(c))
			fmt.Println()
			indentation += 2
			r.indent(indentation)
		} else if output[i] == '}' {
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
}

func (r *Runtime) indent(indentation int) {
	for range indentation {
		fmt.Print(" ")
	}
}
