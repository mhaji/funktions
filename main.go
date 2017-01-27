package main

import (
	"github.com/mhaji/funktions/funktions"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	funktions.RegisterRoutes(r)


	port := 9090
	log.Printf("Serving on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}