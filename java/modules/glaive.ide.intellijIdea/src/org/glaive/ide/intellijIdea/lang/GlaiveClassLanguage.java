package org.glaive.ide.intellijIdea.lang;

import com.intellij.lang.Language;

public class GlaiveClassLanguage extends Language {
    public  static String ID ="GLAIVE_CLASS_LANGUAGE";
    public  static GlaiveClassLanguage INSTANCE = new GlaiveClassLanguage();
    public GlaiveClassLanguage() {
        super(ID);
    }
}
