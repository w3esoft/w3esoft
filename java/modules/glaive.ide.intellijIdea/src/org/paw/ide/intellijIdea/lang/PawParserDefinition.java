package org.paw.ide.intellijIdea.lang;

import com.intellij.lang.ASTNode;
import com.intellij.lang.Language;
import com.intellij.lang.ParserDefinition;
import com.intellij.lang.PsiParser;
import com.intellij.lexer.Lexer;
import com.intellij.openapi.project.Project;
import com.intellij.psi.FileViewProvider;
import com.intellij.psi.PsiElement;
import com.intellij.psi.PsiFile;
import com.intellij.psi.TokenType;
import com.intellij.psi.tree.IFileElementType;
import com.intellij.psi.tree.TokenSet;
import org.paw.ide.intellijIdea.PawLanguage;
import org.jetbrains.annotations.NotNull;
import org.paw.ide.intellijIdea.lang.token.PawToken;

public class PawParserDefinition implements ParserDefinition {

    public static final TokenSet WHITE_SPACES = TokenSet.create(TokenType.WHITE_SPACE);
    public static final TokenSet COMMENTS = TokenSet.create(PawToken.COMMENT);

    public static final IFileElementType FILE = new IFileElementType(Language.<PawLanguage>findInstance(PawLanguage.class));


    @NotNull
    @Override
    public Lexer createLexer(Project project) {
        return null;
    }

    @Override
    public PsiParser createParser(Project project) {
        return null;
    }

    @Override
    public IFileElementType getFileNodeType() {
        return FILE;
    }

    @NotNull
    @Override
    public TokenSet getWhitespaceTokens() {
        return WHITE_SPACES;
    }

    @NotNull
    @Override
    public TokenSet getCommentTokens() {
        return COMMENTS;
    }

    @NotNull
    @Override
    public TokenSet getStringLiteralElements() {
        return null;
    }

    @NotNull
    @Override
    public PsiElement createElement(ASTNode astNode) {
        return null;
    }

    @Override
    public PsiFile createFile(FileViewProvider fileViewProvider) {
        return null;
    }

    @Override
    public SpaceRequirements spaceExistanceTypeBetweenTokens(ASTNode astNode, ASTNode astNode1) {
        return null;
    }
}
