package repository

import (
	"database/sql"

	"github.com/JoaoDallagnol/Go-API-Structure-Example.git/internal/model"
)

// AccountRepository fornece métodos para interagir com os dados da conta
type AccountRepository struct {
	db *sql.DB
}

// NewAccountRepository cria uma nova instância de AccountRepository
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db}
}

// Create insere uma nova conta no banco de dados
func (r *AccountRepository) Create(account *model.Account) error {
	query := "INSERT INTO accounts (number, balance, user_id) VALUES ($1, $2, $3) RETURNING id"
	return r.db.QueryRow(query, account.Number, account.Balance, account.UserID).Scan(&account.ID)
}

// GetAll retorna todas as contas do banco de dados
func (r *AccountRepository) GetAll() ([]model.Account, error) {
	query := "SELECT id, number, balance, user_id FROM accounts"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []model.Account
	for rows.Next() {
		var account model.Account
		if err := rows.Scan(&account.ID, &account.Number, &account.Balance, &account.UserID); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

// GetByID retorna uma conta pelo ID
func (r *AccountRepository) GetByID(id int64) (*model.Account, error) {
	query := "SELECT id, number, balance, user_id FROM accounts WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var account model.Account
	if err := row.Scan(&account.ID, &account.Number, &account.Balance, &account.UserID); err != nil {
		return nil, err
	}
	return &account, nil
}

// Update atualiza uma conta existente no banco de dados
func (r *AccountRepository) Update(account *model.Account) error {
	query := "UPDATE accounts SET number = $1, balance = $2, user_id = $3 WHERE id = $4"
	_, err := r.db.Exec(query, account.Number, account.Balance, account.UserID, account.ID)
	return err
}

// Delete remove uma conta pelo ID
func (r *AccountRepository) Delete(id int64) error {
	query := "DELETE FROM accounts WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
