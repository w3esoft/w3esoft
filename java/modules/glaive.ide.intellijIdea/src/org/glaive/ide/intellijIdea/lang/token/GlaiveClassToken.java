package org.glaive.ide.intellijIdea.lang.token;

import com.intellij.psi.tree.IElementType;
import org.glaive.ide.intellijIdea.lang.GlaiveClassLanguage;
public class GlaiveClassToken {
  public static IElementType COMMENT = new IElementType("COMMENT", GlaiveClassLanguage.INSTANCE);
}
