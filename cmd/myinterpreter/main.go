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
	line      int
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s null", t.tokenType, t.lexeme)
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

type Scanner struct {
	source   []byte
	tokens   []Token
	start    int
	current  int
	line     int
	hadError bool
}

func NewScanner(source []byte) Scanner {
	return Scanner{
		source,
		[]Token{},
		0,
		0,
		1,
		false,
	}
}

func (s *Scanner) ScanToks() []Token {
	for i := 0; i < len(s.source); i++ {
		s.start = s.current
		s.current++
		c := s.source[i]
		switch c {
		case '(':
			s.addToken(LEFT_PAREN)
		case ')':
			s.addToken(RIGHT_PAREN)
		case '{':
			s.addToken(LEFT_BRACE)
		case '}':
			s.addToken(RIGHT_BRACE)
		case ',':
			s.addToken(COMMA)
		case '.':
			s.addToken(DOT)
		case '-':
			s.addToken(MINUS)
		case '+':
			s.addToken(PLUS)
		case ';':
			s.addToken(SEMICOLON)
		case '*':
			s.addToken(STAR)
		case '!':
			var tok TokenType
			if s.match('=') {
				tok = BANG_EQUAL
			} else {
				tok = BANG
			}
			s.addToken(tok)
		case '=':
			var tok TokenType
			if s.match('=') {
				tok = EQUAL_EQUAL
			} else {
				tok = EQUAL
			}
			s.addToken(tok)
		case '<':
			var tok TokenType
			if s.match('=') {
				tok = LESS_EQUAL
			} else {
				tok = LESS
			}
			s.addToken(tok)
		case '>':
			var tok TokenType
			if s.match('=') {
				tok = GREATER_EQUAL
			} else {
				tok = GREATER
			}
			s.addToken(tok)
		case ' ':
			fallthrough
		case '\r':
			fallthrough
		case '\t':
			//do nothing
		case '\n':
			s.line++

		default:
			reportError(s.line, "", "Unexpected character:")
			fmt.Fprintf(os.Stderr, " %c\n", c)
			s.hadError = true
			//panic("Unexpected character")
		}
	}
	s.tokens = append(s.tokens, Token{
		EOF,
		"",
		s.line,
	})
	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != byte(expected) {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return 0
	}
	return rune(s.source[s.current])
}

func (s *Scanner) addToken(tokenType TokenType) {
	text := string(s.source[s.start:s.current])
	s.tokens = append(s.tokens, Token{
		tokenType,
		text,
		s.line,
	})
}

func reportError(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s", line, where, message)
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	//fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	// Uncomment this block to pass the first stage

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
		scanner := NewScanner(fileContents)
		tokens := scanner.ScanToks()
		//fmt.Printf("%v\n", tokens)
		for _, tok := range tokens {
			fmt.Printf("%s\n", tok)
		}
		if scanner.hadError {
			os.Exit(65)
		}
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
	os.Exit(0)
}
