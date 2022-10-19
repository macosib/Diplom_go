package server

import (
	"Diplom_Makarov/internal/handlers"
	parser_service "Diplom_Makarov/internal/parser-service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(r *mux.Router) {
	r.HandleFunc("/", h.handleConnection)
}

func (h *handler) handleConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res []byte
	res, _ = json.Marshal(parser_service.GetResultData())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
