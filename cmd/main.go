package main

import (
	"log"
	"net/http"

	"github.com/JoaoDallagnol/Go-API-Structure-Example.git/internal"
	"github.com/JoaoDallagnol/Go-API-Structure-Example.git/internal/controller"
	"github.com/JoaoDallagnol/Go-API-Structure-Example.git/internal/repository"
	"github.com/JoaoDallagnol/Go-API-Structure-Example.git/internal/service"
	"github.com/JoaoDallagnol/Go-API-Structure-Example.git/pkg"
)

func main() {
	// Configurações do banco de dados
	dataSourceName := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable"

	db, err := pkg.Connect(dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Repositórios
	userRepo := repository.NewUserRepository(db)
	accountRepo := repository.NewAccountRepository(db)

	// Serviços
	userService := service.NewUserService(userRepo)
	accountService := service.NewAccountService(accountRepo)

	// Controladores
	userController := controller.NewUserController(userService)
	accountController := controller.NewAccountController(accountService)

	// Configuração das rotas
	router := internal.SetupRouter(userController, accountController)

	// Inicia o servidor
	log.Println("Iniciando o servidor na porta 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
