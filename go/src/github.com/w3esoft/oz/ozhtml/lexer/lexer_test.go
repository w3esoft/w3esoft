package lexer

import (
	"testing"
	 "github.com/w3esoft/glaive/input/string"
	"github.com/w3esoft/oz/ozhtml/token"
)

func Test(t *testing.T) {
	i :=string.New(`123456789`);
	l:= New(i);
	for{
		tk1:=l.Tokenize()
		if !tk1.Valid {
			t.Log(tk1);
		}else {
			t.Log(tk1.Name,tk1.Value,tk1.Position.Offset,tk1.Position.Len)
		}
		if (tk1.Is([]int{token.EOF},nil)){
			break
		}
	}
}