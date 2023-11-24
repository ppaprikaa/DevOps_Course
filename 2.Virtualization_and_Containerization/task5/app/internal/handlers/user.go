package handlers

import (
	"app/internal/model"
	"app/internal/storage"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

type userServiceHandlers struct {
	storage *storage.UserStorage
}

func NewUserServiceHandler(storage *storage.UserStorage) *userServiceHandlers {
	return &userServiceHandlers{
		storage: storage,
	}
}

func (u *userServiceHandlers) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			offset    int
			limit     int
			offsetStr = strings.TrimSpace(r.URL.Query().Get("offset"))
			limitStr  = strings.TrimSpace(r.URL.Query().Get("limit"))
		)

		if offsetStr == "" {
			offset = 0
		} else {
			offset, _ = strconv.Atoi(offsetStr)
		}

		if limitStr == "" {
			limit = 1000
		} else {
			limit, _ = strconv.Atoi(limitStr)
		}

		queryCtx, queryCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer queryCancel()
		users, err := u.storage.QueryAll(queryCtx, offset, limit)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Some error occured"))
			return
		}

		res := strings.Builder{}
		for _, u := range users {
			res.WriteString(fmt.Sprintf("%d - %s\n", u.ID, u.Username))
		}

		w.Write([]byte(res.String()))
	}
}

func (u *userServiceHandlers) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err   error
			id    int
			idStr = strings.TrimSpace(chi.URLParam(r, "id"))
		)

		if idStr == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("no id given"))
			return
		}

		id, err = strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id not a number"))
			return
		}

		deleteCtx, deleteCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer deleteCancel()
		if err = u.storage.DeleteOne(deleteCtx, id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Some error occured"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User deleted"))
	}
}

func (u *userServiceHandlers) Insert() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			user     *model.User
			err      error
			username = strings.TrimSpace(chi.URLParam(r, "username"))
		)

		if username == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("username cannot be empty"))
			return
		}
		insertCtx, insertCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer insertCancel()

		if user, err = u.storage.InsertOne(insertCtx, username); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("couldn't insert user"))
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User was added. User: %d - %s", user.ID, user.Username)
	}
}
