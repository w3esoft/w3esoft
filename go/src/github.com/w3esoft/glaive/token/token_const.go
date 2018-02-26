package token

const (
	INVALID          = 0
	EOF              = 1
	WHITESPACE       = 2
	NUMERIC          = 3
	DOT              = 4
	AMPERSAT         = 5
	DOUBLEDOT        = 6
	BRACKETOPEN      = 7
	BRACKETClOSE     = 8
	PARENTHESISOPEN  = 9
	PARENTHESISCLOSE = 10
	BRACEOPEN        = 11
	BRACECLOSE       = 12
	MINUS            = 13
	STRING           = 14
	SEMICOLON        = 15
	FORWARDSLASH     = 16
	BACKSLASH        = 17
	COMMA            = 18
	PLUS             = 19
	MULTIPLICATION   = 20
	GREATER          = 21
	MINOR            = 22
	EQUAL            = 23
	BANG             = 24
	AMPERSAND             = 25
	WORD             = 26
)

var NAMES = map[int]string{
	INVALID:    "INVALID",
	EOF:        "EOF",
	WHITESPACE: "WHITESPACE",
	NUMERIC:    "NUMERIC",
	DOT:        "DOT",
}
