package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:]{
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err!= nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisicao: %v", err)
        }
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err!= nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v", err)
		}

		var data ViaCep
		err = json.Unmarshal(res, &data)
		if err!= nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v", err)
		}
		file, err := os.Create("cidade.txt")
		if err!= nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v", err)
		}
		defer file.Close()

		_, err = file.WriteString(data.Logradouro)
	}

}