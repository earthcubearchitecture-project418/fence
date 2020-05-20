package spatial

import (
	"log"
	"regexp"
	"strconv"
	"strings"

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

		if data[i].Type == "GeoShape" && data[i].Box != "" {
			err := Box2Geom(g, data[i].Line)
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

	return nil

}

// Box2Geom converts a box to a geometry
// TODO really just a plygon...
func Box2Geom(g *geom.GeometryCollection, box string) error {
	cs := clnstr(box)
	ls := strings.Split(cs, " ")
	var ca []geom.Coord
	var aca [][]geom.Coord

	// TODO  Is a box always going to be a len(ca) = 4 ?   If so we can test that
	// Box: 39.3280 120.1633 40.445 123.7878
	// There is no box..  need to convert to a series of multilines...
	// NewPolygon(XY).MustSetCoords([][]Coord{{{1, 2}, {3, 4}, {5, 6}}, {{7, 8}, {9, 10}, {11, 12}}}),

	for i := range ls {
		lp := strings.Split(ls[i], ",")
		long, err := strconv.ParseFloat(lp[0], 64)
		if err != nil {
			return err
		}
		lat, err := strconv.ParseFloat(lp[1], 64)
		if err != nil {
			return err
		}

		ca = append(ca, geom.Coord{long, lat})
	}

	aca = append(aca, ca)

	g.MustPush(geom.NewPolygon(geom.XY).MustSetCoords(aca))

	return nil
}

// Line2GJ line to geometry
func Line2GJ(g *geom.GeometryCollection, line string) error {
	cs := clnstr(line)
	ls := strings.Split(cs, " ")
	var ca []geom.Coord

	for i := range ls {
		lp := strings.Split(ls[i], ",")
		long, err := strconv.ParseFloat(lp[0], 64)
		if err != nil {
			return err
		}
		lat, err := strconv.ParseFloat(lp[1], 64)
		if err != nil {
			return err
		}

		ca = append(ca, geom.Coord{long, lat})
	}

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

func clnstr(str string) string {
	// remove more than one space
	space := regexp.MustCompile(`\s+`)
	s0 := space.ReplaceAllString(str, " ")

	// trime leading and ending spaces
	s1 := strings.TrimSpace(s0)

	// convert " ," and ", " to just ","
	s2 := strings.ReplaceAll(s1, " ,", ",")
	s3 := strings.ReplaceAll(s2, ", ", ",")

	return s3
}
