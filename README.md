### Inconsistencies README vs Codebase

- currently does not have a blacklist feature
- log file has a default that writes to a file
- log file currently doesn't write to set logfile

# Template Compiler

A simple template system meant to be used both as a go package and as a terminal executable / command.

The command has various flags, primarily:

- t : Used to specify the directory of the template directory.
- c : Used to specify the directory of the content directory.
- o : Used to specify the directory of the output directory, where compiled HTML files are written to.
- s : Used to specify the file suffix for template files.
- l : Used to specify the name of the log file. By default, leaving it blank means that no log file will be generated.

You do not want to share any suffix namespace between your templates and content files, since there will be no compilation performed on the template files. The program does support having the template and content directories as the same folder, however if you want to call template files html, you must have the template and content folders be distinct.

## System

The system will read from a directory structure, and insert markdown files (after formatting them to HTML), into template files, and write to files in the output directory.

Example:

~~~markdown
content
- index.md
- page2.md
- htmlPage.html
- subfolder/index.md 
- static/global.css

template
- template.tpl
- subfolder/template.tpl
~~~

The above will result in the following output directory.

~~~markdown
docs 
- index.html
- page2.html
- htmlPage.html
- subfolder/index.html
- static/global.css
~~~

Things to note:

- MD files are inserted into the corresponding template within the specified directory.
  - EG a markdown file in the root, is placed into the template file at the root.
  - A markdown file in subfolder10 will be inserted into the corresponding template in subfolder10.
- HTML files are not inserted, they are simply copied straight over.
  - If you want to have them inserted into template files, you can rename them to .md or .markdown files.
- Any miscaleneous files will also be transferred over, unless they are in the specified blacklist.
  - Currently, the file types copied over by default are: jpg jpeg png webp css js
