package org.paw.lang.lexer;

import org.paw.lang.PawScan;

import java.util.Deque;
import java.util.LinkedList;

public class PawLexer {
    static int LEXER_MODE_TEXT = 1;
    static int LEXER_MODE_TAG_DEFINITION = 2;


    int pos = 0;
    final static int CHAR_EOF = 0;
    final static int CHAR_DOLAR = 36;
    final static int CHAR_DOT = 46;
    final static int CHAR_EQUAL = 61;
    final static int CHAR_UNDERSCORE = 95;
    final static int CHAR_MINOR = 60;
    final static int CHAR_GREATER = 62;
    final static int CHAR_QUOTATION = 39;
    final static int CHAR_PARENTHESIS_OPEN = 40;
    final static int CHAR_PARENTHESIS_CLOSE = 41;
    final static int CHAR_COMMA = 44;
    final static int CHAR_COLON = 58;
    final static int CHAR_EXCLAMATION = 33;
    final static int CHAR_SLASH = 47;
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
            return new PawLexerToken(PawLexerToken.TOKEN_EOF, value, offset, 0);
        }else if (this.mode == LEXER_MODE_TEXT) {
            if (c == CHAR_MINOR) {
                c = scan.read();
                this.mode = LEXER_MODE_TAG_DEFINITION;
                if (CHAR_SLASH==c) {
                    return new PawLexerToken(PawLexerToken.TOKEN_TAG_START_END, value, offset, len);
                }
                scan.addCache(c);
                return new PawLexerToken(PawLexerToken.TOKEN_TAG_START, value, offset, len);
            }else {
                return new PawLexerToken(PawLexerToken.TOKEN_TOKEN_TEXT, value, offset, len);
            }
        }else if (this.mode == LEXER_MODE_TAG_DEFINITION ){
            if (isWhitespace(c)){
                while (isWhitespace(c)){
                    c = scan.read();
                }
                scan.addCache(c);
                return new PawLexerToken(PawLexerToken.TOKEN_WHITESPACE,value, offset, len);
            }else if (c ==10){
                while (isWhitespace(c)){
                    c = scan.read();
                }
                scan.addCache(c);
                return new PawLexerToken(PawLexerToken.TOKEN_LINE,value, offset, len);
            }else if (isWord(c)||CHAR_DOLAR==c || CHAR_UNDERSCORE==c){
                value ="";
                while (isWord(c)||isNumeric(c)||CHAR_DOLAR==c|| CHAR_UNDERSCORE==c){
                    value = value+String.valueOf(c);
                    c = scan.read();
                }
                scan.addCache(c);
                return new PawLexerToken(PawLexerToken.TOKEN_TAG_IDENT,value, offset, len);
            }else if (isNumeric(c)){
                value ="";
                while (isNumeric(c)){
                    value = value+String.valueOf(c);
                    c = scan.read();
                }
                scan.addCache(c);
                return new PawLexerToken(PawLexerToken.TOKEN_TAG_NUMERIC,value, offset, len);
            }else if (CHAR_GREATER==c ){
                this.mode = LEXER_MODE_TEXT;
                return new PawLexerToken(PawLexerToken.TOKEN_TAG_END,value, offset, len);
            } else if (CHAR_SLASH==c) {
                c = scan.read();
                this.mode = LEXER_MODE_TAG_DEFINITION;
                if (CHAR_GREATER==c) {
                    return new PawLexerToken(PawLexerToken.TOKEN_TAG_START_END, value, offset, len);
                }
                return new PawLexerToken(PawLexerToken.TOKEN_TAG_START_END, value, offset, len);
            }
        } else {
            return new PawLexerToken(PawLexerToken.TOKEN_TAG_NUMERIC, value, offset, len);
        }
        return null;
    }
}
