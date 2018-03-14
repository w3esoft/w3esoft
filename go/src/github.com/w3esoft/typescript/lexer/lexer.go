package lexer

import (
	"github.com/w3esoft/input"
	"github.com/w3esoft/typescript/token"
)

const (
	ASTERISK         = 42
	DOUBLEDOT        = 58
	GREATER          = 60
	MINOR            = 62
	EQUAL            = 61
	DOUBLEQUOTES     = 34
	QUOTES           = 39
	MULTIPLICATION   = 42
	BRACEOPEN        = 123
	BRACEClOSE       = 125
	BRACKETOPEN      = 91
	BRACKETClOSE     = 93
	COMMA            = 44
	PARENTHESISOPEN  = 40
	PARENTHESISCLOSE = 41
	DOT              = 46
	SEMICOLON        = 59
	BANG             = 33
	MINUS            = 45
	UNDERSCORE       = 95
	HASH             = 35
	DOLLAR           = 36
	AMPERSAT         = 64
	EOF              = 0
	FORWARDSLASH     = 47
	BACKSLASH        = 92
	PLUS             = 43
	PERCENT          = 37
	AMPERSAND        = 38
	PIPE             = 124
	QUESTION_MARK    = 63
	ACUTE			 =96
)
const (
	FSM_OPERATION = 1
	FSM_DECLARE   = 2
)

type Lexer struct {
	CHAR             int
	Input            input.Input
	Tokens           [] *token.Token
	pos              int
	forwardSlashMode int
}

func New(input input.Input) *Lexer {
	lexer := Lexer{}
	lexer.CHAR = -1
	lexer.pos = 0
	lexer.Input = input
	lexer.forwardSlashMode = FSM_DECLARE
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
	tk := lexer.tokenize();
	if (tk != nil) {
		if (tk.Is([]int{
			token.PARENTHESISOPEN,
			token.EQUAL,
			token.BRACKETOPEN,
			token.LINE,
			token.BRACEOPEN,
		}, nil, true)) {
			lexer.forwardSlashMode = FSM_DECLARE
		} else if (tk.Is([]int{
			token.WHITESPACE,
		}, nil, true)) {
			lexer.forwardSlashMode = lexer.forwardSlashMode
		}else {
			lexer.forwardSlashMode = FSM_OPERATION
		}

	}

	return tk;
}

func (lexer *Lexer) tokenize() *token.Token {
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
	} else {
		switch {
		case IsLine(char):
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.LINE, value, position, true)
		case IsWhiteSpace(char) || char == HASH:
			for IsWhiteSpace(char) || char == HASH {
				char = lexer.Read()
			}
			lexer.CHAR = int(char)
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.WHITESPACE, value, position, true)
		case IsWord(char) || char == DOLLAR || char == UNDERSCORE:
			v := ""
			for IsWord(char) || IsNumeric(char) || char == DOLLAR || char == UNDERSCORE {
				v = v + string(char)
				char = lexer.Read()
			}
			lexer.CHAR = int(char)
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			value = &v
			if v=="import"{

				return token.New(token.IMPORT, value, position, true)
			}
			if v=="export"{

				return token.New(token.EXPORT, value, position, true)
			}
			if v=="from"{

				return token.New(token.FROM, value, position, true)
			}
			if v=="let"{

				return token.New(token.LET, value, position, true)
			}
			if v=="class"{

				return token.New(token.CLASS, value, position, true)
			}
			if v=="const"{

				return token.New(token.CONST, value, position, true)
			}



			return token.New(token.WORD, value, position, true)
		case IsNumeric(char):
			v := ""
			b:=true;
			for IsNumeric(char) || (DOT == char&&b) {
				if DOT == char&& b{
					b =false
				}
				v = v + string(char)
				char = lexer.Read()
			}
			lexer.CHAR = int(char)
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			value = &v
			return token.New(token.NUMERIC, value, position, true)
		case char == ASTERISK:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.ASTERISK, value, position, true)
		case char == DOT:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.DOT, value, position, true)
		case char == AMPERSAT:
			char = lexer.Read()
			t:=token.META;
			if (char==DOUBLEDOT){
				t=token.MACRO_AST;
				char = lexer.Read()
			}
			v := ""
			for IsWord(char) || IsNumeric(char) || char == DOLLAR || char == UNDERSCORE {
				v = v + string(char)
				char = lexer.Read()
			}
			lexer.CHAR = int(char)
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			value = &v
			return token.New(t, value, position, true)
		case char == PERCENT:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.OPERATOR_MOD, value, position, true)
		case char == DOUBLEDOT:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.DOUBLEDOT, value, position, true)
		case char == DOUBLEDOT:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.DOUBLEDOT, value, position, true)
		case char == BRACKETOPEN:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.BRACKETOPEN, value, position, true)
		case char == BRACKETClOSE:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.BRACKETClOSE, value, position, true)
		case char == QUESTION_MARK:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.QUESTION_MARK, value, position, true)

		case char == BACKSLASH:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.BACKSLASH, value, position, true)
		case char == FORWARDSLASH:
			char = lexer.Read();
			if char == FORWARDSLASH {
				char = lexer.Read()
				var v = "";
				for char != EOF && !IsLine(char) {
					v = v + string(char)
					char = lexer.Read()
				}
				lexer.CHAR = int(char)
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				return token.New(token.COMMENT_LINE, value, position, true)
			}else if char == ASTERISK {
				char = lexer.Read();
				var uu = char;
				var v = "";
				v = v + string(char)
				char = lexer.Read();
				for !(uu == ASTERISK && char == FORWARDSLASH ) {
					v = v + string(char)
					uu =char
					char = lexer.Read()
				}

				v = v[0:len(v)-2]
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				return token.New(token.COMMENT_BOCK, value, position, true)
			}else if  lexer.forwardSlashMode == FSM_DECLARE{
				var v = "";
				v = v + string(char)
				char = lexer.Read();
				for  {
					if char==BACKSLASH {
						char = lexer.Read()
						v = v + string(char)
						char = lexer.Read()
					}
					if char == FORWARDSLASH{
						break
					}
					v = v + string(char)
					char = lexer.Read()
				}
				char = lexer.Read();
				var v1 = "";
				for(IsWord(char)){

					v1 = v1 + string(char)
					char = lexer.Read()
				}
				lexer.CHAR = int(char)
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				return token.New(token.REGEXP, value, position, true)
			}else {
				lexer.CHAR = int(char)
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				return token.New(token.OPERATOR_DIVISION, value, position, true)

			}

		case char == BRACEOPEN:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.BRACEOPEN, value, position, true)
		case char == BRACEClOSE:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.BRACECLOSE, value, position, true)
		case char == COMMA:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.COMMA, value, position, true)
		case char == PARENTHESISOPEN:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.PARENTHESISOPEN, value, position, true)
		case char == PARENTHESISCLOSE:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.PARENTHESISCLOSE, value, position, true)
		case char == PARENTHESISCLOSE:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.PARENTHESISCLOSE, value, position, true)
		case char == SEMICOLON:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.SEMICOLON, value, position, true)
		case char == PLUS:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.PLUS, value, position, true)
		case char == DOUBLEQUOTES:
			v := ""
			char = lexer.Read()
			for  {
				if char==BACKSLASH {
					char = lexer.Read()
					v = v + string(char)
					char = lexer.Read()
				}
				if char == DOUBLEQUOTES{
					break;
				}
				if char == EOF {
					lexer.CHAR = int(char)
					len := lexer.Pos() - offset
					position := &token.Position{Len: len, Offset: offset}
					value = &v
					return token.New(token.STRING, value, position, false)
				}
				v = v + string(char)
				char = lexer.Read()

			}
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			value = &v
			return token.New(token.STRING, value, position, true)
		case char == QUOTES:
			v := ""
			char = lexer.Read()
			for  {
				if char==BACKSLASH {
					char = lexer.Read()
					v = v + string(char)
					char = lexer.Read()
				}
				if char == QUOTES{
					break;
				}
				if char == EOF {
					lexer.CHAR = int(char)
					len := lexer.Pos() - offset
					position := &token.Position{Len: len, Offset: offset}
					value = &v
					return token.New(token.STRING, value, position, false)
				}
				v = v + string(char)
				char = lexer.Read()

			}
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			value = &v
			return token.New(token.STRING, value, position, true)
		case char == ACUTE:
			v := ""
			char = lexer.Read()
			for  {
				if char==BACKSLASH {
					char = lexer.Read()
					v = v + string(char)
					char = lexer.Read()
				}
				if char == ACUTE{
					break;
				}
				if char == EOF {
					lexer.CHAR = int(char)
					len := lexer.Pos() - offset
					position := &token.Position{Len: len, Offset: offset}
					value = &v
					return token.New(token.STRING_TEMPLATE, value, position, false)
				}
				v = v + string(char)
				char = lexer.Read()

			}
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			value = &v
			return token.New(token.STRING_TEMPLATE, value, position, true)
		case char == MINUS:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.MINUS, value, position, true)
		case char == MULTIPLICATION:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.MULTIPLICATION, value, position, true)
		case char == MINUS:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.MINUS, value, position, true)
		case char == GREATER:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.GREATER, value, position, true)
		case char == MINOR:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.MINOR, value, position, true)
		case char == EQUAL:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.EQUAL, value, position, true)
		case char == BANG:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.BANG, value, position, true)
		case char == PIPE:
			char = lexer.Read()
			if (char == PIPE) {
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				return token.New(token.OPERATOR_CONDITION_OR, value, position, true)

			} else {
				lexer.CHAR = int(char)
				len := lexer.Pos() - offset
				position := &token.Position{Len: len, Offset: offset}
				return token.New(token.PIPE, value, position, true)
			}

		case char == AMPERSAND:
			len := lexer.Pos() - offset
			position := &token.Position{Len: len, Offset: offset}
			return token.New(token.AMPERSAND, value, position, true)

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
func IsLine(char uint8) bool {
	return char == 10
}
func IsWhiteSpace(char uint8) bool {
	return char == 32 || char == 13 || char == 9
}
