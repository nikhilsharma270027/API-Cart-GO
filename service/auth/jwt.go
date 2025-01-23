package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nikhilsharma270027/API-Cart-GO/config"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	// func jwt.NewWithClaims(method jwt.SigningMethod, claims jwt.Claims, opts ...jwt.TokenOption) *jwt.Token
	//NewWithClaims creates a new [Token] with the specified signing method and claims.
	//Additional options can be specified, but are currently unused.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	// we are creating jwt token
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
