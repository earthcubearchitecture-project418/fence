package framing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"../core"

	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
)

// FrameData is the structure holding the JSON-LD framing results
type FrameData struct {
	Name        string
	Description string
	Keywords    string
	Type        string
}

// SpatialFD is the structure holding the JSON-LD framing results
type SpatialFD struct {
	Type      string
	Latitude  string
	Longitude string
	Line      string
	Polygon   string
	Box       string
}

// Frame performs the JSON-LD framing call
func Frame(w http.ResponseWriter, r *http.Request) {
	var err error
	var templateFile, sfr string

	ct := r.Header.Get("Accept")
	vars := mux.Vars(r)
	url := r.FormValue("url")
	fr := r.FormValue("frame")

	sdo, err := core.GetSDO(url)
	if err != nil {
		log.Println(err)
		// todo  need to just return a 4** error to the user, something went wrong...
	}

	log.Printf("ct: %s   vars: %s   urls: %s   fr: %s   ", ct, vars, url, fr)

	var data interface{}

	switch fr {
	case "literals":
		templateFile = "./web/templates/frame.html"
		sfr = DataLiterial(sdo)
		data = toCSVQuick(sfr)
	case "spatial":
		templateFile = "./web/templates/spatialframe.html"
		sfr = SpatialFrame(sdo)
		data = SpatialTabv2(sfr)
	default:
		fmt.Println("three")
	}

	if !strings.Contains(ct, "html") {
		w.Header().Set("Content-Type", "application/ld+json")
		jd, err := json.MarshalIndent(data, "", " ")
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

		err = ht.ExecuteTemplate(w, "Q", data)
		if err != nil {
			log.Printf("Template execution failed: %s", err)
		}

	}
}
