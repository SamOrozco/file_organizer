# File Organizer

File organizer is a CLI application that will sit in the background, watch a directory and organize new files that are added as they are added.
This is typically used as a "service" in linux or "launchd" on mac os. 

```
Watcher a directory and organizes your files. Defaults to organizing by type

Usage:
fo [flags]

Flags:
-x, --exclude-pattern string      if set will exclude file from being organized if it matches regex pattern
-h, --help                        help for fo
-d, --organize-by-date-and-type   this will organize a parent folder into sub folders of dates added and then sub-sub folders of type
-t, --organize-by-type            this will organize a parent folder into sub folders based on file extension
-l, --organize-by-type-and-date   this will organize a parent folder into sub folders of types and then sub-sub folders of date added
-m, --sample-millis int           how often the watcher will check for new files and move them (default 1000)
-f, --type-folder-suffix string   for the type created folders this will be the suffix (default "_files")
```

## Organize by type
Organizing the directory by type or `-t` will put all files in the folder into sub folders based on their file extension. 
If the file extension is empty or unknown the value "unknown" will be used. 

## Organize by date and type
Organizing by date and type or `-d` will put all files into a date sub folder(for the date they are moved) and then sub folders in that
for the extension type. 

## Organize by type and date
Organizing by type and date or `-l` will put all files into a type sub folder and then put the files further into sub folders for the date they are moved.

