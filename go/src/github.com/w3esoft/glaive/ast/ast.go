package ast

import (
	"errors"
)

type AstModule struct {
	Items     []*AstNode
}
type AstImportObj struct {
	Imports     []*string
	From        *string
	Position *Position
}
type AstImportAny struct {
	Alias       *string
	From        *string
	Position *Position
}
type AstNode struct {
	Index    int
	Value    interface{}
	Name     string
	Position *Position
	Valid    bool
	Id       int
}
type Position struct {
	Len    int
	Offset int
}
type Ast interface {
}
var Id = 0

func New(Index int, value interface{}, position *Position) *AstNode {
	var ast = AstNode{}
	ast.Index = Index
	ast.Value = value
	ast.Id = Id
	ast.Name = NAMES[Index]
	Id++
	ast.Position = position
	return &ast
}

func NewAstModule (Items []*AstNode ,Position *Position) *AstNode {
	return New(MODULE, &AstModule{Items:Items},Position)
}
func NewAstImportObj(Imports []*string,From *string ,Position *Position) *AstNode {
	return New(IMPORT_OBJ, &AstImportObj{Imports:Imports,From:From},Position)
}
func NewAstImportAny(Alias       *string,From *string ,Position *Position) *AstNode {
	return New(IMPORT_ANY, &AstImportAny{Alias:Alias,From:From},Position)
}
func NewAstExportObj(Imports []*string,From *string ,Position *Position) *AstNode {
	return New(IMPORT_OBJ, &AstImportObj{Imports:Imports,From:From},Position)
}
func NewAstExportAny(From *string ,Position *Position) *AstNode {
	return New(IMPORT_ANY, &AstImportAny{From:From},Position)
}
func (ast *AstNode) Is(indexs []int, value *Ast) bool {
	for _, index := range indexs {
		b1 := ast.Index==index
		b2 := true;
		if value != nil {
			b2 = ast.Value == value
		}
		if b1 && b2 {
			return true
		}
	}
	return false
}
func (ast *AstNode) IsNot(astIndexs []int, value *Ast) bool {
	return !ast.Is(astIndexs, value)
}
func (ast *AstNode) Expected(astIndexs []int, value *Ast) (r bool, err error) {
	b := ast.Is(astIndexs, value)
	if !b {
		return b, errors.New("unexpected ast " + ast.Name)
	}
	return b, nil
}
func (ast *AstNode) Unexpected(astIndexs []int, value *Ast) (r bool, err error) {
	b := ast.IsNot(astIndexs, value)
	if !b {
		return b, errors.New("unexpected ast " + ast.Name)
	}
	return b, nil
}
