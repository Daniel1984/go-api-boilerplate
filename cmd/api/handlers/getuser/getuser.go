package getuser

import (
	"fmt"
	"net/http"

	"github.com/boilerplate/pkg/application"
	"github.com/julienschmidt/httprouter"
)

func Do(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "hello")
	}
}
