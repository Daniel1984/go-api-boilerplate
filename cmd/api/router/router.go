package router

import (
	"github.com/boilerplate/cmd/api/handlers/getuser"
	"github.com/boilerplate/pkg/application"
	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/", getuser.Do)
	return mux
}
