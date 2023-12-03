package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Work      string `json:"work"`
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		Request(w, r)
	case http.MethodPost:
		Request(w, r)
	default:
		w.Header().Set("Allow", http.MethodGet+", "+http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}

}

func Request(w http.ResponseWriter, r *http.Request) {
	// Retorna um slice de bytes ([]byte, error)
	resp, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	// pega o post json e converte para struct
	var itens Client
	err = json.Unmarshal(resp, &itens)

	if err != nil {
		panic(err)
	}

	// Pegar por form-data: r.PostFormValue("lastName")
	client := Client{FirstName: itens.FirstName,
		Work:     itens.Work,
		LastName: itens.LastName}
	// Pega struct e convert para json
	jsonParse, err := json.Marshal(client)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonParse)
}
