package org.glaive.ide.intellijIdea.fileTypes;

import com.intellij.lexer.Lexer;
import com.intellij.openapi.editor.colors.TextAttributesKey;
import com.intellij.openapi.fileTypes.SyntaxHighlighter;
import com.intellij.openapi.project.Project;
import com.intellij.psi.tree.IElementType;
import org.glaive.ide.intellijIdea.lexer.GlaiveScriptLexer;
import org.jetbrains.annotations.NotNull;

public class GlaiveScriptSyntaxHighlighter implements SyntaxHighlighter {
    public GlaiveScriptSyntaxHighlighter(Project project) {

    }

    @NotNull
    @Override
    public Lexer getHighlightingLexer() {
        return new GlaiveScriptLexer();
    }

    @NotNull
    @Override
    public TextAttributesKey[] getTokenHighlights(IElementType iElementType) {
        return new TextAttributesKey[0];
    }
}
