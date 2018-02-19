package ast

const (
	TAG            = 1
	DOCUMENT       = 2
	COMMENT        = 3
	TAG_ATTR       = 4
	TAG_ATTR_KEY   = 5
	TAG_ATTR_VALUE = 6
)

var NAMES = map[int]string{
	TAG:            "TAG",
	DOCUMENT:       "DOCUMENT",
	COMMENT:        "COMMENT",
	TAG_ATTR:       "TAG_ATTR",
	TAG_ATTR_KEY:   "TAG_ATTR_KEY",
	TAG_ATTR_VALUE: "TAG_ATTR_VALUE",
}
