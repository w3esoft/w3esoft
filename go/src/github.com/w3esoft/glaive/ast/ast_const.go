package ast

const (
	IMPORT_OBJ                   = 0x1
	IMPORT_ANY                   = 0x2
	EXPORT_OBJ                   = 0x3
	EXPORT_ANY                   = 0x4
	MODULE                       = 0x5
)

var NAMES = map[int]string{
	IMPORT_OBJ:"IMPORT_OBJ",
	IMPORT_ANY:"IMPORT_ANY",
	EXPORT_OBJ:"EXPORT_OBJ",
	EXPORT_ANY:"EXPORT_ANY",
	MODULE:"MODULE",
}