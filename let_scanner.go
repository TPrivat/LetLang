// Author: Taylor Privat

package main

import (
	"text/scanner"
	"unicode"
)

const (
	LETTER = 0
	DIGIT = 1
	UNKNOWN = 99
	INT_LIT = 10
	IDENT = 11
	LET = 16
	IN = 17
	IF = 18
	THEN = 19
	ELSE = 20
	EQUAL = 21
	COMMA = 22
	MINUS = 23
	ISZERO = 24
	LEFT_PAREN = 25
	RIGHT_PAREN = 26
)

// Token representation
type Token struct {
	tokenType int
	tokenValue string
}

// Node for Abstract Syntax Tree
type astNode struct {
	parent *astNode		// ptr to parent node, null if root
	ttype string		// token type
	termsym bool		// terminal symbol? (root node)
	contents string
	children []*astNode
}

var (
	charClass int
	lexeme string
	nextChar rune
	nextToken int
)

func addChar() {
	var temp = string(nextChar)
	lexeme += temp
}

func getChar() {
	nextChar = s.Next()
	if nextChar == scanner.EOF {
		charClass = scanner.EOF
	} else {
		if unicode.IsLetter(nextChar) {
			charClass = LETTER
		} else if unicode.IsDigit(nextChar) {
			charClass = DIGIT
		} else {
			charClass = UNKNOWN
		}
	}
}

func getNonBlank() {
	for unicode.IsSpace(nextChar) {
		getChar()

	}
}

func lookup(ch rune) int {
	switch ch {
	case '(':
		addChar()
		nextToken = LEFT_PAREN
		break
	case ')':
		addChar()
		nextToken = RIGHT_PAREN
		break
	case '=':
		addChar()
		nextToken = EQUAL
		break
	case ',':
		addChar()
		nextToken = COMMA
		break
	default:
		addChar()
		nextToken = scanner.EOF
		break
	}
	return nextToken
}

func lex() int {
	lexeme = ""
	getNonBlank()
	switch charClass {
	case LETTER:
		addChar()
		getChar()
		for charClass == LETTER || charClass == DIGIT {
			addChar()
			getChar()
		}
		nextToken = IDENT
		break
	case DIGIT:
		addChar()
		getChar()
		for charClass == DIGIT {
			addChar()
			getChar()
		}
		nextToken = INT_LIT
		break
	case UNKNOWN:
		lookup(nextChar)
		getChar()
		break
	case scanner.EOF:
		nextToken = scanner.EOF
		lexeme = "EOF"
		break
	}

	if (nextToken == IDENT) {
		switch lexeme {
		case "let":
			nextToken = LET
			break
		case "in":
			nextToken = IN
			break
		case "if":
			nextToken = IF
			break
		case "else":
			nextToken = ELSE
			break
		case "then":
			nextToken = THEN
			break
		case "iszero":
			nextToken = ISZERO
			break
		case "minus":
			nextToken = MINUS
			break
		}
	}
	//fmt.Println("Next token is: ", nextToken, ", Next lexeme is: ", lexeme)
	tokenQueue = append(tokenQueue, Token{nextToken, lexeme})
	return nextToken
}



