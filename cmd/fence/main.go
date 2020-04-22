package main

import (
	"log"
	"net/http"
	"os"

	"../../internal/fence/core"
	"../../internal/fence/framing"
	"../../internal/fence/sitemap"
	"../../internal/fence/spatial"

	"github.com/gorilla/mux"
)

// MyServer struct for mux router
type MyServer struct {
	r *mux.Router
}

func main() {
	htmlRouter := mux.NewRouter()
	htmlRouter.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web"))))
	http.Handle("/", &MyServer{htmlRouter})

	fc := mux.NewRouter()
	fc.HandleFunc("/fence", core.Render)
	http.Handle("/fence", fc)

	fp := mux.NewRouter()
	fp.HandleFunc("/fencepull", core.Pull)
	http.Handle("/fencepull", fp)

	cp := mux.NewRouter()
	cp.HandleFunc("/sitemap", sitemap.Check)
	http.Handle("/sitemap", cp)

	// TODO   combine frames into one with routing based on form request
	fr := mux.NewRouter()
	fr.HandleFunc("/frame", framing.Frame)
	http.Handle("/frame", fr)

	gj := mux.NewRouter()
	gj.HandleFunc("/spatial", spatial.GeoJSON)
	http.Handle("/spatial", gj)

	// sf := mux.NewRouter()
	// sf.HandleFunc("/spatialframe", fence.SpatialRes)
	// http.Handle("/spatialframe", sf)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on %s. Go to http://127.0.0.1:%s", port, port)
	err := http.ListenAndServe(":"+port, nil)
	// http 2.0 http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// Let the Gorilla work
	s.r.ServeHTTP(rw, req)
}

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fn(w, r)
	}
}
