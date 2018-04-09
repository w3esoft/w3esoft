package org.juice.ide.intellijIdea;

import com.intellij.lang.Language;

public class JuiceLanguage extends Language {
    public static final JuiceLanguage INSTANCE = new JuiceLanguage();
    public JuiceLanguage() {
        super("JUICE", new String[]{"text/x-juice-source", "text/juice", "application/x-juice", "text/x-juice"});
    }
}
