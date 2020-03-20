package fence

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/alecthomas/template"
	"github.com/araddon/dateparse"
	"github.com/gorilla/mux"
	"github.com/yterajima/go-sitemap"
)

type MapData struct {
	MapLen    int
	CheckLen  int
	ErrorLen  int
	URLs      []string
	ErrorURLs []string
}

// Check NOTE the so slice does nothing..  the template is just static content.
// This function is here only in the thought this could be a dynamic list at some point.
func Check(w http.ResponseWriter, r *http.Request) {
	templateFile := "./web/templates/sitemap.html"

	ct := r.Header.Get("Accept") //  strings.Contains(ct, "text/html")
	log.Println(ct)

	var err error
	vars := mux.Vars(r)
	url := r.FormValue("url")
	dc := r.FormValue("date")

	log.Printf("vars: %v\n", vars)
	log.Printf("url: %s\n", url)
	log.Printf("dc: %s\n", dc)

	smap, err := sitemap.Get(url, nil)
	if err != nil {
		fmt.Println(err)
	}

	log.Print(len(smap.URL))

	c, o, err := dateCheck(smap, dc)
	if err != nil {
		fmt.Println(err)
	}

	// // move sitemap.URL to []string
	// u := make([]string, len(smap.URL))
	// for i := range smap.URL {
	// 	u[i] = smap.URL[i].Loc
	// }

	data := MapData{MapLen: len(smap.URL), CheckLen: len(c), ErrorLen: len(o), URLs: c, ErrorURLs: o}

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

func afterTime(lastmod, check time.Time) bool {
	return lastmod.After(check)
}

func dateCheck(smap sitemap.Sitemap, date string) ([]string, []string, error) {
	var c []string
	var o []string

	for _, URL := range smap.URL {
		if URL.LastMod != "" {
			t, err := dateparse.ParseAny(URL.LastMod)
			if err != nil {
				log.Println(err)
				o = append(o, URL.Loc)
			}
			check, _ := time.Parse(time.RFC822, date)
			q := afterTime(t, check)
			if q {
				c = append(c, URL.Loc)
			}
		} else {
			o = append(o, URL.Loc)
		}
	}

	if len(o) > 0 {
		log.Println("errors in the data parsing")
		return c, o, nil
	}

	return c, o, nil
}
