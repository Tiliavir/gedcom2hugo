package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const wikidataSPARQLEndpoint = "https://query.wikidata.org/sparql"

// wikidataBinding holds a single SPARQL result row from the Wikidata endpoint.
type wikidataBinding struct {
	EventLabel  struct{ Value string `json:"value"` } `json:"eventLabel"`
	Date        struct{ Value string `json:"value"` } `json:"date"`
	CountryLabel struct{ Value string `json:"value"` } `json:"countryLabel"`
	Description struct{ Value string `json:"value"` } `json:"description"`
}

// wikidataResult mirrors the JSON structure returned by the Wikidata SPARQL endpoint.
type wikidataResult struct {
	Results struct {
		Bindings []wikidataBinding `json:"bindings"`
	} `json:"results"`
}

// sparqlQuery is the SPARQL query used to fetch notable historical events from Wikidata.
// It retrieves wars, battles, natural disasters and other significant occurrences that
// have a point-in-time value (P585) and an optional country (P17).
const sparqlQuery = `SELECT DISTINCT ?eventLabel ?date ?countryLabel WHERE {
  VALUES ?type {
    wd:Q198
    wd:Q178561
    wd:Q45382
    wd:Q3839081
    wd:Q11483816
  }
  ?event wdt:P31 ?type .
  ?event wdt:P585 ?date .
  OPTIONAL { ?event wdt:P17 ?country . }
  FILTER(?date >= "1600-01-01T00:00:00Z"^^xsd:dateTime)
  FILTER(?date <= "2030-12-31T23:59:59Z"^^xsd:dateTime)
  SERVICE wikibase:label { bd:serviceParam wikibase:language "en". }
}
ORDER BY ?date
LIMIT 2000`

// fetchAndCacheEvents queries Wikidata SPARQL and writes the results to cachePath as JSON.
// The parent directory of cachePath is created if it does not exist.
func fetchAndCacheEvents(cachePath string) error {
	reqURL := fmt.Sprintf("%s?query=%s&format=json",
		wikidataSPARQLEndpoint, url.QueryEscape(sparqlQuery))

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "gedcom2hugo/1.0 (https://github.com/Tiliavir/gedcom2hugo)")
	req.Header.Set("Accept", "application/sparql-results+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var wikidataResp wikidataResult
	if err := json.Unmarshal(body, &wikidataResp); err != nil {
		return fmt.Errorf("parsing Wikidata response: %w", err)
	}

	seen := make(map[string]bool)
	var events []historicalEvent
	for _, b := range wikidataResp.Results.Bindings {
		title := strings.TrimSpace(b.EventLabel.Value)
		date := normalizeDate(b.Date.Value)
		if title == "" || date == "" {
			continue
		}
		key := date + "|" + title
		if seen[key] {
			continue
		}
		seen[key] = true
		events = append(events, historicalEvent{
			Date:        date,
			Title:       title,
			Description: strings.TrimSpace(b.Description.Value),
			Location:    strings.TrimSpace(b.CountryLabel.Value),
			Source:      "Wikidata",
		})
	}

	if err := os.MkdirAll(filepath.Dir(cachePath), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(cachePath, data, 0644)
}

// normalizeDate converts a Wikidata datetime like "2020-01-15T00:00:00Z" to "2020-01-15".
func normalizeDate(date string) string {
	if idx := strings.Index(date, "T"); idx >= 0 {
		date = date[:idx]
	}
	if len(date) >= 4 {
		return date
	}
	return ""
}
