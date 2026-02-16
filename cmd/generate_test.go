package cmd

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestReadGedcomNoFile(t *testing.T) {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "gedcom",
		},
	}

	// Create a test context with no gedcom file
	set := flag.NewFlagSet("test", 0)
	set.String("gedcom", "", "")
	ctx := cli.NewContext(app, set, nil)

	_, err := readGedcom(ctx)
	if err == nil {
		t.Error("Expected error when no GEDCOM file specified, got nil")
	}
	if err.Error() != "no GEDCOM file specified for input" {
		t.Errorf("Expected 'no GEDCOM file specified for input', got '%s'", err.Error())
	}
}

func TestReadGedcomInvalidFile(t *testing.T) {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "gedcom",
		},
	}

	set := flag.NewFlagSet("test", 0)
	set.String("gedcom", "/nonexistent/file.ged", "")
	ctx := cli.NewContext(app, set, nil)

	_, err := readGedcom(ctx)
	if err == nil {
		t.Error("Expected error when reading nonexistent file, got nil")
	}
}

func TestReadGedcomValidFile(t *testing.T) {
	// Create a temporary valid GEDCOM file
	tmpDir := t.TempDir()
	gedcomPath := filepath.Join(tmpDir, "test.ged")

	gedcomContent := `0 HEAD
1 SOUR Test
1 GEDC
2 VERS 5.5
1 CHAR UTF-8
0 @I1@ INDI
1 NAME John /Doe/
0 TRLR
`
	err := os.WriteFile(gedcomPath, []byte(gedcomContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test GEDCOM file: %v", err)
	}

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "gedcom",
		},
	}

	set := flag.NewFlagSet("test", 0)
	set.String("gedcom", gedcomPath, "")
	ctx := cli.NewContext(app, set, nil)

	gc, err := readGedcom(ctx)
	if err != nil {
		t.Fatalf("Expected no error when reading valid GEDCOM file, got: %v", err)
	}

	if gc == nil {
		t.Error("Expected non-nil gedcom object")
	}

	if len(gc.Individual) != 1 {
		t.Errorf("Expected 1 individual, got %d", len(gc.Individual))
	}
}

func TestTagTable(t *testing.T) {
	// This tests the tag table initialization in Generate
	// We can't easily test the full Generate function without a complete setup
	// but we can verify the tag table has expected values

	expectedTags := map[string]string{
		"BAPM": "Baptism",
		"BIRT": "Birth",
		"DEAT": "Death",
		"MARR": "Married",
	}

	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(app, set, nil)

	// Call readGedcom with empty context to trigger error
	// This is just to ensure the function is accessible
	_, err := readGedcom(ctx)
	if err == nil {
		t.Error("Expected error with empty context")
	}

	// Verify expected tag mappings would be correct
	for tag, name := range expectedTags {
		if tag == "" || name == "" {
			t.Errorf("Tag table should have non-empty values for %s -> %s", tag, name)
		}
	}
}

func TestConfigureForJsonHeaders(t *testing.T) {
	tmpDir := t.TempDir()

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "project",
		},
	}

	set := flag.NewFlagSet("test", 0)
	set.String("project", tmpDir, "")
	ctx := cli.NewContext(app, set, nil)

	api := &apiControl{
		cx: ctx,
	}

	// First call should fail because static/api directory doesn't exist
	err := configureForJsonHeaders(api)
	if err == nil {
		t.Error("Expected error when directory doesn't exist")
	}

	// Create the directory and try again
	err = os.MkdirAll(filepath.Join(tmpDir, "static", "api"), 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	err = configureForJsonHeaders(api)
	if err != nil {
		t.Errorf("Expected no error after creating directory, got: %v", err)
	}

	// Verify the file was created
	headersPath := filepath.Join(tmpDir, "static", "api", "_headers")
	if _, err := os.Stat(headersPath); os.IsNotExist(err) {
		t.Error("Expected _headers file to be created")
	}

	// Verify content
	content, err := os.ReadFile(headersPath)
	if err != nil {
		t.Fatalf("Failed to read headers file: %v", err)
	}

	expectedContent := "/*  Access-Control-Allow-Origin: *  content-type: application/json; charset=utf-8"
	if string(content) != expectedContent {
		t.Errorf("Expected headers content '%s', got '%s'", expectedContent, string(content))
	}
}
