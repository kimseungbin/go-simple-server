package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if player == "Pepper" {
		fmt.Fprintf(w, "20")
		return
	}

	if player == "Floyd" {
		fmt.Fprintf(w, "10")
		return
	}
}

func main() {
	handler := http.HandlerFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":5001", handler))
}
