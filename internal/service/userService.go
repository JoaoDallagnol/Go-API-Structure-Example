package service

import (
	"github.com/JoaoDallagnol/Go-API-Structure-Example.git/internal/model"
	"github.com/JoaoDallagnol/Go-API-Structure-Example.git/internal/repository"
)

// UserService fornece métodos para lógica de negócios relacionada a usuários
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService cria uma nova instância de UserService
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo}
}

// Create cria um novo usuário
func (s *UserService) Create(user *model.User) error {
	return s.userRepo.Create(user)
}

// GetAll retorna todos os usuários
func (s *UserService) GetAll() ([]model.User, error) {
	return s.userRepo.GetAll()
}

// GetByID retorna um usuário pelo ID
func (s *UserService) GetByID(id int64) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

// Update atualiza um usuário existente
func (s *UserService) Update(user *model.User) error {
	return s.userRepo.Update(user)
}

// Delete remove um usuário pelo ID
func (s *UserService) Delete(id int64) error {
	return s.userRepo.Delete(id)
}
