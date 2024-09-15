package main

import (
	"context"
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
	defer res.Body.Close()
	body := make([]byte, res.ContentLength)
	_, _ = res.Body.Read(body)
	log.Println("Response from /endpoint:", res.Status, "Body", string(body))
}

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	})
	http.HandleFunc("/endpoint", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func RequestWithContext() {
	ctx := context.Background()

	client := &http.Client{}

	r1, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/", nil)
	if err != nil {
		panic(err)
	}
	res1, err := client.Do(r1)
	if err != nil {
		panic(err)
	}
	defer res1.Body.Close()
	body := make([]byte, res1.ContentLength)
	_, _ = res1.Body.Read(body)
	log.Println("Response from /:", res1.Status, "Body", string(body))
}
