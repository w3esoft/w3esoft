package token

const (
	INVALID           = 0
	EOF               = 1
	COMMENT           = 2
	TAG_OPEN          = 3
	TEXT              = 4
	DOCTYPE_OPEN      = 5
	TAG_CLOSE         = 100
	KEY               = 110
	TAG_KEY_TEMPLATE  = 102
	TAG_KEY_OUTPUT    = 105
	TAG_KEY_INPUT     = 103
	TAG_VALUE_STRING  = 113
	TAG_VALUE_NUMERIC = 111
	TAG_DOUBLEDOT     = 107
	TAG_DOT           = 108
	TAG_WHITESPACE    = 109
	TAG_EQUAL         = 112
	TAG_END_OPEN      = 114
	TAG_END_CLOSE     = 115
	TAG_END_CLOSE_NOCHILD= 116
)

var NAMES = map[int]string{
	INVALID:           "INVALID",
	EOF:               "EOF",
	COMMENT:           "COMMENT",
	TAG_OPEN:          "TAG_OPEN",
	TEXT:              "TEXT",
	TAG_CLOSE:         "TAG_CLOSE",
	KEY:               "KEY",
	TAG_KEY_TEMPLATE:  "TAG_KEY_TEMPLATE",
	TAG_KEY_INPUT:     "TAG_KEY_INPUT",
	TAG_KEY_OUTPUT:    "TAG_KEY_OUTPUT",
	TAG_DOUBLEDOT:     "TAG_DOUBLEDOT",
	TAG_DOT:           "TAG_DOT",
	TAG_WHITESPACE:    "TAG_WHITESPACE",
	TAG_VALUE_STRING:  "TAG_VALUE_STRING",
	TAG_EQUAL:         "TAG_EQUAL",
	TAG_VALUE_NUMERIC: "TAG_VALUE_NUMERIC",
	TAG_END_OPEN:      "TAG_END_OPEN",
	TAG_END_CLOSE:     "TAG_END_CLOSE",
	DOCTYPE_OPEN:      "DOCTYPE_OPEN",
}
