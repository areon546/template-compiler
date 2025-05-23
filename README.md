# Go Pages

The program made here (currently unmade) will compile the pages in the `templates` folder and write to the `docs` folder.
Template system will be made using [golang html templates](https://pkg.go.dev/html/template)

It will use the `rankings.csv` file to add items to the database, and from the database it will construct the `rankings.html` page.

One thing I have to test:
Can you reference media outside of the docs folder in the docs folder, and still have it display said images?
If so then I can safely delete any files within docs that aren't html pages.
Otherwise I will have to be a bit smarter with my code.

## Plans

- [ ] setup sqlite database in database.db
- [ ] write program to compile html pages based off of the templates in the `templates` folder

- the CLI will:
  - read specified directories for:
    - templates
    - content
  - it will generate a file for each and every file in `content`, using the template named after the folder
  - it will have special names for specific types of templates that will be handled differently
  - given two directories with the same internal directory structure, eg:
   content -> bunnies
   templates -> bunnies
    - the program will thereforth read the template directory, and look for special templates:
    - types of special templates:
      - index.html
      -
