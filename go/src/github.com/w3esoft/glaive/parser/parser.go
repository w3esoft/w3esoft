package parser

import (
	"github.com/w3esoft/glaive/lexer"
	"github.com/w3esoft/glaive/ast"
	"github.com/w3esoft/glaive/token"
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
		if tok.IsNot([]int{token.WHITESPACE,token.LINE}, nil, true) {
			return tok
		}

	}
}

func (par *Parser) TokenPush(tok *token.Token) {
	par.lexer.TokenPush(tok);
}

func (par *Parser) ParseDocumentChild() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	if tk1.Is([]int{token.EXPORT},nil,true){

		tk1 = par.Tokenize()
		print(tk1);
		return nil, nil;
	}

	print(tk1);
	return nil, nil;
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
		if (err != nil) {
			return nil, err;
		}
		value = append(value, child)
	}
	ast := ast.NewAstDOCUMENT(value, &position, true);
	return ast, nil;
}
