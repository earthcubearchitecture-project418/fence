package framing

import (
	"encoding/json"
	"log"

	"github.com/kazarena/json-gold/ld" // kazarena  piprate
	"github.com/tidwall/gjson"
)

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

	// ld.PrintDocument("JSON-LD expansion succeeded", framedDoc)

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
