package web

import (
	"fmt"
	"net/http"
	"strconv"
)

func WebHandler(w http.ResponseWriter, r *http.Request) {
	component := HelloForm("0")
	component.Render(r.Context(), w)
}

func HelloWebHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	name := r.FormValue("name")
	component := HelloPost(name)
	component.Render(r.Context(), w)
}

func IncrementWebHandler(w http.ResponseWriter, r *http.Request) {
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

	fmt.Printf("New Count: %d ... Prev: %d", count+1, count)
	component := CountResult(strconv.Itoa(count + 1))
	component.Render(r.Context(), w)
}
