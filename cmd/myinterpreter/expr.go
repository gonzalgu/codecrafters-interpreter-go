package main

import "fmt"

type ExprType int

const (
	BINARY_EXPR ExprType = iota
	GROUPING_EXPR
	LITERAL
	UNARY
)

type Expr struct {
	exprType ExprType
	left     *Expr
	operator Token
	right    *Expr
	value    interface{}
}

func print_ast(expr *Expr) string {
	if expr == nil {
		return ""
	}
	switch expr.exprType {
	case BINARY_EXPR:
		return parenthesize(expr.operator.lexeme, expr.left, expr.right)
	case GROUPING_EXPR:
		return parenthesize("group", expr.left)
	case LITERAL:
		if expr.value == nil {
			return "nil"
		}

		switch v := expr.value.(type) {
		case float64:
			return printFloat(v)
		default:
			return fmt.Sprintf("%v", expr.value)
		}

	case UNARY:
		return parenthesize(expr.operator.lexeme, expr.right)
	}
	return ""
}

func parenthesize(name string, exprs ...*Expr) string {
	result := "(" + name
	for _, expr := range exprs {
		str := print_ast(expr)
		result += " " + str
	}
	result += ")"
	return result
}
