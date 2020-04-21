package getuser

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/boilerplate/cmd/api/models"
	"github.com/julienschmidt/httprouter"
)

func validateRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		uid := p.ByName("id")

		id, err := strconv.Atoi(uid)
		if err != nil {
			w.WriteHeader(http.StatusPreconditionFailed)
			fmt.Fprintf(w, "malformed id")
			return
		}

		ctx := context.WithValue(r.Context(), models.CtxKey("userid"), id)
		r = r.WithContext(ctx)
		next(w, r, p)
	}
}
