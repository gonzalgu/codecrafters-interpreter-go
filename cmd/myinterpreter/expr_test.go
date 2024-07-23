package main

import "testing"

func TestPrintAst(t *testing.T) {
	tests := []struct {
		name     string
		expr     *Expr
		expected string
	}{
		{
			name: "Literal",
			expr: &Expr{
				exprType: LITERAL,
				value:    123,
			},
			expected: "123",
		},
		{
			name: "Unary",
			expr: &Expr{
				exprType: UNARY,
				operator: Token{lexeme: "-"},
				right: &Expr{
					exprType: LITERAL,
					value:    123,
				},
			},
			expected: "(- 123)",
		},
		{
			name: "Binary",
			expr: &Expr{
				exprType: GROUPING_EXPR,
				left: &Expr{
					exprType: LITERAL,
					value:    45.67,
				},
			},
			expected: "(group 45.67)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := print_ast(tt.expr)
			if result != tt.expected {
				t.Errorf("print_ast() = %s, want %s", result, tt.expected)
			}
		})
	}
}
