package org.glaive.ide.intellijIdea.lang;

import com.intellij.lang.Language;

public class GlaiveScriptLanguage extends Language {
    public  static String ID ="GLAIVE_SCRIPT_LANGUAGE";
    public  static GlaiveScriptLanguage INSTANCE = new GlaiveScriptLanguage();
    public GlaiveScriptLanguage() {
        super(ID);
    }
}
