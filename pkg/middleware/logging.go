package middleware

import (
	"net/http"

	"github.com/boilerplate/pkg/logger"
	"github.com/julienschmidt/httprouter"
)

func LogRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		logger.Info.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next(w, r, p)
	}
}
