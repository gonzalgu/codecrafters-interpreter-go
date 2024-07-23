package main

import (
	"fmt"
	"testing"
)

func TestLex(t *testing.T) {
	text := "false"
	scanner := NewScanner([]byte(text))
	tokens := scanner.ScanToks()
	for _, tok := range tokens {
		fmt.Printf("%s\n", tok)
	}
	parser := NewParser(tokens)
	expr, err := parser.Parse()
	if err != nil {
		t.Errorf("failed to parse %v", err)
	}
	fmt.Printf("expression: [%v]", expr)
	fmt.Printf("pprint: [%s]", print_ast(expr))
}
