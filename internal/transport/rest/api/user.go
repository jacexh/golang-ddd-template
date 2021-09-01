package api

import (
	"encoding/json"
	"net/http"

	"github.com/jacexh/golang-ddd-template/internal/application"
	"github.com/jacexh/golang-ddd-template/internal/transport/dto"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	u := new(dto.User)
	_ = json.NewDecoder(r.Body).Decode(u)
	_ = application.User.CreateUser(r.Context(), u)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(u)
}
