package demo_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// TestDemoFilesExist verifies that all expected demo files are present
func TestDemoFilesExist(t *testing.T) {
	demoDir := "."

	// Check GEDCOM file
	gedcomPath := filepath.Join(demoDir, "sample-family.ged")
	if _, err := os.Stat(gedcomPath); os.IsNotExist(err) {
		t.Errorf("Sample GEDCOM file not found: %s", gedcomPath)
	}

	// Check Hugo site directory
	hugoSiteDir := filepath.Join(demoDir, "hugo-site")
	if _, err := os.Stat(hugoSiteDir); os.IsNotExist(err) {
		t.Errorf("Hugo site directory not found: %s", hugoSiteDir)
	}
}

// TestGeneratedPersonPages verifies that person pages were generated correctly
func TestGeneratedPersonPages(t *testing.T) {
	personDir := filepath.Join("hugo-site", "content", "person")

	expectedPersons := []string{"i1.md", "i2.md", "i3.md", "i4.md", "i5.md", "i6.md"}

	for _, person := range expectedPersons {
		personPath := filepath.Join(personDir, person)
		if _, err := os.Stat(personPath); os.IsNotExist(err) {
			t.Errorf("Person page not found: %s", personPath)
		}
	}
}

// TestGeneratedFamilyPages verifies that family pages were generated correctly
func TestGeneratedFamilyPages(t *testing.T) {
	familyDir := filepath.Join("hugo-site", "content", "family")

	expectedFamilies := []string{"f1.md", "f2.md"}

	for _, family := range expectedFamilies {
		familyPath := filepath.Join(familyDir, family)
		if _, err := os.Stat(familyPath); os.IsNotExist(err) {
			t.Errorf("Family page not found: %s", familyPath)
		}
	}
}

// TestJSONAPIFiles verifies that JSON API files are valid
func TestJSONAPIFiles(t *testing.T) {
	individualDir := filepath.Join("hugo-site", "static", "api", "individual")

	expectedIndividuals := []string{"i1.json", "i2.json", "i3.json", "i4.json", "i5.json", "i6.json"}

	for _, individual := range expectedIndividuals {
		jsonPath := filepath.Join(individualDir, individual)

		// Check file exists
		if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
			t.Errorf("JSON file not found: %s", jsonPath)
			continue
		}

		// Validate JSON format
		data, err := os.ReadFile(jsonPath)
		if err != nil {
			t.Errorf("Failed to read JSON file %s: %v", jsonPath, err)
			continue
		}

		var jsonData map[string]interface{}
		if err := json.Unmarshal(data, &jsonData); err != nil {
			t.Errorf("Invalid JSON in file %s: %v", jsonPath, err)
		}
	}
}

// TestJSONAPIStructure verifies that JSON files have expected structure
func TestJSONAPIStructure(t *testing.T) {
	jsonPath := filepath.Join("hugo-site", "static", "api", "individual", "i1.json")

	data, err := os.ReadFile(jsonPath)
	if err != nil {
		t.Fatalf("Failed to read JSON file: %v", err)
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		t.Fatalf("Invalid JSON: %v", err)
	}

	// Check for expected top-level fields
	expectedFields := []string{"id", "ref", "name", "events"}
	for _, field := range expectedFields {
		if _, ok := jsonData[field]; !ok {
			t.Errorf("Expected field '%s' not found in JSON", field)
		}
	}
}

// TestFamilyJSONFiles verifies that family JSON API files are valid
func TestFamilyJSONFiles(t *testing.T) {
	familyDir := filepath.Join("hugo-site", "static", "api", "family")

	expectedFamilies := []string{"f1.json", "f2.json"}

	for _, family := range expectedFamilies {
		jsonPath := filepath.Join(familyDir, family)

		// Check file exists
		if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
			t.Errorf("Family JSON file not found: %s", jsonPath)
			continue
		}

		// Validate JSON format
		data, err := os.ReadFile(jsonPath)
		if err != nil {
			t.Errorf("Failed to read family JSON file %s: %v", jsonPath, err)
			continue
		}

		var jsonData map[string]interface{}
		if err := json.Unmarshal(data, &jsonData); err != nil {
			t.Errorf("Invalid JSON in family file %s: %v", jsonPath, err)
		}
	}
}
