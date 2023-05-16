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

	defer request.Body.Close() //Atrasar a execucao so executando por ultimo a linha que tem o defer


}