package org.glaive.ide.intellijIdea;

import com.intellij.lang.Language;

public class GlaiveLanguage extends Language {
    public static final GlaiveLanguage INSTANCE = new GlaiveLanguage();
    public GlaiveLanguage() {
        super("GLAIVE", new String[]{"text/x-glaive-source", "text/glaive", "application/x-glaive", "text/x-glaive"});
    }
}
