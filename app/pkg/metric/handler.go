package metric

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handler struct {
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, "/api/heartbeat", h.HeartBeat)
}

func (h *Handler) HeartBeat(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
