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
	EOF              = 0
)
const (
	MODE_TAG  = 1
	MODE_BODY = 2
)

type Lexer struct {
	CHAR   int
	Input  input.Input
	Tokens []token.Token
	pos    int
	mode   int
}

func New(input input.Input) Lexer {
	lexer := Lexer{}
	lexer.CHAR = -1
	lexer.pos = 0
	lexer.mode = MODE_BODY;
	lexer.Input = input
	return lexer
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
		r := &lexer.Tokens[len(lexer.Tokens)-1:][0]
		lexer.Tokens = lexer.Tokens[:len(lexer.Tokens)-1]
		return r
	}
	offset := lexer.Pos()
	var value *string = nil
	char := lexer.Read()
	if char==EOF{
		position := &token.Position{Len: 0, Offset: offset}
		v:=string(char)
		value = &v
		return token.New(token.EOF, value, position, true)
	}else if lexer.mode == MODE_BODY {
		switch {
		case char == GREATER:
			char = lexer.Read()
			switch char {
			case BANG:
				char = lexer.Read()
				char = lexer.Read()
				var ii = 0
				var v = ""
				for {
					v = v + string(char)
					char = lexer.Read()
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
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				return token.New(token.COMMENT, value, position, true)

			default:
				if char!=BACKSLASH{
					lexer.CHAR = int(char)
				}
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				lexer.mode = MODE_TAG;
				return token.New(token.TAG_OPEN, value, position, true)
			}
		case char == EOF:
			position := &token.Position{Len: 0, Offset: offset}
			return token.New(token.EOF, value, position, true)
		default:
			var v =  string(char)
			char =lexer.Read();
			for char != GREATER && char != EOF {
				v = v + string(char)
				char = lexer.Read()
			}
			lexer.CHAR = int(char)
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			value = &v
			return token.New(token.TEXT, value, position, true)
		}
	} else if lexer.mode == MODE_TAG {
		switch {
		case char == MINOR:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			lexer.mode = MODE_BODY
			return token.New(token.TAG_CLOSE, value, position, true)
		case char == BACKSLASH:
			char = lexer.Read()
			if char == MINOR {
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				return token.New(token.TAG_END_CLOSE, value, position, true)
			}else {
				lexer.CHAR = int(char)
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				return token.New(token.TAG_END_CLOSE, value, position, false)
			}
		case char == MULTIPLICATION:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.TAG_MULTIPLICATION, value, position, true)
		case char == BRACKETOPEN:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.TAG_BRACKETOPEN, value, position, true)
		case char == BRACKETClOSE:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.TAG_BRACKETCLOSE, value, position, true)
		case char == PARENTHESISOPEN:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.TAG_PARENTHESISOPEN, value, position, true)
		case char == PARENTHESISCLOSE:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.TAG_PARENTHESISCLOSE, value, position, true)
		case char == EQUAL:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.TAG_EQUAL, value, position, true)
		case char == DOUBLEQUOTES:
			v := ""
			char = lexer.Read()
			for char != DOUBLEQUOTES {
				if char == EOF{
					lexer.CHAR = int(char)
					len := lexer.Pos() - offset
					position := &token.Position{Len: len, Offset: offset}
					value = &v
					return token.New(token.TAG_STRING, value, position, false)
				}
				v = v + string(char)
				char = lexer.Read()

			}
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			value = &v
			return token.New(token.TAG_STRING, value, position, true)
		case IsWhiteSpace(char):
			for IsWhiteSpace(char) {
				char = lexer.Read()
			}
			lexer.CHAR = int(char)
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.TAG_WHITESPACE, value, position, true)
		case char == DOT:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.TAG_DOT, value, position, true)
		case char == DOUBLEDOT:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.TAG_DOUBLEDOT, value, position, true)
		case IsWord(char):
			v := ""
			for IsWord(char) || IsNumeric(char) {
				v = v + string(char)
				char = lexer.Read()
			}
			lexer.CHAR = int(char)
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			value = &v
			return token.New(token.TAG_WORD, value, position, true)
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
			return token.New(token.TAG_NUMERIC, value, position, true)

		default:
		}
	}
	len := lexer.Pos() - offset
	position := &token.Position{Len: len, Offset: offset}
	v:=string(char)
	value = &v
	return token.New(token.INVALID, value, position, false)
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
