package getuser

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/boilerplate/cmd/api/models"
	"github.com/boilerplate/pkg/application"
	"github.com/boilerplate/pkg/middleware"
	"github.com/julienschmidt/httprouter"
)

func getUser(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		id := r.Context().Value(models.CtxKey("userid"))
		user := &models.User{ID: id.(int)}

		if err := user.GetByID(r.Context(), app); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusPreconditionFailed)
				fmt.Fprintf(w, "user does not exist")
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Oops")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(user)
		w.Write(response)
	}
}

func Do(app *application.Application) httprouter.Handle {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		validateRequest,
	}

	return middleware.Chain(getUser(app), mdw...)
}
