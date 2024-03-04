package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{Addr: ":3000"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
		w.Write([]byte("Hello World!"))
	})

	fmt.Println("Server is running at http://localhost:3000")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
