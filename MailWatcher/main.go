package main

import (
	"fmt"
	"log"
	"net/http"
)

func mailWatcher(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Servidor funcionando"))

	codes, ok := r.URL.Query()["code"]

	if !ok || len(codes) < 1 {
		log.Println("Url Param 'code' is missing")
		return
	}

	ControllerStart(codes[0])
}

func setupRoutes() {
	http.HandleFunc("/mainPage", mailWatcher)
}

func main() {
	fmt.Println("Mail Watcher v0.1")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
