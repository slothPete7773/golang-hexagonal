package apihandler

import (
	"bank/errs"
	"fmt"
	"net/http"
)

func handleError(err error, w http.ResponseWriter) {
	switch e := err.(type) {
	case errs.AppError:
		w.WriteHeader(e.Code)
		fmt.Fprintln(w, e.Message)
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, e)
	}

}
