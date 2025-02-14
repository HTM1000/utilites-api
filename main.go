package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HTM1000/table-inss/internal/handlers"
)

func main() {
	http.HandleFunc("/api/inss", handlers.InssHandler)
	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
