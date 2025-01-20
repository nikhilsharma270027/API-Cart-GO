package user

import (
	"database/sql"

	"github.com/nikhilsharma270027/API-Cart-GO/types"
)

type Store struct {
	db *sql.DB
} // we this to create database query

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User) // scannning the user
	for rows.Next() {    // looking over rows

	}

}
