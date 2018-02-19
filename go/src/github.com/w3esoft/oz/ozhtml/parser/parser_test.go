package parser

import (
	"testing"
	"github.com/w3esoft/tool"
	"github.com/damsistemas"
	"path/filepath"
	"github.com/w3esoft/glaive/input/inputfile"
	"github.com/w3esoft/oz/ozhtml/lexer"
	"github.com/w3esoft/oz/ozhtml/ast"
)
func TestAll(t *testing.T) {
	files :=[]string{}
	for _,k :=range damsistemas.GetPackageMap()  {
		files =append(files,k+ "/**.html")
	}
	exclude :=[]string{"**/node_modules"}
	files =tool.Find(damsistemas.RootDir,files,exclude)
	for  _ , file :=range files{
		inputFile ,_ :=inputfile.New(filepath.Join(damsistemas.RootDir,file))
		l:=lexer.New(inputFile)
		build(l);
	}
}

func Test(t *testing.T) {
	inputFile ,_ :=inputfile.New(filepath.Join(damsistemas.RootDir,"packages/atomicburst/packages/calendar/src/view_week.component.html"))
	l:=lexer.New(inputFile)
	build(l);
}


func build(l *lexer.Lexer )  ( a *ast.AstNode, err error){
	p:= New(l);
	return  p.Parse()
}