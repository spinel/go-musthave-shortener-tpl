package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/model"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/repository"
)

const Host = "http://localhost:8080"

func CreateUserHandler(repo *repository.Store) func(w http.ResponseWriter, r *http.Request) {
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
		user := &model.User{
			URL: url,
		}
		code, err := repo.User.SaveUser(user)
		if err != nil {
			http.Error(w, "save user error", http.StatusInternalServerError)
			return
		}
		result := fmt.Sprintf("%s/%s", Host, code)
		w.WriteHeader(201)
		w.Write([]byte(result))
	}
}

func GetUserHandler(repo *repository.Store) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			http.Error(w, "no id", http.StatusBadRequest)
			return
		}
		fmt.Println(id)
		url, err := repo.User.GetUserBy(id)
		if err != nil {
			http.Error(w, "save user error", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}
