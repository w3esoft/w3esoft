package ast

const (
	IMPORT_DESTRUCTURING_ASSIGNMENT                = 0x11
	IMPORT_DESTRUCTURING_ASSIGNMENT_FIELD          = 0x12
	IMPORT_ANY                                     = 0x13
	EXPORT_DESTRUCTURING_ASSIGNMENT                = 0x21
	EXPORT_ANY                                     = 0x22
	MODULE                                         = 0x3
	DESTRUCTURING_ASSIGNMENT_OBJECT_MATCHING_FIELD = 0x4
	META                                           = 0x5
	FIELD                                          = 0x6
	IDENT                                          = 0x8
	OBJECT_ITEM                                    = 0x9
	OBJECT                                         = 0xa
	ARRAY                                          = 0xb
	CALL                                           = 0xc
)

var NAMES = map[int]string{
	IMPORT_DESTRUCTURING_ASSIGNMENT:                "IMPORT_DESTRUCTURING_ASSIGNMENT",
	IMPORT_DESTRUCTURING_ASSIGNMENT_FIELD:          "IMPORT_DESTRUCTURING_ASSIGNMENT_FIELD",
	IMPORT_ANY:                                     "IMPORT_ANY",
	EXPORT_DESTRUCTURING_ASSIGNMENT:                "EXPORT_DESTRUCTURING_ASSIGNMENT",
	EXPORT_ANY:                                     "EXPORT_ANY",
	MODULE:                                         "MODULE",
	DESTRUCTURING_ASSIGNMENT_OBJECT_MATCHING_FIELD: "DESTRUCTURING_ASSIGNMENT_OBJECT_MATCHING_FIELD",
	META:                                           "META",
	FIELD:                                          "FIELD",
	IDENT:                                          "IDENT",
	OBJECT:                                         "OBJECT",
	OBJECT_ITEM:                                    "OBJECT_ITEM",
	ARRAY:                                          "ARRAY",
	CALL:                                           "CALL",
}
