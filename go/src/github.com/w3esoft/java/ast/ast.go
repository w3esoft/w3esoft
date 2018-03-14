package ast

import (
	"errors"
)

type AstModule struct {
	Items     []*AstNode
}
type AstValue struct {
	Value    *AstNode
}
type AstIdent struct {
	Value    * string
}



type AstLet struct {
	Assign    *AstNode
}
type AstMeta struct {
	Field    *AstNode
	Ident    *AstNode
}
type AstField struct {
	Field    *AstNode
	Ident    *AstNode
}



type AstExportValue struct {
	Value    *AstNode
}
type AstExportLet struct {
	Value      *AstNode
}

type AstExportDestructuringAssignmentField struct {
	Imports     *AstNode
	From        *string
}

type AstImportDestructuringAssignment struct {
	Imports     *AstNode
	From        *string
}
type AstExportDestructuringAssignment struct {
	Imports     *AstNode
	From        *string
}
type AstDestructuringAssignmentObjectMatchingField struct {
	Next      *AstNode
	Value     *string
}
type AstDestructuringAssignmentObjectMatching struct {
	Fields      []*AstNode
}
type AstObjectItem struct {
	Value      *AstNode
	Key      *AstNode
}


type AstObject struct {
	Values    []*AstNode
}






type AstImportAny struct {
	Alias       *string
	From        *string
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
func NewAstObjectItem(Key *AstNode,Value *AstNode ,Position *Position) *AstNode {
	return New(OBJECT_ITEM, &AstObjectItem{Key :Key ,Value:Value},Position)
}
func NewAstArray(Values    []*AstNode ,Position *Position) *AstNode {
	return New(OBJECT, &AstObject{Values:Values},Position)
}
func NewAstObject(Values    []*AstNode ,Position *Position) *AstNode {
	return New(OBJECT, &AstObject{Values:Values},Position)
}
func NewAstIdent(Value *string ,Position *Position) *AstNode {
	return New(IDENT, &AstIdent{Value:Value},Position)
}
func NewAstField(Ident * AstNode ,Field *AstNode ,Position *Position) *AstNode {
	return New(FIELD, &AstField{Ident:Ident,Field:Field},Position)
}
func NewAstMeta(Ident *AstNode ,Field *AstNode ,Position *Position) *AstNode {
	return New(META, &AstMeta{Ident:Ident, Field:Field},Position)
}
func NewAstCall(Ident *AstNode ,Parameters []*AstNode ,Position *Position) *AstNode {
	return New(CALL, &AstCall{Ident:Ident, Parameters:Parameters},Position)
}
func NewAstModule (Items []*AstNode ,Position *Position) *AstNode {
	return New(MODULE, &AstModule{Items:Items},Position)
}
func NewAstLet(Assign *AstNode,Position *Position) *AstNode {
	return New(IMPORT_DESTRUCTURING_ASSIGNMENT, &AstLet{Assign:Assign},Position)
}
func NewAstImportDestructuringAssignment(Imports *AstNode,From *string ,Position *Position) *AstNode {
	return New(IMPORT_DESTRUCTURING_ASSIGNMENT, &AstImportDestructuringAssignment{Imports:Imports,From:From},Position)
}
func NewAstImportAny(Alias       *string,From *string ,Position *Position) *AstNode {
	return New(IMPORT_ANY, &AstImportAny{Alias:Alias,From:From},Position)
}
func NewAstExportDestructuringAssignment(Imports *AstNode,From *string ,Position *Position) *AstNode {
	return New(IMPORT_DESTRUCTURING_ASSIGNMENT, &AstExportDestructuringAssignment{Imports:Imports,From:From},Position)
}
func NewAstExportDestructuringAssignmentField(Imports *AstNode,From *string ,Position *Position) *AstNode {
	return New(IMPORT_DESTRUCTURING_ASSIGNMENT_FIELD, &AstExportDestructuringAssignmentField{Imports:Imports,From:From},Position)
}
func NewAstDestructuringAssignmentObjectMatchingField(Next *AstNode,Value     *string,Position *Position) *AstNode {
	return New(
		DESTRUCTURING_ASSIGNMENT_OBJECT_MATCHING_FIELD,
		&AstDestructuringAssignmentObjectMatchingField{Next:Next,Value:Value},
		Position,
	)
}
func NewAstDestructuringAssignmentObjectMatching(Fields     []*AstNode,Position *Position) *AstNode {
	return New(
		DESTRUCTURING_ASSIGNMENT_OBJECT_MATCHING_FIELD,
		&AstDestructuringAssignmentObjectMatching{Fields:Fields},
		Position,
	)
}

func NewAstExportAny(From *string ,Position *Position) *AstNode {
	return New(IMPORT_ANY, &AstImportAny{From:From},Position)
}


func NewAstExportValue(Value *AstNode ,Position *Position) *AstNode {
	return New(IMPORT_ANY, &AstExportValue{Value:Value},Position)
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
