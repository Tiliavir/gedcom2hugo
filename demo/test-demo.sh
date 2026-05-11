#!/bin/bash

# Test script for the demo
# 
# NOTE: Go tests (demo_test.go) are the recommended way to test the demo.
# This bash script is provided for manual validation and convenience.
# 
# To run Go tests instead:
#   cd demo && go test -v
#   or from project root: go test ./demo -v

echo "Testing gedcom2hugo demo..."
echo ""

# Get the directory where the script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

# Check if gedcom2hugo binary exists
if [ ! -f ../gedcom2hugo ]; then
    echo "Building gedcom2hugo..."
    (cd .. && go build -o gedcom2hugo main.go)
fi

echo "✓ gedcom2hugo binary found"
echo ""

# Verify GEDCOM file exists
if [ ! -f sample-family.ged ]; then
    echo "✗ Error: sample-family.ged not found"
    exit 1
fi

echo "✓ Sample GEDCOM file found"
echo ""

# Check generated content
echo "Checking generated content..."
PERSON_COUNT=$(find hugo-site/content/person -name "*.md" 2>/dev/null | wc -l)
FAMILY_COUNT=$(find hugo-site/content/family -name "*.md" 2>/dev/null | wc -l)
JSON_COUNT=$(find hugo-site/static/api/individual -name "*.json" 2>/dev/null | wc -l)

echo "  - Person pages: $PERSON_COUNT"
echo "  - Family pages: $FAMILY_COUNT"
echo "  - JSON API files: $JSON_COUNT"
echo ""

if [ $PERSON_COUNT -eq 6 ] && [ $FAMILY_COUNT -eq 2 ] && [ $JSON_COUNT -eq 6 ]; then
    echo "✓ All expected files generated"
else
    echo "✗ Unexpected file counts"
    exit 1
fi

# Verify JSON is valid
echo "Validating JSON files..."
for json_file in hugo-site/static/api/individual/*.json; do
    if ! python3 -m json.tool "$json_file" > /dev/null 2>&1; then
        echo "✗ Invalid JSON: $json_file"
        exit 1
    fi
done
echo "✓ All JSON files are valid"
echo ""

echo "Demo test completed successfully!"
echo ""
echo "To view the demo:"
echo "  cd hugo-site"
echo "  hugo server"
echo ""
echo "Or to regenerate content:"
echo "  ../gedcom2hugo -gedcom sample-family.ged -project hugo-site"
