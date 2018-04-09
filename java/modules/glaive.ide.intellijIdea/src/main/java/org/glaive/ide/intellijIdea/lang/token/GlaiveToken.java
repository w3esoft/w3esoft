package org.glaive.ide.intellijIdea.lang.token;

import com.intellij.psi.tree.IElementType;
import org.glaive.ide.intellijIdea.GlaiveLanguage;
public class GlaiveToken {
  public static IElementType COMMENT = new IElementType("COMMENT", GlaiveLanguage.INSTANCE);
}
