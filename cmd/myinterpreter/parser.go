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

func (parser *Parser) expression() (*Expr, error) {
	res, err := parser.equality()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (parser *Parser) equality() (*Expr, error) {
	expr, err := parser.comparison()
	if err != nil {
		return nil, err
	}

	for parser.match(BANG_EQUAL, EQUAL_EQUAL) {
		op := parser.previous()
		right, err := parser.comparison()
		if err != nil {
			return nil, err
		}
		expr = &Expr{
			exprType: BINARY_EXPR,
			left:     expr,
			operator: op,
			right:    right,
		}
	}
	return expr, nil
}

func (parser *Parser) comparison() (*Expr, error) {
	expr, err := parser.term()
	if err != nil {
		return nil, err
	}
	for parser.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		operator := parser.previous()
		right, err := parser.term()
		if err != nil {
			return nil, err
		}
		expr = &Expr{
			exprType: BINARY_EXPR,
			left:     expr,
			operator: operator,
			right:    right,
		}
	}
	return expr, nil
}

func (parser *Parser) term() (*Expr, error) {
	expr, err := parser.factor()
	if err != nil {
		return nil, err
	}
	for parser.match(MINUS, PLUS) {
		operator := parser.previous()
		right, err := parser.factor()
		if err != nil {
			return nil, err
		}
		expr = &Expr{
			exprType: BINARY_EXPR,
			left:     expr,
			operator: operator,
			right:    right,
		}
	}
	return expr, nil
}

func (parser *Parser) factor() (*Expr, error) {
	expr, err := parser.unary()
	if err != nil {
		return nil, err
	}
	for parser.match(SLASH, STAR) {
		operator := parser.previous()
		right, err := parser.unary()
		if err != nil {
			return nil, err
		}
		expr = &Expr{
			exprType: BINARY_EXPR,
			left:     expr,
			operator: operator,
			right:    right,
		}
	}
	return expr, nil
}

func (parser *Parser) unary() (*Expr, error) {
	if parser.match(BANG, MINUS) {
		operator := parser.previous()
		right, err := parser.unary()
		if err != nil {
			return nil, err
		}
		return &Expr{
			exprType: UNARY,
			operator: operator,
			right:    right,
		}, nil
	}
	return parser.primary()
}

func (parser *Parser) primary() (*Expr, error) {
	if parser.match(FALSE) {
		return &Expr{
			exprType: LITERAL,
			value:    false,
		}, nil
	}
	if parser.match(TRUE) {
		return &Expr{
			exprType: LITERAL,
			value:    true,
		}, nil
	}
	if parser.match(NIL) {
		return &Expr{
			exprType: LITERAL,
			value:    nil,
		}, nil
	}
	if parser.match(NUMBER, STRING) {
		return &Expr{
			exprType: LITERAL,
			value:    parser.previous().literal,
		}, nil
	}
	if parser.match(LEFT_PAREN) {
		expr, err := parser.expression()
		if err != nil {
			return nil, err
		}

		if _, err := parser.consume(RIGHT_PAREN, "Expect ')' after expression."); err != nil {
			return nil, err
		}
		return &Expr{
			exprType: GROUPING_EXPR,
			left:     expr,
		}, nil
	}
	return nil, parser.error(parser.peek(), "Expect expression.")
	//panic(parser.error(parser.peek(), "Expect expression."))
}

func (parser *Parser) consume(tokenType TokenType, message string) (*Token, error) {
	if parser.check(tokenType) {
		t := parser.advance()
		return &t, nil
	}
	return nil, parser.error(parser.peek(), message)
	//panic(parser.error(parser.peek(), message))
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
	return p.expression()
}
