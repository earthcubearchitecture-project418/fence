package spatial

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"../core"

	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
)

type Position struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type XYZData struct {
	Tags []string `json:"tags"`
}

// GeoJSON is the handler for the geojson call
func GeoJSON(w http.ResponseWriter, r *http.Request) {
	var err error
	var templateFile string

	ct := r.Header.Get("Accept")
	vars := mux.Vars(r)
	url := r.FormValue("url")

	sdo, err := core.GetSDO(url)
	if err != nil {
		log.Println(err)
	}

	// log.Println(sdo)

	log.Printf("ct: %s   vars: %s   urls: %s   ", ct, vars, url)
	templateFile = "./web/templates/geojson.html"

	// g, err := LatLong2GeoJSON(sdo)
	// if err != nil {
	// 	log.Println(err)
	// }

	// TESTING CALL (prints to console for now)

	g, err := SDO2GeoJSON(sdo)
	if err != nil {
		log.Println(err)
	}
	log.Printf("----\n  %s   \n----\n", g)

	if !strings.Contains(ct, "html") {
		w.Header().Set("Content-Type", "application/geo+json")
		jd, err := json.MarshalIndent(string(g), "", " ")
		if err != nil {
			log.Println(err)
		}
		r := bytes.NewReader(jd)

		n, err := io.Copy(w, r)
		if err != nil {
			log.Println("Issue with writing bytes to http response")
			log.Println(err)
		}
		log.Printf("NEW :   Sent %d bytes\n", n)
	} else {
		ht, err := template.New("Template").ParseFiles(templateFile) //open and parse a template text file
		if err != nil {
			log.Printf("template parse failed: %s", err)
		}

		err = ht.ExecuteTemplate(w, "Q", string(g))
		if err != nil {
			log.Printf("Template execution failed: %s", err)
		}
	}
}
