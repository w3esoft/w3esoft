const tokenConst = require("./token_const");
const charConst = require("./char_const");
const {HtmlToken} = require("./token");
const HtmlLexer = function (input){
    let CHAR=-1;
    let me = this;
    let isWord=function(char){
        return ( 96 < char && 123 > char) || ( 64 < char && 91 > char || char === charConst.MINUS || char === 95 || char === 35);
    };
    let isNumeric=function(char){
        return ( 47 < char && 58 > char);
    };
    let isWhiteSpace=function(char){
        return char === 32 || char === 13 || char === 10;
    };
    me.tokens=[];
    me.tokenize =function () {
        if (me.tokens.length)
            return me.tokens.pop();
        let char;
        if (CHAR!==-1) {
            char = CHAR;
            CHAR=-1;
        }else {
            char = input.read();
        }
        if (char===charConst.GREATER){
            char = input.read();
            if (char===charConst.BACKSLASH){
                return new HtmlToken(tokenConst.TAGHEADCLOSE)
            }else if (char===charConst.BANG){
                 input.read();
                 input.read();
                let ii=0;
                let value="";
                while(true){
                    char= input.read();
                    value+= String.fromCharCode(char);
                    if (char===charConst.EOF){
                        break;
                    }else if (ii===0&&char===charConst.MINUS){
                        ii=1;
                    }else if (ii===1&&char===charConst.MINUS){
                        ii=2;
                    }else if (ii===2&&char===charConst.MINOR){
                        break;
                    }else {
                        ii=0;
                    }
                }
                value =value.substring(0,value.length-3);
                return new HtmlToken(tokenConst.TAGCOMMENT,value);
            }else {
                CHAR = char;
                return new HtmlToken(tokenConst.TAGHEADOPEN);
            }
        }else  if (char===charConst.BACKSLASH){
            char = input.read();
            if (char===charConst.MINOR){
                console.log(char);
            }
            return new HtmlToken(tokenConst.TAGHEADCLOSE);
        }else if (char===charConst.MINOR){
            let value="";
            char = input.read();
            while (char !== charConst.GREATER && char !== charConst.EOF){
                value+= String.fromCharCode(char);
                char = input.read();
            }
            CHAR = char;
            return new HtmlToken(tokenConst.TAGBODY,value);
        }else if (char===charConst.PARENTHESISOPEN){
            return new HtmlToken(tokenConst.PARENTHESISOPEN);
        }else if (char===charConst.PARENTHESISCLOSE){
            return new HtmlToken(tokenConst.PARENTHESISCLOSE);
        }else if (char===charConst.EQUAL) {
            return new HtmlToken(tokenConst.EQUAL);
        }else if (char===charConst.MULTIPLICATION) {


            return new HtmlToken(tokenConst.MULTIPLICATION);

        }else if (char===charConst.BRACKETOPEN) {
            return new HtmlToken(tokenConst.BRACKETOPEN);
        }else if (char===charConst.BRACKETClOSE) {
            return new HtmlToken(tokenConst.BRACKETCLOSE);
        }else if (char===charConst.DOT) {
            return new HtmlToken(tokenConst.DOT);
        }else if (char===charConst.DOUBLEDOT) {
            return new HtmlToken(tokenConst.DOUBLEDOT);
        }else if (char===charConst.DOUBLEQUOTES) {
            let value="";
            char = input.read();
            while (char !== charConst.DOUBLEQUOTES && char !== charConst.EOF){
                value+= String.fromCharCode(char);
                char = input.read();
            }
            return new HtmlToken(tokenConst.STRING,value);
        }else if (isWhiteSpace(char)) {
            while (isWhiteSpace(char)){
                char = input.read();
            }
            CHAR = char;
            return new HtmlToken(tokenConst.WHITESPACE);
        }else if (isWord(char)) {
            let value="";
            while (isWord(char)||isNumeric(char)){
                value+= String.fromCharCode(char);
                char = input.read();
            }
            CHAR = char;
            return new HtmlToken(tokenConst.WORD,value);
        }else if (isNumeric(char)) {
            let value="";
            while (isNumeric(char)){
                value+= String.fromCharCode(char);
                char = input.read();
            }
            CHAR = char;
            return new HtmlToken(tokenConst.NUMERIC,value);
        }else if (char===charConst.EOF) {
            return new HtmlToken(tokenConst.EOF);
        }else {
            throw "do'nt know charcode "+char;
        }

    };
};
module.exports = {
    HtmlLexer
};



