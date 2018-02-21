package lexer

import (
	"github.com/w3esoft/glaive/input"
	"github.com/w3esoft/oz/ozhtml/token"
	"strings"
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
const (
	MODE_TAG  = 1
	MODE_BODY = 2
)

type Lexer struct {
	CHAR   int
	Input  input.Input
	Tokens [] *token.Token
	pos    int
	mode   int
}

func New(input input.Input) *Lexer {
	lexer := Lexer{}
	lexer.CHAR = -1
	lexer.pos = 0
	lexer.mode = MODE_BODY;
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
	} else if lexer.mode == MODE_BODY {
		switch {
		case char == GREATER:
			char = lexer.Read()
			switch char {
			case BANG:
				char = lexer.Read()
				if (IsWord(char)) && char != MINUS {
					var v = ""
					for IsWord(char) {
						v = v + string(char)
						char = lexer.Read()
					}
					if strings.Compare(strings.ToLower(v), "doctype") != 0 {
						len := lexer.Pos() - offset
						position := &token.Position{Len: len, Offset: offset}
						return token.New(token.DOCTYPE_OPEN, value, position, false)
					}
					lexer.CHAR = int(char)
					len := lexer.Pos() - offset
					position := &token.Position{Len: len, Offset: offset}
					lexer.mode = MODE_TAG;
					return token.New(token.DOCTYPE_OPEN, value, position, true)
				}
				if char != MINUS {
					len := lexer.Pos() - offset
					position := &token.Position{Len: len, Offset: offset}
					return token.New(token.COMMENT, value, position, false)
				}
				char = lexer.Read()
				if char != MINUS {
					len := lexer.Pos() - offset
					position := &token.Position{Len: len, Offset: offset}
					return token.New(token.COMMENT, value, position, false)
				}

				var ii = 0
				var v = ""
				for {
					v = v + string(char)
					char = lexer.Read()
					b := false;
					switch {
					case char == EOF:
						len := lexer.Pos() - offset
						position := &token.Position{Len: len, Offset: offset}
						return token.New(token.COMMENT, value, position, false)
					case ii == 0 && char == MINUS:
						ii = 1
					case ii == 1 && char == MINUS:
						ii = 2
					case ii == 2 && char == MINOR:
						b = true;
					default:
						ii = 0
					}
					if (b) {
						break
					}

				}
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				return token.New(token.COMMENT, value, position, true)
			case BACKSLASH:
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				lexer.mode = MODE_TAG;
				return token.New(token.TAG_END_OPEN, value, position, true)
			default:
				lexer.CHAR = int(char)
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				lexer.mode = MODE_TAG;
				return token.New(token.TAG_OPEN, value, position, true)

			}
		case char == EOF:
			position := &token.Position{Len: 0, Offset: offset}
			return token.New(token.EOF, value, position, true)
		default:
			var v = string(char)
			char = lexer.Read();
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
				lexer.mode = MODE_BODY
				return token.New(token.TAG_END_CLOSE, value, position, true)
			} else {
				lexer.CHAR = int(char)
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				lexer.mode = MODE_BODY
				return token.New(token.TAG_END_CLOSE, value, position, false)
			}
		case IsWord(char) || char == UNDERSCORE || HASH == char:
			v := ""
			if !(IsWord(char) || char == UNDERSCORE || HASH == char) {
				lexer.CHAR = int(char)
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				return token.New(token.KEY, value, position, false)
			}
			v = v + string(char)
			char = lexer.Read()
			for IsWord(char) || char == MINUS || char == UNDERSCORE || char == DOT || char == DOUBLEDOT || IsNumeric(char) {
				v = v + string(char)
				char = lexer.Read()
			}
			lexer.CHAR = int(char)
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			value = &v
			return token.New(token.KEY, value, position, true)
		case char == MULTIPLICATION:
			char = lexer.Read()
			v := ""
			if !(IsWord(char) || char == UNDERSCORE || HASH == char ) {
				lexer.CHAR = int(char)
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				return token.New(token.TAG_KEY_OUTPUT, value, position, false)
			}
			v = v + string(char)
			char = lexer.Read()
			for IsWord(char) || char == MINUS || char == UNDERSCORE || char == DOT || char == DOUBLEDOT || IsNumeric(char) {
				v = v + string(char)
				char = lexer.Read()
			}
			lexer.CHAR = int(char)
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			value = &v
			return token.New(token.TAG_KEY_OUTPUT, value, position, true)
		case char == BRACKETOPEN:
			v := ""
			char = lexer.Read()
			if !(IsWord(char) || char == UNDERSCORE || HASH == char) {
				lexer.CHAR = int(char)
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				return token.New(token.TAG_KEY_INPUT, value, position, false)
			}
			v = v + string(char)
			char = lexer.Read()
			for IsWord(char) || char == MINUS || char == UNDERSCORE || char == DOT || char == DOUBLEDOT || IsNumeric(char) || IsWhiteSpace(char) {
				v = v + string(char)
				char = lexer.Read()
			}
			if char != BRACKETClOSE {
				lexer.CHAR = int(char)
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				return token.New(token.TAG_KEY_INPUT, value, position, false)
			}
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			value = &v
			return token.New(token.TAG_KEY_INPUT, value, position, true)
		case char == PARENTHESISOPEN:
			v := ""
			char = lexer.Read()
			if !(IsWord(char) || char == UNDERSCORE || HASH == char) {
				lexer.CHAR = int(char)
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				return token.New(token.TAG_KEY_OUTPUT, value, position, false)
			}
			v = v + string(char)
			char = lexer.Read()
			for IsWord(char) || char == MINUS || char == UNDERSCORE || char == DOT || char == DOUBLEDOT || IsNumeric(char) || IsWhiteSpace(char) {
				v = v + string(char)
				char = lexer.Read()
			}
			if char != PARENTHESISCLOSE {
				lexer.CHAR = int(char)
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				return token.New(token.TAG_KEY_OUTPUT, value, position, false)
			}
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			value = &v
			return token.New(token.TAG_KEY_OUTPUT, value, position, true)
		case char == EQUAL:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.TAG_EQUAL, value, position, true)
		case char == DOUBLEQUOTES:
			v := ""
			char = lexer.Read()
			for char != DOUBLEQUOTES {
				if char == EOF {
					lexer.CHAR = int(char)
					len := lexer.Pos() - offset
					position := &token.Position{Len: len, Offset: offset}
					value = &v
					return token.New(token.TAG_VALUE_STRING, value, position, false)
				}
				v = v + string(char)
				char = lexer.Read()

			}
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			value = &v
			return token.New(token.TAG_VALUE_STRING, value, position, true)
		case IsWhiteSpace(char):
			for IsWhiteSpace(char) {
				char = lexer.Read()
			}
			lexer.CHAR = int(char)
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.TAG_WHITESPACE, value, position, true)
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
			return token.New(token.TAG_VALUE_NUMERIC, value, position, true)

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
