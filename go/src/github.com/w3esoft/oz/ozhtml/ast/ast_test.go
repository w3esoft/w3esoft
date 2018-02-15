package ast

import "testing"

func Test(t *testing.T) {
	var ast = New(TAG,"123");
	b1 := ast.Is([]int{TAG},"123");
	t.Log("Is",b1)
	b2 := ast.IsNot([]int{TAG},nil);
	t.Log("IsNot",b2)
	b3, e1 := ast.Expected([]int{TAG},nil);
	t.Log("Expected",b3,e1)
	b4, e2 := ast.Unexpected([]int{TAG},nil);
	t.Log("Unexpected",b4,e2)
}