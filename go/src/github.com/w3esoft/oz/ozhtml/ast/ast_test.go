package ast

import "testing"

func TestAverage(t *testing.T) {
	var ast Ast = New(TAG,"dfsdfsdf");
	b1, _ := ast.Expected([]int{TAG},nil);
	t.Log(" Expected ",b1)
	b2, _ := ast.Expected([]int{TAG},nil);
	t.Log(" Expected ",b2)
	b3, _ := ast.Expected([]int{TAG},nil);
	t.Log(" Expected ",b3)
	b4, _ := ast.Expected([]int{TAG},nil);
	t.Log(" Expected ",b4)
	b5, _ := ast.Expected([]int{TAG},nil);
	t.Log(" Expected ",b5,)
}