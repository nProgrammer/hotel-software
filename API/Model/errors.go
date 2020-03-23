package Model

import (
	"fmt"
	"net/http"
)

func BadToken(w http.ResponseWriter) {
	fmt.Fprintf(w, "Bad token verification")
}
