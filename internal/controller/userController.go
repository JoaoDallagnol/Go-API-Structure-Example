package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JoaoDallagnol/Go-API-Structure-Example.git/internal/model"
	"github.com/JoaoDallagnol/Go-API-Structure-Example.git/internal/service"
)

// UserController fornece métodos para lidar com requisições de usuários
type UserController struct {
	userService *service.UserService
}

// NewUserController cria uma nova instância de UserController
func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService}
}

// Create cria um novo usuário
func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.userService.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetAll retorna todos os usuários
func (c *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := c.userService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetByID retorna um usuário pelo ID
func (c *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := c.userService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Update atualiza um usuário existente
func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.userService.Update(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Delete remove um usuário pelo ID
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.userService.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
