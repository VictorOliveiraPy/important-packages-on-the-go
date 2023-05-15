package main

import (
	"fmt"
	"net/http"
)


func main() {
	request, err := http.Get("https://www.google.com")
	if err!= nil {
        panic(err)
    }
	fmt.Println(request.StatusCode)

	request.Body.Close()

}