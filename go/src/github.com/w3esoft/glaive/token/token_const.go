package token

const (
	INVALID               = 0x01
	EOF                   = 0x02
	WHITESPACE            = 0x03
	NUMERIC               = 0x04
	DOT                   = 0x05
	AMPERSAT              = 0x06
	DOUBLEDOT             = 0x07
	BRACKETOPEN           = 0x08
	BRACKETClOSE          = 0x09
	PARENTHESISOPEN       = 0x0a
	PARENTHESISCLOSE      = 0x0b
	BRACEOPEN             = 0x0c
	BRACECLOSE            = 0x0d
	MINUS                 = 0x0f
	STRING                = 0x1110
	STRING_TEMPLATE       = 0x1120
	SEMICOLON             = 0x12
	BACKSLASH             = 0x14
	COMMA                 = 0x15
	PLUS                  = 0x16
	MULTIPLICATION        = 0x17
	GREATER               = 0x18
	MINOR                 = 0x19
	EQUAL                 = 0x1a
	BANG                  = 0x1b
	AMPERSAND             = 0x1c
	WORD                  = 0x1d
	LINE                  = 0x1f
	OPERATOR_CONDITION_OR = 0x2111
	OPERATOR_DIVISION	  = 0x2121
	OPERATOR_MOD     	  = 0x2131
	COMMENT_LINE          = 0x22
	COMMENT_BOCK          = 0x23
	QUESTION_MARK         = 0x24
	ASTERISK         	  = 0x25
	REGEXP         	      = 0x26
	PIPE         	      = 0x27
	IMPORT         	      = 0x28
	EXPORT         	      = 0x29
	FROM         	      = 0x2a
	LET         	      = 0x2b
)

var NAMES = map[int]string{
	INVALID:               "INVALID",
	EOF:                   "EOF",
	WHITESPACE:            "WHITESPACE",
	NUMERIC:               "NUMERIC",
	DOT:                   "DOT",
	AMPERSAT:              "AMPERSAT",
	DOUBLEDOT:             "DOUBLEDOT",
	BRACKETOPEN:           "BRACKETOPEN",
	BRACKETClOSE:          "BRACKETClOSE",
	PARENTHESISOPEN:       "PARENTHESISOPEN",
	PARENTHESISCLOSE:      "PARENTHESISCLOSE",
	BRACEOPEN:             "BRACEOPEN",
	BRACECLOSE:            "BRACECLOSE",
	MINUS:                 "MINUS",
	STRING:                "STRING",
	SEMICOLON:             "SEMICOLON",
	BACKSLASH:             "BACKSLASH",
	COMMA:                 "COMMA",
	PLUS:                  "PLUS",
	MULTIPLICATION:        "MULTIPLICATION",
	GREATER:               "GREATER",
	MINOR:                 "MINOR",
	EQUAL:                 "EQUAL",
	BANG:                  "BANG",
	AMPERSAND:             "AMPERSAND",
	WORD:                  "WORD",
	LINE:                  "LINE",
	OPERATOR_CONDITION_OR: "OPERATOR_CONDITION_OR",
	COMMENT_LINE:          "COMMENT_LINE",
	COMMENT_BOCK:          "COMMENT_BOCK",
	QUESTION_MARK:         "QUESTION_MARK",
	ASTERISK:              "ASTERISK",
	OPERATOR_DIVISION:     "OPERATOR_DIVISION",
	REGEXP:     		   "REGEXP",
	STRING_TEMPLATE:       "STRING_TEMPLATE",
	IMPORT:       			"IMPORT",
	EXPORT:       			"EXPORT",
	FROM:       			"FROM",
	LET:       			    "LET",
}
