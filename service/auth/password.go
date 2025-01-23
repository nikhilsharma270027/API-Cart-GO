package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// we will find the logged user by its hashedpassword
// comparing the user and hased pass in db
func ComparePasswords(hashed string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)
	return err == nil
}

// example
// err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)
// if err == nil {
//     // Passwords match
//     return true
// } else {
//     // Passwords don't match
//     return false
// }
