package web

import (
	"net/http"
	"strconv"
)

func HelloWebHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	name := r.FormValue("name")
	component := HelloPost(name)
	component.Render(r.Context(), w)
}

func HelloWebIncrementHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	var previousCount = r.FormValue("count")
	if previousCount == "" {
		previousCount = "0"
	}

	count, err := strconv.Atoi(previousCount)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	component := CountResult(strconv.Itoa(count + 1))
	component.Render(r.Context(), w)
}
