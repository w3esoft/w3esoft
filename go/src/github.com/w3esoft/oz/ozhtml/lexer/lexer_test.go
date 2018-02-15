package lexer

import (
	"testing"
	 "github.com/w3esoft/glaive/input/string"
)

func Test(t *testing.T) {
	i :=string.New(`
		<html>
		</html>
	`);
	l:= New(i);
	for{
		tk1,e1:=l.Tokenize();
		if (e1!=nil){

			t.Log(e1);
		}else {

			t.Log(tk1);
			if (tk1.Is([]int{EOF},nil)){
				break;
			}
		}
	}
}