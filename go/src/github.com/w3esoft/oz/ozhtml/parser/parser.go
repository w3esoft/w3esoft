package parser

import (
	"github.com/w3esoft/oz/ozhtml/lexer"
	"github.com/w3esoft/oz/ozhtml/ast"
	"github.com/w3esoft/oz/ozhtml/token"
)

type Parser struct {
	lexer *lexer.Lexer
}


func New(lex *lexer.Lexer) *Parser {
	par := Parser{}
	par.lexer=lex;
	return &par
}
func (par *Parser ) Tokenize() *token.Token{
	for{
		tok :=par.lexer.Tokenize()
		if tok.IsNot([]int{token.TAG_WHITESPACE},nil,true){
			return  tok
		}

	}
}

func (par *Parser ) ParseAttrKey()  ( a *ast.AstNode, err error) {
	tk1:=par.Tokenize()
	if _ , err =tk1.Expected([]int{
		token.KEY,
		token.TAG_KEY_INPUT,
		token.TAG_KEY_OUTPUT,
		token.TAG_KEY_TEMPLATE,
	},nil,true);err!=nil{
		return nil,err
	}
	operators:=[]int{}
	if (tk1.Is([]int{token.TAG_KEY_TEMPLATE},nil,true)){
		operators =append(operators,ast.AST_TAG_ATTR_OP_TEMPLATE)
	}else {
		for {
			if (tk1.Is([]int{token.TAG_KEY_INPUT},nil,true)){
				operators =append(operators,ast.AST_TAG_ATTR_OP_INPUT)
				tk1=par.Tokenize()
			}else if (tk1.Is([]int{token.TAG_KEY_OUTPUT},nil,true)){
				operators =append(operators,ast.AST_TAG_ATTR_OP_OUTPUT)
				tk1=par.Tokenize()
			}else {
				break
			}
		}

	}
	_,err =tk1.Expected([]int{token.KEY},nil,true)
	if err!=nil {
		return nil,err
	}
	var nameTok = tk1

}

func (par *Parser ) ParseAttrValue()  ( a *ast.AstNode, err error) {


}
func (par *Parser ) ParseAttr()  ( a *ast.AstNode, err error) {
	tk1:=par.Tokenize()
	var tk2= tk1;
	par.TokenPush(tk1)
	par.ParseChild();
	var key * ast.AstNode;
	var value * ast.AstNode;
	if (tk1.Is([]int{token.TAG_EQUAL},nil,true)){

	}else {
		par.TokenPush(tk1)
	}
	var tk3=par.Tokenize()
	par.TokenPush(tk3)
	return ast.NewAstTagAttr(key,value,&ast.Position{
		Len:tk3.Position.Offset+tk3.Position.Len-tk2.Position.Offset,
		Offset:tk2.Position.Offset,
	}),nil
}
func (par *Parser ) ParseChild()  ( a *ast.AstNode, err error) {
	tk1:=par.Tokenize()
	if (tk1.Is([]int{token.COMMENT},nil,true)){
		return ast.NewAstCOMMENT(tk1),nil;
	}else if (tk1.Is([]int{token.TEXT},nil,true)){
		return ast.NewAstTEXT(tk1),nil;
	}else if (tk1.Is([]int{token.TAG_OPEN},nil,true)){
		tk1:=par.Tokenize()
		for{
			if (tk1.IsNot([]int{
					token.KEY,
					token.TAG_KEY_INPUT,
					token.TAG_KEY_OUTPUT,
					token.TAG_KEY_TEMPLATE,
				},nil,true)){
				break;
			}
			par.TokenPush(tk1);
			par.ParseAttr()
		}

		_,err=tk1.Unexpected([]int{tk1.Index},nil,true)
		return nil,err
	}else {
		_,err=tk1.Unexpected([]int{tk1.Index},nil,true)
		return nil,err
	}
}


func (par *Parser ) TokenPush( tok *token.Token ){
	par.lexer.TokenPush(tok);
}

func (par *Parser ) Parse()  ( a *ast.AstNode, err error) {
	var value  =[]*ast.AstNode{};
	var position =ast.Position{};
	for{
		tk1:=par.Tokenize()
		if (tk1.Is([]int{token.EOF},nil,true)){
			break
		}
		par.TokenPush(tk1)
		var child,err = par.ParseChild()
		if (err!=nil){
			return  nil,err;
		}
		value =append(value,child)
	}
	ast:=ast.NewAstDOCUMENT(value,&position);
	return ast,nil;
}
