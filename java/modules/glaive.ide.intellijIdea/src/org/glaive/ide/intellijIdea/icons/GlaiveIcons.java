package org.glaive.ide.intellijIdea.icons;

import com.intellij.openapi.util.IconLoader;

import javax.swing.*;

public class GlaiveIcons {
    private static Icon load(String path) {
        return IconLoader.getIcon(path, GlaiveIcons.class);
    }
    public static final Icon C_Glaive = load("/icons/C_Logo.png");              //16x16
    public static final Icon E_Glaive = load("/icons/E_Logo.png");              //16x16
    public static final Icon Glaive_16 = load("/icons/GlaiveLogo_16.png");      //16x16
    public static final Icon Glaive_24 = load("/icons/GlaiveLogo_24.png");      //24x24
    public static final Icon I_Glaive = load("/icons/I_Logo.png");              //16x16
    public static final Icon Nmml_16 = load("/icons/nmml_16.png");              //16x16
    public static final Icon Method_Glaive = load("/icons/method-logo.png");    //16x16
    public static final Icon Field_Glaive = load("/icons/field-logo.png");      //16x16
}
