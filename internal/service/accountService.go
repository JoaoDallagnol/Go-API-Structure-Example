package service

import (
	"github.com/JoaoDallagnol/Go-API-Structure-Example.git/internal/model"
	"github.com/JoaoDallagnol/Go-API-Structure-Example.git/internal/repository"
)

// AccountService fornece métodos para lógica de negócios relacionada a contas
type AccountService struct {
	accountRepo *repository.AccountRepository
}

// NewAccountService cria uma nova instância de AccountService
func NewAccountService(accountRepo *repository.AccountRepository) *AccountService {
	return &AccountService{accountRepo}
}

// Create cria uma nova conta
func (s *AccountService) Create(account *model.Account) error {
	return s.accountRepo.Create(account)
}

// GetAll retorna todas as contas
func (s *AccountService) GetAll() ([]model.Account, error) {
	return s.accountRepo.GetAll()
}

// GetByID retorna uma conta pelo ID
func (s *AccountService) GetByID(id int64) (*model.Account, error) {
	return s.accountRepo.GetByID(id)
}

// Update atualiza uma conta existente
func (s *AccountService) Update(account *model.Account) error {
	return s.accountRepo.Update(account)
}

// Delete remove uma conta pelo ID
func (s *AccountService) Delete(id int64) error {
	return s.accountRepo.Delete(id)
}
