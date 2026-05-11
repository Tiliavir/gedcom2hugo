package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// --------------------------------------------------------------------------
// extractYear
// --------------------------------------------------------------------------

func TestExtractYear_ISO(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{"2020-07-15", 2020},
		{"1969-07-20T00:00:00Z", 1969},
		{"1900", 1900},
		{"", 0},
		{"unknown", 0},
	}
	for _, tc := range cases {
		got := extractYear(tc.input)
		if got != tc.want {
			t.Errorf("extractYear(%q) = %d; want %d", tc.input, got, tc.want)
		}
	}
}

func TestExtractYear_GEDCOM(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{"1 JAN 1950", 1950},
		{"15 JUN 2020", 2020},
		{"MAR 1975", 1975},
		{"APR 2005", 2005},
	}
	for _, tc := range cases {
		got := extractYear(tc.input)
		if got != tc.want {
			t.Errorf("extractYear(%q) = %d; want %d", tc.input, got, tc.want)
		}
	}
}

// --------------------------------------------------------------------------
// historicalEventRepository.loadFromFile
// --------------------------------------------------------------------------

func TestLoadFromFile_Valid(t *testing.T) {
	events := []historicalEvent{
		{Date: "1969-07-20", Title: "Moon Landing", Location: "Moon", Source: "test"},
		{Date: "1989-11-09", Title: "Fall of Berlin Wall", Location: "Germany", Source: "test"},
	}
	data, _ := json.Marshal(events)

	tmpDir := t.TempDir()
	cachePath := filepath.Join(tmpDir, "events.json")
	if err := os.WriteFile(cachePath, data, 0644); err != nil {
		t.Fatalf("setup: %v", err)
	}

	repo := &historicalEventRepository{}
	if err := repo.loadFromFile(cachePath); err != nil {
		t.Fatalf("loadFromFile: %v", err)
	}
	if len(repo.events) != 2 {
		t.Errorf("expected 2 events, got %d", len(repo.events))
	}
}

func TestLoadFromFile_Missing(t *testing.T) {
	repo := &historicalEventRepository{}
	err := repo.loadFromFile("/nonexistent/path/events.json")
	if err == nil {
		t.Error("expected error for missing file, got nil")
	}
}

func TestLoadFromFile_InvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	cachePath := filepath.Join(tmpDir, "bad.json")
	if err := os.WriteFile(cachePath, []byte("not json"), 0644); err != nil {
		t.Fatalf("setup: %v", err)
	}

	repo := &historicalEventRepository{}
	if err := repo.loadFromFile(cachePath); err == nil {
		t.Error("expected error for invalid JSON, got nil")
	}
}

// --------------------------------------------------------------------------
// historicalEventRepository.eventsForPerson
// --------------------------------------------------------------------------

func sampleRepo() *historicalEventRepository {
	return &historicalEventRepository{
		events: []historicalEvent{
			{Date: "1960-01-01", Title: "Event Before Birth", Location: "Nowhere"},
			{Date: "1970-06-15", Title: "Event During Life", Location: "Somewhere"},
			{Date: "1985-03-10", Title: "Another During Life", Location: "Elsewhere"},
			{Date: "2025-12-31", Title: "Event After Death", Location: "Future"},
		},
	}
}

func TestEventsForPerson_WithinRange(t *testing.T) {
	repo := sampleRepo()
	events := repo.eventsForPerson(1969, 2000)
	if len(events) != 2 {
		t.Errorf("expected 2 events, got %d", len(events))
	}
	for _, e := range events {
		y := extractYear(e.Date)
		if y < 1969 || y > 2000 {
			t.Errorf("event %q with year %d is outside [1969,2000]", e.Title, y)
		}
	}
}

func TestEventsForPerson_NoBounds(t *testing.T) {
	repo := sampleRepo()
	events := repo.eventsForPerson(0, 0)
	if len(events) != len(repo.events) {
		t.Errorf("expected all %d events, got %d", len(repo.events), len(events))
	}
}

func TestEventsForPerson_ChronologicallySorted(t *testing.T) {
	repo := sampleRepo()
	events := repo.eventsForPerson(0, 0)
	for i := 1; i < len(events); i++ {
		if events[i].Date < events[i-1].Date {
			t.Errorf("events are not sorted: %q comes after %q", events[i-1].Date, events[i].Date)
		}
	}
}

func TestEventsForPerson_EmptyWhenNoMatch(t *testing.T) {
	repo := sampleRepo()
	events := repo.eventsForPerson(2100, 2200)
	if len(events) != 0 {
		t.Errorf("expected 0 events for future range, got %d", len(events))
	}
}

// --------------------------------------------------------------------------
// normalizeDate
// --------------------------------------------------------------------------

func TestNormalizeDate(t *testing.T) {
	cases := []struct {
		input, want string
	}{
		{"2020-07-15T00:00:00Z", "2020-07-15"},
		{"1900-01-01", "1900-01-01"},
		{"2000", "2000"},
		{"", ""},
		{"123", ""},
	}
	for _, tc := range cases {
		got := normalizeDate(tc.input)
		if got != tc.want {
			t.Errorf("normalizeDate(%q) = %q; want %q", tc.input, got, tc.want)
		}
	}
}

// --------------------------------------------------------------------------
// renderHistoricalEvents
// --------------------------------------------------------------------------

func TestRenderHistoricalEvents_Empty(t *testing.T) {
	html := renderHistoricalEvents(nil)
	if html != "" {
		t.Errorf("expected empty string for nil events, got %q", html)
	}

	html = renderHistoricalEvents([]historicalEvent{})
	if html != "" {
		t.Errorf("expected empty string for empty events, got %q", html)
	}
}

func TestRenderHistoricalEvents_ContainsExpectedStructure(t *testing.T) {
	events := []historicalEvent{
		{Date: "1969-07-20", Title: "Moon Landing", Location: "Moon", Source: "test"},
	}
	html := string(renderHistoricalEvents(events))

	checks := []string{
		`<details class="historical-events">`,
		`<summary`,
		`</details>`,
		`<table`,
		`<thead>`,
		`<tbody>`,
		"1969-07-20",
		"Moon Landing",
		"Moon",
	}
	for _, c := range checks {
		if !strings.Contains(html, c) {
			t.Errorf("rendered HTML missing expected fragment: %q", c)
		}
	}
}

func TestRenderHistoricalEvents_EscapesSpecialChars(t *testing.T) {
	events := []historicalEvent{
		{Date: "2000-01-01", Title: "<script>alert('xss')</script>", Location: "& Land"},
	}
	html := string(renderHistoricalEvents(events))

	if strings.Contains(html, "<script>") {
		t.Error("rendered HTML should escape <script> tag")
	}
	if !strings.Contains(html, "&amp;") || !strings.Contains(html, "&lt;") {
		t.Error("rendered HTML should contain HTML-escaped entities for & and <")
	}
}

func TestRenderHistoricalEvents_EmptyLocation(t *testing.T) {
	events := []historicalEvent{
		{Date: "2000-01-01", Title: "Some Event", Location: ""},
	}
	html := string(renderHistoricalEvents(events))
	// Empty location should show em-dash placeholder
	if !strings.Contains(html, "—") {
		t.Error("empty location should render as em-dash")
	}
}
