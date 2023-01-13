package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Iniciando Web Server...")

	http.HandleFunc("/", serveIndex)
	err := http.ListenAndServe(":8000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}

func serveIndex(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Teste"))
}
