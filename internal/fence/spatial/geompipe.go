package spatial

import (
	"log"
	"strconv"

	"../framing"
	"github.com/twpayne/go-geom"
	geomgj "github.com/twpayne/go-geom/encoding/geojson"
)

// An all in one test for the SDO to GeoJSON flow

// SDO2GeoJSON a generic SDO spatial to GeoJSON function call
func SDO2GeoJSON(jsonld string) (string, error) {
	g := geom.NewGeometryCollection()

	// frame to get the spatial info
	sf := framing.SpatialFrame(jsonld)

	data := framing.SpatialTab(sf)

	for i := range data {
		log.Println("-------------")
		log.Println(data[i].Type)
		log.Println(data[i].Latitude)
		log.Println(data[i].Longitude)
		log.Println(data[i].Line)
		log.Println(data[i].Box)
		log.Println(data[i].Polygon)

		if data[i].Type == "GeoCoordinates" {
			var la, lo []string
			la = append(la, data[i].Latitude)
			lo = append(lo, data[i].Longitude)

			err := GeoCoordinates2GJ(g, la, lo)
			if err != nil {
				log.Println(err)
			}
		}

		if data[i].Type == "GeoShape" && data[i].Line != "" {
			err := Line2GJ(g, data[i].Line)
			if err != nil {
				log.Println(err)
			}
		}
	}

	s, err := geomgj.Marshal(g)

	var fc geomgj.Feature

	fc.Geometry = g
	fc.ID = "test"
	m := make(map[string]interface{})
	m["URL"] = "http://example.org"

	fc.Properties = m

	fgj, err := fc.MarshalJSON()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(fgj))
	log.Println("-------  fgj about  ------")

	if err != nil {
		log.Println(err)
	}

	return string(s), nil
}

// AppProp line to geometry
func AppProp(g *geom.GeometryCollection, prop string) error {

	// need an array of geom coordinates
	// 39.3280,120.1633 40.445,123.7878

	// g.MustPush(gom

	return nil

}

// Line2GJ line to geometry
func Line2GJ(g *geom.GeometryCollection, line string) error {

	// need an array of geom coordinates
	// 39.3280,120.1633 40.445,123.7878
	var ca []geom.Coord
	ca = append(ca, geom.Coord{120.1633, 39.3280})
	ca = append(ca, geom.Coord{123.7878, 40.445})

	g.MustPush(geom.NewLineString(geom.XY).MustSetCoords(ca))

	return nil

}

// GeoCoordinates2GJ converts an array of lats and longs to geometry
func GeoCoordinates2GJ(g *geom.GeometryCollection, slat, slong []string) error {
	for x := range slat {
		lat, err := strconv.ParseFloat(slat[x], 64)
		if err != nil {
			return err
		}
		long, err := strconv.ParseFloat(slong[x], 64)
		if err != nil {
			return err
		}
		g.MustPush(geom.NewPoint(geom.XY).MustSetCoords(geom.Coord{long, lat}))
	}
	return nil
}
