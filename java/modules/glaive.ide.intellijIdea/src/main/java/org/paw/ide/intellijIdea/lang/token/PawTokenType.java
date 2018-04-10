package org.paw.ide.intellijIdea.lang.token;

import com.intellij.psi.TokenType;
import com.intellij.psi.tree.IElementType;
import com.intellij.psi.tree.TokenSet;
import org.paw.ide.intellijIdea.PawLanguage;
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
    public static IElementType IDENT = new IElementType("IDENT", PawLanguage.INSTANCE);
    public static IElementType HASH = new IElementType("HASH", PawLanguage.INSTANCE);
    public static IElementType DOT = new IElementType("DOT", PawLanguage.INSTANCE);
    public static IElementType META_TAG_END = new IElementType("META_TAG_END", PawLanguage.INSTANCE);
    public static IElementType META_TAG_START = new IElementType("META_TAG_START", PawLanguage.INSTANCE);
    public static IElementType BRACE_OPEN = new IElementType("BRACE_OPEN", PawLanguage.INSTANCE);
    public static IElementType BRACE_CLOSE = new IElementType("BRACE_CLOSE", PawLanguage.INSTANCE);
    public static IElementType EQUAL = new IElementType("EQUAL", PawLanguage.INSTANCE);
    public static IElementType PARENTHESIS_CLOSE = new IElementType("PARENTHESIS_CLOSE", PawLanguage.INSTANCE);
    public static IElementType PARENTHESIS_OPEN = new IElementType("PARENTHESIS_OPEN", PawLanguage.INSTANCE);
    public static IElementType TEXT = new IElementType("TEXT", PawLanguage.INSTANCE);
    public static IElementType WHITESPACE = new IElementType("WHITESPACE", PawLanguage.INSTANCE);
    public static IElementType PERCENT_SIGN = new IElementType("PERCENT_SIGN", PawLanguage.INSTANCE);
    public static IElementType BRACKET_CLOSE = new IElementType("BRACKET_CLOSE", PawLanguage.INSTANCE);
    public static IElementType BRACKET_OPEN = new IElementType("BRACKET_OPEN", PawLanguage.INSTANCE);
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
            return  ERROR_ELEMENT;
        } else if (PawLexerTokenType.LINE.name().contentEquals(t.name())) {
            return  NEW_LINE_INDENT;
        } else if (PawLexerTokenType.EOF.name().contentEquals(t.name())) {
            return  EOF;
        } else if (PawLexerTokenType.IDENT.name().contentEquals(t.name())) {
            return  IDENT;
        } else if (PawLexerTokenType.HASH.name().contentEquals(t.name())) {
            return  HASH;
        } else if (PawLexerTokenType.BRACKET_OPEN.name().contentEquals(t.name())) {
            return  BRACKET_OPEN;
        } else if (PawLexerTokenType.BRACKET_CLOSE.name().contentEquals(t.name())) {
            return  BRACKET_CLOSE;
        } else if (PawLexerTokenType.PERCENT_SIGN.name().contentEquals(t.name())) {
            return  PERCENT_SIGN;
        } else if (PawLexerTokenType.STRING.name().contentEquals(t.name())) {
            return  STRING;
        } else if (PawLexerTokenType.WHITESPACE.name().contentEquals(t.name())) {
            return  WHITESPACE;
        } else if (PawLexerTokenType.NUMERIC.name().contentEquals(t.name())) {
            return  NUMERIC;
        } else if (PawLexerTokenType.TEXT.name().contentEquals(t.name())) {
            return  TEXT;
        } else if (PawLexerTokenType.TAG_END.name().contentEquals(t.name())) {
            return  TAG_END;
        } else if (PawLexerTokenType.TAG_START.name().contentEquals(t.name())) {
            return  TAG_START;
        } else if (PawLexerTokenType.TAG_START_END.name().contentEquals(t.name())) {
            return  TAG_START_END;
        } else if (PawLexerTokenType.TAG_ATTR_TEMPLATE.name().contentEquals(t.name())) {
            return  TAG_ATTR_TEMPLATE;
        } else if (PawLexerTokenType.TAG_ATTR_INPUT.name().contentEquals(t.name())) {
            return  TAG_ATTR_INPUT;
        } else if (PawLexerTokenType.TAG_ATTR_OUTPUT.name().contentEquals(t.name())) {
            return  TAG_ATTR_OUTPUT;
        } else if (PawLexerTokenType.TAG_ATTR_INPUT_OUTPUT.name().contentEquals(t.name())) {
            return  TAG_ATTR_INPUT_OUTPUT;
        } else if (PawLexerTokenType.META_TAG_END.name().contentEquals(t.name())) {
            return  META_TAG_END;
        } else if (PawLexerTokenType.META_TAG_START.name().contentEquals(t.name())) {
            return  META_TAG_START;
        } else if (PawLexerTokenType.EQUAL.name().contentEquals(t.name())) {
            return  EQUAL;
        } else if (PawLexerTokenType.BRACE_OPEN.name().contentEquals(t.name())) {
            return  BRACE_OPEN;
        } else if (PawLexerTokenType.BRACE_CLOSE.name().contentEquals(t.name())) {
            return  BRACE_CLOSE;
        } else if (PawLexerTokenType.PARENTHESIS_OPEN.name().contentEquals(t.name())) {
            return  PARENTHESIS_OPEN;
        } else if (PawLexerTokenType.PARENTHESIS_CLOSE.name().contentEquals(t.name())) {
            return  PARENTHESIS_CLOSE;
        } else if (PawLexerTokenType.DOT.name().contentEquals(t.name())) {
            return  DOT;
        }
        return  ERROR_ELEMENT;
    }
}
