package token

const (
	INVALID             = 0
	EOF                 =1
	WHITESPACE 			=2
	NUMERIC				=3
	DOT					=4
)

var NAMES = map[int]string{
	INVALID:"INVALID",
	EOF:"EOF",
	WHITESPACE:"WHITESPACE",
	NUMERIC:"NUMERIC",
	DOT:"DOT",
}
