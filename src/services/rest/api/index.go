package api

import (
	"cupcake-core/src/logger"
	"fmt"
	"net/http"
)

func Index(l logger.HTTPRequestLogger, w http.ResponseWriter, r *http.Request) {

	l.Response200(r, "", "Ok")
	fmt.Fprintf(w, "Ok")
}
