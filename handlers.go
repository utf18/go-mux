package main

import "net/http"

// YourHandler displays sample handler output here: Gorilla!
func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
