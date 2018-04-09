package org.paw.ide.intellijIdea;
        import com.intellij.lang.Language;
public class PawLanguage extends  Language {
    public static final PawLanguage INSTANCE = new PawLanguage();
    protected PawLanguage() {
        super("PAW");
    }
}
