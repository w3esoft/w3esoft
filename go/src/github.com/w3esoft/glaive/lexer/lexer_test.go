package lexer
import (
	"testing"
	"github.com/w3esoft/glaive/token"
	"github.com/w3esoft/tool"
	"github.com/damsistemas"
	"path/filepath"
	"fmt"
	"github.com/w3esoft/glaive/input/inputfile"
)
func TestAll(t *testing.T) {
	files :=[]string{}
	for _,k :=range damsistemas.GetPackageMap()  {
		files =append(files,k+ "/**.ts")
	}
	exclude :=[]string{"**/node_modules"}
	files =tool.Find(damsistemas.RootDir,files,exclude)
	for  ii , file :=range files{
		rfile:=filepath.Join(damsistemas.RootDir,file)
		fmt.Println(ii,rfile)
		build(rfile);
	}
}

func Test(t *testing.T) {
	file:=filepath.Join(damsistemas.RootDir,"packages/atomicburst/packages/core/src/atomicburst.ts")
	build(file);
}


func build(file string ){
	fmt.Println(file)
	inputFile ,_ :=inputfile.New(file)
	l:= New(inputFile);
	for{
		tk1:=l.Tokenize()

		value := ""
		if tk1.Value!=nil {
			value = *tk1.Value
		}


		if !tk1.Valid{
			fmt.Println(file)
			fmt.Println(tk1.Valid,tk1.Name,tk1.Id,value,tk1.Position.Offset,tk1.Position.Len)
		}else {
			fmt.Println(tk1.Valid,tk1.Name,tk1.Id,value,tk1.Position.Offset,tk1.Position.Len)
		}
		if (tk1.Is([]int{token.EOF},nil,true)){
			break
		}
	}
	fmt.Println("....................... end")

}