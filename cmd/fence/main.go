package main

import (
	"log"
	"net/http"
	"os"

	"../../internal/fence"
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
	fc.HandleFunc("/fence", fence.Render)
	http.Handle("/fence", fc)

	fp := mux.NewRouter()
	fp.HandleFunc("/fencepull", fence.Pull)
	http.Handle("/fencepull", fp)

	cp := mux.NewRouter()
	cp.HandleFunc("/sitemap", fence.Check)
	http.Handle("/sitemap", cp)

	// TODO   combine frames into one with routing based on form request
	fr := mux.NewRouter()
	fr.HandleFunc("/frame", fence.Frame)
	http.Handle("/frame", fr)

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
