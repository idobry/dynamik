package api

import (
	"github.com/idobry/dynamik/app"
	"net/http"
)

func (a *API) ByeHandler(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	_, err := w.Write([]byte(`{"bye" : "world"}`))
	return err
}
