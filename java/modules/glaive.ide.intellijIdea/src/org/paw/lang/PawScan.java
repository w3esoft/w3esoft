package org.paw.lang;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.nio.charset.StandardCharsets;
import java.util.LinkedList;

public class PawScan {
    private final InputStream input;
    private int position=0;
    LinkedList<Integer> cacheChar = new LinkedList();
    public PawScan(CharSequence charSequence) {
        this(new ByteArrayInputStream(charSequence.toString().getBytes(StandardCharsets.UTF_8)));
    }
    public PawScan(InputStream input) {
        this.input = input;
    }
    public int read() {
        if (cacheChar.size()>0){
           return cacheChar.pollLast();
        }
        try {
            int ii = input.read();


            if (ii==-1){

                return 0;

            }

            position++;
            return  ii;
        } catch (IOException e) {
            e.printStackTrace();
        }
        return 0;
    }
    public int getPosition() {
        return position - cacheChar.size();
    }
    public void addCache(int c) {
        cacheChar.add(c);
    }
}
