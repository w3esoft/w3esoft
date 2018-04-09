package org.glaive.ide.intellijIdea.fileTypes;

import com.intellij.openapi.fileTypes.FileTypeConsumer;
import com.intellij.openapi.fileTypes.FileTypeFactory;
import org.jetbrains.annotations.NotNull;

import static org.glaive.ide.intellijIdea.fileTypes.GlaiveLanguageFileType.DEFAULT_EXTENSION;

public class GlaiveFileTypeFactory extends FileTypeFactory {
    @Override
    public void createFileTypes(@NotNull FileTypeConsumer fileTypeConsumer) {
        fileTypeConsumer.consume(GlaiveLanguageFileType.INSTANCE, DEFAULT_EXTENSION);
    }
}
