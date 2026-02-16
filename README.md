# gedcom2hugo

Generate Hugo content from a GEDCOM 5.5 file.

## Usage

```bash
$ go run . -gedcom /path/to/family.ged -project /path/to/target/
```

```bash
$ go build
$ ./gedcom2hugo -gedcom /path/to/family.ged -project /path/to/target/
```

## Demo

A complete working demo is available in the `/demo` directory! This demonstrates how gedcom2hugo converts GEDCOM genealogy data into a Hugo static site.

The demo includes:
- A sample GEDCOM file with a fictional family tree (6 individuals across 2 generations)
- A complete Hugo site with generated content
- Custom theme and styling for displaying genealogical data
- Interactive JavaScript-based display of family relationships

To try the demo:

```bash
# Generate the demo content (already done, but you can regenerate)
./gedcom2hugo -gedcom demo/sample-family.ged -project demo/hugo-site

# View the site with Hugo
cd demo/hugo-site
hugo server
```

See `/demo/README.md` for detailed information about the demo and how to use it.

## Contribution
Contribution is highly welcome! There is just one rule: use `go fmt` before committing. I don't want to discuss code
style, it's boring. There is a standard, follow it.

## Credits
This project is based on [the implementation](https://github.com/tektsu/gedcom2hugo) of [@tektsu](https://github.com/tektsu/)
done for his [idrisproject](https://www.idrisproject.com/) and uses the [gedcom parser](https://github.com/iand/gedcom)
from [@iand](https://github.com/iand/).
