# Go Pages

## Template System

There are three fundamental directories that need to be specified.

- `template` directory - specifies where the command will look for templates to use when compiling, see [templates](#templates)
- `content` directory - specifies from where Markdown files will be read from, see [content](#content)
- `output` directory - specifies where html files will be generated

### Templates

The template system looks at files with the `gohtml` extension.

### Content

The content displayed is based off of markdown files that are present in the `content` directory

~~~
The command will read the files in the template directory specified (or `templates`) and will perform actions based on them.

Special files:

- a file named 'index.html' when found in the `content` directory will be copied over directly to the output directory.

The program made here (currently unmade) will compile the pages in the `templates` folder and write to the `docs` folder.
Template system will be made using [golang html templates](https://pkg.go.dev/html/template)

One thing I have to test:
Can you reference media outside of the docs folder in the docs folder, and still have it display said images?
If so then I can safely delete any files within docs that aren't html pages.
Otherwise I will have to be a bit smarter with my code.

## Plans

- the CLI will:
  - read specified directories for:
    - templates
    - content
  - it will generate a file for each and every file in `content`, using the template named after the folder
  - it will have special names for specific types of templates that will be handled differently
  - given two directories with the same internal directory structure, eg:
   content -> bunnies
   templates -> bunnies
    - the program will therefore read the template directory, and look for special templates:
    - types of special templates:
      - index.html
      -
