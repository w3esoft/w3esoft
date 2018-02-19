package ast

import (
	"errors"
	"github.com/w3esoft/oz/ozhtml/token"
)

type AstNode struct {
	Index int
	Value interface{}
	Name  string
	Position * Position
	Valid bool
	Id int
}


const (
	AST_TAG_ATTR_OP_INPUT              = 1
	AST_TAG_ATTR_OP_TEMPLATE           = 2
	AST_TAG_ATTR_OP_OUTPUT             = 3
)
type AstTag struct {
	Attrs[]*AstNode
	Items[]*AstNode
}
type AstTagAttr struct {
	Key*AstNode
	Value*AstNode
}
type AstTagAttrKey struct {
	value*string
}
type AstTagAttrValue struct {
	key*AstNode
	value*AstNode
}

type Ast interface {

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




type AstUnexpected struct {
	items []*token.Token
}



type Position struct {
	Len int
	Offset int
}
var Id =0
func New(astIndex int, value *Ast, position * Position ) *AstNode {
	var ast = AstNode{}
	ast.Index = astIndex
	ast.Value = value
	ast.Position = position
	ast.Name = NAMES[astIndex]
	return &ast
}

func NewAstDOCUMENT(value []*AstNode, position * Position ) *AstNode {
	var astDOCUMENT  = AstDOCUMENT{}
	astDOCUMENT.Items=value
	var ast Ast = astDOCUMENT
	return  New(DOCUMENT , &ast , position  )
}
func NewAstCOMMENT(tk *token.Token) *AstNode {
	var astCOMMENT  = AstCOMMENT{}
	astCOMMENT.Value=tk.Value
	var ast Ast = astCOMMENT
	return  New(COMMENT , &ast , &Position{Len:tk.Position.Len,Offset:tk.Position.Offset} )
}
func NewAstTEXT(tk *token.Token) *AstNode {
	var astTEXT  = AstTEXT{}
	astTEXT.Value=tk.Value
	var ast Ast = astTEXT
	return New(COMMENT , &ast , &Position{Len:tk.Position.Len,Offset:tk.Position.Offset} )
}

func NewAstTagAttr(key * AstNode,value * AstNode, position * Position ) *AstNode {
	var astTagAttr  = AstTagAttr{}
	astTagAttr.Key = key
	astTagAttr.Value=value
	var ast Ast = astTagAttr
	return  New(TAG_ATTR , &ast , position )
}



func (ast *AstNode) Is(astIndexs []int,value *Ast) bool {
	for _,astIndex := range astIndexs {
		b1 := ast.Index == astIndex
		b2 := true;
		if value!=nil{
			b2 =ast.Value == value
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
