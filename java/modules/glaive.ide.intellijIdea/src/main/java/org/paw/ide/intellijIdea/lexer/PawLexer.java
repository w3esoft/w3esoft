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

import java.util.ArrayList;
import java.util.List;

public class PawLexer extends Lexer {
    private PawSyntaxHighlighter highlighter;
    PawLexerToken token;
    private org.paw.lang.lexer.PawLexer lexer;
    private List tokens = new ArrayList();
    public PawLexer(PawSyntaxHighlighter highlighter) {
        this.highlighter = highlighter;
    }
    @Override
    public void start(@NotNull CharSequence charSequence, int i, int i1, int i2) {
        lexer = new org.paw.lang.lexer.PawLexer(new PawScan(charSequence));
        while (true){
            PawLexerToken token = lexer.tokinize();
            tokens.add(token);
            if (token.is(PawLexerToken.TOKEN_EOF)){
                break;
            }
        }
        tokens.size();
    }
    @Override
    public int getState() {
        return 0;
    }
    @Nullable
    @Override
    public IElementType getTokenType() {
        return PawTokenType.getIElementTypeFromIndex(token.getIndex());
    }
    @Override
    public int getTokenStart() {
        return token.getPositionStart();
    }

    @Override
    public int getTokenEnd() {
        return token.getPositionEnd();
    }

    @Override
    public void advance() {
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
        return null;
    }

    @Override
    public int getBufferEnd() {
        return 0;
    }
}
