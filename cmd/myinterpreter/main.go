package main

import (
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	//fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]
	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	/*
		if len(fileContents) == 0 {
			fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
		}*/

	switch command {
	case "tokenize":
		scanner := NewScanner(fileContents)
		tokens := scanner.ScanToks()
		//fmt.Printf("%v\n", tokens)
		for _, tok := range tokens {
			fmt.Printf("%s\n", tok)
		}
		if scanner.hadError {
			os.Exit(65)
		}
	case "parse":
		scanner := NewScanner(fileContents)
		tokens := scanner.ScanToks()
		parser := NewParser(tokens)
		result, err := parser.Parse()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Parsing error.")
			os.Exit(1)
		}
		fmt.Printf("%s\n", print_ast(result))
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	// Uncomment this block to pass the first stage

	os.Exit(0)
}
