package org.buildcraft.ide.intellijIdea.fileTypes;

import com.intellij.openapi.fileTypes.LanguageFileType;
import com.intellij.openapi.vfs.CharsetToolkit;
import com.intellij.openapi.vfs.VirtualFile;
import org.glaive.ide.intellijIdea.GlaiveLanguage;
import org.glaive.ide.intellijIdea.icons.GlaiveIcons;
import org.jetbrains.annotations.NonNls;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import javax.swing.*;

public class BuildCraftLanguageFileType extends LanguageFileType{
    public  static BuildCraftLanguageFileType INSTANCE = new BuildCraftLanguageFileType();
    public BuildCraftLanguageFileType(){
        super(GlaiveLanguage.INSTANCE);
    }

    @NonNls
    public static final String DEFAULT_EXTENSION = "bc";

    @NotNull
    @Override
    public String getName() {
        return "Glaive Toolkit Module";
    }

    @NotNull
    @Override
    public String getDescription() {
        return "A module that contains AMD modules written to take advantage of the Glaive Toolkit and used by Needs More Glaive. " +
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
        return GlaiveIcons.Glaive_16;
    }

    @Override
    public String getCharset(@NotNull VirtualFile file, byte[] content) {
        return CharsetToolkit.UTF8;
    }

}
