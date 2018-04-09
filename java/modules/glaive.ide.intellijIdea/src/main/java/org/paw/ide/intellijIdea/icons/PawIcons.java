package org.paw.ide.intellijIdea.icons;

import com.intellij.openapi.util.IconLoader;
import org.glaive.ide.intellijIdea.icons.GlaiveIcons;

import javax.swing.*;

public class PawIcons {
    private static Icon load(String path) {
        return IconLoader.getIcon(path, GlaiveIcons.class);
    }

    public static final Icon C_Paw = load("/icons/C_Logo.png");              //16x16
    public static final Icon E_Paw = load("/icons/E_Logo.png");              //16x16
    public static final Icon Paw_16 = load("/icons/GlaiveLogo_16.png");      //16x16
    public static final Icon Paw_24 = load("/icons/GlaiveLogo_24.png");      //24x24
    public static final Icon I_Paw = load("/icons/I_Logo.png");              //16x16
    public static final Icon Nmml_16 = load("/icons/nmml_16.png");              //16x16
    public static final Icon Method_Paw = load("/icons/method-logo.png");    //16x16
    public static final Icon Field_Paw = load("/icons/field-logo.png");      //16x16
}
