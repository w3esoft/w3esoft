/* 
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

const HtmlInputFile = function (c){



};

const HtmlToken = function (tokenIndex,value){
    let me = this;
    me.tokenIndex = tokenIndex;
    me.value = value;
    me.is=function(tokenIndex){
        return (me.tokenIndex===tokenIndex);
    }
    me.isNot=function(tokenIndex){
        let me = this;
        return !me.is.apply(me,arguments);
    }
    me.toString =function () {
        if (value){
            return "HtmlToken."+TOKEN_NAMES[tokenIndex]+"("+ JSON.stringify(value)+");"
        }else {
            return "HtmlToken."+TOKEN_NAMES[tokenIndex]+"();"
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
    this.tokenize =function () {
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
                   }else if (char==CHAR_MINUS){
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
           while (isWord(char)){
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
           return new HtmlToken(value);
       }else if (char===CHAR_EOF) {
           return new HtmlToken(TOKEN_EOF);
       }else {
           throw "do'nt know charcode "+char;
       }

    };
};



const HtmlParser = function (lexer ){
    let me =this;
    me.lexer = lexer;
    me.parse =function () {
        while (true){
            let tk = me.lexer.tokenize();
            console.log(tk.toString())
            if (tk.is(TOKEN_EOF)){
                console.log("HtmlToken.......................")
                break;
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