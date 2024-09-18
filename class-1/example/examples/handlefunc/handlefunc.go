package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	go func() {
		<-time.After(2 * time.Second)
		SimpleRequest()
		RequestWithAClient()
		RequestWithContext()
	}()

	LoadHandlers()
	StartServer()
}

func RequestWithAClient() {
	client := &http.Client{}
	r, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		panic(err)
	}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body := make([]byte, res.ContentLength)
	_, _ = res.Body.Read(body)
	log.Println("Response from /endpoint:", res.Status, "Body", string(body))
}

func SimpleRequest() {
	res, err := http.Get("http://localhost:8080/endpoint")
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		_, _ = io.Copy(io.Discard, Body)
		_ = Body.Close()
	}(res.Body)
	body := make([]byte, res.ContentLength)
	_, _ = res.Body.Read(body)
	log.Println("Response from /endpoint:", res.Status, "Body", string(body))
}

type Poster struct {
	Name   string `json:"name"`
	School string `json:"school"`
	Scar   bool   `json:"scar"`
}

func LoadHandlers() {
	http.HandleFunc("GET /hello/{name...}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		if name == "" {
			name = "World"
		}
		_, _ = io.WriteString(w, fmt.Sprintf("Hello %s from a HandleFunc #1!\n", name))
	})
	http.HandleFunc("POST /hello/poster", func(w http.ResponseWriter, r *http.Request) {
		defer func(Body io.ReadCloser) {
			_, _ = io.Copy(io.Discard, Body)
		}(r.Body)

		var poster Poster
		if err := json.NewDecoder(r.Body).Decode(&poster); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, _ = io.WriteString(w, fmt.Sprintf("Hello %v from a HandleFunc #2!\n", poster))
	})
}

func StartServer() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func RequestWithContext() {
	client := &http.Client{}
	poster := Poster{
		Name:   "Harry",
		School: "Hogwarts",
		Scar:   true,
	}
	posterJSON, err := json.Marshal(poster)
	if err != nil {
		panic(err)
	}
	reader := bytes.NewReader(posterJSON)

	ctx := context.Background()
	r1, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:8080/hello/poster", reader)
	if err != nil {
		panic(err)
	}
	res1, err := client.Do(r1)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		_, _ = io.Copy(io.Discard, Body)
	}(res1.Body)
	body := make([]byte, res1.ContentLength)
	_, _ = res1.Body.Read(body)
	log.Println("Response from /:", res1.Status, "Body", string(body))
}
