package main

import (
	"fmt"
	"net/http"
)

func main() {
	health := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello World\n")
	}

	http.Handle("/health", Logger(http.HandlerFunc(health)))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
