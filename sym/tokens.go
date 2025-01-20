package sym

type Token struct {
	TokenType string
	Lexeme    string
	Literal   interface{}
	Line      int
}

func NewToken(tokenType string, lexeme string, literal interface{}, line int) Token {
	return Token{
		TokenType: tokenType,
		Lexeme:    lexeme,
		Literal:   literal,
		Line:      line,
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
	BREAK  = "Ɵ"
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
