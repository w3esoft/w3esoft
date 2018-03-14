package parser

import (
	"github.com/w3esoft/java/lexer"
	"github.com/w3esoft/java/ast"
	"github.com/w3esoft/java/token"
)

type Parser struct {
	lexer *lexer.Lexer
}

func New(lex *lexer.Lexer) *Parser {
	par := Parser{}
	par.lexer = lex;
	return &par
}
func (par *Parser) Tokenize() *token.Token {
	for {
		tok := par.lexer.Tokenize()
		if tok.IsNot([]int{token.WHITESPACE, token.LINE}, nil, true) {
			return tok
		}

	}
}

func (par *Parser) TokenPush(tok *token.Token) {
	par.lexer.TokenPush(tok);
}

func (par *Parser) ParseDestructuringAssignmentArrayMatching() (a *ast.AstNode, err error) {

	tk1 := par.Tokenize()
	if _, err := tk1.Expected([]int{token.BRACKETOPEN}, nil, true); err != nil {
		return nil, err;
	}
	return nil, nil;
}
func (par *Parser) ParseDestructuringAssignmentObjectMatchingField() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	name := tk1

	if _, err := tk1.Expected([]int{token.WORD}, nil, true); err != nil {
		return nil, err;
	}
	var next *ast.AstNode = nil
	tk1 = par.Tokenize()
	if tk1.Is([]int{token.DOT}, nil, true) {
		var er error
		next, er = par.ParseDestructuringAssignmentObjectMatchingField();
		if er != nil {
			return nil, err;
		}
		tk1 = par.Tokenize()
	}
	par.TokenPush(tk1)
	return ast.NewAstDestructuringAssignmentObjectMatchingField(next, name.Value, &ast.Position{Offset: tk1.Position.Offset, Len: tk1.Position.Len}), nil
}

func (par *Parser) ParseDestructuringAssignmentObjectMatching() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	start := tk1
	if _, err := tk1.Expected([]int{token.BRACEOPEN}, nil, true); err != nil {
		return nil, err;
	}
	fields := []*ast.AstNode{};
	for {
		if tk1.Is([]int{token.EOF}, nil, true) {
			return nil, nil;
		}
		if tk1.Is([]int{token.BRACECLOSE}, nil, true) {
			break
		}
		item, er := par.ParseDestructuringAssignmentObjectMatchingField();
		if (er != nil) {
			return nil, er;
		}
		fields = append(fields, item)
		tk1 = par.Tokenize()
		if tk1.IsNot([]int{token.COMMA}, nil, true) {
			break
		}
	}
	if _, err := tk1.Expected([]int{token.BRACECLOSE}, nil, true); err != nil {
		return nil, err;
	}
	end := par.Tokenize()
	par.TokenPush(end)
	return ast.NewAstDestructuringAssignmentObjectMatching(fields, &ast.Position{Offset: start.Position.Offset, Len: start.Position.Len}), nil;
}
func (par *Parser) ParseLet() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	start :=tk1;
	tk1 = par.Tokenize()
	tk2 := par.Tokenize()
	par.TokenPush(tk2)
	return ast.NewAstLet(nil, &ast.Position{Len: start.Position.Len, Offset: start.Position.Len,}), nil
}
func (par *Parser) ParseConst() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	par.TokenPush(tk1)
	return nil, nil;
}
func (par *Parser) ParseImport() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	if _, err := tk1.Expected([]int{token.IMPORT}, nil, true); err != nil {
		return nil, err;
	}
	start := tk1
	tk1 = par.Tokenize()
	if tk1.Is([]int{token.ASTERISK}, nil, true) {
		tk1 = par.Tokenize()
		if _, err := tk1.Expected([]int{token.FROM}, nil, true); err != nil {
			return nil, err;
		}
		tk1 = par.Tokenize()
		if _, err := tk1.Expected([]int{token.STRING}, nil, true); err != nil {
			return nil, err;
		}
		name := tk1
		tk1 = par.Tokenize()
		if tk1.Is([]int{token.SEMICOLON}, nil, true) {
			tk1 = par.Tokenize()
		}
		par.TokenPush(tk1);
		return ast.NewAstExportAny(name.Value, &ast.Position{Len: start.Position.Len, Offset: start.Position.Len,}), nil
	} else if tk1.Is([]int{token.BRACEOPEN, token.BRACKETOPEN}, nil, true) {
		par.TokenPush(tk1);
		nn, ee := par.ParseDestructuringAssignmentObjectMatching();
		if ee != nil {

			return nil, ee;
		}
		tk1 = par.Tokenize()

		if _, err := tk1.Expected([]int{token.FROM}, nil, true); err != nil {
			return nil, err;
		}
		tk1 = par.Tokenize()
		if _, err := tk1.Expected([]int{token.STRING}, nil, true); err != nil {
			return nil, err;
		}
		name := tk1
		tk1 = par.Tokenize()
		if tk1.Is([]int{token.SEMICOLON}, nil, true) {
			tk1 = par.Tokenize()
		}
		par.TokenPush(tk1);
		return ast.NewAstExportDestructuringAssignment(nn,name.Value, &ast.Position{Len: start.Position.Len, Offset: start.Position.Len,}), nil
	}else {
		print(tk1);
	}
	return nil, nil;
}

func (par *Parser) ParseExport() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	if _, err := tk1.Expected([]int{token.EXPORT}, nil, true); err != nil {
		return nil, err;
	}
	start := tk1
	tk1 = par.Tokenize()
	if tk1.Is([]int{token.ASTERISK}, nil, true) {
		tk1 = par.Tokenize()
		if tk1.Is([]int{token.FROM}, nil, true) {

			tk1 = par.Tokenize()
			if _, err := tk1.Expected([]int{token.STRING}, nil, true); err != nil {
				return nil, err;
			}
			name := tk1
			tk1 = par.Tokenize()
			if tk1.Is([]int{token.SEMICOLON}, nil, true) {
				tk1 = par.Tokenize()
			}
			par.TokenPush(tk1);
			return ast.NewAstExportAny(name.Value, &ast.Position{Len: start.Position.Len, Offset: start.Position.Len,}), nil
		} else {

			print(tk1);
		}
		return nil, nil;
	} else if tk1.Is([]int{token.CLASS}, nil, true)  {
		par.TokenPush(tk1);
		item, er := par.ParseClass();
		if (er != nil) {
			return nil, er;
		}
		return ast.NewAstExportValue(item, &ast.Position{Len: start.Position.Len, Offset: start.Position.Len,}), nil
	} else if tk1.Is([]int{token.CONST}, nil, true)  {
		par.TokenPush(tk1);
		par.ParseConst()
	}else if tk1.Is([]int{token.LET}, nil, true)  {
		par.TokenPush(tk1);
		par.ParseLet()
	}else  {
		print(tk1);
	}
	return nil, nil;
}

func (par *Parser) ParseClass() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	print(tk1);
	return nil, nil;
}

func (par *Parser) ParseCall(aa*ast.AstNode ) (a *ast.AstNode, err error) {
	par.Tokenize()
	//start := tk1
	parameters , _:=par.ParseEmpression(nil);
	return ast.NewAstMeta(aa,parameters, &ast.Position{Len: start.Position.Len, Offset: start.Position.Len}), nil;
}

func (par *Parser) ParseMeta() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	start := tk1
	var field *ast.AstNode = nil
	tk1 = par.Tokenize()
	if (tk1.Is([]int{token.DOT}, nil, true)){
		tk1 = par.Tokenize()
	}else {

	}
	par.TokenPush(tk1)
	return ast.NewAstMeta(par.ParseIdent(start),field, &ast.Position{Len: start.Position.Len, Offset: start.Position.Len}), nil;
}

func (par *Parser) ParseIdent(tk * token.Token) (a *ast.AstNode) {
	return ast.NewAstIdent(tk.Value, &ast.Position{Len: tk.Position.Len, Offset: tk.Position.Len});
}

func (par *Parser) ParseObjectItem() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	start := tk1
	par.TokenPush(tk1)
	var value * ast.AstNode= nil;
	var key,er1 =par.ParseEmpression(nil);
	if er1!=nil{
		return nil, err;
	}
	tk1 = par.Tokenize()
	if   tk1.Is([]int{token.DOUBLEDOT}, nil, true) {
		var v,er2 =par.ParseEmpression(nil);
		if er2!=nil{
			return nil, err;
		}
		value= v;
	}
	return ast.NewAstObjectItem(key,value, &ast.Position{Len: start.Position.Len, Offset: start.Position.Len}), nil;
}

func (par *Parser) ParseObject() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	start := tk1
	if _, err := tk1.Expected([]int{token.BRACEOPEN}, nil, true); err != nil {
		return nil, err;
	}
	values:= []* ast.AstNode{}
	tk1 = par.Tokenize()
	for {
		if (tk1.Is([]int{token.BRACECLOSE}, nil, true)){
			break;
		}
		par.TokenPush(tk1)
		var item,er =par.ParseObjectItem()
		if (er!=nil){
			return nil, err;
		}
		values = append(values, item)
		tk1 = par.Tokenize()
		if (tk1.IsNot([]int{token.COMMA}, nil, true)){
			break;
		}else {
			tk1 = par.Tokenize()
		}
	}

	if _, err := tk1.Expected([]int{token.BRACECLOSE}, nil, true); err != nil {
		return nil, err;
	}

	return ast.NewAstObject(values, &ast.Position{Len: start.Position.Len, Offset: start.Position.Len}), nil;
}



func (par *Parser) ParseArray() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	start := tk1
	if _, err := tk1.Expected([]int{token.BRACKETOPEN}, nil, true); err != nil {
		return nil, err;
	}
	values:= []* ast.AstNode{}
	tk1 = par.Tokenize()
	for {
		if (tk1.Is([]int{token.BRACKETClOSE}, nil, true)){
			break;
		}
		par.TokenPush(tk1)
		var item,er =par.ParseEmpression(nil);
		if (er!=nil){
			return nil, err;
		}
		values = append(values, item)
		tk1 = par.Tokenize()
		if (tk1.IsNot([]int{token.COMMA}, nil, true)){
			break;
		}else {

			tk1 = par.Tokenize()
		}
	}

	if _, err := tk1.Expected([]int{token.BRACKETClOSE}, nil, true); err != nil {
		return nil, err;
	}
	return ast.NewAstArray(values, &ast.Position{Len: start.Position.Len, Offset: start.Position.Len}), nil;
}


func (par *Parser) ParseField() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	start := tk1
	var field *ast.AstNode = nil
	tk1 = par.Tokenize()
	if (tk1.Is([]int{token.DOT}, nil, true)){
		par.TokenPush(tk1)
		field , err = par.ParseField();
		if (err!=nil){
			return nil, err;
		}
		tk1 = par.Tokenize()
	}
	par.TokenPush(tk1)
	return ast.NewAstField(par.ParseIdent(start),field, &ast.Position{Len: start.Position.Len, Offset: start.Position.Len}), nil;
}


func (par *Parser) ParseEmpression(aa *ast.AstNode ) (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	if  aa==nil && tk1.Is([]int{token.LET}, nil, true) {
		par.TokenPush(tk1)
		return par.ParseLet();
	}else if aa==nil && tk1.Is([]int{token.META}, nil, true) {
		par.TokenPush(tk1)
		aa ,ee :=  par.ParseMeta();
		if (ee!=nil){
			return aa, ee;
		}
		return par.ParseEmpression(aa);
	}else if aa!=nil && aa.Is([]int{ast.META,ast.FIELD}, nil)&& tk1.Is([]int{token.PARENTHESISOPEN}, nil, true) {
		par.TokenPush(tk1)
		return par.ParseCall(aa);
	}else if aa==nil &&tk1.Is([]int{token.BRACEOPEN}, nil, true)  {
		par.TokenPush(tk1)
		return par.ParseObject();
	}else if aa==nil && tk1.Is([]int{token.BRACKETOPEN}, nil, true)  {
		par.TokenPush(tk1)
		return par.ParseArray();
	}else if aa==nil && tk1.Is([]int{token.WORD}, nil, true)  {
		par.TokenPush(tk1)
		return par.ParseField();

	}else {
		println(tk1)
	}
	return nil, nil;
}

func (par *Parser) ParseDocumentChild() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	if tk1.Is([]int{token.EXPORT}, nil, true) {
		par.TokenPush(tk1)
		return par.ParseExport();
	} else if tk1.Is([]int{token.IMPORT}, nil, true) {
		par.TokenPush(tk1)
		return par.ParseImport();
	} else{
		par.TokenPush(tk1)
		return par.ParseEmpression(nil)
	}
}
func (par *Parser) Parse() (a *ast.AstNode, err error) {
	var value = []*ast.AstNode{};
	var position = ast.Position{};
	for {
		tk1 := par.Tokenize()
		if (tk1.Is([]int{token.EOF}, nil, true)) {
			break
		}
		par.TokenPush(tk1)
		var child, err = par.ParseDocumentChild()
		if (err != nil||child==nil) {
			return nil, err;
		}
		value = append(value, child)
	}
	ast := ast.NewAstModule(value, &position);
	return ast, nil;
}
