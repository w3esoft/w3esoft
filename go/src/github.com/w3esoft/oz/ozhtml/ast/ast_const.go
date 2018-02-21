package ast

const (
	TEXT        					= 1
	TAG           					= 2
	DOCUMENT       					= 3
	COMMENT        					= 4
	TAG_ATTR       					= 5
	TAG_ATTR_VALUE_STRING           = 511
	TAG_ATTR_VALUE_NUMERIC          = 512
	TAG_ATTR_KEY_NORMAL             = 521
	TAG_ATTR_KEY_INPUT              = 522
	TAG_ATTR_KEY_TEMPLATE           = 523
	TAG_ATTR_KEY_OUTPUT             = 524
	TAG_ATTR_VALUE 					= 7
	TAG_NAME						= 6
)

var NAMES = map[int]string{
	TEXT:						 "TEXT",
	TAG:          				 "TAG",
	DOCUMENT:     				 "DOCUMENT",
	TAG_NAME:					 "TAG_NAME",
	COMMENT:      				 "COMMENT",
	TAG_ATTR:     				 "TAG_ATTR",
	TAG_ATTR_KEY_INPUT:          "TAG_ATTR_KEY_INPUT",
	TAG_ATTR_KEY_OUTPUT:       	 "TAG_ATTR_KEY_OUTPUT",
	TAG_ATTR_KEY_NORMAL:         "TAG_ATTR_KEY_NORMAL",
	TAG_ATTR_KEY_TEMPLATE:       "TAG_ATTR_KEY_TEMPLATE",
	TAG_ATTR_VALUE_STRING:       "TAG_ATTR_VALUE_STRING",
	TAG_ATTR_VALUE_NUMERIC:      "TAG_ATTR_VALUE_NUMERIC",
	TAG_ATTR_VALUE:              "TAG_ATTR_VALUE",
}
