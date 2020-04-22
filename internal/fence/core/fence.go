package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
	"github.com/piprate/json-gold/ld"
)

type PageData struct {
	SDO string
	URL string
}

// Render NOTE the so slice does nothing..  the template is just static content.
// This function is here only in the thought this could be a dynamic list at some point.
func Render(w http.ResponseWriter, r *http.Request) {
	log.Println("in fence render")
	templateFile := "./web/templates/fence.html"

	var err error
	vars := mux.Vars(r)
	url := r.FormValue("url")

	log.Println(vars)
	log.Println(url)

	sdo := ""

	// here is where I check for the isJSONLD flag and just
	// read the URL if it is
	if url != "" {
		if strings.HasSuffix(url, ".jsonld") {
			resp, err := http.Get(url)
			if err != nil {
				sdo = "{}" // crappy way to deal with it..  need to return info to the user..   TODO!
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				sdo = "{}" // crappy way to deal with it..  need to return info to the user..   TODO!
			}
			sdo = string(body)

		} else {
			sdo, err = GetSDO(url)
			// sdo, err = headless(url)
			// sdo, err = getLink(url)
			if err != nil {
				sdo = "{}"
			}
		}
	}

	data := PageData{SDO: sdo, URL: url}

	ht, err := template.New("Template").ParseFiles(templateFile) //open and parse a template text file
	if err != nil {
		log.Printf("template parse failed: %s", err)
	}

	err = ht.ExecuteTemplate(w, "Q", data)
	if err != nil {
		log.Printf("Template execution failed: %s", err)
	}
}

func isValid(jsonld string) error {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Format = "application/nquads"

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
		return err
	}

	_, err = proc.ToRDF(myInterface, options) // returns triples but toss them, just validating
	if err != nil {
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return err
	}

	return err
}
