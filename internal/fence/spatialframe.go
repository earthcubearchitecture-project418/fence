package fence

import (
	"encoding/json"
	"log"

	"github.com/kazarena/json-gold/ld"
	"github.com/tidwall/gjson"
)

// SpatialFrame is a simple testing frame
func SpatialFrame(jsonld string) string {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	frame := map[string]interface{}{
		"@context":  "http://schema.org/",
		"@explicit": true,
		// "@type":     "Dataset",
		"spatialCoverage": map[string]interface{}{
			"@type": "Place",
			"geo":   map[string]interface{}{},
		},
	}

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

	// log.Println("From Frame function")
	// log.Println(string(jsonm))

	return string(jsonm)
}

// Testing urls
// curl "http://localhost:8080/spatialframe?url=http://opencoredata.org/doc/dataset/b8d7bd1b-ef3b-4b08-a327-e28e1420adf0"
// curl "http://localhost:8080/spatialframe?url=https://gist.githubusercontent.com/fils/8738793069ae18fc368f04b2ace7118d/raw/24a89814b2e807d6abeb092b5f7c6626b33bca97/spatialtest.jsonld"

func spatiaTab(records string) []SpatialFD {
	fda := []SpatialFD{}

	if gjson.Get(records, "0.spatialCoverage.geo.#.type").Exists() {
		println("Array of spatial elements mode")
		result := gjson.Get(records, "0.spatialCoverage.geo")

		result.ForEach(func(key, value gjson.Result) bool {
			println(value.Get("type").String())
			fd := SpatialFD{}
			fd.Type = value.Get("type").String()
			fd.Latitude = value.Get("latitude").String()
			fd.Longitude = value.Get("longitude").String()
			fd.Line = value.Get("line").String()
			fd.Polygon = value.Get("polygon").String()
			fd.Box = value.Get("box").String()
			fda = append(fda, fd)
			return true // keep iterating
		})
		// for _, geo := range result.Array() {
		// 	println(geo.String())
	} else if gjson.Get(records, "0.spatialCoverage.geo.type").Exists() {
		println("Single spatial element mode")
		fd := SpatialFD{}
		fd.Type = gjson.Get(records, "0.spatialCoverage.geo.type").String()
		fd.Latitude = gjson.Get(records, "0.spatialCoverage.geo.latitude").String()
		fd.Longitude = gjson.Get(records, "0.spatialCoverage.geo.longitude").String()
		fd.Line = gjson.Get(records, "0.spatialCoverage.geo.line").String()
		fd.Polygon = gjson.Get(records, "0.spatialCoverage.geo.polygon").String()
		fd.Box = gjson.Get(records, "0.spatialCoverage.geo.box").String()
		fda = append(fda, fd)
	}
	// log.Println(fd)
	return fda
}
