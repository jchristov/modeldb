// THIS FILE IS AUTO-GENERATED. DO NOT EDIT
package ai.verta.modeldb.versioning.autogenerated._public.modeldb.versioning.model;

import ai.verta.modeldb.versioning.*;
import ai.verta.modeldb.versioning.blob.diff.*;
import com.pholser.junit.quickcheck.generator.*;
import com.pholser.junit.quickcheck.generator.java.util.*;
import com.pholser.junit.quickcheck.random.*;
import java.util.*;

public class AutogenCodeBlobGen extends Generator<AutogenCodeBlob> {
  public AutogenCodeBlobGen() {
    super(AutogenCodeBlob.class);
  }

  @Override
  public AutogenCodeBlob generate(SourceOfRandomness r, GenerationStatus status) {
    // if (r.nextBoolean())
    //     return null;

    AutogenCodeBlob obj = new AutogenCodeBlob();
    if (r.nextBoolean()) {
      obj.setGit(Utils.removeEmpty(gen().type(AutogenGitCodeBlob.class).generate(r, status)));
    }
    if (r.nextBoolean()) {
      obj.setNotebook(
          Utils.removeEmpty(gen().type(AutogenNotebookCodeBlob.class).generate(r, status)));
    }
    return obj;
  }
}
