package user

import (
	"database/sql"
	"fmt"

	"github.com/nikhilsharma270027/API-Cart-GO/types"
)

type Store struct {
	db *sql.DB
}

// CreateUser implements types.UserStore.
func (s *Store) CreateUser(types.User) error {
	panic("unimplemented")
}

// GetUserByID implements types.UserStore.
func (s *Store) GetUserByID(id int) (*types.User, error) {
	panic("unimplemented")
}

// we this to create database query

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
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		// u = {ID: 0,Name: "",Email: "",Age: 0} if we get empty
		return nil, fmt.Errorf("user not found")
	}
	fmt.Print(u)
	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {

	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, err
}

//
