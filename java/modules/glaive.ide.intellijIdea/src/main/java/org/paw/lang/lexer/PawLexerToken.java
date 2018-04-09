package org.paw.lang.lexer;

import com.intellij.lexer.LexerPosition;

public class PawLexerToken {
    private PawLexerTokenType index;
    private LexerPosition currentPosition;
    private String value;
    private int offset;
    private int len;

    public PawLexerToken(PawLexerTokenType index, String value, int offset, int len) {
        this.index = index;
        this.value = value;
        this.offset = offset;
        this.len = len;
    }

    public int getPositionStart() {
        return offset;

    }

    public int getPositionEnd() {
        return offset + len;
    }

    public PawLexerTokenType getIndex() {
        return index;
    }

    public LexerPosition getCurrentPosition() {
        return currentPosition;
    }

    public boolean is(PawLexerTokenType index) {
        return this.index.name().contentEquals(index.name());
    }

    public boolean not(PawLexerTokenType index) {
        return !is(index);
    }

    public boolean is(PawLexerTokenType i, String v) {
        return is(i, value) && v.contentEquals(this.value);
    }

    public boolean not(PawLexerTokenType index, String value) {
        return !is(index, value);
    }


    public boolean is(PawLexerTokenType[] indexs, String v) {
        boolean b = is(indexs);
        return b && v.contentEquals(this.value);
    }

    public boolean not(PawLexerTokenType indexs[], String value) {
        return !is(indexs, value);
    }


    public boolean is(PawLexerTokenType[] indexs) {
        for (PawLexerTokenType index : indexs) {
            boolean b = is(index, value) && value.contentEquals(this.value);
            ;
            if (b) {
                return b;
            }
        }
        return false;
    }

    public boolean not(PawLexerTokenType indexs[]) {
        return is(indexs);
    }


}
