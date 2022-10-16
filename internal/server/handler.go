package server

import (
	"Diplom_Makarov/internal/handlers"
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
	//if err != nil {
	//	res, _ = json.Marshal(map[string]string{"status": err.Error()})
	//	w.WriteHeader(http.StatusBadRequest)
	//	w.Write(res)
	//	return
	//}
	res, _ = json.Marshal("ok")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
