package org.paw.lang;

import java.io.ByteArrayInputStream;
import java.io.InputStream;
import java.nio.charset.StandardCharsets;

public class PawScan {
    public PawScan(CharSequence charSequence) {
        this(new ByteArrayInputStream(charSequence.toString().getBytes(StandardCharsets.UTF_8)));
    }

    public PawScan(InputStream inputStream) {

    }
}
