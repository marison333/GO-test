package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Cookies()) // dit zijn de cookies die de website terug geeft
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body)) // dit is de response die word opgehaald van de url
}
