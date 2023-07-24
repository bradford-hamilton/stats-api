package server

import "net/http"

func (a *API) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
