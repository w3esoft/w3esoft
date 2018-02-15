package ast

import (
	"errors"
)

type Ast struct {
	Index int
	Value interface{}
	Name  string
}

func New(astIndex int, value interface{}) *Ast {
	var ast Ast = Ast{}
	ast.Index = astIndex
	ast.Value = value
	ast.Name = NAMES[astIndex]
	return &ast
}

func (ast *Ast) Is(astIndexs []int, value interface{}) bool {
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
func (ast *Ast) IsNot(astIndexs []int, value interface{}) bool {
	return !ast.Is(astIndexs, value)
}
func (ast *Ast) Expected(astIndexs []int, value interface{}) (r bool, err error) {
	b := ast.Is(astIndexs, value)
	if !b {
		return b, errors.New("unexpected ast " + ast.Name)
	}
	return b, nil
}
func (ast *Ast) Unexpected(astIndexs []int, value interface{}) (r bool, err error) {
	b := ast.IsNot(astIndexs, value)
	if !b {
		return b, errors.New("unexpected ast " + ast.Name)
	}
	return b, nil
}
