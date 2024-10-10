package model

// Account representa o modelo de uma conta
type Account struct {
	ID      int64   `json:"id"`
	Number  string  `json:"number"`
	Balance float64 `json:"balance"`
	UserID  int64   `json:"user_id"` // Relacionamento com User
}
