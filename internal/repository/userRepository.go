package repository

import (
	"database/sql"

	"github.com/JoaoDallagnol/Go-API-Structure-Example.git/internal/model"
)

// UserRepository fornece métodos para interagir com os dados do usuário
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository cria uma nova instância de UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

// Create insere um novo usuário no banco de dados
func (r *UserRepository) Create(user *model.User) error {
	query := "INSERT INTO users (name, email, age, password) VALUES ($1, $2, $3, $4) RETURNING id"
	return r.db.QueryRow(query, user.Name, user.Email, user.Age, user.Password).Scan(&user.ID)
}

// GetAll retorna todos os usuários do banco de dados
func (r *UserRepository) GetAll() ([]model.User, error) {
	query := "SELECT id, name, email, age, password FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetByID retorna um usuário pelo ID
func (r *UserRepository) GetByID(id int64) (*model.User, error) {
	query := "SELECT id, name, email, age, password FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var user model.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
}

// Update atualiza um usuário existente no banco de dados
func (r *UserRepository) Update(user *model.User) error {
	query := "UPDATE users SET name = $1, email = $2, age = $3, password = $4 WHERE id = $5"
	_, err := r.db.Exec(query, user.Name, user.Email, user.Age, user.Password, user.ID)
	return err
}

// Delete remove um usuário pelo ID
func (r *UserRepository) Delete(id int64) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
