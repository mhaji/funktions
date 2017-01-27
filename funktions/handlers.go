package funktions

import (
	"net/http"
	"fmt"
)

func CreateFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func ListFunctions(w http.ResponseWriter, r *http.Request) {

}

func GetFunction(w http.ResponseWriter, r *http.Request) {

}

func  DeleteFunction(w http.ResponseWriter, r *http.Request) {

}

func UpdateFunction(w http.ResponseWriter, r *http.Request) {

}