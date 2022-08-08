package main

import (
	"log"
	"fmt"
	"net/http"
	"os"
)

func HandleHealthz(w http.ResponseWriter, r *http.Request) {
	log.Printf("Namaste World!!!")
	fmt.Fprint(w, "Ok!")
}

func main() {
	http.HandleFunc("/healthz", HandleHealthz)

	listenAddress := ":8080"
	if os.Getenv("PORT") != "" {
		listenAddress = ":" + os.Getenv("PORT")
	}

	log.Printf("listening on %v", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
