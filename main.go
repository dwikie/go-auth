package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dwikie/go-auth/config"
)

func main() {
	config.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server up!"))
	})

	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", config.Cfg.App.Port),
		Handler: mux,
	}

	defer server.Close()
	fmt.Printf("server running on: http://%v", server.Addr)
	log.Fatal(server.ListenAndServe())
}
