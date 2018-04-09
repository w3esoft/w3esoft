package org.glaive.ide.intellijIdea.projectWizard;

import com.intellij.ide.util.projectWizard.ModuleBuilder;
import com.intellij.ide.util.projectWizard.ModuleBuilderListener;
import com.intellij.ide.util.projectWizard.SourcePathsBuilder;
import com.intellij.openapi.module.Module;
import com.intellij.openapi.module.ModuleType;
import com.intellij.openapi.options.ConfigurationException;
import com.intellij.openapi.roots.CompilerModuleExtension;
import com.intellij.openapi.roots.ContentEntry;
import com.intellij.openapi.roots.ModifiableRootModel;
import com.intellij.openapi.util.Pair;
import com.intellij.openapi.vfs.LocalFileSystem;
import com.intellij.openapi.vfs.VirtualFile;
import org.glaive.ide.intellijIdea.module.GlaiveModuleType;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.io.File;
import java.util.List;

public class GlaiveModuleBuilder  extends ModuleBuilder implements SourcePathsBuilder, ModuleBuilderListener {
    @Override
    public void setupRootModel(ModifiableRootModel modifiableRootModel) throws ConfigurationException {
        this.doAddContentEntry(modifiableRootModel);
    }

    @Nullable
    protected ContentEntry doAddContentEntry(ModifiableRootModel modifiableRootModel) {
        String contentEntryPath = this.getContentEntryPath();
        if (contentEntryPath == null) {
            return null;
        } else {
            (new File(contentEntryPath)).mkdirs();
            VirtualFile moduleContentRoot = LocalFileSystem.getInstance().refreshAndFindFileByPath(contentEntryPath.replace('\\', '/'));
            return moduleContentRoot == null ? null : modifiableRootModel.addContentEntry(moduleContentRoot);
        }
    }

    @Override
    public ModuleType getModuleType() {
        return GlaiveModuleType.getInstance();
    }

    @Override
    public void moduleCreated(@NotNull Module module) {
        final CompilerModuleExtension model = (CompilerModuleExtension) CompilerModuleExtension.getInstance(module).getModifiableModel(true);
        model.setCompilerOutputPath(model.getCompilerOutputUrl());
        model.inheritCompilerOutputPath(false);
        model.commit();
    }

    @Override
    public List<Pair<String, String>> getSourcePaths() throws ConfigurationException {
        return null;
    }

    @Override
    public void setSourcePaths(List<Pair<String, String>> list) {

    }

    @Override
    public void addSourcePath(Pair<String, String> pair) {

    }
}
