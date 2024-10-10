package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JoaoDallagnol/Go-API-Structure-Example.git/internal/model"
	"github.com/JoaoDallagnol/Go-API-Structure-Example.git/internal/service"
)

// AccountController fornece métodos para lidar com requisições de contas
type AccountController struct {
	accountService *service.AccountService
}

// NewAccountController cria uma nova instância de AccountController
func NewAccountController(accountService *service.AccountService) *AccountController {
	return &AccountController{accountService}
}

// Create cria uma nova conta
func (c *AccountController) Create(w http.ResponseWriter, r *http.Request) {
	var account model.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.accountService.Create(&account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

// GetAll retorna todas as contas
func (c *AccountController) GetAll(w http.ResponseWriter, r *http.Request) {
	accounts, err := c.accountService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

// GetByID retorna uma conta pelo ID
func (c *AccountController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := c.accountService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

// Update atualiza uma conta existente
func (c *AccountController) Update(w http.ResponseWriter, r *http.Request) {
	var account model.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.accountService.Update(&account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Delete remove uma conta pelo ID
func (c *AccountController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.accountService.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
