package main

import (
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/apache/skywalking-go"
	"github.com/apache/skywalking-go/toolkit/trace"
)

func ServerHTTP(writer http.ResponseWriter, request *http.Request) {
	time.Sleep(time.Duration(500) * time.Millisecond)
	req, err := http.NewRequest(http.MethodPost, "upstream-service", http.NoBody)
	trace.SetCorrelation("key", "value")
	client := &http.Client{}
	resp, err := client.Do(req)
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
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
