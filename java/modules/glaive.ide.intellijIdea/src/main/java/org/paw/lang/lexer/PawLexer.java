package org.paw.lang.lexer;

import org.paw.lang.PawScan;

import java.util.Deque;
import java.util.LinkedList;

public class PawLexer {
    static int LEXER_MODE_TEXT = 1;
    static int LEXER_MODE_TAG = 2;
    static int LEXER_MODE_STATEMENT = 3;
    static int LEXER_MODE_MACRO = 4;


    int pos = 0;
    final static int CHAR_EOF = 0;
    final static int CHAR_DOLAR = 36;
    final static int CHAR_DOT = 46;
    final static int CHAR_MINUS = 45;
    final static int CHAR_EQUAL = 61;
    final static int CHAR_UNDERSCORE = 95;
    final static int CHAR_MINOR = 60;
    final static int CHAR_GREATER = 62;
    final static int CHAR_PARENTHESIS_OPEN = 40;
    final static int CHAR_PARENTHESIS_CLOSE = 41;
    final static int CHAR_COMMA = 44;
    final static int CHAR_COLON = 58;
    final static int CHAR_EXCLAMATION = 33;
    final static int CHAR_SLASH = 47;
    final static int CHAR_BRACE_OPEN =  123;
    final static int CHAR_BRACE_CLOSE = 125;
    final static int CHAR_BRACKET_OPEN =  91;
    final static int CHAR_BRACKET_CLOSE = 93;
    final static int CHAR_PERCENT_SIGN = 37;
    final static int CHAR_DOUBLE_QUOTATION_MARK = 34;
    final static int CHAR_QUOTATION_MARK = 39;
    final static int CHAR_HASH = 35;
    private String source;
    Deque<PawLexerToken> tokens = new LinkedList<PawLexerToken>();


    boolean isWord(int c) {
        return (c >= 65 && c <= 90) || (c >= 97 && c <= 122);
    }
    boolean isWhitespace(int c) {
        return  c == 32;
    }
    boolean isNumeric(int c) {
        return (c >= 48 && c <= 57);
    }
    private int mode;
    private PawScan scan;
    public PawLexer(PawScan scan) {
        this.scan = scan;
        this.mode = LEXER_MODE_TEXT;
    }

    public PawLexerToken tokinize() {
        int c = scan.read();
        String value = null;
        int offset = scan.getPosition();
        int len = 1;
        if (c == CHAR_EOF) {
            return new PawLexerToken(PawLexerTokenType.EOF, value, offset, 0);
        }else if (this.mode == LEXER_MODE_TEXT) {
            String text = "";
            if (c == CHAR_MINOR) {
                c = scan.read();
                this.mode = LEXER_MODE_TAG;
                if (CHAR_SLASH==c) {
                    return new PawLexerToken(PawLexerTokenType.TAG_START_END, value, offset, len);
                }
                scan.addCache(c);
                return new PawLexerToken(PawLexerTokenType.TAG_START, value, offset, len);
            } else if(c == CHAR_BRACE_OPEN){
                c = scan.read();
                if (CHAR_PERCENT_SIGN==c) {
                    this.mode = LEXER_MODE_STATEMENT;
                    return new PawLexerToken(PawLexerTokenType.META_TAG_START, value, offset, len);
                }else if (CHAR_HASH==c) {
                    this.mode = LEXER_MODE_MACRO;
                    return new PawLexerToken(PawLexerTokenType.META_TAG_START, value, offset, len);
                }
                text= "{" +((char)c);
            }
            while (true){
                c = scan.read();
                if (c == CHAR_MINOR) {
                    scan.addCache(c);
                    break;
                } if (c == CHAR_EOF) {
                    scan.addCache(c);
                    break;
                }else if(c == CHAR_BRACE_OPEN){
                    c = scan.read();
                    if (CHAR_PERCENT_SIGN==c) {
                        scan.addCache(c);
                        scan.addCache(CHAR_BRACE_OPEN);
                        break;
                    }
                }else {
                    text= text +((char)c);
                }

            }


            return new PawLexerToken(PawLexerTokenType.TOKEN_TEXT, text, offset, len);

        }else if (this.mode == LEXER_MODE_STATEMENT|| this.mode == LEXER_MODE_MACRO ||  this.mode == LEXER_MODE_TAG){
            if (isWhitespace(c)){
                while (isWhitespace(c)){
                    c = scan.read();
                }
                scan.addCache(c);
                return new PawLexerToken(PawLexerTokenType.WHITESPACE,value, offset, len);
            }else if (c ==10){
                return new PawLexerToken(PawLexerTokenType.LINE,value, offset, len);
            }else if (isWord(c)||CHAR_DOLAR==c || CHAR_UNDERSCORE==c){
                value =""+ ((char)c);
                while (true){
                    boolean b;
                    c = scan.read();
                    if (  this.mode == LEXER_MODE_TAG){
                        b = isWord(c) || isNumeric(c) || CHAR_DOLAR == c || CHAR_UNDERSCORE == c ||   c == CHAR_MINUS;
                    }else {
                        b = isWord(c) || isNumeric(c) || CHAR_DOLAR == c || CHAR_UNDERSCORE == c ;
                    }
                    if(!b){
                        break;
                    }
                    value = value+((char)c);
                }
                scan.addCache(c);
                return new PawLexerToken(PawLexerTokenType.IDENT,value, offset, len);
            } else if (isNumeric(c)){
                value ="";
                while (isNumeric(c)){
                    value = value+((char)c);
                    c = scan.read();
                }
                scan.addCache(c);
                return new PawLexerToken(PawLexerTokenType.NUMERIC,value, offset, len);
            }else if (CHAR_GREATER==c && this.mode == LEXER_MODE_TAG){
                this.mode = LEXER_MODE_TEXT;
                return new PawLexerToken(PawLexerTokenType.TAG_END,value, offset, len);
            } else if (CHAR_SLASH==c  && this.mode == LEXER_MODE_TAG) {
                c = scan.read();
                if (CHAR_GREATER==c) {
                    this.mode = LEXER_MODE_TEXT;
                    return new PawLexerToken(PawLexerTokenType.TAG_START_END, value, offset, len);
                }
                scan.addCache(c);
                return new PawLexerToken(PawLexerTokenType.ERROR, value, offset, len);
            }else if (CHAR_EQUAL==c) {
                return new PawLexerToken(PawLexerTokenType.EQUAL, value, offset, len);
            } else if (CHAR_BRACE_OPEN ==c) {
                return new PawLexerToken(PawLexerTokenType.BRACE_OPEN, value, offset, len);
            }else if (CHAR_BRACE_CLOSE ==c) {
                return new PawLexerToken(PawLexerTokenType.BRACE_CLOSE, value, offset, len);
            }else if (CHAR_BRACKET_OPEN ==c) {
                return new PawLexerToken(PawLexerTokenType.BRACKET_OPEN, value, offset, len);
            }else if (CHAR_BRACKET_CLOSE ==c) {
                return new PawLexerToken(PawLexerTokenType.BRACKET_CLOSE, value, offset, len);
            }else if (CHAR_DOT ==c) {
                return new PawLexerToken(PawLexerTokenType.DOT, value, offset, len);
            }else if (CHAR_PARENTHESIS_OPEN ==c) {
                return new PawLexerToken(PawLexerTokenType.PARENTHESIS_OPEN, value, offset, len);
            }else if (CHAR_PARENTHESIS_CLOSE ==c) {
                return new PawLexerToken(PawLexerTokenType.PARENTHESIS_CLOSE, value, offset, len);
            }else if (CHAR_PERCENT_SIGN==c ) {
                c = scan.read();
                if (CHAR_BRACE_CLOSE ==c&& this.mode == LEXER_MODE_STATEMENT) {
                    this.mode = LEXER_MODE_TEXT;
                    return new PawLexerToken(PawLexerTokenType.META_TAG_END, value, offset, len);
                }
                scan.addCache(c);
                return new PawLexerToken(PawLexerTokenType.PERCENT_SIGN, value, offset, len);
            }else if (CHAR_HASH ==c ) {
                c = scan.read();
                if (CHAR_BRACE_CLOSE ==c&& this.mode == LEXER_MODE_MACRO) {
                    this.mode = LEXER_MODE_TEXT;
                    return new PawLexerToken(PawLexerTokenType.META_TAG_END, value, offset, len);
                }
                while (isWord(c)||isNumeric(c)||CHAR_DOLAR==c|| CHAR_UNDERSCORE==c){
                    value = value+((char)c);
                    c = scan.read();
                }
                scan.addCache(c);
                return new PawLexerToken(PawLexerTokenType.HASH,value, offset, len);
            }else if (CHAR_DOUBLE_QUOTATION_MARK==c ||CHAR_QUOTATION_MARK==c ){
                int cc = c;
                value ="";
                c = scan.read();
                while (true){
                    if (cc==c)
                        break;
                    if (CHAR_EOF==c)
                        return new PawLexerToken(PawLexerTokenType.STRING_ERROR, value, offset, len);
                    value = value+((char)c);
                    c = scan.read();
                }
                return new PawLexerToken(PawLexerTokenType.STRING, value, offset, len);
            }else {
                return new PawLexerToken(PawLexerTokenType.ERROR, value, offset, len);
            }
        } else {
            return new PawLexerToken(PawLexerTokenType.ERROR, value, offset, len);
        }
    }
}
