package fence

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
	"github.com/kazarena/json-gold/ld"
	"github.com/tidwall/gjson"
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

	ct := r.Header.Get("Accept")
	log.Println(ct)

	var err error
	vars := mux.Vars(r)
	url := r.FormValue("url")
	fr := r.FormValue("frame")

	log.Printf("vars: %v\n", vars)
	log.Printf("url: %s\n", url)
	log.Printf("fr: %s\n", fr)

	sdo := ""

	// here is where I check for the isJSONLD flag and just
	// read the URL if it is
	if url != "" {
		if strings.HasSuffix(url, ".jsonld") {
			resp, err := http.Get(url)
			if err != nil {
				log.Println(err)
				sdo = "{}" // crappy way to deal with it..  need to return info to the user..   TODO!
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
				sdo = "{}" // crappy way to deal with it..  need to return info to the user..   TODO!
			}
			sdo = string(body)

		} else {
			log.Println("do getsdo")
			sdo, err = GetSDO(url)
			// sdo, err = headless(url)
			// sdo, err = getLink(url)
			if err != nil {
				log.Println(err)
				sdo = "{}"
			}
		}
	}

	if fr == "literals" {
		templateFile := "./web/templates/frame.html"

		// do framing here
		sfr := DataLiterial(sdo)
		data := toCSVQuick(sfr)

		// // check out return type here..

		ct := r.Header.Get("Accept") //  strings.Contains(ct, "text/html")
		log.Println(ct)

		// set output if no text/html
		// set default to octet stream?  but use stored if I have it
		if !strings.Contains(ct, "html") {
			w.Header().Set("Content-Type", "application/ld+json")
			// send the bytes
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
			// data literal frame
			ht, err := template.New("Template").ParseFiles(templateFile) //open and parse a template text file
			if err != nil {
				log.Printf("template parse failed: %s", err)
			}

			err = ht.ExecuteTemplate(w, "Q", data)
			if err != nil {
				log.Printf("Template execution failed: %s", err)
			}

		}

	} else {
		templateFile := "./web/templates/spatialframe.html"

		sfr := SpatialFrame(sdo)
		data := spatiaTab(sfr)

		// // check out return type here..

		ct := r.Header.Get("Accept") //  strings.Contains(ct, "text/html")
		log.Println(ct)

		// set output if no text/html
		// set default to octet stream?  but use stored if I have it
		if !strings.Contains(ct, "html") {
			w.Header().Set("Content-Type", "application/ld+json")
			// send the bytes
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
}

// DataLiterial is a simple testing frame
func DataLiterial(jsonld string) string {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	frame := map[string]interface{}{
		"@context":    "http://schema.org/",
		"@explicit":   true,
		"@type":       "Dataset",
		"description": "",
		"name":        "",
		"keywords":    ""}

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
	}

	framedDoc, err := proc.Frame(myInterface, frame, options) // do I need the options set in order to avoid the large context that seems to be generated?
	if err != nil {
		log.Println("Error when trying to frame document", err)
	}

	graph := framedDoc["@graph"]
	// ld.PrintDocument("JSON-LD graph section", graph) // debug print....

	jsonm, err := json.MarshalIndent(graph, "", " ")
	if err != nil {
		log.Println("Error trying to marshal data", err)
	}

	return string(jsonm)
}

func toCSVQuick(records string) FrameData {
	fd := FrameData{}

	log.Println("------------------------")
	log.Println(records)

	fd.Description = gjson.Get(records, "0.description").String()
	fd.Keywords = gjson.Get(records, "0.keywords").String()
	fd.Name = gjson.Get(records, "0.name").String()
	fd.Type = gjson.Get(records, "0.type").String()

	return fd
}
