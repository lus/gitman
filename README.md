# gitman
A CLI which allows you to use presets as a Git config for your current repository

---

### 1.) Installation
1. Download the latest executable from the [releases tab](https://github.com/Lukaesebrot/gitman/releases)
2. Put the executable in a directory which is part of your `PATH` environment variable
3. Restart your PC (Sometimes it is enough to restart your console window)

### 2.) Usage
GitMan currently supports two usage modes:
1. Environment mode:
   ```
   1.) Set the 'GITMAN_PRESETS' environment variable to the directory where you will save your presets
   2.) Execute GitMan like this: 'gitman --name=PRESET_NAME'
   
   GitMan will automatically look for a preset file named 'PRESET_NAME.txt' inside the specified folder. 
   
   Tip: If you execute GitMan like this: 'gitman' (without any parameters),
   then it will search for a file named 'default.txt' inside your preset directory.
   ```
2. Standalone mode:
   ```
   1.) Execute GitMan like this: 'gitman --path=path/to/preset.txt'
   
   GitMan will use the given file as preset file.
   ```
   
### 3.) File syntax
A simple preset file will look like this:
```
user.name||Your Name
user.email||email@example.com
test.value||test value here
```
I think this is self-explanatory.