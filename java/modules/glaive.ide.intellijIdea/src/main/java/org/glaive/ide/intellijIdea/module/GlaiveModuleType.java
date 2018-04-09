package org.glaive.ide.intellijIdea.module;

import com.intellij.openapi.module.ModuleType;
import com.intellij.openapi.module.ModuleTypeManager;
import org.glaive.ide.intellijIdea.icons.GlaiveIcons;
import org.glaive.ide.intellijIdea.projectWizard.GlaiveModuleBuilder;
import org.jetbrains.annotations.NonNls;
import org.jetbrains.annotations.NotNull;

import javax.swing.*;

public class GlaiveModuleType extends ModuleType<GlaiveModuleBuilder> {

    @NonNls
    private static final String ID = "GLAIVE_MODULE";

    public static GlaiveModuleType getInstance() {
        return (GlaiveModuleType) ModuleTypeManager.getInstance().findByID(ID);
    }

    public GlaiveModuleType() {
        super(ID);
    }

    @NotNull
    @Override
    public GlaiveModuleBuilder createModuleBuilder() {
        return new GlaiveModuleBuilder();
    }

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

    public Icon getBigIcon() {
        return GlaiveIcons.Glaive_24;
    }

    @Override
    public Icon getNodeIcon(boolean b) {
        return GlaiveIcons.Glaive_16;
    }
}
