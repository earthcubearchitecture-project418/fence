package fence

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
	"github.com/piprate/json-gold/ld"
)

// Render NOTE the so slice does nothing..  the template is just static content.
// This function is here only in the thought this could be a dynamic list at some point.
func Render(w http.ResponseWriter, r *http.Request) {
	log.Println("in fence render")
	templateFile := "./web/templates/fence.html"

	//url := r.URL.Query().Get("url")
	vars := mux.Vars(r)
	url := r.FormValue("url")
	//url := vars["url"]
	sdo := ""
	var err error

	log.Println(vars)
	log.Println(url)

	if url != "" {
		sdo, err = getSDO(url)
		if err != nil {
			sdo = "{}"
		}
	}

	ht, err := template.New("Template").ParseFiles(templateFile) //open and parse a template text file
	if err != nil {
		log.Printf("template parse failed: %s", err)
	}

	err = ht.ExecuteTemplate(w, "Q", sdo)
	if err != nil {
		log.Printf("Template execution failed: %s", err)
	}
}

// getSDO needs to take a URL and then get the SDO from it
// if it can.
func getSDO(urlloc string) (string, error) {

	// sdo := sdo()
	// urlloc := "http://opencoredata.org/doc/dataset/b8d7bd1b-ef3b-4b08-a327-e28e1420adf0"

	var client http.Client
	req, err := http.NewRequest("GET", urlloc, nil)
	if err != nil {
		// not even being able to make a req instance..  might be a fatal thing?
		log.Printf("------ error making request------ \n %s", err)
	}

	req.Header.Set("User-Agent", "EarthCube_DataBot/1.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error reading location: %s", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Printf("Error doc from resp: %v", err)
	}

	var jsonld string
	if err == nil {
		doc.Find("script").Each(func(i int, s *goquery.Selection) {
			val, _ := s.Attr("type")
			if val == "application/ld+json" {
				err = isValid(s.Text())
				if err != nil {
					log.Printf("ERROR: At %s JSON-LD is NOT valid: %s", urlloc, err)
				}
				jsonld = s.Text()
			}
		})
	}

	// return sdo, nil
	return jsonld, nil
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

// sdo is a testing function to return a sample scheme.org
// type dataset for testing and initial development
// func sdo() string {

// 	sdo := `
// {
//     "@context": {
//         "@vocab": "http://schema.org/",
//         "re3data": "http://example.org/re3data/0.1/"
//     },
//     "@id": "http://opencoredata.org/pkg/id/a4c8a2794619d25846c00d9ce43ed1a166b9ccd84aab964d8353307cae7743e1",
//     "@type": "Dataset",
//     "description": "A CSDCO data package for  project HAWS (HAWS (Hawaii Soils))",
//     "distribution": {
//         "@type": "DataDownload",
//         "contentUrl": "http://opencoredata.org/pkg/id/a4c8a2794619d25846c00d9ce43ed1a166b9ccd84aab964d8353307cae7743e1.zip",
//         "fileFormat": "application/vnd.datapackage+json"
//     },
//     "keywords": "CSDCO, Continental Scientific Drilling",
//     "license": "https://creativecommons.org/publicdomain/zero/1.0/",
//     "name": "a4c8a2794619d25846c00d9ce43ed1a166b9ccd84aab964d8353307cae7743e1.zip",
//     "publisher": {
//         "@type": "Organization",
//         "description": "Continental Scientific Drilling Coordination Office",
//         "name": "CSDCO",
//         "url": "https://csdco.umn.edu/"
//     },
//     "spatialCoverage": {
//         "@type": "Place",
//         "geo": {
//             "@type": "GeoCoordinates",
//             "latitude": "0.0",
//             "longitude": "0.0"
//         }
//     },
//     "url": "http://opencoredata.org/pkg/id/a4c8a2794619d25846c00d9ce43ed1a166b9ccd84aab964d8353307cae7743e1",
//     "variableMeasured": null
// }
// 	`

// 	return sdo
// }
