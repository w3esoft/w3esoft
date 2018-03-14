package parser

import (
	"github.com/w3esoft/oz/ozcss/lexer"
	"github.com/w3esoft/oz/ozcss/ast"
	"github.com/w3esoft/oz/ozcss/token"
	"errors"
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
		if tok.IsNot([]int{token.TAG_WHITESPACE}, nil, true) {
			return tok
		}

	}
}

func (par *Parser) ParseAttrKey() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	if _, err = tk1.Expected([]int{
		token.KEY,
		token.TAG_KEY_INPUT,
		token.TAG_KEY_OUTPUT,
		token.TAG_KEY_TEMPLATE,
	}, nil, true); err != nil {
		return nil, err
	}
	if tk1.Is([]int{
		token.KEY,
	}, nil, true) {
		return ast.NewAstTagAttrKeyNormal(tk1.Value, &ast.Position{Offset: tk1.Position.Offset, Len: tk1.Position.Len}, true), nil
	} else if tk1.Is([]int{
		token.TAG_KEY_INPUT,
	}, nil, true) {
		return ast.NewAstTagAttrKeyInput(tk1.Value, &ast.Position{Offset: tk1.Position.Offset, Len: tk1.Position.Len}, true), nil
	} else if tk1.Is([]int{
		token.TAG_KEY_OUTPUT,
	}, nil, true) {
		return ast.NewAstTagAttrKeyOutput(tk1.Value, &ast.Position{Offset: tk1.Position.Offset, Len: tk1.Position.Len}, true), nil
	} else if tk1.Is([]int{
		token.TAG_KEY_TEMPLATE,
	}, nil, true) {
		return ast.NewAstTagAttrKeyTemplate(tk1.Value, &ast.Position{Offset: tk1.Position.Offset, Len: tk1.Position.Len}, true), nil
	}
	return nil, errors.New("ParseAttrKey")
}

func (par *Parser) ParseAttrValue() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	if _, err = tk1.Expected([]int{
		token.TAG_VALUE_NUMERIC,
		token.TAG_VALUE_STRING,
	}, nil, true); err != nil {
		return nil, err
	}
	if tk1.Is([]int{
		token.TAG_VALUE_NUMERIC,
	}, nil, true) {
		return ast.NewAstTagAttrValueNumeric(tk1.Value, &ast.Position{Offset: tk1.Position.Offset, Len: tk1.Position.Len}, true), nil
	} else if tk1.Is([]int{
		token.TAG_VALUE_STRING,
	}, nil, true) {
		return ast.NewAstTagAttrValueString(tk1.Value, &ast.Position{Offset: tk1.Position.Offset, Len: tk1.Position.Len}, true), nil
	}
	return nil, errors.New("ParseAttrKey")

}
func (par *Parser) ParseAttr() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	var tk2 = tk1;
	par.TokenPush(tk1)
	var key, _ = par.ParseAttrKey();
	var value *ast.AstNode = nil;
	tk1 = par.Tokenize()
	if (tk1.Is([]int{token.TAG_EQUAL}, nil, true)) {
		var v, err1 = par.ParseAttrValue();
		value = v;
		if err1 != nil {
			return nil, err1
		}
	} else {
		par.TokenPush(tk1)
	}
	var tk3 = par.Tokenize()
	par.TokenPush(tk3)
	return ast.NewAstTagAttr(key, value, &ast.Position{
		Len:    tk3.Position.Offset + tk3.Position.Len - tk2.Position.Offset,
		Offset: tk2.Position.Offset,
	}, true), nil
}
func (par *Parser) ParseChild() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	if (tk1.Is([]int{token.COMMENT}, nil, true)) {
		return ast.NewAstCOMMENT(tk1, true), nil;
	} else if (tk1.Is([]int{token.TEXT}, nil, true)) {
		return ast.NewAstTEXT(tk1, true), nil;
	} else if (tk1.Is([]int{token.TAG_OPEN}, nil, true)) {
		par.TokenPush(tk1);
		return par.ParseTag();
	} else if (tk1.Is([]int{token.DOCTYPE_OPEN}, nil, true)) {
		par.TokenPush(tk1);
		return par.ParseDoctype();
	} else {
		_, err = tk1.Unexpected([]int{tk1.Index}, nil, true)
		return nil, err
	}
}
func (par *Parser) ParseTagName() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	nameTk := tk1
	if _, eei := tk1.Expected([] int{token.KEY}, nil, true); eei != nil {
		return nil, eei
	}
	return ast.NewAstTagName(nameTk, true), nil
}
func (par *Parser) ParseDoctypeValue() (a *ast.AstNode, err error) {
	tk1 := par.Tokenize()
	if _, err = tk1.Expected([]int{
		token.TAG_VALUE_NUMERIC,
		token.TAG_VALUE_STRING,
		token.KEY,
	}, nil, true); err != nil {
		return nil, err
	}
	if tk1.Is([]int{
		token.KEY,
	}, nil, true) {
		return ast.NewAstDoctypeValueKey(tk1.Value, &ast.Position{Offset: tk1.Position.Offset, Len: tk1.Position.Len}, true), nil
	} else if tk1.Is([]int{
		token.TAG_VALUE_STRING,
	}, nil, true) {
		return ast.NewAstDoctypeValueString(tk1.Value, &ast.Position{Offset: tk1.Position.Offset, Len: tk1.Position.Len}, true), nil
	} else if tk1.Is([]int{
		token.TAG_VALUE_NUMERIC,
	}, nil, true) {
		return ast.NewAstDoctypeValueNumeric(tk1.Value, &ast.Position{Offset: tk1.Position.Offset, Len: tk1.Position.Len}, true), nil
	}
	return nil, errors.New("ParseAttrKey")

}

func (par *Parser) ParseDoctype() (a *ast.AstNode, err error)  {
	tk12 := par.Tokenize()
	if _, err = tk12.Expected([]int{
		token.DOCTYPE_OPEN,
	}, nil, true); err != nil {
		return nil, err
	}
	values := []*ast.AstNode{}
	for {
		tk1 := par.Tokenize()
		if (tk1.Is([]int{
			token.EOF,
			token.TAG_CLOSE,
		}, nil, true)) {
			par.TokenPush(tk1);
			break;
		}
		par.TokenPush(tk1);
		var v, err1 = par.ParseDoctypeValue();
		if (err1 != nil) {
			return nil, err1
		}
		values = append(values, v)
	}
	tk1 := par.Tokenize()
	if _, eei := tk1.Expected([] int{token.TAG_CLOSE}, nil, true); eei != nil {
		return nil, eei
	}
	return ast.NewAstDoctype(values, &ast.Position{Offset: tk12.Position.Offset, Len: tk12.Position.Len}, true), nil

}

func (par *Parser) ParseTag() (a *ast.AstNode, err error) {
	tk12 := par.Tokenize()
	if _, err = tk12.Expected([]int{
		token.TAG_OPEN,
	}, nil, true); err != nil {
		return nil, err
	}

	nameTk := par.Tokenize()
	par.TokenPush(nameTk)
	nameAst, _err := par.ParseTagName()
	if (_err != nil) {

	}
	items := []*ast.AstNode{}
	attrs := []*ast.AstNode{}
	for {
		tk1 := par.Tokenize()
		if (tk1.IsNot([]int{
			token.KEY,
			token.TAG_KEY_INPUT,
			token.TAG_KEY_OUTPUT,
			token.TAG_KEY_TEMPLATE,
		}, nil, true)) {
			par.TokenPush(tk1);
			break;
		}
		par.TokenPush(tk1);
		var a, err = par.ParseAttr()
		if (err != nil) {
			return nil, err
		}
		attrs = append(attrs, a)
	}
	tk1 := par.Tokenize()


	if (tk1.Is([]int{
		token.TAG_END_CLOSE_NOCHILD,
	}, nil, true)) {
		return ast.NewAstTag(nameAst, items, attrs, false, &ast.Position{Offset: tk12.Position.Offset, Len: tk12.Position.Len}, true), nil

	}


	if _, eei := tk1.Expected([] int{token.TAG_CLOSE}, nil, true); eei != nil {
		return nil, eei
	}
	for {
		tk4 := par.Tokenize()
		if _, eei := tk4.Unexpected([] int{token.EOF}, nil, true); eei != nil {
			return nil, eei
		}
		if (tk4.Is([]int{token.TAG_END_OPEN}, nil, true)) {

			tk6 := par.Tokenize()
			if _, eei := tk6.Expected([] int{token.KEY}, nil, true); eei != nil {
				return nil, eei
			}

			tk7 := par.Tokenize()
			if _, eei := tk7.Expected([] int{token.TAG_CLOSE}, nil, true); eei != nil {
				return nil, eei
			}

			if (*tk6.Value != *nameTk.Value) {
				par.TokenPush(tk7)
				par.TokenPush(tk6)
				par.TokenPush(tk4)
				return ast.NewAstTag(nameAst, items, attrs, true, &ast.Position{Offset: nameTk.Position.Offset, Len: nameTk.Position.Len}, true), nil
			}

			break
		}
		par.TokenPush(tk4);
		chil, rr1 := par.ParseChild();
		if (rr1 != nil) {
			return nil, rr1
		}
		items = append(items, chil)
		if (chil.Is([]int{ast.TAG}, nil)) {
			var t = chil.Value.(*ast.AstTag)
			if (t.NoItems) {
				for _, item := range t.Items {
					items = append(items, item)
				}
				t.Items = []*ast.AstNode{};
			}
		}

	}
	return ast.NewAstTag(nameAst, items, attrs, false, &ast.Position{Offset: tk12.Position.Offset, Len: tk12.Position.Len}, true), nil
}

func (par *Parser) TokenPush(tok *token.Token) {
	par.lexer.TokenPush(tok);
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
		var child, err = par.ParseChild()
		if (err != nil) {
			return nil, err;
		}
		value = append(value, child)
	}
	ast := ast.NewAstDOCUMENT(value, &position, true);
	return ast, nil;
}
