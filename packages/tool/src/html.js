/* 
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

const HtmlInputFile = function (c){



};



const AST_NAMES={};
const AST_TAG                  = 1;
AST_NAMES[AST_TAG]             = "TAG";
const AST_VALUE                = 2;
AST_NAMES[AST_VALUE]           = "VALUE";
const AST_ATTRKEY              = 3;
AST_NAMES[AST_ATTRKEY]         = "ATTRKEY";
const AST_TAGCOMMENT           = 3;
AST_NAMES[AST_TAGCOMMENT]      = "TAGCOMMENT";


const HtmlAst = function (astIndex,value){
    let me = this;
    me.astIndex = astIndex;
    me.value = value;
    me.is=function(astIndex,value){
        if (Array.isArray(astIndex)){
            let b =false;
            for ( let iii of tokenIndex){
                b =me.is(iii,value);
                if (b){
                    b = true;
                    break;
                }
            }
            return b;
        }

        return (me.astIndex===astIndex);
    }
    me.isNot=function(astIndex,value){
        let me = this;
        return !me.is.apply(me,arguments);
    }
    me.expected=function (astIndex,value) {
        let me = this;
        let b = me.is.apply(me,arguments) ;
        if (!b){
            let tk =new HtmlAst(astIndex,value);
            throw new Error("unexpected ast " + me.name);
        }
        return b;
    }
    me.unexpected=function (astIndex,value) {
        let me = this;
        let b = me.isNot.apply(me,arguments);
        if (!b){
            let tk =new HtmlAst(astIndex,value);
            throw new Error("unexpected ast " + me.name);
        }
        return b;
    }
    me.name= AST_NAMES[astIndex];
    me.toString =function () {
        if (value){
            return "HtmlAst."+me.name+"("+ JSON.stringify(value)+");";
        }else {
            return "HtmlAst."+me.name+"();";
        }
    };

};

let TOKENID= 1;
const HtmlToken = function (tokenIndex,value){
    let me = this;
    me.id=100+(TOKENID++);
    me.tokenIndex = tokenIndex;
    me.value = value;
    me.is=function(tokenIndex,value){
        if (Array.isArray(tokenIndex)){
            let b =false;
            for ( let iii of tokenIndex){
                b =me.is(iii,value);
                if (b){
                    b = true;
                    break;
                }
            }
            return b;
        }
        let b1 = (me.tokenIndex===tokenIndex);

        let b2 = (me.value===value||value===undefined);

        return b1 && b2;
    }
    me.isNot=function(tokenIndex,value){
        let me = this;
        return !me.is.apply(me,arguments);
    }
    me.expected=function (tokenIndex,value) {
        let me = this;
        let b = me.is.apply(me,arguments) ;
        if (!b){
            let tk =new HtmlToken(tokenIndex,value);
            let msg = "not expect token " + me.name + "  ";
            if (me.value){
                msg+=" value = " + me.value;
            }
            throw new Error(msg);
        }
        return b;
    }
    me.unexpected=function (tokenIndex,value) {
        let me = this;
        let b = me.isNot.apply(me,arguments);
        if (!b){
            let tk =new HtmlToken(tokenIndex,value);
            throw new Error("not expect token " +me.name);
        }
        return b;
    }
    me.name= TOKEN_NAMES[tokenIndex];
    me.toString =function () {
        if (value){
            return "HtmlToken."+me.name+"("+ JSON.stringify(value)+");"
        }else {
            return "HtmlToken."+me.name+"();"
        }
    };

};

const HtmlInputString = function (str){
    let me = this;
    me.pos=0;
    me.read=function(){
       let char = str[me.pos++];
       if (!char)
           return -1;
       else
           return char.charCodeAt(0);

    }
};

const TOKEN_NAMES={};
const TOKEN_EOF                         = 1;
TOKEN_NAMES[TOKEN_EOF]                  = "EOF";
const TOKEN_TAGHEADOPEN                 = 2;
TOKEN_NAMES[TOKEN_TAGHEADOPEN]          = "TAGHEADOPEN";
const TOKEN_TAGHEADCLOSE                = 4;
TOKEN_NAMES[TOKEN_TAGHEADCLOSE]         = "TAGHEADCLOSE";
const TOKEN_TAGBODY                     = 5;
TOKEN_NAMES[TOKEN_TAGBODY]              = "TAGBODY";
const TOKEN_WORD                        = 6;
TOKEN_NAMES[TOKEN_WORD]                 = "WORD";
const TOKEN_WHITESPACE                  = 7;
TOKEN_NAMES[TOKEN_WHITESPACE]           = "WHITESPACE";
const TOKEN_EQUAL                       = 8;
TOKEN_NAMES[TOKEN_EQUAL]                = "EQUAL";
const TOKEN_STRING                      = 9;
TOKEN_NAMES[TOKEN_STRING]               = "STRING";
const TOKEN_MULTIPLICATION              = 10;
TOKEN_NAMES[TOKEN_MULTIPLICATION]       = "MULTIPLICATION";
const TOKEN_PARENTHESISOPEN             = 11;
TOKEN_NAMES[TOKEN_PARENTHESISOPEN]      = "PARENTHESISOPEN";
const TOKEN_PARENTHESISCLOSE            = 12;
TOKEN_NAMES[TOKEN_PARENTHESISCLOSE]     = "PARENTHESISCLOSE";
const TOKEN_BRACKETOPEN                 = 13;
TOKEN_NAMES[TOKEN_BRACKETOPEN]          = "BRACKETOPEN";
const TOKEN_BRACKETCLOSE                = 14;
TOKEN_NAMES[TOKEN_BRACKETCLOSE]         = "BRACKETCLOSE";
const TOKEN_DOUBLEDOT                   = 16;
TOKEN_NAMES[TOKEN_DOUBLEDOT]            = "DOUBLEDOT";
const TOKEN_DOT                         = 17;
TOKEN_NAMES[TOKEN_DOT]                  = "DOT";
const TOKEN_TAGCOMMENT                  = 18;
TOKEN_NAMES[TOKEN_TAGCOMMENT]           = "TAGCOMMENT";
const TOKEN_NUMERIC                     = 19;
TOKEN_NAMES[TOKEN_NUMERIC]              = "NUMERIC";


const HtmlLexer = function (input){
    const CHAR_DOUBLEDOT            = 58;
    const CHAR_GREATER              = 60;
    const CHAR_MINOR                = 62;
    const CHAR_EQUAL                = 61;
    const CHAR_DOUBLEQUOTES         = 34;
    const CHAR_EOF                  = -1;
    const CHAR_MULTIPLICATION       = 42;
    const CHAR_BRACKETOPEN          = 91;
    const CHAR_BRACKETClOSE         = 93;
    const CHAR_PARENTHESISOPEN      = 40;
    const CHAR_PARENTHESISCLOSE     = 41;
    const CHAR_BACKSLASH            = 47;
    const CHAR_DOT                  = 46;
    const CHAR_BANG                 = 33;
    const CHAR_MINUS                = 45;
    let CHAR=-1;
    let me = this;
    let isWord=function(char){
       if (( 96 < char && 123 > char)||( 64 < char && 91> char|| char===CHAR_MINUS||  char===95|| char===35)){
           return true;
       }else {
           return false;
       }
    }
    let isNumeric=function(char){
        if (( 47 < char && 58 > char)){
            return true;
        }else {
            return false;
        }
    };
    let isWhiteSpace=function(char){
        if (char==32||char==13||char==10){
            return true;
        }else {
            return false;
        }
    };
    me.tokens=[];
    me.tokenize =function () {
       if (me.tokens.length)
            return me.tokens.pop();
       let char;
       if (CHAR!=-1) {
           char = CHAR;
           CHAR=-1;
       }else {
           char = input.read();
       }
       if (char===CHAR_GREATER){
           char = input.read();
           let bb= false;
           if (char==CHAR_BACKSLASH){
               return new HtmlToken(TOKEN_TAGHEADCLOSE)
           }else if (char==CHAR_BANG){
               char =input.read();
               char = input.read();
               let ii=0;
               let value="";
               while(true){
                   char= input.read();
                   value+= String.fromCharCode(char);
                   if (char==CHAR_EOF){
                       break;
                   }else if (ii==0&&char==CHAR_MINUS){
                       ii=1;
                   }else if (ii==1&&char==CHAR_MINUS){
                       ii=2;
                   }else if (ii==2&&char==CHAR_MINOR){
                       break;
                   }else {
                       ii=0;
                   }
                }
               value =value.substring(0,value.length-3)
               return new HtmlToken(TOKEN_TAGCOMMENT,value);
           }else {
               CHAR = char;
               return new HtmlToken(TOKEN_TAGHEADOPEN);
           }
       }else  if (char===CHAR_BACKSLASH){
          char = input.read();
          if (char===CHAR_MINOR){
              console.log(char);
          }
          return new HtmlToken(TOKEN_TAGHEADCLOSE);
       }else if (char===CHAR_MINOR){
           let value="";
           char = input.read();
           while (char != CHAR_GREATER && char != CHAR_EOF){
               value+= String.fromCharCode(char);
               char = input.read();
           }
           CHAR = char;
           return new HtmlToken(TOKEN_TAGBODY,value);
       }else if (char===CHAR_PARENTHESISOPEN){
            return new HtmlToken(TOKEN_PARENTHESISOPEN);
       }else if (char===CHAR_PARENTHESISCLOSE){
            return new HtmlToken(TOKEN_PARENTHESISCLOSE);
       }else if (char===CHAR_EQUAL) {
            return new HtmlToken(TOKEN_EQUAL);
       }else if (char===CHAR_MULTIPLICATION) {


            return new HtmlToken(TOKEN_MULTIPLICATION);

       }else if (char===CHAR_BRACKETOPEN) {
           return new HtmlToken(TOKEN_BRACKETOPEN);
       }else if (char===CHAR_BRACKETClOSE) {
           return new HtmlToken(TOKEN_BRACKETCLOSE);
       }else if (char===CHAR_DOT) {
           return new HtmlToken(TOKEN_DOT);
       }else if (char===CHAR_DOUBLEDOT) {
           return new HtmlToken(TOKEN_DOUBLEDOT);
       }else if (char===CHAR_DOUBLEQUOTES) {
           let value="";
           char = input.read();
           while (char != CHAR_DOUBLEQUOTES && char != CHAR_EOF){
               value+= String.fromCharCode(char);
               char = input.read();
           }
           return new HtmlToken(TOKEN_STRING,value);
       }else if (isWhiteSpace(char)) {
           while (isWhiteSpace(char)){
               char = input.read();
           }
           CHAR = char;
           return new HtmlToken(TOKEN_WHITESPACE);
       }else if (isWord(char)) {
           let value="";
           while (isWord(char)||isNumeric(char)){
               value+= String.fromCharCode(char);
               char = input.read();
           }
           CHAR = char;
           return new HtmlToken(TOKEN_WORD,value);
       }else if (isNumeric(char)) {
           let value="";
           while (isNumeric(char)){
               value+= String.fromCharCode(char);
               char = input.read();
           }
           CHAR = char;
           return new HtmlToken(TOKEN_NUMERIC,value);
       }else if (char===CHAR_EOF) {
           return new HtmlToken(TOKEN_EOF);
       }else {
           throw "do'nt know charcode "+char;
       }

    };
    me.showAllToken=function(){
        let p = [];
        while(true){
           let tk =  me.tokenize();
           console.log(tk.toString());
           if (tk.is(TOKEN_EOF)) break;
        }
    }

};



const HtmlParser = function (lexer ){
    let me =this;
    me.lexer = lexer;
    const tokenize =function () {
        while (true){
            let tk = me.lexer.tokenize();
           // console.log( (tk.id) +  "  " +tk.toString()  );
            if (tk.isNot(TOKEN_WHITESPACE))
                return tk;
        }
    };
    const pushToken=function(tk){
        me.lexer.tokens.push(tk);
    };

    const tokenToAstValue = function (tk) {
        if (tk.is(TOKEN_TAGCOMMENT)){
            return new HtmlAst(AST_TAGCOMMENT,tk.value)
        }


        return new HtmlAst(AST_VALUE,tk.value)
    }
    const attrValue =function () {
        let tk = tokenize();
        tk.expected([TOKEN_NUMERIC,TOKEN_STRING]);
        return tokenToAstValue(tk)
    }
    const attrKey =function () {
        let tk = tokenize();
        let fields = [];
        let hasTokenMultiplication=false;
        let isGlobal=false;
        let closeAttr=[];
        if (tk.is(TOKEN_MULTIPLICATION)){
            hasTokenMultiplication=true;
            tk = tokenize();
        }
        let operation=[];
        while (true){
            if (tk.is(TOKEN_PARENTHESISOPEN)){
                closeAttr.push(TOKEN_PARENTHESISCLOSE);
                operation.push("OUTPUT");
                tk = tokenize();
            }else if (tk.is(TOKEN_BRACKETOPEN)){
                operation.push("INPUT");
                closeAttr.push(TOKEN_BRACKETCLOSE);
                tk = tokenize();
            }else {
                break;
            }
        }
        let  value = tk;
        tk.expected(TOKEN_WORD);
        tk = tokenize();
        if (tk.is(TOKEN_DOUBLEDOT)){
            isGlobal=true;
            tk = tokenize();
            tk.expected(TOKEN_WORD);
            fields.push(tokenToAstValue(tk));
            tk = tokenize();

        }
        while (tk.is(TOKEN_DOT)){
            tk = tokenize();
            tk.expected(TOKEN_WORD);
            fields.push(tokenToAstValue(tk));
            tk = tokenize();
        }
        while (closeAttr.length){
            let index =closeAttr.pop();
            tk.expected(index);
            tk = tokenize();
        }
        pushToken(tk);

        return new HtmlAst(AST_ATTRKEY,{
            value:value,
            fields:fields,
            isGlobal:isGlobal,
            operation:operation
        });
    }


    const tagBody = function () {
        let tk = tokenize();
        if (tk.is(TOKEN_TAGHEADOPEN)){
            let tkName = tokenize();
            tkName.expected(TOKEN_WORD);
            let attrs =[];
            let body =[];
            while (true){
                let tk = tokenize();
                while(tk.is([TOKEN_WORD,TOKEN_MULTIPLICATION,TOKEN_PARENTHESISOPEN,TOKEN_BRACKETOPEN])){
                    let _attr ={};
                    pushToken(tk);
                    _attr.key =attrKey();
                    tk = tokenize();
                    if (tk.is(TOKEN_EQUAL)){
                        _attr.key =attrValue();
                        tk = tokenize();
                    }
                }
                if (tk.is(TOKEN_NUMERIC)){
                    console.log (tk);
                    TOKEN_NUMERIC;
                }
                if (tk.expected([TOKEN_TAGBODY,TOKEN_TAGHEADOPEN])){
                    pushToken(tk);
                    break;
                }
            }
            while (true){
               tk = tokenize();
               if (tk.is(TOKEN_TAGHEADCLOSE)){
                   let tk1 =tk;
                   tk.toString();
                   let tk2 = tk = tokenize();
                   tk.expected(TOKEN_WORD)
                   if (!tk.is(TOKEN_WORD,tkName.value)){
                       pushToken(tk2);
                       pushToken(tk1);
                   }
                   break;
               }
               if (tk.is(TOKEN_EOF)){
                   tk;
               }
               pushToken(tk);
               let b = tagBody();
                body.push(b)
            }
            return new HtmlAst(AST_TAG,{
                name: tokenToAstValue(tkName),
                attrs:attrs,
                body:body
            });

        }else if (tk.is(TOKEN_TAGBODY)){
            return tokenToAstValue(tk);
        }else if (tk.is(TOKEN_TAGCOMMENT)){
            return tokenToAstValue(tk);
        }else {
            tk.unexpected(tk.tokenIndex)
        }
    }
    me.parse =function () {
        while (true){
            let tk = tokenize();
            if (tk.is(TOKEN_EOF)){
                console.log("HtmlToken.......................")
                break;
            }else {
                pushToken(tk);
               let  b= tagBody();
            }
        }
    };
};

module.exports = {
    HtmlParser,
    HtmlInputFile,
    HtmlInputString,
    HtmlLexer
};