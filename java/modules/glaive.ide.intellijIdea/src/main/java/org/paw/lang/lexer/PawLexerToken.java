package org.paw.lang.lexer;

import com.intellij.lexer.LexerPosition;

public class PawLexerToken {
    private int index;
    private LexerPosition currentPosition;
    private String value;
    private int offset;
    private int len;

    public PawLexerToken(int index, String value, int offset, int len) {
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

    public int getIndex() {
        return index;
    }

    public LexerPosition getCurrentPosition() {
        return currentPosition;
    }

    public boolean is(int index) {
        return this.index == index;
    }

    public boolean not(int index) {
        return !is(index);
    }

    public boolean is(int i, String v) {
        return is(i, value) && v.contentEquals(this.value);
    }

    public boolean not(int index, String value) {
        return !is(index, value);
    }


    public boolean is(int[] indexs, String v) {
        boolean b = is(indexs);
        return b && v.contentEquals(this.value);
    }

    public boolean not(int indexs[], String value) {
        return !is(indexs, value);
    }


    public boolean is(int[] indexs) {
        for (int index : indexs) {
            boolean b = is(index, value) && value.contentEquals(this.value);
            ;
            if (b) {
                return b;
            }
        }
        return false;
    }

    public boolean not(int indexs[]) {
        return is(indexs);
    }


}
