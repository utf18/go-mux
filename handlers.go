package main

import (
	"net/http"
)

// YourHandler displays the string below
func YourHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("this is not the endpoint you are looking for... try /metrics"))

}
