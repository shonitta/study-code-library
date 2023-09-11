package main

import (
	"fmt"
	"io" // io/ioutil is deprecated: https://future-architect.github.io/articles/20210210/
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	url := "http://localhost:" + port + "/health"
	req, _ := http.NewRequest("GET", url, nil)

	Client := new(http.Client)
	resp, _ := Client.Do(req)
	defer resp.Body.Close() // Lazy evaluation

	byteArray, _ := io.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}
