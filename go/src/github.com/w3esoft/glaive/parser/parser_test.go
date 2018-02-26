package parser

import (
	"testing"
	"github.com/w3esoft/tool"
	"github.com/damsistemas"
	"path/filepath"
	"github.com/w3esoft/glaive/input/inputfile"
	"github.com/w3esoft/glaive/lexer"
	"github.com/w3esoft/glaive/ast"
	"fmt"
)
func TestAll(t *testing.T) {
	files :=[]string{}
	for _,k :=range damsistemas.GetPackageMap()  {
		files =append(files,k+ "/**.html")
	}
	exclude :=[]string{"**/node_modules"}
	files =tool.Find(damsistemas.RootDir,files,exclude)
	for  _ , file :=range files{
		d:=filepath.Join(damsistemas.RootDir,file)
		fmt.Println(d)
		inputFile ,_ :=inputfile.New(d)
		l:=lexer.New(inputFile)
		build(l);
	}
}

func Test(t *testing.T) {
	d:=filepath.Join(damsistemas.RootDir,"packages/damsistemas/packages/damonline/dist/index.html")
	fmt.Println(d)
	inputFile ,_ :=inputfile.New(d)
	l:=lexer.New(inputFile)
	build(l);
}


func build(l *lexer.Lexer )  ( a *ast.AstNode, err error){
	p:= New(l);
	return  p.Parse()
}