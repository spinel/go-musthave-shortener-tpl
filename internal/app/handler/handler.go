package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/helper"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/storage"
)

const Host = "http://localhost:8080"

func Create(s storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "wrong body", http.StatusBadRequest)
			return
		}
		url := string(body)
		if url == "" {
			http.Error(w, "no body", http.StatusBadRequest)
			return
		}
		generatedString, err := helper.NewStringGenerator()
		if err != nil {
			http.Error(w, "generate code error", http.StatusBadRequest)
			return
		}
		s[generatedString.Value] = url
		result := fmt.Sprintf("%s/%s", Host, generatedString.Value)
		w.Write([]byte(result))
	}
}

func Get(s storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			http.Error(w, "no id", http.StatusBadRequest)
			return
		}
		if val, ok := s[id]; ok {
			http.Redirect(w, r, val, http.StatusTemporaryRedirect)
			return
		}
		http.Error(w, "not exists", http.StatusBadRequest)
	}
}
