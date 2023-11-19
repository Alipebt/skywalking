package main

import (
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/apache/skywalking-go"
)

func ServerHTTP(writer http.ResponseWriter, request *http.Request) {
	time.Sleep(time.Duration(500) * time.Millisecond)
	//req, err := http.NewRequest(http.MethodPost, "oap:11800", http.NoBody)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	_, _ = writer.Write(body)
}

func main() {
	http.HandleFunc("/", ServerHTTP)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
