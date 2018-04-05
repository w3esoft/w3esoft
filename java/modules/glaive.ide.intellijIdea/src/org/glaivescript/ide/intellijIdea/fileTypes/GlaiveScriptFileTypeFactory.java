package org.glaivescript.ide.intellijIdea.fileTypes;

import com.intellij.openapi.fileTypes.FileTypeConsumer;
import com.intellij.openapi.fileTypes.FileTypeFactory;
import org.jetbrains.annotations.NotNull;
public class GlaiveScriptFileTypeFactory extends FileTypeFactory {
    @Override
    public void createFileTypes(@NotNull FileTypeConsumer fileTypeConsumer) {
        fileTypeConsumer.consume(GlaiveScriptLanguageFileType.INSTANCE, GlaiveScriptLanguageFileType.DEFAULT_EXTENSION);
    }
}
