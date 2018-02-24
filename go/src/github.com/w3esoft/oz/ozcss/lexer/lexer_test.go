package lexer
import (
	"testing"
	"github.com/w3esoft/oz/ozhtml/token"
	"github.com/w3esoft/tool"
	"github.com/damsistemas"
	"path/filepath"
	"fmt"
	"github.com/w3esoft/glaive/input/inputfile"
	"github.com/w3esoft/glaive/input"
)
func TestAll(t *testing.T) {
	files :=[]string{}
	for _,k :=range damsistemas.GetPackageMap()  {
		files =append(files,k+ "/**.sass")
		files =append(files,k+ "/**.scss")
	}
	exclude :=[]string{"**/node_modules"}
	files =tool.Find(damsistemas.RootDir,files,exclude)
	for  ii , file :=range files{

		fmt.Println(ii,file);
		inputFile ,_ :=inputfile.New(filepath.Join(damsistemas.RootDir,file))
		build(inputFile);
	}
}

func Test(t *testing.T) {
	file:=filepath.Join(damsistemas.RootDir,"packages/damsistemas/packages/damgerenciador/src/index.sass")
	fmt.Println(file);
	inputFile ,_ :=inputfile.New(file)
	build(inputFile);
}


func build(inputFile input.Input ){
	l:= New(inputFile);
	for{
		tk1:=l.Tokenize()

		value := ""
		if tk1.Value!=nil {
			value = *tk1.Value
		}


		if !tk1.Valid{
			fmt.Println(tk1.Valid,tk1.Name,tk1.Id,value,tk1.Position.Offset,tk1.Position.Len)
		}else {
			fmt.Println(tk1.Valid,tk1.Name,tk1.Id,value,tk1.Position.Offset,tk1.Position.Len)
		}
		if (tk1.Is([]int{token.EOF},nil,true)){
			break
		}
	}

}