package main

import (
	"fmt"
	"os"
)

type TokenType int

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	//One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	//Literals.
	IDENTIFIER
	STRING
	NUMBER
	//Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	EOF
)

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   interface{}
	line      int
}

func (t TokenType) String() string {
	return [...]string{
		"LEFT_PAREN",
		"RIGHT_PAREN",
		"LEFT_BRACE",
		"RIGHT_BRACE",
		"COMMA",
		"DOT",
		"MINUS",
		"PLUS",
		"SEMICOLON",
		"SLASH",
		"STAR",
		//One or two character tokens.
		"BANG",
		"BANG_EQUAL",
		"EQUAL",
		"EQUAL_EQUAL",
		"GREATER",
		"GREATER_EQUAL",
		"LESS",
		"LESS_EQUAL",
		//Literals.
		"IDENTIFIER",
		"STRING",
		"NUMBER",
		//Keywords.
		"AND",
		"CLASS",
		"ELSE",
		"FALSE",
		"FUN",
		"FOR",
		"IF",
		"NIL",
		"OR",
		"PRINT",
		"RETURN",
		"SUPER",
		"THIS",
		"TRUE",
		"VAR",
		"WHILE",
		"EOF",
	}[t]
}

// keywords map
var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

func (t Token) String() string {
	switch t.tokenType {
	case STRING:
		return fmt.Sprintf("%s \"%s\" %s", t.tokenType, t.lexeme, t.literal)
	case NUMBER:
		switch num := t.literal.(type) {
		case float64:
			return fmt.Sprintf("%s %s %s", t.tokenType, t.lexeme, printFloat(num))
		default:
			fmt.Fprintf(os.Stderr, "error printing float.")
			os.Exit(65)
		}
	default:
		return fmt.Sprintf("%s %s %s", t.tokenType, t.lexeme, "null")
	}
	return ""
}
