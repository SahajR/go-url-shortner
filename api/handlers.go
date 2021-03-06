package api

import (
	"net/http"
	"encoding/json"
	"URL-Shortner/utils"
	"github.com/gorilla/mux"
	netURL "net/url"
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

// The Handler for the `shorten` api route.
// The URL to be shortened must be passed as a query parameter.
func (handler *APIHandler) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	longURL, ok := r.URL.Query()["url"]
	if ok {
		_, urlParseError := netURL.ParseRequestURI(longURL[0])
		if urlParseError != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ResultFailure{
				ErrorMessage: "Invalid URL",
			})
			return
		}
		// Check if the long URL is already in the database.
		// Proceed to retrieve its ID and generate the short code if it does.
		existingURL, existError := handler.dbHandler.GetURLByLongURL(longURL[0])
		if existError != nil {
			// The given URL doesn't exist in the Db. Add it and retrieve the ID.
			nextId, idError := handler.dbHandler.GetNextId()
			if nextId == -1 {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(ResultFailure{
					ErrorMessage: "Internal server error. " + idError.Error(),
				})
				return
			}
			err := handler.dbHandler.AddURL(URL{
				ID:nextId,
				LongURL:longURL[0],
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
			ErrorMessage: "Bad request: Please provide a valid URL",
		})
	}
}


// The Handler for redirection.
func (handler *APIHandler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if ok {
		url, err := handler.dbHandler.GetURLById(utils.Decode(id))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ResultFailure{
				ErrorMessage: "There's no such URL",
			})
		} else {
			http.Redirect(w, r, url.LongURL, http.StatusMovedPermanently)
			// Silently update click count of the URL
			handler.dbHandler.UpdateClickCount(url.ID)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResultFailure{
			ErrorMessage: "Bad request: No ID found",
		})
	}
}
