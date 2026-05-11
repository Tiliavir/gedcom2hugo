# CHANGELOG

## [Unreleased] - 2026-02-16

### Added
- Go modules support (go.mod, go.sum) for proper dependency management
- Comprehensive README with:
  - Installation instructions
  - Detailed usage examples
  - Fork rationale and comparison with original project
  - GEDCOM version support information (5.5 supported, 7.0 future plans)
  - Demo section with examples
  - Output structure documentation
- Demo GEDCOM file (`examples/demo.ged`) with sample family tree
- Examples README with usage instructions
- Basic test coverage:
  - Main package tests
  - GEDCOM file reading tests
  - Configuration tests
- Bounds checking for array access to prevent panics
- Fallback handling for individuals without names
- Warning logs for individuals without names

### Changed
- Updated from deprecated cli package to cli/v2
- Replaced deprecated `io/ioutil.ReadFile` with `os.ReadFile`
- Improved error logging throughout:
  - Photo file opening/decoding errors now use log.Printf with context
  - Replaced generic fmt.Printf with proper log.Printf calls
- Simplified defer statements (removed unnecessary anonymous functions)
- Fixed file permissions from insecure 0777 to secure 0755
- Reorganized imports for consistency

### Fixed
- Critical nil check bug in `configureForJsonHeaders` that could cause incorrect behavior
- Error parameter shadowing in `configureForJsonHeaders`
- Potential panic from accessing Name[0] without bounds check
- Potential panic from accessing Media[0] without bounds check
- Insecure file permissions on created directories

### Security
- All dependencies scanned - no vulnerabilities found
- CodeQL security analysis passed with 0 alerts
- Fixed directory permissions from 0777 to 0755 (prevents unauthorized access)

## Project Goals

This fork aims to:
1. Modernize the codebase for current Go best practices
2. Fix bugs and improve code robustness
3. Add comprehensive documentation and examples
4. Provide test coverage for reliability
5. Plan for future GEDCOM 7.0 support

## Credits

Original project: [tektsu/gedcom2hugo](https://github.com/tektsu/gedcom2hugo)
GEDCOM parser: [iand/gedcom](https://github.com/iand/gedcom)
