package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	health := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello World\n")
	}

	http.Handle("/health", Logger(http.HandlerFunc(health)))

	port := os.Getenv("PORT")
	fmt.Printf("Listening on port " + port + "...\n")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
