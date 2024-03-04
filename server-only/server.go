package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"time"
)

type jsonPayload struct {
	petName string
}

var pet = jsonPayload{
	petName: "random",
}

type CustomHandler struct{}

func (c *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URL received: %s", r.URL.String())
	client := http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	resp, err := client.Get("https://catfact.ninja/fact/")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	io.WriteString(w, string(body))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func randomJsonHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, string("{\"Received url\":\""+r.URL.String()+"\"}"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func main() {

	r := mux.NewRouter()
	r.Handle("/cat-fact", &CustomHandler{})
	r.HandleFunc("/random-endpoint/{id}", randomJsonHandler)

	log.Println("Starting to listen on 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
