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
	for _, token := range tokens {
		fmt.Printf("%#v\n", token)
	}
}
