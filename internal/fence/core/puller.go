package core

import (
	"io"
	"log"
	"net/http"
	"strings"
)

// Pull NOTE the so slice does nothing..  the template is just static content.
// This function is here only in the thought this could be a dynamic list at some point.
func Pull(w http.ResponseWriter, r *http.Request) {
	log.Println("in fence puller")
	url := r.FormValue("url")
	sdo := ""
	var err error

	if url != "" {
		sdo, err = GetSDO(url)
		if err != nil {
			sdo = "{}"
		}
	}

	sr := strings.NewReader(sdo)
	n, err := io.Copy(w, sr)
	if err != nil {
		log.Println(err)
	}

	log.Printf("Sent schema.org package %d bytes\n", n)
}
