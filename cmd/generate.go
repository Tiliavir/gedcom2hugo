package cmd

import (
	"bytes"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/iand/gedcom"
	"github.com/urfave/cli/v2"
)

var tagTable map[string]string

// Generate reads the GEDCOM file and builds the Hugo input files.
// noinspection GoUnusedExportedFunction
func Generate(cx *cli.Context) error {
	tagTable = map[string]string{
		"BAPM":  "Baptism",
		"BIRT":  "Birth",
		"BURI":  "Buried",
		"CENS":  "Census",
		"CHR":   "Christening",
		"DEAT":  "Death",
		"DIV":   "Divorced",
		"DIVF":  "Divorce Filed",
		"EMIG":  "Emigrated",
		"ENGA":  "Engaged",
		"GRAD":  "Graduated",
		"MARB":  "Marriage Bann",
		"MARL":  "Marriage License",
		"MARR":  "Married",
		"NATU":  "Naturalized",
		"OCCU":  "Occupation",
		"RELI":  "Religion",
		"RESI":  "Residence",
		"_MILT": "Military Service",
	}

	gc, err := readGedcom(cx)
	if err != nil {
		return cli.Exit(err, 1)
	}

	api := newAPIControl(cx)

	// Load historical events cache (fetch from Wikidata if the cache is absent and
	// --fetch-history is requested).
	if histRepo, loadErr := loadHistoricalEvents(cx); loadErr != nil {
		log.Printf("Warning: historical events not available: %v\n", loadErr)
	} else {
		api.histRepo = histRepo
	}

	err = api.buildFromGedcom(gc)
	if err != nil {
		return cli.Exit(err, 1)
	}

	err = api.exportSourceAPI()
	if err != nil {
		return cli.Exit(err, 1)
	}

	err = api.exportSourcePages()
	if err != nil {
		return cli.Exit(err, 1)
	}

	err = api.exportIndividualAPI()
	if err != nil {
		return cli.Exit(err, 1)
	}

	err = api.exportIndividualPages()
	if err != nil {
		return cli.Exit(err, 1)
	}

	err = api.exportFamilyAPI()
	if err != nil {
		return cli.Exit(err, 1)
	}

	err = api.exportFamilyPages()
	if err != nil {
		return cli.Exit(err, 1)
	}

	err = api.exportPhotoAPI()
	if err != nil {
		return cli.Exit(err, 1)
	}

	err = api.exportPhotoPages()
	if err != nil {
		return cli.Exit(err, 1)
	}

	err = configureForJsonHeaders(api)
	if err != nil {
		return cli.Exit(err, 1)
	}

	return nil
}

// loadHistoricalEvents resolves the cache path, optionally fetches events from Wikidata,
// and returns a ready repository. An error is returned only when the cache cannot be read
// after a potential fetch attempt.
func loadHistoricalEvents(cx *cli.Context) (*historicalEventRepository, error) {
	cachePath := cx.String("history-cache")
	if cachePath == "" {
		cachePath = filepath.Join(cx.String("project"), "data", "history-events.json")
	}

	// Fetch from Wikidata when explicitly requested and the cache is absent.
	if cx.Bool("fetch-history") {
		if _, err := os.Stat(cachePath); errors.Is(err, os.ErrNotExist) {
			log.Printf("Fetching historical events from Wikidata → %s\n", cachePath)
			if fetchErr := fetchAndCacheEvents(cachePath); fetchErr != nil {
				log.Printf("Warning: could not fetch historical events: %v\n", fetchErr)
			}
		}
	}

	repo := &historicalEventRepository{}
	if err := repo.loadFromFile(cachePath); err != nil {
		return nil, err
	}
	return repo, nil
}

func configureForJsonHeaders(api *apiControl) error {
	headers := filepath.Join(api.cx.String("project"), "/static/api/_headers")
	file, err := os.Create(headers)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte("/*\n  Access-Control-Allow-Origin: *\n  Content-Type: application/json; charset=utf-8"))
	return err
}

// readGedcom reads the GEDCOM file specified in the context into memory.
func readGedcom(cx *cli.Context) (*gedcom.Gedcom, error) {
	var gc *gedcom.Gedcom

	if cx.String("gedcom") == "" {
		return gc, errors.New("no GEDCOM file specified for input")
	}

	data, err := os.ReadFile(cx.String("gedcom"))
	if err != nil {
		return gc, err
	}

	decoder := gedcom.NewDecoder(bytes.NewReader(data))

	gc, err = decoder.Decode()
	if err != nil {
		return gc, err
	}
	return gc, nil
}

