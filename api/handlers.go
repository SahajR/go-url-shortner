package api

import (
	"net/http"
	"encoding/json"
	"URL-Shortner/utils"
	"github.com/gorilla/mux"
)

type APIHandler struct {
	dbHandler DbHandler
}

type ResultOk struct {
	ShortURL string
}

type ResultFailure struct {
	ErrorMessage string
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
	url, ok := r.URL.Query()["url"]
	if ok {
		existingURL, existError := handler.dbHandler.GetURLByLongURL(url[0])
		if existError != nil {
			nextId, _ := handler.dbHandler.GetNextId()
			err := handler.dbHandler.AddURL(URL{
				ID:nextId,
				LongURL:url[0],
				Clicks:0,
			})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(ResultFailure{
					ErrorMessage: "Internal server error",
				})
			} else {
				json.NewEncoder(w).Encode(ResultOk{
					ShortURL:utils.GetURL(r.Host, utils.Encode(nextId)),
				})
			}
		} else {
			json.NewEncoder(w).Encode(ResultOk{
				ShortURL:utils.GetURL(r.Host, utils.Encode(existingURL.ID)),
			})
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResultFailure{
			ErrorMessage: "Bad request: Please provide URL",
		})
	}
}

func (handler *APIHandler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if ok {
		url, err := handler.dbHandler.GetURLById(utils.Decode(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ResultFailure{
				ErrorMessage: "There's no such URL",
			})
		} else {
			handler.dbHandler.UpdateClickCount(url.ID)
			http.Redirect(w, r, url.LongURL, http.StatusMovedPermanently)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResultFailure{
			ErrorMessage: "Bad request: No ID found",
		})
	}
}
