package org.glaive.ide.intellijIdea;

import com.intellij.lang.Language;

public class GlaiveLanguage extends Language {
    public  static String ID ="GLAIVE_LANGUAGE";
    public  static GlaiveLanguage INSTANCE = new GlaiveLanguage();
    public GlaiveLanguage() {
        super(ID);
    }
}
