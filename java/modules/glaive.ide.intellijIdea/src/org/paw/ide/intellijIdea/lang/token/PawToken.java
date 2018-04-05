package org.paw.ide.intellijIdea.lang.token;

import com.intellij.psi.tree.IElementType;
import org.paw.ide.intellijIdea.PawLanguage;

public class PawToken {
  public static IElementType COMMENT = new IElementType("COMMENT", PawLanguage.INSTANCE );
}
