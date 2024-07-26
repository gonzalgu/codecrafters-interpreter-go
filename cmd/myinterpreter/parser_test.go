package main

import (
	"fmt"
	"testing"
)

func TestParseLiteral(t *testing.T) {
	tokens := []Token{
		{
			tokenType: NUMBER,
			lexeme:    "123",
			literal:   123,
			line:      1,
		},
		{
			tokenType: EOF,
		},
	}
	parser := NewParser(tokens)
	expr, err := parser.Parse()
	if err == nil {
		fmt.Printf("%s", print_ast(expr))
	}
}

func TestParseSum(t *testing.T) {
	tokens := []Token{
		{
			tokenType: NUMBER,
			lexeme:    "2.0",
			literal:   2.0,
			line:      1,
		},
		{
			tokenType: PLUS,
			lexeme:    "+",
			line:      1,
		},
		{
			tokenType: NUMBER,
			lexeme:    "3.0",
			literal:   3.0,
			line:      1,
		},
		{
			tokenType: EOF,
		},
	}
	parser := NewParser(tokens)
	expr, err := parser.Parse()
	if err == nil {
		fmt.Printf("%s", print_ast(expr))
	}
}

func TestParseEmptyGroup(t *testing.T) {
	tokens := []Token{
		{
			tokenType: LEFT_PAREN,
			lexeme:    "(",
			line:      1,
		},
		{
			tokenType: RIGHT_PAREN,
			lexeme:    ")",
			line:      1,
		},
		{
			tokenType: EOF,
		},
	}
	parser := NewParser(tokens)
	expr, err := parser.Parse()
	if err == nil {
		fmt.Printf("%v", expr)
		fmt.Printf("%s", print_ast(expr))
	}

}
