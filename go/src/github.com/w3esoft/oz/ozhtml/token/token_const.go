package token

const (
	EOF              = 1
	TAGHEADOPEN      = 2
	TAGHEADCLOSE     = 4
	TAGBODY          = 5
	STRING           = 9
	MULTIPLICATION   = 10
	PARENTHESISOPEN  = 11
	PARENTHESISCLOSE = 12
	BRACKETOPEN      = 13
	BRACKETCLOSE     = 14
	DOUBLEDOT        = 16
	DOT              = 17
	TAGCOMMENT       = 18
	NUMERIC          = 19
	WORD             = 6
	WHITESPACE       = 7
	EQUAL            = 8
)

var NAMES = map[int]string{
	EOF              : "EOF",
	TAGHEADOPEN      : "TAGHEADOPEN",
	TAGHEADCLOSE     : "TAGHEADCLOSE",
	TAGBODY          : "TAGBODY",
	STRING           : "STRING",
	MULTIPLICATION   : "MULTIPLICATION",
	PARENTHESISOPEN  : "PARENTHESISOPEN",
	PARENTHESISCLOSE : "PARENTHESISCLOSE",
	BRACKETOPEN      : "BRACKETOPEN",
	BRACKETCLOSE     : "BRACKETCLOSE",
	DOUBLEDOT        : "DOUBLEDOT",
	DOT              : "DOT",
	TAGCOMMENT       : "TAGCOMMENT",
	NUMERIC          : "NUMERIC",
	WORD             : "WORD",
	WHITESPACE       : "WHITESPACE",
	EQUAL            : "EQUAL",

}
