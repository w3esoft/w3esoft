package org.paw.ide.intellijIdea.fileTypes;

import com.intellij.openapi.fileTypes.LanguageFileType;
import com.intellij.openapi.vfs.CharsetToolkit;
import com.intellij.openapi.vfs.VirtualFile;
import org.paw.ide.intellijIdea.icons.PawIcons;
import org.jetbrains.annotations.NonNls;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
import org.paw.ide.intellijIdea.PawLanguage;

import javax.swing.*;

public class PawLanguageFileType extends LanguageFileType{
    public  static PawLanguageFileType INSTANCE = new PawLanguageFileType();
    public PawLanguageFileType(){
        super( PawLanguage.INSTANCE);
    }

    @NonNls
    public static final String DEFAULT_EXTENSION = "pw";

    @NotNull
    @Override
    public String getName() {
        return "Paw Toolkit Module";
    }

    @NotNull
    @Override
    public String getDescription() {
        return "A module that contains AMD modules written to take advantage of the Paw Toolkit and used by Needs More Paw. " +
                "The module's content root should be the root of your most top level packages so that they can be imported correctly.";

    }

    @NotNull
    @Override
    public String getDefaultExtension() {
        return DEFAULT_EXTENSION;
    }
    @Nullable
    @Override
    public Icon getIcon() {
        return PawIcons.Paw_16;
    }

    @Override
    public String getCharset(@NotNull VirtualFile file, byte[] content) {
        return CharsetToolkit.UTF8;
    }

}
