# gedcom2hugo

Generate Hugo content from a GEDCOM 5.5 file.

## About This Fork

This repository is a fork of [tektsu/gedcom2hugo](https://github.com/tektsu/gedcom2hugo) by [@tektsu](https://github.com/tektsu/).

**Reason for forking:** This fork was created to modernize the codebase, fix bugs, add better documentation, tests, and examples. The original project provided a solid foundation, but this fork aims to make it more maintainable and easier to use with modern Go tooling and better feature support.

## Features

- Converts GEDCOM 5.5 genealogical data files to Hugo static site content
- Generates individual pages for people, families, sources, and photos
- Creates JSON API endpoints for dynamic access to genealogical data
- Supports Hugo static site generation for publishing family history online
- Handles relationships, events, citations, and media

## GEDCOM Version Support

**Currently Supported:** GEDCOM 5.5

**GEDCOM 7.0 Status:** The underlying parser ([iand/gedcom](https://github.com/iand/gedcom)) currently supports GEDCOM 5.5/5.5.1. GEDCOM 7.0 introduces significant changes including embedded media, enhanced Unicode support, structured citations, and change tracking. Support for GEDCOM 7.0 would require either waiting for the parser to be updated or switching to a GEDCOM 7.0-compatible parser (e.g., [funwithbots/go-gedcom](https://github.com/funwithbots/go-gedcom)).

## Installation

### Prerequisites

- Go 1.20 or higher
- A GEDCOM 5.5 file
- Hugo (for generating the final static site)

### Install from source

```bash
git clone https://github.com/Tiliavir/gedcom2hugo.git
cd gedcom2hugo
go build
```

## Usage

### Basic Usage

```bash
# Run directly with go
go run . -gedcom /path/to/family.ged -project /path/to/hugo-project/

# Or build and run the binary
go build
./gedcom2hugo -gedcom /path/to/family.ged -project /path/to/hugo-project/
```

### Command-Line Options

```bash
  -gedcom, -g     Specify the input GEDCOM file (required)
  -project, -p    Specify the top level directory of the Hugo project (default: ".")
  -verbose, -v    Enable verbose output
  -debug, -d      Enable debugging output
  -version, -V    Print version information
```

### Example

```bash
# Process a GEDCOM file and output to Hugo project
./gedcom2hugo -gedcom examples/royal.ged -project /path/to/my-hugo-site/

# This will create:
# - /path/to/my-hugo-site/data/api/*.json (API data files)
# - /path/to/my-hugo-site/content/individual/*.md (individual pages)
# - /path/to/my-hugo-site/content/family/*.md (family pages)
# - /path/to/my-hugo-site/content/source/*.md (source pages)
# - /path/to/my-hugo-site/content/photo/*.md (photo pages)
```

## Demo

A simple demo is available in the `examples/` directory:

1. **Sample GEDCOM file:** `examples/demo.ged` - A small example family tree
2. **Generated output:** See how the tool transforms GEDCOM data into Hugo content

To try the demo:

```bash
# Create a test Hugo site
mkdir -p /tmp/demo-hugo
cd /tmp/demo-hugo
hugo new site .

# Process the demo GEDCOM file
/path/to/gedcom2hugo -gedcom /path/to/gedcom2hugo/examples/demo.ged -project .

# View the generated files
ls -la data/api/
ls -la content/individual/
```

## How It Works

1. **Parse GEDCOM:** Reads and parses the input GEDCOM file using the [iand/gedcom](https://github.com/iand/gedcom) parser
2. **Build Data Structures:** Processes individuals, families, sources, and media objects
3. **Generate JSON APIs:** Creates JSON files in `static/api/` for dynamic access
4. **Create Markdown Pages:** Generates Hugo-compatible markdown files with front matter in `content/`
5. **Configure Headers:** Sets up CORS headers for API access

## Output Structure

The tool generates the following structure in your Hugo project:

```
your-hugo-project/
├── content/
│   ├── individual/     # Individual person pages
│   ├── family/         # Family unit pages
│   ├── source/         # Source citation pages
│   └── photo/          # Photo/media pages
├── static/
│   └── api/
│       ├── individual/ # JSON API for individuals
│       ├── family/     # JSON API for families
│       ├── source/     # JSON API for sources
│       ├── photo/      # JSON API for photos
│       └── _headers    # CORS configuration
```

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build
```

### Code Formatting

This project follows standard Go formatting conventions:

```bash
go fmt ./...
```

## Demo

A complete working demo is available in the `/demo` directory! This demonstrates how gedcom2hugo converts GEDCOM genealogy data into a Hugo static site.

The demo includes:
- A sample GEDCOM file with a fictional family tree (6 individuals across 2 generations)
- A complete Hugo site with generated content
- Custom theme and styling for displaying genealogical data
- Interactive JavaScript-based display of family relationships
- Go tests to validate the demo functionality

To try the demo:

```bash
# Run the tests to verify everything works
go test ./demo -v

# Generate the demo content (already done, but you can regenerate)
./gedcom2hugo -gedcom demo/sample-family.ged -project demo/hugo-site

# View the site with Hugo
cd demo/hugo-site
hugo server
```

See `/demo/README.md` for detailed information about the demo and how to use it.

## Contribution

Contribution is highly welcome! There is just one rule: use `go fmt` before committing. I don't want to discuss code style, it's boring. There is a standard, follow it.

## Credits

This project is based on [the implementation](https://github.com/tektsu/gedcom2hugo) of [@tektsu](https://github.com/tektsu/) done for his [idrisproject](https://www.idrisproject.com/) and uses the [gedcom parser](https://github.com/iand/gedcom) from [@iand](https://github.com/iand/).

## License

MIT License - see LICENSE file for details
