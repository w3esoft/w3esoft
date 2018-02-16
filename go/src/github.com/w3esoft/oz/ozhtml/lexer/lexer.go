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
	var char uint8 = lexer.Read();
	switch {
	case char == GREATER:
		char = lexer.Input.Read()
		switch char {
		case BACKSLASH:
			len:=lexer.Pos() - offset
			tok:=token.New(token.TAGHEADCLOSE, nil, token.Position{Len: len , Offset: offset})
			return tok, nil
		case BANG:
			char = lexer.Input.Read()
			char = lexer.Input.Read()
			var ii = 0
			var value = ""
			for {
				value = value + string(char)
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
			return token.New(token.TAGCOMMENT, &value, token.Position{Len: len, Offset: offset}), nil
		default:
			lexer.CHAR = int(char)
			len:=lexer.Pos() - offset
			return token.New(token.TAGHEADOPEN, nil, token.Position{Len: len, Offset: offset}), nil
		}
	case char == BACKSLASH:
		len:=lexer.Pos() - offset
		return token.New(token.TAGHEADCLOSE, nil, token.Position{Len: len, Offset: offset}), nil
	case char == MINOR:
		var value = ""
		char = lexer.Input.Read()
		for char != GREATER && char != EOF {
			value = value + string(char)
			char = lexer.Input.Read()
			lexer.pos++
		}
		lexer.CHAR = int(char)

		len:=lexer.Pos() - offset
		return token.New(token.TAGBODY, &value, token.Position{Len: len, Offset: offset}), nil
	case char == PARENTHESISOPEN:
		len:=lexer.Pos() - offset
		return token.New(token.PARENTHESISOPEN, nil, token.Position{Len: len, Offset: offset}), nil
	case char == PARENTHESISCLOSE:
		len:=lexer.Pos() - offset
		return token.New(token.PARENTHESISCLOSE, nil, token.Position{Len: len, Offset: offset}), nil
	case char == EQUAL:
		len:=lexer.Pos() - offset
		return token.New(token.EQUAL, nil, token.Position{Len: len, Offset: offset}), nil
	case char == MULTIPLICATION:
		len:=lexer.Pos() - offset
		return token.New(token.MULTIPLICATION, nil, token.Position{Len: len, Offset: offset}), nil
	case char == BRACKETOPEN:
		len:=lexer.Pos() - offset
		return token.New(token.BRACKETOPEN, nil, token.Position{Len: len, Offset: offset}), nil
	case char == BRACKETClOSE:
		len:=lexer.Pos() - offset
		return token.New(token.BRACKETCLOSE, nil, token.Position{Len: len, Offset: offset}), nil
	case char == DOT:
		len:=lexer.Pos() - offset
		return token.New(token.DOT, nil, token.Position{Len: len, Offset: offset}), nil
	case char == DOUBLEDOT:
		len:=lexer.Pos() - offset
		return token.New(token.DOUBLEDOT, nil, token.Position{Len: len, Offset: offset}), nil
	case char == DOUBLEQUOTES:
		value := ""
		char = lexer.Input.Read()
		for char != DOUBLEQUOTES && char != EOF {
			value = value + string(char)
			char = lexer.Input.Read()
		}
		len:=lexer.Pos() - offset
		return token.New(token.STRING, nil, token.Position{Len: len, Offset: offset}), nil
	case char == EOF:
		len:=lexer.Pos() - offset
		return token.New(token.EOF, nil, token.Position{Len: len, Offset: offset}), nil

	case IsWhiteSpace(char):
		for IsWhiteSpace(char) {
			char = lexer.Input.Read()
		}
		lexer.CHAR = int(char)
		len:=lexer.Pos() - offset
		return token.New(token.WHITESPACE, nil, token.Position{Len: len, Offset: offset}), nil
	case IsWord(char):
		value := ""
		for IsWord(char) || IsNumeric(char) {
			value = value + string(char)
			char = lexer.Input.Read()
		}
		lexer.CHAR = int(char)
		len:=lexer.Pos() - offset
		return token.New(token.WORD, &value, token.Position{Len: len, Offset: offset}), nil
	case IsNumeric(char):
		value := ""
		for IsNumeric(char) {
			value = value + string(char)
			char = lexer.Input.Read()
		}
		lexer.CHAR = int(char)
		len:=lexer.Pos() - offset
		return token.New(token.NUMERIC, &value, token.Position{Len: len, Offset: offset}), nil
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
