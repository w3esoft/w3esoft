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
		if tok.IsNot([]int{token.WHITESPACE, token.LINE}, nil, true) {
			return tok
		}

	}
}

func (par *Parser) TokenPush(tok *token.Token) {
	par.lexer.TokenPush(tok);
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
	}else if tk1.Is([]int{token.LET}, nil, true) {

		print(tk1);
	}else {

		print(tk1);
	}
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
	}else if tk1.Is([]int{token.LET}, nil, true) {

		print(tk1);
	}else {

		print(tk1);
	}
	return nil, nil;
}
func (par *Parser) ParseDocumentChild() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	if tk1.Is([]int{token.EXPORT}, nil, true) {
		par.TokenPush(tk1)
		return  par.ParseExport();
	}else if tk1.Is([]int{token.IMPORT}, nil, true) {
		par.TokenPush(tk1)
		return  par.ParseImport();
	}
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
	ast := ast.NewAstModule(value, &position);
	return ast, nil;
}
