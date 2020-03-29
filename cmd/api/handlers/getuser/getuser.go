package getuser

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Do(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello")
}
