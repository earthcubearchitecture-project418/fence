package sitemap

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/alecthomas/template"
	"github.com/araddon/dateparse"
	"github.com/yterajima/go-sitemap"
)

type MapData struct {
	MapLen    int    `json:"urlcount"`
	CheckLen  int    `json:"urlspastdate"`
	ErrorLen  int    `json:"datesnotfound"`
	Date      string `json:"filterdate"`
	URLs      []string
	ErrorURLs []string
}

// Check NOTE the so slice does nothing..  the template is just static content.
// This function is here only in the thought this could be a dynamic list at some point.
func Check(w http.ResponseWriter, r *http.Request) {
	var err error
	ct := r.Header.Get("Accept") //  strings.Contains(ct, "text/html")
	// vars := mux.Vars(r)
	url := r.FormValue("url")
	dc := r.FormValue("date")

	smap, err := sitemap.Get(url, nil)
	if err != nil {
		log.Println(err)
	}

	c, o, err := DateCheck(smap, dc)
	if err != nil {
		log.Println(err)
	}

	data := MapData{MapLen: len(smap.URL), CheckLen: len(c), Date: dc, ErrorLen: len(o), URLs: c, ErrorURLs: o}

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
		log.Printf("Sent %d bytes\n", n)

	} else {
		templateFile := "./web/templates/sitemap.html"
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

func afterTime(lastmod, check time.Time) bool {
	return lastmod.After(check)
}

// DateCheck uses a sitemap and date to filter URLs.  If date is null then all
// URLs are added.  If date is not nil and date is not found in sitemap the URL
// is added to the "error" array for the user to decide on.
func DateCheck(smap sitemap.Sitemap, date string) ([]string, []string, error) {
	var c []string
	var o []string

	for _, URL := range smap.URL {
		if URL.LastMod != "" && date != "" {
			t, err := dateparse.ParseAny(URL.LastMod)
			if err != nil {
				log.Println(err)
				o = append(o, URL.Loc)
			}
			check, err := time.Parse(time.RFC822, date)
			if err != nil {
				log.Println(err)
			}
			q := afterTime(t, check)
			if q {
				c = append(c, URL.Loc)
			}
		} else {
			c = append(c, URL.Loc)
		}
	}

	return c, o, nil
}
