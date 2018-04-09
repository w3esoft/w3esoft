package org.glaivescript.ide.intellijIdea.lang.token;

import com.intellij.psi.tree.IElementType;
import org.glaive.ide.intellijIdea.GlaiveLanguage;

public class GlaiveScriptToken {
  public static IElementType COMMENT = new IElementType("COMMENT", GlaiveLanguage.INSTANCE );
}
