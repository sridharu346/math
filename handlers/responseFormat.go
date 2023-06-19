package handlers

import (
	"encoding/json"
	"net/http"
	schema "sridharu346/calculations/schema"
)

func ResponseFormat(w http.ResponseWriter, Prime string, Sum int, Product int, Message string, StatusCode int) {
	response := schema.Response{
		Prime:      Prime,
		Sum:        Sum,
		Product:    Product,
		Message:    Message,
		StatusCode: StatusCode,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		ResponseFormat(w, "NA", 0, 0, "StatusMethodNotAllowed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	w.Write(jsonResponse)

}
