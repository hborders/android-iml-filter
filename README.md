androidIml
==================

A git clean filter that normalizes innocuous changes (dumb maven paths, repeated libraries, newlines) from iml files

Currently only dumb maven exclude-folder paths are supported. 

```xml
<excludeFolder url="file://$MODULE_DIR$/target/surefire-reports" />
<excludeFolder url="file://$MODULE_DIR$/target/idea-classes" />
<excludeFolder url="file://$MODULE_DIR$/target/idea-test-classes" />
```

Please open an issue (preferably with an iml file diff) if you encounter any iml file changes that IntelliJ or Android Studio repeatedly creates and then deletes.

# Installation
Add the following to your .git/config:

```
[filter "androidIml"]
        clean = androidIml
```

Add the following to your root .gitattributes:

```
*.iml filter=androidIml
```

# License

This filter is licensed under the GPLv3. Since you'll only run it locally, and it won't actually be part of your source, it won't virally infect any of your real projects. Don't freak out. If I'm wrong about this, I'm reasonable and I'll change it. Open an issue about it.
