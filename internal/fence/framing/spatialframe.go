package framing

import (
	"encoding/json"
	"log"

	"cuelang.org/go/pkg/strings"
	"github.com/piprate/json-gold/ld"
	"github.com/tidwall/gjson"
)

// SpatialFrame is a simple testing frame
func SpatialFrame(jsonld string) string {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	// frame := map[string]interface{}{
	// 	"@context":  "http://schema.org/",
	// 	"@explicit": true,
	// 	"spatialCoverage": map[string]interface{}{
	// 		"@type": "Place",
	// 		"geo":   map[string]interface{}{},
	// 	},
	// }

	// TODO check the JSON-LD for http or https and alter the frame

	frame := map[string]interface{}{}

	if strings.Contains(jsonld, "@vocab") { // brittle !!!!!!!!!
		log.Println("https context")
		frame = map[string]interface{}{
			"@context":  map[string]interface{}{"@vocab": "https://schema.org/"},
			"@explicit": true,
			"geo":       map[string]interface{}{},
		}
	} else {
		log.Println("http context")
		frame = map[string]interface{}{
			"@context":  "http://schema.org/",
			"@explicit": true,
			"geo":       map[string]interface{}{},
		}
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

// SpatialTabv2 testing new frame
func SpatialTabv2(records string) []SpatialFD {
	fda := []SpatialFD{}

	log.Println("SpatialTabv2")
	log.Println(records)

	if strings.Contains(records, "@value") {
		// We are a typed graph
		if gjson.Get(records, "0.spatialCoverage.geo.#.type").Exists() {
			println("TYPED:  Array of spatial elements mode")
			result := gjson.Get(records, "0.spatialCoverage.geo")

			result.ForEach(func(key, value gjson.Result) bool {
				println(value.Get("type").String())
				fd := SpatialFD{}
				fd.Type = value.Get("type").String()
				fd.Latitude = value.Get("latitude.@value").String()
				fd.Longitude = value.Get("longitude.@value").String()
				fd.Line = value.Get("line").String()
				fd.Polygon = value.Get("polygon").String()
				fd.Box = value.Get("box").String()
				fda = append(fda, fd)
				return true // keep iterating
			})
			// for _, geo := range result.Array() {
			// 	println(geo.String())
		} else if gjson.Get(records, "0.spatialCoverage.geo.type").Exists() {
			println("TYPED:  Single spatial element mode")
			fd := SpatialFD{}
			fd.Type = gjson.Get(records, "0.spatialCoverage.geo.type").String()
			fd.Latitude = gjson.Get(records, "0.spatialCoverage.geo.latitude.@value").String()
			fd.Longitude = gjson.Get(records, "0.spatialCoverage.geo.longitude.@value").String()
			fd.Line = gjson.Get(records, "0.spatialCoverage.geo.line").String()
			fd.Polygon = gjson.Get(records, "0.spatialCoverage.geo.polygon").String()
			fd.Box = gjson.Get(records, "0.spatialCoverage.geo.box").String()
			fda = append(fda, fd)
		}
	} else {
		if gjson.Get(records, "0.geo.#.type").Exists() {
			println("Array of spatial elements mode")
			result := gjson.Get(records, "0.geo")

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
		} else if gjson.Get(records, "0.geo.type").Exists() {
			println("Single spatial element mode")
			fd := SpatialFD{}
			fd.Type = gjson.Get(records, "0.geo.type").String()
			fd.Latitude = gjson.Get(records, "0.geo.latitude").String()
			fd.Longitude = gjson.Get(records, "0.geo.longitude").String()
			fd.Line = gjson.Get(records, "0.geo.line").String()
			fd.Polygon = gjson.Get(records, "0.geo.polygon").String()
			fd.Box = gjson.Get(records, "0.geo.box").String()
			fda = append(fda, fd)
		}
	}

	log.Println(fda)
	return fda
}

// SpatialTab return a struct of geometry elements from the flattend SDO
func SpatialTab(records string) []SpatialFD {
	fda := []SpatialFD{}

	log.Println("SpatialTab")
	log.Println(records)

	if strings.Contains(records, "@value") {
		// We are a typed graph
		if gjson.Get(records, "0.spatialCoverage.geo.#.type").Exists() {
			println("TYPED:  Array of spatial elements mode")
			result := gjson.Get(records, "0.spatialCoverage.geo")

			result.ForEach(func(key, value gjson.Result) bool {
				println(value.Get("type").String())
				fd := SpatialFD{}
				fd.Type = value.Get("type").String()
				fd.Latitude = value.Get("latitude.@value").String()
				fd.Longitude = value.Get("longitude.@value").String()
				fd.Line = value.Get("line").String()
				fd.Polygon = value.Get("polygon").String()
				fd.Box = value.Get("box").String()
				fda = append(fda, fd)
				return true // keep iterating
			})
			// for _, geo := range result.Array() {
			// 	println(geo.String())
		} else if gjson.Get(records, "0.spatialCoverage.geo.type").Exists() {
			println("TYPED:  Single spatial element mode")
			fd := SpatialFD{}
			fd.Type = gjson.Get(records, "0.spatialCoverage.geo.type").String()
			fd.Latitude = gjson.Get(records, "0.spatialCoverage.geo.latitude.@value").String()
			fd.Longitude = gjson.Get(records, "0.spatialCoverage.geo.longitude.@value").String()
			fd.Line = gjson.Get(records, "0.spatialCoverage.geo.line").String()
			fd.Polygon = gjson.Get(records, "0.spatialCoverage.geo.polygon").String()
			fd.Box = gjson.Get(records, "0.spatialCoverage.geo.box").String()
			fda = append(fda, fd)
		}
	} else {
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
	}

	log.Println(fda)
	return fda
}
