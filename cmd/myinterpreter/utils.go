package main

import (
	"fmt"
	"math"
	"os"
)

func printFloat(value float64) string {
	_, fracPart := math.Modf(value)
	if fracPart == 0 {
		return fmt.Sprintf("%.1f", value)
	}
	return fmt.Sprintf("%g", value)
}

func reportError(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s", line, where, message)
}
