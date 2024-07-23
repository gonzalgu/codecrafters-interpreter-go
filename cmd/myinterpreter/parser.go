package main

import (
	"fmt"
)

type ParseError struct {
	token   Token
	message string
}

type Parser struct {
	tokens  []Token
	current int
}

func (e ParseError) Error() string {
	return fmt.Sprintf("Parse error at %v: %s", e.token, e.message)
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

func (parser *Parser) expression() *Expr {
	return parser.equality()
}

func (parser *Parser) equality() *Expr {
	expr := parser.comparison()
	for parser.match(BANG_EQUAL, EQUAL_EQUAL) {
		op := parser.previous()
		right := parser.comparison()
		expr = &Expr{
			exprType: BINARY_EXPR,
			left:     expr,
			operator: op,
			right:    right,
		}
	}
	return expr
}

func (parser *Parser) comparison() *Expr {
	expr := parser.term()
	for parser.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		operator := parser.previous()
		right := parser.term()
		expr = &Expr{
			exprType: BINARY_EXPR,
			left:     expr,
			operator: operator,
			right:    right,
		}
	}
	return expr
}

func (parser *Parser) term() *Expr {
	expr := parser.factor()
	for parser.match(MINUS, PLUS) {
		operator := parser.previous()
		right := parser.factor()
		expr = &Expr{
			exprType: BINARY_EXPR,
			left:     expr,
			operator: operator,
			right:    right,
		}
	}
	return expr
}

func (parser *Parser) factor() *Expr {
	expr := parser.unary()
	for parser.match(SLASH, STAR) {
		operator := parser.previous()
		right := parser.unary()
		expr = &Expr{
			exprType: BINARY_EXPR,
			left:     expr,
			operator: operator,
			right:    right,
		}
	}
	return expr
}

func (parser *Parser) unary() *Expr {
	if parser.match(BANG, MINUS) {
		operator := parser.previous()
		right := parser.unary()
		return &Expr{
			exprType: UNARY,
			operator: operator,
			right:    right,
		}
	}
	return parser.primary()
}

func (parser *Parser) primary() *Expr {
	if parser.match(FALSE) {
		return &Expr{
			exprType: LITERAL,
			value:    false,
		}
	}
	if parser.match(TRUE) {
		return &Expr{
			exprType: LITERAL,
			value:    true,
		}
	}
	if parser.match(NIL) {
		return &Expr{
			exprType: LITERAL,
			value:    nil,
		}
	}
	if parser.match(NUMBER, STRING) {
		return &Expr{
			exprType: LITERAL,
			value:    parser.previous().literal,
		}
	}
	if parser.match(LEFT_PAREN) {
		expr := parser.expression()
		parser.consume(RIGHT_PAREN, "")
		return &Expr{
			exprType: GROUPING_EXPR,
			left:     expr,
		}
	}
	panic(parser.error(parser.peek(), "Expect expression."))
}

func (parser *Parser) consume(tokenType TokenType, message string) Token {
	if parser.check(tokenType) {
		return parser.advance()
	}
	panic(parser.error(parser.peek(), message))
}

func (parser *Parser) peek() Token {
	return parser.tokens[parser.current]
}

func (parser *Parser) advance() Token {
	if !parser.isAtEnd() {
		parser.current++
	}
	return parser.previous()
}

func (parser *Parser) isAtEnd() bool {
	return parser.peek().tokenType == EOF
}

func (parser *Parser) previous() Token {
	return parser.tokens[parser.current-1]
}

func (parser *Parser) check(tokenType TokenType) bool {
	if parser.isAtEnd() {
		return false
	}
	return parser.peek().tokenType == tokenType
}

func (parser *Parser) match(tokenTypes ...TokenType) bool {
	for _, tt := range tokenTypes {
		if parser.check(tt) {
			parser.advance()
			return true
		}
	}
	return false
}

func (p *Parser) error(t Token, message string) ParseError {
	return ParseError{token: t, message: message}
}

// Parse attempts to parse the tokens into an expression
func (p *Parser) Parse() (*Expr, error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(ParseError); ok {
				return
			}
			panic(r)
		}
	}()

	return p.expression(), nil
}
