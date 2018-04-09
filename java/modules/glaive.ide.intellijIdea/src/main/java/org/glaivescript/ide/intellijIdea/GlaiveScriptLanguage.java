package org.glaivescript.ide.intellijIdea;

import com.intellij.lang.Language;

public class GlaiveScriptLanguage  extends Language {
    public static final GlaiveScriptLanguage INSTANCE = new GlaiveScriptLanguage();
    public GlaiveScriptLanguage() {
        super("GLAIVESCRIPT", new String[]{"text/x-glaivescript-source", "text/glaivescript", "application/x-glaivescript", "text/x-glaivescript"});
    }
}
