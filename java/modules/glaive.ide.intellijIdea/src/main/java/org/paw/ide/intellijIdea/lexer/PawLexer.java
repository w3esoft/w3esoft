package org.paw.ide.intellijIdea.lexer;

import com.intellij.lexer.Lexer;
import com.intellij.lexer.LexerPosition;
import com.intellij.psi.tree.IElementType;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
import org.paw.ide.intellijIdea.fileTypes.PawSyntaxHighlighter;
import org.paw.ide.intellijIdea.lang.token.PawTokenType;
import org.paw.lang.PawScan;
import org.paw.lang.lexer.PawLexerToken;
import org.paw.lang.lexer.PawLexerTokenType;

import java.util.ArrayList;
import java.util.List;

public class PawLexer extends Lexer {
    private PawSyntaxHighlighter highlighter;
    PawLexerToken token;
    private org.paw.lang.lexer.PawLexer lexer;
    CharSequence charSequence;

    public PawLexer(PawSyntaxHighlighter highlighter) {
        this.highlighter = highlighter;
    }
    @Override
    public void start(@NotNull CharSequence charSequence, int s, int pEnd, int i2) {
        this.charSequence=charSequence;
        lexer = new org.paw.lang.lexer.PawLexer(new PawScan(charSequence));
        token = lexer.tokinize();
//        while (true){
//            tokens.add(token);
//            if (token.is(PawLexerTokenType.EOF)){
//                break;
//            }
//        }
//        tokens.size();
    }
    @Override
    public int getState() {
        return 0;
    }
    @Nullable
    @Override
    public IElementType getTokenType() {

        if (token.is(PawLexerTokenType.EOF)){

            return null;
        }else {
            return PawTokenType.getIElementTypeFromIndex(token.getIndex());
        }
    }
    @Override
    public int getTokenStart() {
        return token.getOffset()+1;
    }

    @Override
    public int getTokenEnd() {
       return token.getOffset() + token.getLength()+1;

    }

    @Override
    public void advance() {
        token = lexer.tokinize();
    }

    @NotNull
    @Override
    public LexerPosition getCurrentPosition() {
        return  token.getCurrentPosition();
    }

    @Override
    public void restore(@NotNull LexerPosition lexerPosition) {

    }

    @NotNull
    @Override
    public CharSequence getBufferSequence() {
        return charSequence;
    }

    @Override
    public int getBufferEnd() {
        return 0;
    }
}
