package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jhon-Henkel/go_lang_api_example/tree/main/internal/dto"
	"github.com/Jhon-Henkel/go_lang_api_example/tree/main/internal/entity"
	"github.com/Jhon-Henkel/go_lang_api_example/tree/main/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB       database.UserInterface
	JwtExpiresIn int
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: db}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpires := r.Context().Value("jwtExpires").(int)
	var jwtInput dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&jwtInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := h.UserDB.FindByEmail(jwtInput.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !user.ComparePassword(jwtInput.Password) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, token, _ := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpires)).Unix(),
	})
	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userToInsert, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.UserDB.Create(userToInsert)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
