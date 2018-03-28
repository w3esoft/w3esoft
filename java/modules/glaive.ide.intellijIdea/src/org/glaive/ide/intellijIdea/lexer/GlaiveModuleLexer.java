package org.glaive.ide.intellijIdea.lexer;

import com.intellij.lexer.Lexer;
import com.intellij.lexer.LexerPosition;
import com.intellij.psi.tree.IElementType;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

public class GlaiveModuleLexer extends Lexer {
    @Override
    public void start(@NotNull CharSequence charSequence, int i, int i1, int i2) {

    }

    @Override
    public int getState() {
        return 0;
    }

    @Nullable
    @Override
    public IElementType getTokenType() {
        return null;
    }

    @Override
    public int getTokenStart() {
        return 0;
    }

    @Override
    public int getTokenEnd() {
        return 0;
    }

    @Override
    public void advance() {

    }

    @NotNull
    @Override
    public LexerPosition getCurrentPosition() {
        return null;
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
