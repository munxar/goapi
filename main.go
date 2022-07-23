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
	port := uint(3001)
	flag.BoolVar(&devMode, "dev", devMode, "enable dev mode")
	flag.UintVar(&port, "p", port, "server port")
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

	fmt.Printf("api served at http://localhost:%d\ngui served at http://localhost:%d/admin\n", port, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
