package routers

import (
	"encoding/json"
	"guolmal/internal/bd"
	"guolmal/internal/models"
	"net/http"
)

func SearchProducts(w http.ResponseWriter, r *http.Request) {
	// get params from the url
	sStr := r.URL.Query().Get("s")
	sId := r.URL.Query().Get("id")
	var kind string
	var s string

	if len(sId) < 1 && len(sStr) < 3 {
		encodeResponseAsJSON(w, models.Error{Message: "mistake committed in query params"}, http.StatusBadRequest)
		return
	}
	if len(sId) > 0 && len(sStr) < 3 {
		kind = "id"
		s = sId
	}
	if len(sId) < 1 && len(sStr) > 3 {
		kind = "s"
		s = sStr
	}

	if len(kind) == 0 {
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
	json.NewEncoder(w).Encode(res)

}
