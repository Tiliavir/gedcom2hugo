package cmd

import (
	"encoding/json"
	"html/template"
	"image"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/iand/gedcom"
)

func (api *apiControl) addPhoto(o *gedcom.MediaRecord) *photoResponse {
	key := getPhotoKeyFromObject(o)
	if _, ok := api.photos[key]; !ok {
		for _, record := range o.File {
			api.photos[key] = &photoResponse{
				ID:    key,
				File:  filepath.Base(record.Name),
				Title: record.Title,
			}

			for i := range o.UserDefined {
				if o.UserDefined[i].Tag == "_TEXT" {
					api.photos[key].Description = o.UserDefined[i].Value
					break
				}
			}

			for _, note := range o.Note {
				api.photos[key].Notes = append(api.photos[key].Notes, note.Note)
			}

			file, err := os.Open(record.Name)
			if err != nil {
				log.Printf("Warning: Unable to open photo file %s: %v\n", record.Name, err)
				return api.photos[key]
			}
			defer file.Close()

			img, _, err := image.DecodeConfig(file)
			if err != nil {
				log.Printf("Warning: Unable to decode image %s: %v\n", record.Name, err)
				return api.photos[key]
			}

			api.photos[key].Width = img.Width
			api.photos[key].Height = img.Height
		}
	}

	return api.photos[key]
}

func (api *apiControl) addPhotoForIndividual(o *gedcom.MediaRecord, i *individualResponse) *photoResponse {
	response := api.addPhoto(o)

	ir, err := api.getIndividualIndexEntry(i.ID)
	if err != nil {
		log.Printf("Warning: could not build individual photo index entry for %s: %v\n", i.ID, err)
		return response
	}
	response.People = append(response.People, ir)

	return response
}

func (api *apiControl) addPhotoForFamily(o *gedcom.MediaRecord, f *familyResponse) *photoResponse {
	response := api.addPhoto(o)

	fr, err := api.getFamilyIndexEntry(f.ID)
	if err != nil {
		log.Printf("Warning: could not build family photo index entry for %s: %v\n", f.ID, err)
		return response
	}
	response.Families = append(response.Families, fr)

	return response
}

func (api *apiControl) buildFromGedcom(g *gedcom.Gedcom) error {
	api.gc = g

	var err error

	err = api.addSources()
	if err != nil {
		return err
	}

	err = api.addIndividuals()
	if err != nil {
		return err
	}

	err = api.addFamilies()
	if err != nil {
		return err
	}

	return nil
}

func (api *apiControl) exportPhotoAPI() error {
	photoAPIDir := filepath.Join(api.cx.String("project"), "static", "api", "photo")
	err := os.MkdirAll(photoAPIDir, 0755)
	if err != nil {
		return err
	}

	var photoIDs []string
	for id, photo := range api.photos {
		photoIDs = append(photoIDs, id)
		file := filepath.Join(photoAPIDir, strings.ToLower(id+".json"))
		if err := writeJSONFile(file, photo); err != nil {
			return err
		}
	}
	sort.Strings(photoIDs)

	file := filepath.Join(photoAPIDir, strings.ToLower("list.json"))
	if err := writeJSONFile(file, photoIDs); err != nil {
		return err
	}

	return nil
}

func (api *apiControl) exportPhotoPages() error {
	const photoPageTemplate = `---
url: "/{{ .ID }}/"
categories:
  - Photo
lead_photo: "{{ .File }}"
photo_key: "{{ .ID  }}"
---
<script src="../js/jquery.min.js"></script>
<script src="../js/photodisplay.js"></script>
<script>
$(document).ready(function(){
    photodisplay("{{ .ID }}")
});
</script>

<div id="display"></div>

<div id="raw"></div>
`

	photoDir := filepath.Join(api.cx.String("project"), "content", "media")
	err := os.MkdirAll(photoDir, 0755)
	if err != nil {
		return err
	}

	for key, photo := range api.photos {
		file := filepath.Join(photoDir, key+".md")

		fh, err := os.Create(file)
		if err != nil {
			return err
		}

		tpl := template.New("photo")
		tpl, err = tpl.Parse(photoPageTemplate)
		if err != nil {
			_ = fh.Close()
			return err
		}
		if err := tpl.Execute(fh, photo); err != nil {
			_ = fh.Close()
			return err
		}
		if err := fh.Close(); err != nil {
			return err
		}
	}

	return nil
}

func writeJSONFile(path string, v interface{}) error {
	fh, err := os.Create(path)
	if err != nil {
		return err
	}

	data, err := json.Marshal(v)
	if err != nil {
		_ = fh.Close()
		return err
	}

	if _, err := fh.Write(data); err != nil {
		_ = fh.Close()
		return err
	}

	return fh.Close()
}

func getPhotoKeyFromObject(o *gedcom.MediaRecord) string {
	if len(o.File) > 0 {
		return "p" + strings.ToLower(strings.Replace(filepath.Base(o.File[0].Name), ".", "", -1))
	}
	return ""
}
