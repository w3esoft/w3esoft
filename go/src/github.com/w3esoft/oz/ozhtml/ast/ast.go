package ast

import (
	"errors"
)
type Ast struct {
	Index int
	Value    *string
	Name     *string
}

func New(astIndex int, value *string) * Ast {
	var ast Ast = Ast{}
	ast.Index = astIndex;
	ast.Value = value;
	ast.Name = "";
	return  &ast;
}

func (ast Ast) Is(astIndexs [] int, value *string) bool {
	for astIndex := range astIndexs {
		b1 := ast.Index==astIndex
		b2 := ast.Value== value
		if b1 && b2 {
			return  true
		}
	}
	return  false
}
func (ast Ast)  IsNot( astIndexs[] int, value *string) bool {
	return !ast.Is( astIndexs, value)
}
func (ast Ast)  expected( astIndexs[] int, value *string) (r bool, err error) {
	b := ast.Is( astIndexs, value)
	if b {
		return b, errors.New("unexpected ast " + ast.Name)
	}
	return b, nil
}
func (ast Ast)  unexpected(  astIndexs[] int, value *string) (r bool, err error) {
	b := ast.IsNot( astIndexs, value)
	if b {
		return b, errors.New("unexpected ast " + ast.Name)
	}
	return b, nil
}