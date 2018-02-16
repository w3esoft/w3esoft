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
	pos int
}

func New(input input.Input) Lexer {
	lexer := Lexer{}
	lexer.CHAR = -1
	lexer.pos = 0
	lexer.Input = input
	return lexer
}

func (lexer *Lexer) Tokenize() (r *token.Token, err error) {
	if len(lexer.Tokens) > 0 {
		r = &lexer.Tokens[len(lexer.Tokens)-1:][0]
		lexer.Tokens = lexer.Tokens[:len(lexer.Tokens)-1]
		return r, nil
	}
	var char uint8
	if lexer.CHAR != -1 {
		char = uint8(lexer.CHAR)
		lexer.CHAR = -1
	} else {
		char = lexer.Input.Read()
	}
	switch {
	case char == GREATER:
		char = lexer.Input.Read()
		switch char {
		case BACKSLASH:
			lexer.pos++;
			return token.New(token.TAGHEADCLOSE, nil ), nil
		case BANG:
			char = lexer.Input.Read()
			char = lexer.Input.Read()
			var ii = 0
			var value = ""
			for {
				value = value + string(char)
				char = lexer.Input.Read()
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
			return token.New(token.TAGCOMMENT, &value), nil
		default:
			lexer.CHAR = int(char)
			return token.New(token.TAGHEADOPEN, nil), nil
		}
	case char == BACKSLASH:
		return token.New(token.TAGHEADCLOSE, nil), nil
	case char == MINOR:
		var value = ""
		char = lexer.Input.Read()
		for char != GREATER && char != EOF {
			value = value + string(char)
			char = lexer.Input.Read()
		}
		lexer.CHAR = int(char)
		return token.New(token.TAGBODY, &value), nil
	case char == PARENTHESISOPEN:
		return token.New(token.PARENTHESISOPEN, nil), nil

	case char == PARENTHESISCLOSE:
		return token.New(token.PARENTHESISCLOSE, nil), nil
	case char == EQUAL:
		return token.New(token.EQUAL, nil), nil
	case char == MULTIPLICATION:
		return token.New(token.MULTIPLICATION, nil), nil
	case char == BRACKETOPEN:
		return token.New(token.BRACKETOPEN, nil), nil
	case char == BRACKETClOSE:
		return token.New(token.BRACKETCLOSE, nil), nil
	case char == DOT:
		return token.New(token.DOT, nil), nil
	case char == DOUBLEDOT:
		return token.New(token.DOUBLEDOT, nil), nil
	case char == DOUBLEQUOTES:
		value := ""
		char = lexer.Input.Read()
		for char != DOUBLEQUOTES && char != EOF {
			value = value + string(char)
			char = lexer.Input.Read()
		}
		return token.New(token.STRING, nil), nil
	case char == EOF:
		return token.New(token.EOF, nil), nil

	case IsWhiteSpace(char):
		for IsWhiteSpace(char) {
			char = lexer.Input.Read()
		}
		lexer.CHAR = int(char)
		return token.New(token.WHITESPACE, nil), nil
	case IsWord(char):
		value := ""
		for IsWord(char) || IsNumeric(char) {
			value = value + string(char)
			char = lexer.Input.Read()
		}
		lexer.CHAR = int(char)
		return token.New(token.WORD, &value), nil
	case IsNumeric(char):
		value := ""
		for IsNumeric(char) {
			value = value + string(char)
			char = lexer.Input.Read()
		}
		lexer.CHAR = int(char)
		return token.New(token.NUMERIC, &value), nil
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
	return char == 32 || char == 13 || char == 10||char == 9
}
