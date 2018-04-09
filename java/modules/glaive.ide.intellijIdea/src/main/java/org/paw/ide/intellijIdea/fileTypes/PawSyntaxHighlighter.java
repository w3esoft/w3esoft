package org.paw.ide.intellijIdea.fileTypes;

import com.intellij.lexer.Lexer;
import com.intellij.openapi.editor.colors.TextAttributesKey;
import com.intellij.openapi.fileTypes.SyntaxHighlighter;
import com.intellij.openapi.project.Project;
import com.intellij.psi.tree.IElementType;
import org.jetbrains.annotations.NotNull;
import org.paw.ide.intellijIdea.lexer.PawLexer;

public class PawSyntaxHighlighter implements SyntaxHighlighter {
    private Project project;

    public PawSyntaxHighlighter(Project project) {

        this.project = project;
    }

    @NotNull
    @Override
    public Lexer getHighlightingLexer() {
        return new PawLexer(this);
    }

    @NotNull
    @Override
    public TextAttributesKey[] getTokenHighlights(IElementType iElementType) {
        return new TextAttributesKey[0];
    }
}
