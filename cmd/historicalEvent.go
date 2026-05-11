package cmd

import (
	"encoding/json"
	"os"
	"sort"
	"strconv"
	"strings"
)

// historicalEvent represents a single historical event to be displayed alongside a person's record.
type historicalEvent struct {
	Date        string `json:"date"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Source      string `json:"source"`
}

// historicalEventRepository holds loaded events and provides filtering for individual persons.
type historicalEventRepository struct {
	events []historicalEvent
}

// loadFromFile reads historical events from a JSON cache file.
func (r *historicalEventRepository) loadFromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &r.events)
}

// eventsForPerson returns events that fall within the person's lifespan,
// sorted chronologically. If birthYear or deathYear are 0 the respective
// bound is not applied.
func (r *historicalEventRepository) eventsForPerson(birthYear, deathYear int) []historicalEvent {
	var result []historicalEvent
	for _, e := range r.events {
		year := extractYear(e.Date)
		if year <= 0 {
			continue
		}
		if birthYear > 0 && year < birthYear {
			continue
		}
		if deathYear > 0 && year > deathYear {
			continue
		}
		result = append(result, e)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Date < result[j].Date
	})
	return result
}

// extractYear parses a year from a date string.
// Supported formats: "YYYY", "YYYY-MM-DD", "D MON YYYY", "MON YYYY".
func extractYear(dateStr string) int {
	if dateStr == "" {
		return 0
	}
	// Try ISO format: first segment before '-' may be the year
	isoStr := strings.TrimSpace(dateStr)
	if idx := strings.Index(isoStr, "T"); idx >= 0 {
		isoStr = isoStr[:idx]
	}
	parts := strings.Split(isoStr, "-")
	if len(parts[0]) == 4 {
		if y, err := strconv.Atoi(parts[0]); err == nil {
			return y
		}
	}
	// Try GEDCOM format: last numeric field with 4 digits is the year
	fields := strings.Fields(dateStr)
	for i := len(fields) - 1; i >= 0; i-- {
		y, err := strconv.Atoi(fields[i])
		if err == nil && y > 999 {
			return y
		}
	}
	return 0
}
