package token

import "testing"

func Test(t *testing.T) {
	p := &Position{}
	var b ="123"
	var token = New(TAG_OPEN,&b,p,true);
	b1 := token.Is([]int{TAG_OPEN},&b,true);
	t.Log("Is",b1)
	b2 := token.IsNot([]int{TAG_OPEN},nil,true);
	t.Log("IsNot",b2)
	b3, e1 := token.Expected([]int{TAG_OPEN},nil,true);
	t.Log("Expected",b3,e1)
	b4, e2 := token.Unexpected([]int{TAG_OPEN},nil,true);
	t.Log("Unexpected",b4,e2)
}