package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spinel/go-musthave-shortener-tpl/internal/app/helper"
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
		var code helper.GeneratedString
		for {
			code, err = helper.NewGeneratedString()
			if err != nil {
				http.Error(w, "save user error", http.StatusInternalServerError)
				return
			}
			if !repo.User.IncludesCode(string(code)) {
				break
			}
		}
		codeString := string(code)
		err = repo.User.SaveUser(codeString, user)
		if err != nil {
			http.Error(w, "user exists", http.StatusInternalServerError)
			return
		}
		result := fmt.Sprintf("%s/%s", Host, code)
		w.Header().Add("Content-type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(result))
	}
}

func GetUserHandler(repo *repository.Store) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()
		id := url[1:]
		if id == "" {
			http.Error(w, "no id", http.StatusBadRequest)
			return
		}
		user, err := repo.User.GetUserBy(id)
		if err != nil {
			http.Error(w, "get user error", http.StatusInternalServerError)
			return
		}
		if user == nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, user.URL, http.StatusTemporaryRedirect)
	}
}
