package sym

import (
	"fmt"
	"strconv"
)

var keywords = map[string]string{
	"&": AND,
	"←": ASSIGN,
	"Ɵ": BREAK,
	"○": FALSE,
	"ƒ": FUNC,
	"¿": IF,
	"∞": LOOP,
	"ø": NIL,
	"|": OR,
	"✉": PRINT,
	"↵": RETURN,
	"●": TRUE,
	"•": VAR,
}

type Lexer struct {
	source  []rune
	tokens  []Token
	start   int
	current int
	line    int
}

func NewLexer(source string) Lexer {
	return Lexer{
		source:  []rune(source),
		tokens:  []Token{},
		start:   0,
		current: 0,
		line:    1,
	}
}

func (l *Lexer) scanTokens() []Token {
	for !l.isAtEnd() {
		l.start = l.current
		l.scanToken()
	}
	l.tokens = append(l.tokens, NewToken(EOF, "", nil, l.line))
	return l.tokens
}

func (l *Lexer) scanToken() {
	c := l.advance()
	switch c {
	case '(':
		l.addToken(LEFTPARENTHESIS)
	case ')':
		l.addToken(RIGHTPARENTHESIS)
	case '{':
		l.addToken(LEFTBRACE)
	case '}':
		l.addToken(RIGHTBRACE)
	case ',':
		l.addToken(COMMA)
	case '.':
		l.addToken(DOT)
	case ';':
		l.addToken(SEMICOLON)
	case '-':
		l.addToken(MINUS)
	case '+':
		l.addToken(PLUS)
	case '÷':
		l.addToken(DIVIDE)
	case '×':
		l.addToken(MULTIPLY)
	case '!':
		l.addToken(BANG)
	case '=':
		l.addToken(EQUAL)
	case '≠':
		l.addToken(NOTEQUAL)
	case '>':
		l.addToken(GREATER)
	case '≥':
		l.addToken(GREATEREQUAL)
	case '<':
		l.addToken(LESS)
	case '≤':
		l.addToken(LESSEQUAL)
	case ' ', '\r', '\t':
		break
	case '\n':
		l.line++
	case '"':
		l.string()
	default:
		if l.isDigit(c) {
			l.number()
		} else if l.isAlpha(c) {
			l.identifier()
		} else {
			fmt.Println(fmt.Errorf("Unexpected character %s at line %d.", string(c), l.line))
		}
	}
}

func (l *Lexer) addToken(tokenType string) {
	l.addTokenLiteral(tokenType, nil)
}

func (l *Lexer) addTokenLiteral(tokenType string, literal interface{}) {
	text := string(l.source[l.start:l.current])
	token := NewToken(tokenType, text, literal, l.line)
	l.tokens = append(l.tokens, token)
}

func (l *Lexer) advance() rune {
	c := l.source[l.current]
	l.current++
	return c
}

func (l *Lexer) isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_' ||
		c == '&' ||
		c == '←' ||
		c == 'Ɵ' ||
		c == '○' ||
		c == 'ƒ' ||
		c == '¿' ||
		c == '∞' ||
		c == 'ø' ||
		c == '|' ||
		c == '✉' ||
		c == '↵' ||
		c == '●' ||
		c == '•'
}

func (l *Lexer) identifier() {
	for l.isAlphaNumeric(l.peek()) {
		l.advance()
	}
	text := string(l.source[l.start:l.current])
	tokenType, ok := keywords[text]
	if ok {
		l.addToken(tokenType)
	} else {
		l.addTokenLiteral(IDENTIFIER, text)
	}
}

func (l *Lexer) isAlphaNumeric(c rune) bool {
	return l.isAlpha(c) || l.isDigit(c)
}

func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.source)
}

func (l *Lexer) isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func (l *Lexer) match(expected rune) bool {
	if l.isAtEnd() {
		return false
	}
	if l.source[l.current] != expected {
		return false
	}
	l.current++
	return true
}

func (l *Lexer) number() {
	for l.isDigit(l.peek()) {
		l.advance()
	}
	if l.peek() == '.' && l.isDigit(l.peekNext()) {
		l.advance()
		for l.isDigit(l.peek()) {
			l.advance()
		}
	}
	number, err := strconv.ParseFloat(string(l.source[l.start:l.current]), 64)
	if err != nil {
		fmt.Println(fmt.Errorf("%s at line %d.", err.Error(), l.line))
	}
	l.addTokenLiteral(NUMBER, number)
}

func (l *Lexer) peek() rune {
	if l.isAtEnd() {
		return '\000'
	}
	return l.source[l.current]
}

func (l *Lexer) peekNext() rune {
	if l.current+1 >= len(l.source) {
		return '\000'
	}
	return l.source[l.current]
}

func (l *Lexer) string() {
	for l.peek() != '"' && !l.isAtEnd() {
		if l.peek() == '\n' {
			l.line++
		}
		l.advance()
	}
	if l.isAtEnd() {
		fmt.Println(fmt.Errorf("Unterminated string at line %d.", l.line))
	}
	l.advance()
	value := string(l.source[l.start+1 : l.current-1])
	l.addTokenLiteral(STRING, value)
}
