package funcktions

import "github.com/gorilla/mux"

type Funktions struct {

}


func RegisterRoutes(r *mux.Router) *mux.Router{
	r.HandleFunc("/v1/funktions", CreateFunction).Methods("POST")
	r.HandleFunc("/v1/funktions", ListFunctions).Methods("GET")
	r.HandleFunc("/v1/funktions/{funktion-id}", GetFunction).Methods("GET")
	r.HandleFunc("/v1/funktions/{funktion-id}", DeleteFunction).Methods("DELETE")
	r.HandleFunc("/v1/funktions/{funktion-id}", UpdateFunction).Methods("PUT")

	return r
}
