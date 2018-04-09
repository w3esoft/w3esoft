package org.glaive.ide.intellijIdea.components;

import com.intellij.openapi.components.AbstractProjectComponent;
import com.intellij.openapi.project.Project;
import org.glaive.ide.intellijIdea.impl.components.GlaiveProjectComponentImpl;

public class GlaiveProjectComponent extends AbstractProjectComponent implements GlaiveProjectComponentImpl {
    protected GlaiveProjectComponent(Project project) {
        super(project);
    }
}
