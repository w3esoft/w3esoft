package org.paw.ide.intellijIdea.lang.token;

import com.intellij.psi.TokenType;
import com.intellij.psi.tree.IElementType;
import com.intellij.psi.tree.TokenSet;
import org.paw.ide.intellijIdea.PawLanguage;
import org.paw.lang.lexer.PawLexerToken;

public class PawTokenType {

    public static final TokenSet WHITE_SPACES = TokenSet.create(PawTokenType.WHITE_SPACE);
    public static final TokenSet COMMENTS = TokenSet.create(PawTokenType.COMMENT);
    public static final TokenSet STRINGS = TokenSet.create(
            PawTokenType.TAG_TOP_VALUE_STRING ,
            PawTokenType.TAG_TOP_ATTR_NORMAL
    );

    public static IElementType EOF = new IElementType("EOF", PawLanguage.INSTANCE);
    public static IElementType TAG_TOP_END = new IElementType("TAG_TOP_END", PawLanguage.INSTANCE);
    public static IElementType TAG_TOP_START = new IElementType("TAG_TOP_START", PawLanguage.INSTANCE);
    public static IElementType TAG_TOP_BOTTOM_END = new IElementType("TAG_TOP_BOTTOM_END", PawLanguage.INSTANCE);
    public static IElementType TAG_TOP_ATTR_NORMAL = new IElementType("TAG_TOP_ATTR_NORMAL", PawLanguage.INSTANCE);
    public static IElementType TAG_TOP_ATTR_TEMPLATE = new IElementType("TAG_TOP_ATTR_TEMPLATE", PawLanguage.INSTANCE);
    public static IElementType TAG_TOP_ATTR_INPUT = new IElementType("TAG_TOP_ATTR_INPUT", PawLanguage.INSTANCE);
    public static IElementType TAG_TOP_ATTR_OUTPUT = new IElementType("TAG_TOP_ATTR_OUTPUT", PawLanguage.INSTANCE);
    public static IElementType TAG_TOP_ATTR_INPUT_OUTPUT = new IElementType("TAG_TOP_ATTR_INPUT_OUTPUT", PawLanguage.INSTANCE);
    public static IElementType TAG_TOP_VALUE_STRING = new IElementType("TAG_TOP_VALUE_STRING", PawLanguage.INSTANCE);
    public static IElementType TAG_TOP_VALUE_NUMERIC = new IElementType("TAG_TOP_VALUE_NUMERIC", PawLanguage.INSTANCE);
    public static IElementType COMMENT = new IElementType("COMMENT", PawLanguage.INSTANCE);
    public static IElementType WHITE_SPACE = TokenType.WHITE_SPACE;
    public static IElementType BAD_CHARACTER = TokenType.BAD_CHARACTER;
    public static IElementType NEW_LINE_INDENT = TokenType.NEW_LINE_INDENT;
    public static IElementType ERROR_ELEMENT = TokenType.ERROR_ELEMENT;
    public static IElementType CODE_FRAGMENT = TokenType.CODE_FRAGMENT;
    public static IElementType DUMMY_HOLDER = TokenType.DUMMY_HOLDER;

    public static IElementType getIElementTypeFromIndex(int  index) {
        switch (index) {
            case PawLexerToken.TOKEN_EOF:
                return PawTokenType.EOF;
            case PawLexerToken.TOKEN_TAG_TOP_END:
                return PawTokenType.TAG_TOP_END;
            case PawLexerToken.TOKEN_TAG_TOP_START:
                return PawTokenType.TAG_TOP_START;
            case PawLexerToken.TOKEN_TAG_TOP_BOTTOM_END:
                return PawTokenType.TAG_TOP_BOTTOM_END;
            case PawLexerToken.TOKEN_TAG_IDENT:
                return PawTokenType.TAG_TOP_ATTR_NORMAL;
            case PawLexerToken.TOKEN_TAG_TOP_ATTR_TEMPLATE:
                return PawTokenType.TAG_TOP_ATTR_TEMPLATE;
            case PawLexerToken.TOKEN_TAG_TOP_ATTR_INPUT:
                return PawTokenType.TAG_TOP_ATTR_INPUT;
            case PawLexerToken.TOKEN_TAG_TOP_ATTR_OUTPUT:
                return PawTokenType.TAG_TOP_ATTR_OUTPUT;
            case PawLexerToken.TOKEN_TAG_TOP_ATTR_INPUT_OUTPUT:
                return PawTokenType.TAG_TOP_ATTR_INPUT_OUTPUT;
            case PawLexerToken.TOKEN_TAG_STRING:
                return PawTokenType.TAG_TOP_VALUE_STRING;
            case PawLexerToken.TOKEN_TAG_NUMERIC:
                return PawTokenType.TAG_TOP_VALUE_NUMERIC;
            default:
                return null;
        }
    }
}
