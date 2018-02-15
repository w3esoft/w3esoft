package token

import "testing"

func Test(t *testing.T) {
	var token = New(TAGHEADOPEN,"123");
	b1 := token.Is([]int{TAGHEADOPEN},"123");
	t.Log("Is",b1)
	b2 := token.IsNot([]int{TAGHEADOPEN},nil);
	t.Log("IsNot",b2)
	b3, e1 := token.Expected([]int{TAGHEADOPEN},nil);
	t.Log("Expected",b3,e1)
	b4, e2 := token.Unexpected([]int{TAGHEADOPEN},nil);
	t.Log("Unexpected",b4,e2)
}