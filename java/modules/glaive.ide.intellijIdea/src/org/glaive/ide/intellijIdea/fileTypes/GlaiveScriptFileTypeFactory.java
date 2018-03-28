package org.glaive.ide.intellijIdea.fileTypes;

import com.intellij.openapi.fileTypes.FileTypeConsumer;
import com.intellij.openapi.fileTypes.FileTypeFactory;
import org.jetbrains.annotations.NotNull;

import static org.glaive.ide.intellijIdea.fileTypes.GlaiveScriptLanguageFileType.DEFAULT_EXTENSION;


public class GlaiveScriptFileTypeFactory extends FileTypeFactory {
    @Override
    public void createFileTypes(@NotNull FileTypeConsumer fileTypeConsumer) {
        fileTypeConsumer.consume(GlaiveClassLanguageFileType.INSTANCE, DEFAULT_EXTENSION);
    }
}
