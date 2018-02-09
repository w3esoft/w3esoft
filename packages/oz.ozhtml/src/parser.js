const  astConst = require("./ast_const");
const  tokenConst = require("./token_const");
const {HtmlAst} = require("./ast");
const HtmlParser = function (lexer){
    let me =this;
    me.lexer = lexer;
    const tokenize =function () {
        while (true){
            let tk = me.lexer.tokenize();
            // console.log( (tk.id) +  "  " +tk.toString()  );
            if (tk.isNot(tokenConst.WHITESPACE))
                return tk;
        }
    };
    const pushToken=function(tk){
        me.lexer.tokens.push(tk);
    };
    const tokenToAstValue = function (tk) {
        if (tk.is(tokenConst.TAGCOMMENT)){
            return new HtmlAst(astConst.TAGCOMMENT,tk.value)
        }
        return new HtmlAst(astConst.VALUE,tk.value)
    };
    const attrValue =function () {
        let tk = tokenize();
        tk.expected([tokenConst.NUMERIC,tokenConst.STRING]);
        return tokenToAstValue(tk)
    };
    const attrKey =function () {
        let tk = tokenize();
        let fields = [];
        let isTemplate=false;
        let isGlobal=false;
        let closeAttr=[];
        if (tk.is(tokenConst.MULTIPLICATION)){
            isTemplate=true;
            tk = tokenize();
        }
        let operation=[];
        while (true){
            if (tk.is(tokenConst.PARENTHESISOPEN)){
                closeAttr.push(tokenConst.PARENTHESISCLOSE);
                operation.push("OUTPUT");
                tk = tokenize();
            }else if (tk.is(tokenConst.BRACKETOPEN)){
                operation.push("INPUT");
                closeAttr.push(tokenConst.BRACKETCLOSE);
                tk = tokenize();
            }else {
                break;
            }
        }
        let  value = tk;
        tk.expected(tokenConst.WORD);
        tk = tokenize();
        if (tk.is(tokenConst.DOUBLEDOT)){
            isGlobal=true;
            tk = tokenize();
            tk.expected(tokenConst.WORD);
            fields.push(tokenToAstValue(tk));
            tk = tokenize();

        }
        while (tk.is(tokenConst.DOT)){
            tk = tokenize();
            tk.expected(tokenConst.WORD);
            fields.push(tokenToAstValue(tk));
            tk = tokenize();
        }
        while (closeAttr.length){
            let index =closeAttr.pop();
            tk.expected(index);
            tk = tokenize();
        }
        pushToken(tk);

        return new HtmlAst(astConst.ATTRKEY,{
            value: tokenToAstValue(value),
            fields:fields,
            isGlobal:isGlobal,
            isTemplate:isTemplate,
            operation:operation
        });
    };
    const tagBody = function () {
        let tk = tokenize();
        if (tk.is(tokenConst.TAGHEADOPEN)){
            let tkName = tokenize();
            tkName.expected(tokenConst.WORD);
            let attrs =[];
            let body =[];
            while (true){
                let tk = tokenize();
                while(tk.is([tokenConst.WORD,tokenConst.MULTIPLICATION,tokenConst.PARENTHESISOPEN,tokenConst.BRACKETOPEN])){
                    let _attr ={};
                    pushToken(tk);
                    _attr.key =attrKey();
                    tk = tokenize();
                    if (tk.is(tokenConst.EQUAL)){
                        _attr.value =attrValue();
                        tk = tokenize();
                    }
                    attrs.push(new HtmlAst(astConst.ATTR,_attr));
                }
                if (tk.expected([tokenConst.TAGBODY,tokenConst.TAGHEADOPEN])){
                    pushToken(tk);
                    break;
                }
            }
            while (true){
                tk = tokenize();
                if (tk.is(tokenConst.TAGHEADCLOSE)){
                    let tk1 =tk;
                    tk.toString();
                    let tk2 = tk = tokenize();
                    tk.expected(tokenConst.WORD);
                    if (!tk.is(tokenConst.WORD,tkName.value)){
                        pushToken(tk2);
                        pushToken(tk1);
                    }
                    break;
                }
                pushToken(tk);
                let b = tagBody();
                body.push(b)
            }
            return new HtmlAst(astConst.TAG,{
                name: tokenToAstValue(tkName),
                attrs:attrs,
                body:body
            });

        }else if (tk.is(tokenConst.TAGBODY)){
            return tokenToAstValue(tk);
        }else if (tk.is(tokenConst.TAGCOMMENT)){
            return tokenToAstValue(tk);
        }else {
            tk.unexpected(tk.tokenIndex)
        }
    };
    me.parse =function () {
        let items = [];
        while (true){
            let tk = tokenize();
            if (tk.is(tokenConst.EOF)){
                console.log("HtmlToken.......................");
                break;
            }else {
                pushToken(tk);
                let  b= tagBody();
                items.push(b)
            }
        }
        return new HtmlAst(astConst.DOCUMENT,items);
    };
};
module.exports = {
    HtmlParser
};

