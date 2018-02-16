package lexer

import (
	"errors"
	"github.com/w3esoft/glaive/input"
	"github.com/w3esoft/oz/ozhtml/token"
	"fmt"
)

type Lexer struct {
	CHAR   int
	Input  input.Input
	Tokens []token.Token
	pos    int
}

func New(input input.Input) Lexer {
	lexer := Lexer{}
	lexer.CHAR = -1
	lexer.pos = 0
	lexer.Input = input
	return lexer
}

func (lexer *Lexer) Read() uint8 {
	var char uint8;
	if lexer.CHAR != -1 {
		char = uint8(lexer.CHAR)
		lexer.CHAR = -1
	} else {
		char = lexer.Input.Read()
		lexer.pos++
	}
	return char;
}
func (lexer *Lexer) Pos() int {
	if lexer.CHAR == -1 {
		return lexer.pos
	} else {
		return lexer.pos-1;
	}
}

func (lexer *Lexer) Tokenize() (r *token.Token, err error) {
	if len(lexer.Tokens) > 0 {
		r = &lexer.Tokens[len(lexer.Tokens)-1:][0]
		lexer.Tokens = lexer.Tokens[:len(lexer.Tokens)-1]
		return r, nil
	}
	offset := lexer.Pos()
	var value *string = nil
	char:= lexer.Read();
	switch {
	case char == GREATER:
		char = lexer.Input.Read()
		switch char {
		case BACKSLASH:
			len:=lexer.Pos() - offset
			position := &token.Position{Len: len , Offset: offset}
			tok:=token.New(token.TAGHEADCLOSE, value, position)
			return tok, nil
		case BANG:
			char = lexer.Input.Read()
			char = lexer.Input.Read()
			var ii = 0
			var v = ""
			for {
				v = v + string(char)
				char = lexer.Input.Read()
				lexer.pos++
				switch {
				case char == EOF:
					break
				case ii == 0 && char == MINUS:
					ii = 1
				case ii == 1 && char == MINUS:
					ii = 2
				case ii == 2 && char == MINOR:
					break
				default:
					ii = 0
				}
			}
			len:=lexer.Pos() - offset
			position := &token.Position{Len: len , Offset: offset}
			return token.New(token.TAGCOMMENT, value, position), nil
		default:
			lexer.CHAR = int(char)
			len:=lexer.Pos() - offset
			position := &token.Position{Len: len , Offset: offset}
			return token.New(token.TAGHEADOPEN, value, position), nil
		}
	case char == BACKSLASH:
		len:=lexer.Pos() - offset
		position := &token.Position{Len: len , Offset: offset}
		return token.New(token.TAGHEADCLOSE, value, position), nil
	case char == MINOR:
		var v = ""
		char = lexer.Input.Read()
		for char != GREATER && char != EOF {
			v = v + string(char)
			char = lexer.Input.Read()
			lexer.pos++
		}
		lexer.CHAR = int(char)
		len:=lexer.Pos() - offset
		position := &token.Position{Len: len , Offset: offset}
		value = &v
		return token.New(token.TAGBODY, value, position), nil
	case char == PARENTHESISOPEN:
		len:=lexer.Pos() - offset
		position := &token.Position{Len: len , Offset: offset}
		return token.New(token.PARENTHESISOPEN, value, position), nil
	case char == PARENTHESISCLOSE:
		len:=lexer.Pos() - offset
		position := &token.Position{Len: len , Offset: offset}
		return token.New(token.PARENTHESISCLOSE, value, position), nil
	case char == EQUAL:
		len:=lexer.Pos() - offset
		position := &token.Position{Len: len , Offset: offset}
		return token.New(token.EQUAL, value, position), nil
	case char == MULTIPLICATION:
		len:=lexer.Pos() - offset
		position := &token.Position{Len: len , Offset: offset}
		return token.New(token.MULTIPLICATION, value, position), nil
	case char == BRACKETOPEN:
		len:=lexer.Pos() - offset
		position := &token.Position{Len: len , Offset: offset}
		return token.New(token.BRACKETOPEN, value, position), nil
	case char == BRACKETClOSE:
		len:=lexer.Pos() - offset
		position := &token.Position{Len: len , Offset: offset}
		return token.New(token.BRACKETCLOSE, value, position), nil
	case char == DOT:
		len:=lexer.Pos() - offset
		position := &token.Position{Len: len , Offset: offset}
		return token.New(token.DOT, value, position), nil
	case char == DOUBLEDOT:
		len:=lexer.Pos() - offset
		position := &token.Position{Len: len , Offset: offset}
		return token.New(token.DOUBLEDOT, value, position), nil
	case char == DOUBLEQUOTES:
		v := ""
		char = lexer.Input.Read()
		for char != DOUBLEQUOTES && char != EOF {
			v = v + string(char)
			char = lexer.Input.Read()
		}
		len:=lexer.Pos() - offset
		position := &token.Position{Len: len , Offset: offset}
		value = &v
		return token.New(token.STRING, value, position), nil
	case char == EOF:
		len:=lexer.Pos() - offset
		position := &token.Position{Len: len , Offset: offset}
		return token.New(token.EOF, value, position), nil

	case IsWhiteSpace(char):
		for IsWhiteSpace(char) {
			char = lexer.Input.Read()
		}
		lexer.CHAR = int(char)
		len:=lexer.Pos() - offset
		position := &token.Position{Len: len , Offset: offset}
		return token.New(token.WHITESPACE, value, position), nil
	case IsWord(char):
		v := ""
		for IsWord(char) || IsNumeric(char) {
			v = v + string(char)
			char = lexer.Input.Read()
		}
		lexer.CHAR = int(char)
		len:=lexer.Pos() - offset
		position := &token.Position{Len: len , Offset: offset}
		value = &v
		return token.New(token.WORD, value, position), nil
	case IsNumeric(char):
		v := ""
		for IsNumeric(char) {
			v = v + string(char)
			char = lexer.Input.Read()
		}
		lexer.CHAR = int(char)
		len:=lexer.Pos() - offset
		position := &token.Position{Len: len , Offset: offset}
		value = &v
		return token.New(token.NUMERIC, value, position), nil
	default:
		return nil, errors.New(fmt.Sprintf("%s %s ", "unexpected char ", int(char)))
	}
}
func IsWord(char uint8) bool {
	return (96 < char && 123 > char) || (64 < char && 91 > char || char == MINUS || char == 95 || char == 35)
}
func IsNumeric(char uint8) bool {
	return 47 < char && 58 > char
}

func IsWhiteSpace(char uint8) bool {
	return char == 32 || char == 13 || char == 10 || char == 9
}
