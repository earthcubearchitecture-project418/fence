package spatial

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"../framing"

	geojson "github.com/paulmach/go.geojson"
	"github.com/piprate/json-gold/ld"
	"github.com/tidwall/gjson"
)

// LatLong2GeoJSON convert a jsonld to geojson if it has spatial data
func LatLong2GeoJSON(jsonld string) ([]byte, error) {
	// start by flattening the JLD
	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
	}

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	context := map[string]interface{}{
		"@context": "http://schema.org/",
	}

	f, err := proc.Flatten(myInterface, context, options)
	if err != nil {
		log.Println("Error trying to flatten data", err)
	}

	j, err := json.MarshalIndent(f, "", " ")
	if err != nil {
		log.Println("Error trying to marshal data", err)
	}

	// map the elements in jld
	m := make(map[string]string)
	result := gjson.Get(string(j), "@graph.0")
	result.ForEach(func(key, value gjson.Result) bool {
		//fmt.Printf("key: %s   value: %s \n", key.String(), value.String())
		if !strings.Contains(value.String(), "@id") && !strings.Contains(key.String(), "@id") {
			m[key.String()] = value.String()
		}
		return true // keep iterating
	})

	// frame to get the spatial info
	sf := framing.SpatialFrame(jsonld)
	data := framing.SpatialTab(sf)

	la, _ := strconv.ParseFloat(data[0].Latitude, 64)
	lo, _ := strconv.ParseFloat(data[0].Longitude, 64)

	// make the geojson
	featureCollection := geojson.NewFeatureCollection()
	feature := geojson.NewPointFeature([]float64{lo, la})

	// loop on passed map to add more properties
	for k, v := range m {
		feature.SetProperty(k, v)
	}

	featureCollection.AddFeature(feature)

	return featureCollection.MarshalJSON()
}
