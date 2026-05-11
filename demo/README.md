# Demo Hugo Site

This directory contains a complete demonstration of gedcom2hugo in action.

## Contents

- `sample-family.ged` - A sample GEDCOM file containing a fictional family tree with 6 individuals
- `hugo-site/` - A complete Hugo site with generated genealogy content
- `demo_test.go` - Go tests to validate the demo (recommended)
- `test-demo.sh` - Bash script for manual validation

## Testing the Demo

### Using Go Tests (Recommended)

Run the Go tests to verify the demo is working correctly:

```bash
cd demo
go test -v
```

Or from the project root:

```bash
go test ./demo -v
```

The tests verify:
- All expected files exist
- Person and family pages are generated
- JSON API files are valid and well-formed
- JSON structure contains expected fields

### Using the Bash Script

Alternatively, you can use the bash script for manual validation:

```bash
cd demo
./test-demo.sh
```

## The Sample Family Tree

The demo includes the Smith family:

**Generation 1:**
- John Smith (1950-2020) - Engineer
- Mary Johnson (b. 1952)

**Generation 2:**
- Robert Smith (b. 1975) - Software Developer, married to Sarah Brown
  - Emily Smith (b. 2005)
  - James Smith (b. 2008)

## How the Demo Was Generated

1. The sample GEDCOM file was created with fictional family data
2. The `gedcom2hugo` tool was run with:
   ```bash
   ./gedcom2hugo -gedcom demo/sample-family.ged -project demo/hugo-site
   ```
3. This generated:
   - Markdown content files in `content/person/` and `content/family/`
   - JSON API files in `static/api/individual/` and `static/api/family/`
   - Supporting files for the Hugo site

## Viewing the Demo

### Option 1: Using Hugo (Recommended)

If you have Hugo installed:

```bash
cd demo/hugo-site
hugo server
```

Then open http://localhost:1313 in your browser.

## Hosting on GitHub Pages

This repository includes a workflow at `.github/workflows/demo-pages.yml` that builds and deploys
`demo/hugo-site` to GitHub Pages.

Required GitHub repo setting:

1. Open repository settings.
2. Go to `Pages`.
3. Set `Build and deployment` source to `GitHub Actions`.

Deployment behavior:

- The workflow runs on pushes to `master` (and manual dispatch).
- It builds with a repository-aware base URL:
   `https://<owner>.github.io/<repo>/`
- It publishes the generated `demo/hugo-site/public` directory.

### Option 2: Using Python's HTTP Server

If you don't have Hugo but have Python:

```bash
cd demo/hugo-site
hugo  # Build the site first (if Hugo is installed)
# OR if Hugo is not installed:
python3 -m http.server 8000
```

Then open http://localhost:8000 in your browser.

Note: Without Hugo, you'll need to navigate directly to the generated files.

### Option 3: View Generated Files Directly

You can explore the generated content:

- **Markdown files**: `demo/hugo-site/content/person/*.md` and `demo/hugo-site/content/family/*.md`
- **JSON data**: `demo/hugo-site/static/api/individual/*.json` and `demo/hugo-site/static/api/family/*.json`

## Site Structure

```
hugo-site/
├── config.toml              # Hugo configuration
├── content/                 # Generated content
│   ├── _index.md           # Home page
│   ├── person/             # Individual person pages
│   └── family/             # Family unit pages
├── static/
│   ├── api/                # JSON API files
│   │   ├── individual/     # Person data
│   │   └── family/         # Family data
│   ├── css/
│   │   └── style.css       # Styling
│   └── js/
│       ├── jquery.min.js
│       ├── idrisutil.js
│       ├── individualdisplay.js
│       └── familydisplay.js
└── themes/
    └── genealogy/          # Custom theme
        └── layouts/        # Hugo templates
```

## What You'll See

When you view the demo site, you can:

1. **Browse individuals** - Click on any person to see their:
   - Life events (birth, death, marriage, etc.)
   - Parents and children
   - Spouse(s) and family relationships

2. **Explore families** - View family units showing:
   - Parents (husband and wife)
   - Marriage events
   - Children in the family

3. **View raw data** - Each page includes a collapsible section showing the raw JSON data structure

## Customizing the Demo

You can modify the demo by:

1. **Editing the GEDCOM file** (`sample-family.ged`) to add more people or events
2. **Customizing the theme** in `hugo-site/themes/genealogy/`
3. **Updating styles** in `hugo-site/static/css/style.css`
4. **Regenerating content** by running gedcom2hugo again

## Learn More

- See the main README.md for more information about gedcom2hugo
- The GEDCOM 5.5 standard is used for the sample data
- Hugo documentation: https://gohugo.io/documentation/
