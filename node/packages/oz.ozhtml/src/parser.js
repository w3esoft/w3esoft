const  astConst = require("./ast_const");
const  tokenConst = require("./token_const");
const {Ast} = require("./ast");
class Parser {
    public contructor (lexer){
        let me =this;
        me.lexer = lexer;
    }
    tokenize() {
        let me =this;
        while (true){
            let tk = me.lexer.tokenize();
            // console.log( (tk.id) +  "  " +tk.toString()  );
            if (tk.isNot(tokenConst.WHITESPACE))
                return tk;
        }
    };
    pushToken(tk){
        let me =this;
        me.lexer.tokens.push(tk);
    };
    tokenToAstValue(tk) {
        if (tk.is(tokenConst.TAGCOMMENT)){
            return new Ast(astConst.TAGCOMMENT,tk.value)
        }
        return new Ast(astConst.VALUE,tk.value)
    };
    attrValue() {
        let me =this;
        let tk = tokenize();
        tk.expected([tokenConst.NUMERIC,tokenConst.STRING]);
        return me.tokenToAstValue(tk)
    };
    attrKey() {
        let me =this;
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
            fields.push(me.tokenToAstValue(tk));
            tk = tokenize();

        }
        while (tk.is(tokenConst.DOT)){
            tk = tokenize();
            tk.expected(tokenConst.WORD);
            fields.push(  me.tokenToAstValue(tk));
            tk = tokenize();
        }
        while (closeAttr.length){
            let index =closeAttr.pop();
            tk.expected(index);
            tk = tokenize();
        }
        me.pushToken(tk);

        return new Ast(astConst.ATTRKEY,{
            value: me.tokenToAstValue(value),
            fields:fields,
            isGlobal:isGlobal,
            isTemplate:isTemplate,
            operation:operation
        });
    };
    tagBody () {
        let me =this;
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
                    me. pushToken(tk);
                    _attr.key =me.attrKey();
                    tk = tokenize();
                    if (tk.is(tokenConst.EQUAL)){
                        _attr.value =me.attrValue();
                        tk = tokenize();
                    }
                    attrs.push(new Ast(astConst.ATTR,_attr));
                }
                if (tk.expected([tokenConst.TAGBODY,tokenConst.TAGHEADOPEN])){
                    me.pushToken(tk);
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
                        me.pushToken(tk2);
                        me.pushToken(tk1);
                    }
                    break;
                }
                me.pushToken(tk);
                let b = me.tagBody();
                body.push(b)
            }
            return new Ast(astConst.TAG,{
                name: me.tokenToAstValue(tkName),
                attrs:attrs,
                body:body
            });

        }else if (tk.is(tokenConst.TAGBODY)){
            return   me.tokenToAstValue(tk);
        }else if (tk.is(tokenConst.TAGCOMMENT)){
            return   me.tokenToAstValue(tk);
        }else {
            tk.unexpected(tk.tokenIndex)
        }
    };
    parse() {
        let me =this;
        let items = [];
        while (true){
            let tk = me.tokenize();
            if (tk.is(tokenConst.EOF)){
                console.log("Token.......................");
                break;
            }else {
                me.pushToken(tk);
                let  b=   me.tagBody();
                items.push(b)
            }
        }
        return new Ast(astConst.DOCUMENT,items);
    };
}

