import * as tokenConst from "./token_const";
import * as charConst from "./char_const";
import {Token} from "./token";
export class Lexer {
    constructor(input){
        this.CHAR=-1;
        this.input = input;
        this.tokens=[];
    }
    static isWord(char){
        return ( 96 < char && 123 > char) || ( 64 < char && 91 > char || char === charConst.MINUS || char === 95 || char === 35);
    };
    static isNumeric(char){
        return ( 47 < char && 58 > char);
    };
    static isWhiteSpace(char){
        return char === 32 || char === 13 || char === 10;
    };
    tokenize () {
        let me = this;
        if (me.tokens.length)
            return me.tokens.pop();
        let char;
        if (me.CHAR!==-1) {
            char = me.CHAR;
            me.CHAR=-1;
        }else {
            char = me.input.read();
        }
        if (char===charConst.GREATER){
            char = me.input.read();
            if (char===charConst.BACKSLASH){
                return new Token(tokenConst.TAGHEADCLOSE)
            }else if (char===charConst.BANG){
            	me.input.read();
                me.input.read();
                let ii=0;
                let value="";
                while(true){
                    char= me.input.read();
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
                return new Token(tokenConst.TAGCOMMENT,value);
            }else {
                me.CHAR = char;
                return new Token(tokenConst.TAGHEADOPEN);
            }
        }else  if (char===charConst.BACKSLASH){
            char = input.read();
            if (char===charConst.MINOR){
                console.log(char);
            }
            return new Token(tokenConst.TAGHEADCLOSE);
        }else if (char===charConst.MINOR){
            let value="";
            char = input.read();
            while (char !== charConst.GREATER && char !== charConst.EOF){
                value+= String.fromCharCode(char);
                char = input.read();
            }
            me.CHAR = char;
            return new Token(tokenConst.TAGBODY,value);
        }else if (char===charConst.PARENTHESISOPEN){
            return new Token(tokenConst.PARENTHESISOPEN);
        }else if (char===charConst.PARENTHESISCLOSE){
            return new Token(tokenConst.PARENTHESISCLOSE);
        }else if (char===charConst.EQUAL) {
            return new Token(tokenConst.EQUAL);
        }else if (char===charConst.MULTIPLICATION) {


            return new Token(tokenConst.MULTIPLICATION);

        }else if (char===charConst.BRACKETOPEN) {
            return new Token(tokenConst.BRACKETOPEN);
        }else if (char===charConst.BRACKETClOSE) {
            return new Token(tokenConst.BRACKETCLOSE);
        }else if (char===charConst.DOT) {
            return new Token(tokenConst.DOT);
        }else if (char===charConst.DOUBLEDOT) {
            return new Token(tokenConst.DOUBLEDOT);
        }else if (char===charConst.DOUBLEQUOTES) {
            let value="";
            char = input.read();
            while (char !== charConst.DOUBLEQUOTES && char !== charConst.EOF){
                value+= String.fromCharCode(char);
                char = input.read();
            }
            return new Token(tokenConst.STRING,value);
        }else if (Lexer.isWhiteSpace(char)) {
            while (Lexer.isWhiteSpace(char)){
                char = input.read();
            }
            me.CHAR = char;
            return new Token(tokenConst.WHITESPACE);
        }else if (Lexer.isWord(char)) {
            let value="";
            while (Lexer.isWord(char)||Lexer.isNumeric(char)){
                value+= String.fromCharCode(char);
                char = input.read();
            }
            me.CHAR = char;
            return new Token(tokenConst.WORD,value);
        }else if (Lexer.isNumeric(char)) {
            let value="";
            while (Lexer.isNumeric(char)){
                value+= String.fromCharCode(char);
                char = input.read();
            }
            me.CHAR = char;
            return new Token(tokenConst.NUMERIC,value);
        }else if (char===charConst.EOF) {
            return new Token(tokenConst.EOF);
        }else {
            throw "do'nt know charcode "+char;
        }
    }
}
