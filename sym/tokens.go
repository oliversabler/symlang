package sym

type Token struct {
	tokenType string
	lexeme    string
	literal   interface{}
	line      int
}

func NewToken(tokenType string, lexeme string, literal interface{}, line int) Token {
	return Token{
		tokenType: tokenType,
		lexeme:    lexeme,
		literal:   literal,
		line:      line,
	}
}

const (
	LEFTPARENTHESIS  = "("
	RIGHTPARENTHESIS = ")"
	LEFTBRACE        = "{"
	RIGHTBRACE       = "}"

	COMMA     = ","
	DOT       = "."
	SEMICOLON = ";"

	MINUS    = "-"
	PLUS     = "+"
	DIVIDE   = "÷"
	MULTIPLY = "×"

	BANG         = "!"
	EQUAL        = "="
	NOTEQUAL     = "≠"
	GREATER      = ">"
	GREATEREQUAL = "≥"
	LESS         = "<"
	LESSEQUAL    = "≤"

	IDENTIFIER = "Identifier"
	STRING     = "String"
	NUMBER     = "Number"

	AND    = "&"
	ASSIGN = "←"
	BREAK  = "θƟ"
	FALSE  = "○"
	FUNC   = "ƒ"
	IF     = "¿"
	LOOP   = "∞"
	NIL    = "ø"
	OR     = "|"
	PRINT  = "✉"
	RETURN = "↵"
	TRUE   = "●"
	VAR    = "•"

	EOF = "EOF"
)