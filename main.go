package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/enrichman")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user map[string]any
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User '%v' (%v) found\n", user["login"], user["name"])
}
