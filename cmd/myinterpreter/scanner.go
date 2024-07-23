package main

import (
	"fmt"
	"os"
	"strconv"
)

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

func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current++
	return c
}

func (s *Scanner) ScanToks() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		c := s.advance()
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
		case '/':
			if s.match('/') {
				for s.peek() != '\n' && !s.isAtEnd() {
					s.advance()
				}
			} else {
				s.addToken(SLASH)
			}
		case ' ':
			fallthrough
		case '\r':
			fallthrough
		case '\t':
			//do nothing
		case '\n':
			s.line++
		case '"':
			s.lexString()

		default:
			if isDigit(c) {
				s.lexNumber()
			} else if isAlpha(c) {
				s.identifier()
			} else {
				reportError(s.line, "", "Unexpected character:")
				fmt.Fprintf(os.Stderr, " %c\n", c)
				s.hadError = true
			}

			//panic("Unexpected character")
		}
	}
	s.tokens = append(s.tokens, Token{
		EOF,
		"",
		nil,
		s.line,
	})
	return s.tokens
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	tokenType, exists := keywords[string(text)]
	if !exists {
		tokenType = IDENTIFIER
	}
	s.addToken(tokenType)
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

func (s *Scanner) lexNumber() {
	for isDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}
	text := string(s.source[s.start:s.current])
	num, _ := strconv.ParseFloat(text, 64)
	s.tokens = append(s.tokens, Token{
		NUMBER,
		text,
		num,
		s.line,
	})
}

func (s *Scanner) lexString() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		reportError(s.line, "", "Unterminated string.")
		s.hadError = true
		return
	}
	s.advance()
	value := string(s.source[s.start+1 : s.current-1])
	s.tokens = append(s.tokens, Token{
		STRING,
		value,
		value,
		s.line,
	})
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

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}

func (s *Scanner) addToken(tokenType TokenType) {
	text := string(s.source[s.start:s.current])
	s.tokens = append(s.tokens, Token{
		tokenType,
		text,
		nil,
		s.line,
	})
}
