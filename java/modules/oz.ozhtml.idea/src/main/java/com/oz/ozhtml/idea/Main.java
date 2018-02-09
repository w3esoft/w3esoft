package com.oz.ozhtml.idea;

import io.apigee.trireme.core.NodeEnvironment;
import io.apigee.trireme.core.NodeException;
import io.apigee.trireme.core.NodeScript;
import io.apigee.trireme.core.ScriptStatus;

import java.io.File;
import java.util.concurrent.ExecutionException;

public class Main {
    public static void main(String[] args) throws NodeException, ExecutionException, InterruptedException {
// The NodeEnvironment controls the environment for many scripts
        NodeEnvironment env = new NodeEnvironment();

// Pass in the script file name, a File pointing to the actual script, and an Object[] containg "argv"
        NodeScript script = env.createScript("my-test-script.js",
                new File("F:\\desenvolvimento\\w3esoft\\index.js"), null);

// Wait for the script to complete
        ScriptStatus status = script.execute().get();

// Check the exit code
        System.exit(status.getExitCode());

    }
}
