package main

import (
	"flag"
	"fmt"

	"log"
	"net/http"

	"github.com/munxar/goapi/frontend"
)

func cors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		handler.ServeHTTP(w, r)
	})
}

func main() {
	devMode := false
	flag.BoolVar(&devMode, "dev", devMode, "enable dev mode")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"version": "1.0.0"}`))
	})

	mux.Handle("/admin/", frontend.SvelteKitHandler("/admin"))

	var handler http.Handler = mux

	if devMode {
		handler = cors(handler)
		fmt.Println("server running in dev mode")
	}

	log.Fatal(http.ListenAndServe(":3001", handler))
}
