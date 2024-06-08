package main

import (
	"fmt"
	"log"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)
	http.HandleFunc("/hello", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("content-type", "application/json")
		_, _ = resp.Write([]byte("hello"))
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
