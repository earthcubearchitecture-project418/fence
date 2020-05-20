package core

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetSDO needs to take a URL and then get the SDO from it
// if it can.
func GetSDO(urlloc string) (string, error) {
	var jsonld string

	if strings.HasSuffix(urlloc, ".jsonld") {
		resp, err := http.Get(urlloc)
		if err != nil {
			log.Println(err)
			jsonld = "{}" // crappy way to deal with it..  need to return info to the user..   TODO!
			return jsonld, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			jsonld = "{}" // crappy way to deal with it..  need to return info to the user..   TODO!
			return jsonld, err
		}
		jsonld = string(body)
		return jsonld, err

	}

	var client http.Client
	req, err := http.NewRequest("GET", urlloc, nil)
	if err != nil {
		// not even being able to make a req instance..  might be a fatal thing?
		log.Printf("------ error making request------ \n %s", err)
	}

	req.Header.Set("User-Agent", "EarthCube_DataBot/1.0")
	req.Header.Set("Accept", "text/html") // TODO we need to think about content negotiation request json-Ld too!

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error reading location: %s", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Printf("Error doc from resp: %v", err)
	}

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

	log.Println(jsonld)

	return jsonld, err
}
