package user

import (
	"database/sql"
	"fmt"

	"github.com/nikhilsharma270027/API-Cart-GO/types"
)

type Store struct {
	db *sql.DB
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
		return nil, fmt.Errorf("user not found by Email")
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

// CreateUser implements types.UserStore.
func (s *Store) CreateUser(user types.User) error {
	// panic("unimplemented")
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email,password) VALUES (?,?,?,?)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil //if works fine
}

// GetUserByID implements types.UserStore.
func (s *Store) GetUserByID(id int) (*types.User, error) {
	// panic("unimplemented")
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
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
		return nil, fmt.Errorf("user not found by ID")
	}
	fmt.Print(u)
	return u, nil
}
