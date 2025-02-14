package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HTM1000/table-inss/internal/services"
)

func InssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	tabelas := services.ScrapeINSS()
	if len(tabelas) == 0 {
		log.Println("Nenhum dado foi encontrado")
		http.Error(w, "Nenhum dado encontrado", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(tabelas)
	if err != nil {
		log.Println("Erro ao codificar JSON:", err)
		http.Error(w, "Erro ao processar dados", http.StatusInternalServerError)
		return
	}
}
