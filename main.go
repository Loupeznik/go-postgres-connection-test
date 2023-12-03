package main

import (
	"fmt"
	"io"
	"net/http"

	database "github.com/loupeznik/go-postgres-connection-test/src"
)

func main() {
	http.HandleFunc("/connect", connect)
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		panic(err)
	}
}

func connect(w http.ResponseWriter, r *http.Request) {
	connectionTestResult, err := database.Connect()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, fmt.Sprintf("Connection test failed: %v", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fmt.Sprintf("Connection test result: %d", connectionTestResult))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Healthy")
}
