package api

import (
	"net/http"
	"encoding/json"
)

type APIHandler struct {
	dbHandler DbHandler
}

func NewAPIHandler(dbHandler DbHandler) *APIHandler {
	return &APIHandler{
		dbHandler,
	}
}

type RouteHandler interface {
	RedirectHandler(http.ResponseWriter, *http.Request)
	ShortenURLHandler(http.ResponseWriter, *http.Request)
}

func (handler *APIHandler) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Add Shorten logic
	json.NewEncoder(w).Encode(Counter{
		ID:4,
		Seq:78,
	})
}

func (handler *APIHandler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Add redirect logic
	nextId, _ := handler.dbHandler.GetNextId()
	json.NewEncoder(w).Encode(Counter{
		ID:4,
		Seq:nextId,
	})
}