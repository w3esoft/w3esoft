package org.juice.ide.intellijIdea.fileTypes;

import com.intellij.lexer.Lexer;
import com.intellij.openapi.editor.colors.TextAttributesKey;
import com.intellij.openapi.fileTypes.SyntaxHighlighter;
import com.intellij.openapi.project.Project;
import com.intellij.psi.tree.IElementType;
import org.jetbrains.annotations.NotNull;

public class JuiceSyntaxHighlighter implements SyntaxHighlighter {
    public JuiceSyntaxHighlighter(Project project) {

    }

    @NotNull
    @Override
    public Lexer getHighlightingLexer() {
        return null;
    }

    @NotNull
    @Override
    public TextAttributesKey[] getTokenHighlights(IElementType iElementType) {
        return new TextAttributesKey[0];
    }
}
