package parser

import (
	"github.com/w3esoft/oz/ozhtml/ast"
	"github.com/w3esoft/oz/ozhtml/lexer"
	"github.com/w3esoft/oz/ozhtml/token"
)

type Parser struct {
	lexer *lexer.Lexer
}

func (p Parser) Tokenize() (r *token.Token, err error) {
	for {
		tk, err := p.lexer.Tokenize()
		if tk.Is([]int{token.WHITESPACE}, nil) {
			return tk, err
		}
	}
}
func PushToken(p *Parser, tk token.Token) {
	p.lexer.Tokens = append(p.lexer.Tokens, tk)
	return
}
func TokenToAstValue(tk *token.Token) *ast.Ast {
	if tk.Is([]int{token.TAGCOMMENT}, nil) {
		return ast.New(ast.TAG_COMMENT, tk.Value)
	} else {
		return ast.New(ast.VALUE, tk.Value)
	}
}
func (p Parser) attrValue() *ast.Ast {
	tk, _ := p.Tokenize()
	tk.Expected([]int{token.NUMERIC, token.STRING}, nil)
	return TokenToAstValue(tk)
}

//func (p Parser ) attrKey() *ast.Ast {
//	tk, _ := p.Tokenize()
//	var fields [] *ast.Ast
//	isTemplate := false
//	isGlobal := false
//	closeAttr := []int{}
//	if (tk.Is([]int{token.MULTIPLICATION}, nil)) {
//		isTemplate = true
//		tk, _ = p.Tokenize()
//	}
//	operation := []string{}
//	for {
//		if (tk.Is([]int{token.PARENTHESISOPEN}, nil)) {
//			closeAttr = append(closeAttr, token.PARENTHESISCLOSE)
//			operation = append(operation, "OUTPUT")
//			tk, _ = p.Tokenize()
//		} else if (tk.Is([]int{token.BRACKETOPEN}, nil)) {
//			closeAttr = append(closeAttr, token.BRACKETCLOSE)
//			operation = append(operation, "INPUT")
//			tk, _ = p.Tokenize()
//		} else {
//			break
//		}
//	}
//	value := tk
//	tk.Expected([]int{token.WORD}, nil)
//	tk, _ = p.Tokenize()
//	if (tk.Is([]int{token.DOUBLEDOT}, nil)) {
//		isGlobal = true
//		tk, _ = p.Tokenize()
//		tk.Expected([]int{token.WORD}, nil)
//		fields = append(fields, TokenToAstValue(tk))
//		tk, _ = p.Tokenize()
//	}
//	for (tk.Is([]int{token.DOT}, nil)) {
//		tk, _ = p.Tokenize()
//		tk.Expected([]int{token.WORD}, nil)
//		fields = append(fields, TokenToAstValue(tk))
//		tk, _ = p.Tokenize()
//	}
//
//}
//attrKey() {
//let me =this;
//let tk = tokenize();
//let fields = [];
//let isTemplate=false;
//let isGlobal=false;
//let closeAttr=[];
//if (tk.is(tokenConst.MULTIPLICATION)){
//isTemplate=true;
//tk = tokenize();
//}
//let operation=[];
//while (true){
//if (tk.is(tokenConst.PARENTHESISOPEN)){
//closeAttr.push(tokenConst.PARENTHESISCLOSE);
//operation.push("OUTPUT");
//tk = tokenize();
//}else if (tk.is(tokenConst.BRACKETOPEN)){
//operation.push("INPUT");
//closeAttr.push(tokenConst.BRACKETCLOSE);
//tk = tokenize();
//}else {
//break;
//}
//}
//let  value = tk;
//tk.expected(tokenConst.WORD);
//tk = tokenize();
//if (tk.is(tokenConst.DOUBLEDOT)){
//isGlobal=true;
//tk = tokenize();
//tk.expected(tokenConst.WORD);
//fields.push(me.tokenToAstValue(tk));
//tk = tokenize();
//
//}
//while (tk.is(tokenConst.DOT)){
//tk = tokenize();
//tk.expected(tokenConst.WORD);
//fields.push(  me.tokenToAstValue(tk));
//tk = tokenize();
//}
//while (closeAttr.length){
//let index =closeAttr.pop();
//tk.expected(index);
//tk = tokenize();
//}
//me.pushToken(tk);
//
//return new Ast(astConst.ATTRKEY,{
//value: me.tokenToAstValue(value),
//fields:fields,
//isGlobal:isGlobal,
//isTemplate:isTemplate,
//operation:operation
//});
//};
//tagBody () {
//let me =this
//let tk = tokenize()
//if (tk.is(tokenConst.TAGHEADOPEN)){
//let tkName = tokenize()
//tkName.expected(tokenConst.WORD)
//let attrs =[]
//let body =[]
//while (true){
//let tk = tokenize()
//while(tk.is([tokenConst.WORD,tokenConst.MULTIPLICATION,tokenConst.PARENTHESISOPEN,tokenConst.BRACKETOPEN])){
//let _attr ={}
//me. pushToken(tk)
//_attr.key =me.attrKey()
//tk = tokenize()
//if (tk.is(tokenConst.EQUAL)){
//_attr.value =me.attrValue()
//tk = tokenize()
//}
//attrs.push(new Ast(astConst.ATTR,_attr))
//}
//if (tk.expected([tokenConst.TAGBODY,tokenConst.TAGHEADOPEN])){
//me.pushToken(tk)
//break
//}
//}
//while (true){
//tk = tokenize()
//if (tk.is(tokenConst.TAGHEADCLOSE)){
//let tk1 =tk
//tk.toString()
//let tk2 = tk = tokenize()
//tk.expected(tokenConst.WORD)
//if (!tk.is(tokenConst.WORD,tkName.value)){
//me.pushToken(tk2)
//me.pushToken(tk1)
//}
//break
//}
//me.pushToken(tk)
//let b = me.tagBody()
//body.push(b)
//}
//return new Ast(astConst.TAG,{
//name: me.tokenToAstValue(tkName),
//attrs:attrs,
//body:body
//})
//
//}else if (tk.is(tokenConst.TAGBODY)){
//return   me.tokenToAstValue(tk)
//}else if (tk.is(tokenConst.TAGCOMMENT)){
//return   me.tokenToAstValue(tk)
//}else {
//tk.unexpected(tk.tokenIndex)
//}
//}
//parse() {
//let me =this
//let items = []
//while (true){
//let tk = me.tokenize()
//if (tk.is(tokenConst.EOF)){
//console.log("Token.......................")
//break
//}else {
//me.pushToken(tk)
//let  b=   me.tagBody()
//items.push(b)
//}
//}
//return new Ast(astConst.DOCUMENT,items)
//}
