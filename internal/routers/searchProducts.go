package routers

import (
	"cchallenge/internal/bd"
	"cchallenge/internal/models"
	"encoding/json"
	"net/http"
)

func SearchProducts(w http.ResponseWriter, r *http.Request) {
	// get params from the url
	sStr := r.URL.Query().Get("s")
	sId := r.URL.Query().Get("id")
	var kind int
	var s string

	if len(sId) < 1 && len(sStr) < 3 {
		kind = 3
		s = ""
	}
	if len(sId) > 0 && len(sStr) < 3 {
		kind = 1
		s = sId
	}
	if len(sId) < 1 && len(sStr) > 3 {
		kind = 2
		s = sStr
	}

	if len(sId) > 0 && len(sStr) > 3 {
		kind = 2
		s = sStr
	}

	if kind == 0 {
		encodeResponseAsJSON(w, models.Error{Message: "mistake committed in query params"}, http.StatusBadRequest)
		return
	}

	res, err := bd.SearchProductsByDescriptionBrand(s, kind)
	if err != nil {
		encodeResponseAsJSON(w, nil, http.StatusNoContent)
		return
	}
	encodeResponseAsJSON(w, res, http.StatusOK)
}

func encodeResponseAsJSON(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}

}
