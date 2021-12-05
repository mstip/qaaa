package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getRouteParamAsUint64(r *http.Request, paramName string) (uint64, error) {
	params := mux.Vars(r)
	param, err := strconv.ParseUint(string(params[paramName]), 10, 64)

	if err != nil {
		return 0, err
	}

	return param, nil
}

func errorResponse(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
