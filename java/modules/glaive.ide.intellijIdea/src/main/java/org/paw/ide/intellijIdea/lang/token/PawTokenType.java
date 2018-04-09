package org.paw.ide.intellijIdea.lang.token;

import com.intellij.psi.TokenType;
import com.intellij.psi.tree.IElementType;
import com.intellij.psi.tree.TokenSet;
import org.paw.ide.intellijIdea.PawLanguage;
import org.paw.lang.lexer.PawLexerToken;
import org.paw.lang.lexer.PawLexerTokenType;

public class PawTokenType {

    public static final TokenSet WHITE_SPACES = TokenSet.create(PawTokenType.WHITE_SPACE);
    public static final TokenSet COMMENTS = TokenSet.create(PawTokenType.COMMENT);
    public static final TokenSet STRINGS = TokenSet.create(
            PawTokenType.STRING
    );

    public static IElementType EOF = new IElementType("EOF", PawLanguage.INSTANCE);
    public static IElementType TAG_END = new IElementType("TAG_END", PawLanguage.INSTANCE);
    public static IElementType TAG_START = new IElementType("TAG_START", PawLanguage.INSTANCE);
    public static IElementType TAG_START_END = new IElementType("TAG_START_END", PawLanguage.INSTANCE);
    public static IElementType TAG_ATTR_TEMPLATE = new IElementType("TAG_ATTR_TEMPLATE", PawLanguage.INSTANCE);
    public static IElementType TAG_ATTR_INPUT = new IElementType("TAG_ATTR_INPUT", PawLanguage.INSTANCE);
    public static IElementType TAG_ATTR_OUTPUT = new IElementType("TAG_ATTR_OUTPUT", PawLanguage.INSTANCE);
    public static IElementType TAG_ATTR_INPUT_OUTPUT = new IElementType("TAG_ATTR_INPUT_OUTPUT", PawLanguage.INSTANCE);
    public static IElementType STRING = new IElementType("TAG_TOP_VALUE_STRING", PawLanguage.INSTANCE);
    public static IElementType NUMERIC = new IElementType("TAG_TOP_VALUE_NUMERIC", PawLanguage.INSTANCE);
    public static IElementType COMMENT = new IElementType("COMMENT", PawLanguage.INSTANCE);
    public static IElementType WHITE_SPACE = TokenType.WHITE_SPACE;
    public static IElementType BAD_CHARACTER = TokenType.BAD_CHARACTER;
    public static IElementType NEW_LINE_INDENT = TokenType.NEW_LINE_INDENT;
    public static IElementType ERROR_ELEMENT = TokenType.ERROR_ELEMENT;
    public static IElementType CODE_FRAGMENT = TokenType.CODE_FRAGMENT;
    public static IElementType DUMMY_HOLDER = TokenType.DUMMY_HOLDER;

    public static IElementType getIElementTypeFromIndex(PawLexerTokenType t) {
        if (PawLexerTokenType.ERROR.name().contentEquals(t.name())) {
            return  ERROR_ELEMENT;
        } else if (PawLexerTokenType.STRING_ERROR.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.LINE.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.IDENT.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.HASH.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.BRACKET_OPEN.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.BRACKET_CLOSE.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.PERCENT_SIGN.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.STRING.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.WHITESPACE.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.NUMERIC.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.TOKEN_TEXT.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.TAG_END.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.TAG_START.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.TAG_START_END.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.TAG_ATTR_TEMPLATE.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.TAG_ATTR_INPUT.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.TAG_ATTR_OUTPUT.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.TAG_ATTR_INPUT_OUTPUT.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.META_TAG_END.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.META_TAG_START.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.EQUAL.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.BRACE_OPEN.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.BRACE_CLOSE.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.PARENTHESIS_OPEN.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.PARENTHESIS_CLOSE.name().contentEquals(t.name())) {
        } else if (PawLexerTokenType.DOT.name().contentEquals(t.name())) {
        }
    }
}
