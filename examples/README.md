# Examples

This directory contains example GEDCOM files and demonstrations of gedcom2hugo functionality.

## demo.ged

A simple GEDCOM 5.5 file containing a small family tree with:

- **John Smith** (1950-2020) - Software Engineer
- **Mary Johnson** (b. 1952) - Teacher
- **Jane Smith** (b. 1975) - Doctor (daughter of John and Mary)
- **Robert Smith** (b. 1978) - Architect (son of John and Mary)

This demo file illustrates:
- Individual records with birth, death, and occupation information
- Family relationships (marriage, parent-child)
- Source citations
- Various event types

## Usage

To process this demo file:

```bash
# Create a test Hugo site
mkdir -p /tmp/demo-hugo
cd /tmp/demo-hugo

# Initialize a basic Hugo site structure
mkdir -p content data static

# Process the demo GEDCOM file
/path/to/gedcom2hugo -gedcom /path/to/gedcom2hugo/examples/demo.ged -project .

# View the generated files
echo "API files:"
ls -la static/api/individual/

echo "Content pages:"
ls -la content/individual/
```

## Expected Output

After processing `demo.ged`, you should see:

1. **JSON API files** in `static/api/`:
   - `individual/*.json` - Data for each person
   - `family/*.json` - Data for each family unit
   - `source/*.json` - Source citation data

2. **Hugo content files** in `content/`:
   - `individual/*.md` - Markdown pages for each person
   - `family/*.md` - Markdown pages for each family
   - `source/*.md` - Markdown pages for sources

Each file will contain structured data in Hugo front matter format with relevant genealogical information.
