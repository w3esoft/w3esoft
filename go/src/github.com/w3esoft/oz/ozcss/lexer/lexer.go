package lexer

import (
	"github.com/w3esoft/glaive/input"
	"github.com/w3esoft/oz/ozhtml/token"
)

const (
	DOUBLEDOT        = 58
	GREATER          = 60
	MINOR            = 62
	EQUAL            = 61
	DOUBLEQUOTES     = 34
	MULTIPLICATION   = 42
	BRACKETOPEN      = 91
	BRACKETClOSE     = 93
	PARENTHESISOPEN  = 40
	PARENTHESISCLOSE = 41
	BACKSLASH        = 47
	DOT              = 46
	BANG             = 33
	MINUS            = 45
	UNDERSCORE       = 95
	HASH             = 35
	EOF              = 0
)

type Lexer struct {
	CHAR   int
	Input  input.Input
	Tokens [] *token.Token
	pos    int
}

func New(input input.Input) *Lexer {
	lexer := Lexer{}
	lexer.CHAR = -1
	lexer.pos = 0
	lexer.Input = input
	return &lexer
}

func (lexer *Lexer) Read() uint8 {
	var char uint8
	if lexer.CHAR != -1 {
		char = uint8(lexer.CHAR)
		lexer.CHAR = -1
	} else {
		lexer.pos++
		char = lexer.Input.Read()

	}
	return char
}
func (lexer *Lexer) Pos() int {
	if lexer.CHAR == -1 {
		return lexer.pos
	} else {
		return lexer.pos - 1
	}
}
func (lexer *Lexer) Tokenize() *token.Token {
	if len(lexer.Tokens) > 0 {
		r := lexer.Tokens[len(lexer.Tokens)-1:][0]
		lexer.Tokens = lexer.Tokens[:len(lexer.Tokens)-1]
		return r
	}
	offset := lexer.Pos()
	var value *string = nil
	char := lexer.Read()
	if char == EOF {
		position := &token.Position{Len: 0, Offset: offset}
		v := string(char)
		value = &v
		return token.New(token.EOF, value, position, true)
	}else {
		switch {
		case IsWhiteSpace(char):
			for IsWhiteSpace(char) {
				char = lexer.Read()
			}
			lexer.CHAR = int(char)
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.WHITESPACE, value, position, true)
		case IsNumeric(char):
			v := ""
			for IsNumeric(char) {
				v = v + string(char)
				char = lexer.Read()
			}
			lexer.CHAR = int(char)
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			value = &v
			return token.New(token.NUMERIC, value, position, true)
		case char==DOT:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.DOT, value, position, true)
		case char==DOT:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.DOT, value, position, true)

		default:
		}
	}
	len := lexer.Pos() - offset
	position := &token.Position{Len: len, Offset: offset}
	v := string(char)
	value = &v
	return token.New(token.INVALID, value, position, false)
}
func (l *Lexer) TokenPush(i2 *token.Token) {
	l.Tokens = append(l.Tokens, i2)
}

func IsWord(char uint8) bool {
	return (96 < char && 123 > char) || (64 < char && 91 > char)
}
func IsNumeric(char uint8) bool {
	return 47 < char && 58 > char
}

func IsWhiteSpace(char uint8) bool {
	return char == 32 || char == 13 || char == 10 || char == 9
}
