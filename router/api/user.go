package api

import (
	"encoding/json"
	"net/http"

	"github.com/jacexh/golang-ddd-template/internal/application"
	"github.com/jacexh/golang-ddd-template/router/dto"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	u := new(dto.User)
	_ = json.NewDecoder(r.Body).Decode(u)
	_ = application.User.CreateUser(r.Context(), u)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(u)
}
