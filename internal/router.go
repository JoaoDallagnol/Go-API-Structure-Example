package internal

import (
	"github.com/JoaoDallagnol/Go-API-Structure-Example.git/internal/controller"
	"github.com/gorilla/mux"
)

// SetupRouter configura as rotas da aplicação
func SetupRouter(userController *controller.UserController, accountController *controller.AccountController) *mux.Router {
	router := mux.NewRouter()

	// Rotas de Usuário
	router.HandleFunc("/users", userController.Create).Methods("POST")
	router.HandleFunc("/users", userController.GetAll).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", userController.GetByID).Methods("GET")
	router.HandleFunc("/users", userController.Update).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", userController.Delete).Methods("DELETE")

	// Rotas de Conta
	router.HandleFunc("/accounts", accountController.Create).Methods("POST")
	router.HandleFunc("/accounts", accountController.GetAll).Methods("GET")
	router.HandleFunc("/accounts/{id:[0-9]+}", accountController.GetByID).Methods("GET")
	router.HandleFunc("/accounts", accountController.Update).Methods("PUT")
	router.HandleFunc("/accounts/{id:[0-9]+}", accountController.Delete).Methods("DELETE")

	return router
}
