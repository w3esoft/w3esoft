package ozhtml

import (
	"errors"
	"oz\ast_const"
)

type Ast struct {
	astIndex int
	value    string
	name     string
}



func is(ast *Ast, astIndexs [] int, value string) bool {
	for astIndex := range astIndexs {
		b1 := ast.astIndex==astIndex
		b2 := ast.value== value
		if b1 && b2 {
			return  true
		}
	}
	return  false
}
func isNot(ast *Ast, astIndexs[] int, value string) bool {
	return !is(ast, astIndexs, value)
}
func expected(ast *Ast, astIndexs[] int, value string) (r bool, err error) {
	b := is(ast, astIndexs, value)
	if b {
		return b, errors.New("unexpected ast " + ast.name)
	}
	return b, nil
}
func unexpected(ast *Ast,  astIndexs[] int, value string) (r bool, err error) {
	b := isNot(ast, astIndexs, value)
	if b {
		return b, errors.New("unexpected ast " + ast.name)
	}
	return b, nil
}
func toString(ast *Ast) string {
	if ast.value == "" {
		return "Ast." + ast.name + "();";
	} else {
		return "Ast." + ast.name + "();";
	}
}