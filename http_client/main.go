package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)



func main() {
	postClien()
	client := http.Client{Timeout: time.Second}
	resp, err := client.Get("http://google.com")
	if err!= nil {
        panic(err)
    }
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err!= nil {
        panic(err)
    }

	println(string(body))

}

func postClien(){
	client := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`"{name:": "victor"}`))
	resp, err := client.Post("http://google.com", "application/json", jsonVar)
	if err!= nil {
        panic(err)
    }
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil) // pegar os dados e escolhe onde vai jogar, vai jogar no stdou e vai copia no body
	// buffer e um espaco cheio de informacao
}

func postCustom(){
	cliente := http.Client{}
	req, err := http.NewRequest("GET", "http.google.com", nil) 
	if err!= nil {
        panic(err)
    }
	req.Header.Set("Accept", "Application/json")
	resp, err := cliente.Do(req) //Custom http client
	if err!= nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err!= nil {
        panic(err)
    }
	println(string(body))

}