package waffel

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetRouteParamAsUint64(r *http.Request, paramName string) (uint64, error) {
	params := mux.Vars(r)
	param, err := strconv.ParseUint(string(params[paramName]), 10, 64)

	if err != nil {
		return 0, err
	}

	return param, nil
}

func ErrorResponse(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func RequestParams(r *http.Request, names ...string) map[string]string {
	fp := map[string]string{}
	params := mux.Vars(r)
	for _, name := range names {
		if r.FormValue(name) != "" {
			fp["form_"+name] = r.FormValue(name)
		}

		if p, ok := params[name]; ok {
			fp["url_"+name] = p
		}
	}
	return fp
}
