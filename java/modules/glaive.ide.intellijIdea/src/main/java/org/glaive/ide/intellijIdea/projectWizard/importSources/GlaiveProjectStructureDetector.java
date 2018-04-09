package org.glaive.ide.intellijIdea.projectWizard.importSources;

import com.intellij.ide.util.projectWizard.importSources.DetectedProjectRoot;
import com.intellij.ide.util.projectWizard.importSources.ProjectStructureDetector;
import org.jetbrains.annotations.NotNull;

import java.io.File;
import java.util.List;

public class GlaiveProjectStructureDetector extends ProjectStructureDetector {
    @NotNull
    @Override
    public DirectoryProcessingResult detectRoots(@NotNull File file, @NotNull File[] files, @NotNull File file1, @NotNull List<DetectedProjectRoot> list) {
        return null;
    }
}
