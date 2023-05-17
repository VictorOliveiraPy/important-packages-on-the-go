package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	http.HandleFunc("/", SearchCepHandler)
	http.ListenAndServe(":8080", nil) // 
}


func SearchCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return

	}
	cepParams := r.URL.Query().Get("cep")
	if cepParams == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		
	cep, error := SearchCep(cepParams)
	if error!= nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
}

func SearchCep(cep string) (*ViaCep, error) {
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error!= nil {
			return nil, error
		}
	
		defer resp.Body.Close()
		body, error := ioutil.ReadAll(resp.Body)
		if error!= nil {
            return nil, error
        }
		var c ViaCep
		error = json.Unmarshal(body, &c) // apontamento na var c &
		if error!= nil {
            return nil, error
        }
		
		return &c, nil
}
