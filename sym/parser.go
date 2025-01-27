package sym

import "fmt"

type Parser struct {
	tokens  []Token
	current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) parse() []Stmt {
	var statements []Stmt
	for !p.isAtEnd() {
		statements = append(statements, p.declaration())
	}
	return statements
}

func (p *Parser) declaration() (declaration Stmt) {
	defer func() {
		err := recover()
		if err != nil {
			p.synchronize()
			declaration = nil
			return
		}
	}()
	if p.match(VAR) {
		return p.varDeclaration()
	} else {
		return p.statement()
	}
}

func (p *Parser) varDeclaration() Stmt {
	name := p.consume(IDENTIFIER, "Expect variable name.")
	var initializer Expr
	if p.match(ASSIGN) {
		initializer = p.expression()
	}
	p.consume(SEMICOLON, "Expect ';' after variable declaration.")
	return NewVarStmt(name, initializer)
}

func (p *Parser) statement() Stmt {
	if p.match(LEFTBRACE) {
		block := p.block()
		return NewBlockStmt(block)
	} else if p.match(BREAK) {
		return p.breakStatement()
	} else if p.match(IF) {
		return p.ifStatement()
	} else if p.match(LOOP) {
		return p.loopStatement()
	} else if p.match(PRINT) {
		return p.printStatement()
	} else {
		return p.expressionStatement()
	}
}

func (p *Parser) breakStatement() Stmt {
	keyword := p.previous()
	p.consume(SEMICOLON, fmt.Sprintf("Expect ';' after '%s.'", BREAK))
	return NewBreakStmt(keyword)
}

func (p *Parser) ifStatement() Stmt {
	p.consume(LEFTPARENTHESIS, fmt.Sprintf("Expect '(' after '%s'", IF))
	condition := p.expression()
	p.consume(RIGHTPARENTHESIS, fmt.Sprintf("Expect ')' after '%s' condition", IF))
	then := p.statement()
	return NewIfStmt(condition, then)
}

func (p *Parser) loopStatement() Stmt {
	body := p.statement()
	return NewLoopStmt(body)
}

func (p *Parser) printStatement() Stmt {
	expression := p.expression()
	p.consume(SEMICOLON, "Expect ';' after expression.")
	return NewPrintStmt(expression)
}

func (p *Parser) block() []Stmt {
	var statements []Stmt
	for !p.check(RIGHTBRACE) && !p.isAtEnd() {
		declaration := p.declaration()
		statements = append(statements, declaration)
	}
	p.consume(RIGHTBRACE, "Expect '}' after block.")
	return statements
}

func (p *Parser) expressionStatement() Stmt {
	expr := p.expression()
	p.consume(SEMICOLON, "Expect ';' after expression.")
	return NewExpressionStmt(expr)
}

func (p *Parser) expression() Expr {
	return p.assignment()
}

func (p *Parser) assignment() Expr {
	expr := p.or()
	if p.match(ASSIGN) {
		assign := p.previous()
		value := p.or()
		varExpr, ok := expr.(*VarExpr)
		if ok {
			return NewAssignExpr(varExpr.Name, value)
		}
		panic(fmt.Sprintf("Invalid assignment target '%v'.", assign))
	}
	return expr
}

func (p *Parser) or() Expr {
	expr := p.and()
	for p.match(OR) {
		operator := p.previous()
		right := p.and()
		expr = NewLogicalExpr(expr, operator, right)
	}
	return expr
}

func (p *Parser) and() Expr {
	expr := p.equality()
	for p.match(AND) {
		operator := p.previous()
		right := p.equality()
		expr = NewLogicalExpr(expr, operator, right)
	}
	return expr
}

func (p *Parser) equality() Expr {
	expr := p.comparison()
	for p.match(NOTEQUAL, EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = NewBinaryExpr(expr, operator, right)
	}
	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.term()
	for p.match(GREATER, GREATEREQUAL, LESS, LESSEQUAL) {
		operator := p.previous()
		right := p.term()
		expr = NewBinaryExpr(expr, operator, right)
	}
	return expr
}

func (p *Parser) term() Expr {
	expr := p.factor()
	for p.match(MINUS, PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = NewBinaryExpr(expr, operator, right)
	}
	return expr
}

func (p *Parser) factor() Expr {
	expr := p.unary()
	for p.match(DIVIDE, MULTIPLY) {
		operator := p.previous()
		right := p.unary()
		expr = NewBinaryExpr(expr, operator, right)
	}
	return expr
}

func (p *Parser) unary() Expr {
	if p.match(BANG, MINUS) {
		operator := p.previous()
		right := p.unary()
		return NewUnaryExpr(operator, right)
	}
	return p.primary()
}

func (p *Parser) primary() Expr {
	if p.match(FALSE) {
		return NewLiteralExpr(false)
	} else if p.match(TRUE) {
		return NewLiteralExpr(true)
	} else if p.match(NIL) {
		return NewLiteralExpr(nil)
	} else if p.match(NUMBER, STRING) {
		return NewLiteralExpr(p.previous().Literal)
	} else if p.match(IDENTIFIER) {
		return NewVarExpr(p.previous())
	} else {
		panic(fmt.Sprintf("Expected expression at line %d.", p.peek().Line))
	}
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) check(tokenType string) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().TokenType == tokenType
}

func (p *Parser) consume(tokenType string, message string) Token {
	if p.check(tokenType) {
		return p.advance()
	}
	panic(fmt.Sprintf("%s at line %d.\n", message, p.peek().Line))
}

func (p *Parser) match(tokenTypes ...string) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == EOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().TokenType == SEMICOLON {
			return
		}
		switch p.peek().TokenType {
		case FUNC:
		case IF:
		case LOOP:
		case PRINT:
		case RETURN:
		case VAR:
			return
		}
		p.advance()
	}
}
