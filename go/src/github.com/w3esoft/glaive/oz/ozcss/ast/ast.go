package ast

import (
	"errors"
	"github.com/w3esoft/oz/ozcss/token"
)

type AstNode struct {
	Index    int
	Value    interface{}
	Name     string
	Position *Position
	Valid    bool
	Id       int
}

type AstTag struct {
	Name  *AstNode
	Attrs []*AstNode
	Items []*AstNode
	NoItems bool
}
type AstDoctype struct {
	Values  []*AstNode
}


type AstTagAttr struct {
	Key   *AstNode
	Value *AstNode
}
type AstTagAttrKey struct {
	Value string
}
type AstTagAttrValue struct {
	Value string
}

type AstDOCUMENT struct {
	Items []*AstNode
}
type AstCOMMENT struct {
	Value *string
}
type AstTEXT struct {
	Value *string
}

type AstTagName struct {
	Value *string
}
type AstUnexpected struct {
	items []*token.Token
}

type Position struct {
	Len    int
	Offset int
}
type Ast interface {
}
var Id = 0

func New(Index int, value interface{}, position *Position,Valid bool) *AstNode {
	var ast = AstNode{}
	ast.Index = Index
	ast.Value = value
	ast.Id = Id
	ast.Valid = Valid
	ast.Name = NAMES[Index]
	Id++
	ast.Position = position
	return &ast
}

func NewAstDoctypeValueKey(Value *string, position *Position,Valid bool) *AstNode {
	return New(DOCTYPE_VALUE_KEY, &AstTagAttrValue{Value:*Value},  position,Valid)
}

func NewAstDoctype( Values []*AstNode,position *Position,Valid bool) *AstNode {
	return New(DOCTYPE, &AstDoctype{Values:Values}, position, Valid)
}

func NewAstDoctypeValueNumeric(Value *string, position *Position,Valid bool) *AstNode {
	return New(DOCTYPE_VALUE_NUMERIC, &AstTagAttrKey{Value:*Value},  position,Valid)
}


func NewAstDoctypeValueString(Value *string, position *Position,Valid bool) *AstNode {
	return New(DOCTYPE_VALUE_STRING, &AstTagAttrKey{Value:*Value},  position,Valid)
}


func NewAstDOCUMENT(Items []*AstNode, position *Position,Valid bool) *AstNode {
	return New(DOCUMENT, &AstDOCUMENT{Items:Items},  position,Valid)
}
func NewAstCOMMENT(tk *token.Token,Valid bool) *AstNode {
	return New(COMMENT, &AstCOMMENT{Value:tk.Value}, &Position{Len: tk.Position.Len, Offset: tk.Position.Offset},Valid)
}
func NewAstTEXT(tk *token.Token,Valid bool) *AstNode {
	return New(TEXT, &AstTEXT{Value:tk.Value},  &Position{Len: tk.Position.Len, Offset: tk.Position.Offset},Valid)
}
func  NewAstTagName(tk *token.Token,Valid bool) *AstNode {
	return New(TAG_NAME, &AstTagName{Value:tk.Value},  &Position{Len: tk.Position.Len, Offset: tk.Position.Offset},Valid)
}

func NewAstTagAttr(Key *AstNode, Value  *AstNode, position *Position,Valid bool) *AstNode {
	return New(TAG_ATTR, &AstTagAttr{Key:Key,Value:Value},  position,Valid)
}
func NewAstTag(Name *AstNode,Items []*AstNode, Attrs []*AstNode,NoItems bool, position *Position,Valid bool) *AstNode {
	return New(TAG, &AstTag{Name:Name,Items:Items,Attrs:Attrs,NoItems:NoItems}, position, Valid)
}



func NewAstTagAttrKeyInput(Value *string, position *Position,Valid bool) *AstNode {
	return New(TAG_ATTR_KEY_INPUT, &AstTagAttrKey{Value:*Value},  position,Valid)
}
func NewAstTagAttrKeyOutput(Value *string, position *Position,Valid bool) *AstNode {
	return New(TAG_ATTR_KEY_OUTPUT, &AstTagAttrKey{Value:*Value},  position,Valid)
}
func NewAstTagAttrKeyNormal(Value *string, position *Position,Valid bool) *AstNode {
	return New(TAG_ATTR_KEY_NORMAL, &AstTagAttrKey{Value:*Value},  position,Valid)
}
func NewAstTagAttrValueString(Value *string, position *Position,Valid bool) *AstNode {
	return New(TAG_ATTR_VALUE_STRING, &AstTagAttrValue{Value:*Value},  position,Valid)
}
func NewAstTagAttrValueNumeric(Value *string, position *Position,Valid bool) *AstNode {
	return New(TAG_ATTR_VALUE_NUMERIC, &AstTagAttrValue{Value:*Value},  position,Valid)
}

func NewAstTagAttrKeyTemplate(Value *string, position *Position,Valid bool) *AstNode {
	return New(TAG_ATTR_KEY_TEMPLATE, &AstTagAttrKey{Value:*Value},  position,Valid)
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
