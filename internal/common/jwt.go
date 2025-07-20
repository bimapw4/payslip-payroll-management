package common

import (
	"errors"
	"fmt"
	"os"
	"payslips/internal/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCode struct {
	secretToken string
	lifetime    string
}

func NewJwt() JwtCode {
	return JwtCode{
		secretToken: os.Getenv("JWT_SECRET_KEY"),
		lifetime:    os.Getenv("JWT_LIFESPAN"),
	}
}

func (s *JwtCode) GenerateAuthorizartionCode(payload entity.Claim) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = payload.UserID
	claims["username"] = payload.Username
	exp, _ := time.ParseDuration(s.lifetime)
	claims["exp"] = time.Now().Add(exp).Unix() // Token expires in 1 hour

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(s.secretToken))
	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}

	// Print the encoded token
	return tokenString, nil
}

func (s *JwtCode) DecodeAccessToken(accesstoken string) (*entity.Claim, error) {

	// Parse the token
	token, err := jwt.Parse(accesstoken, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretToken), nil
	})

	// Check for errors
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		fmt.Println("Invalid token")
		return nil, errors.New("invalid_token")
	}

	// Access claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Error accessing claims")
		return nil, errors.New("error_claim")
	}

	// Check expiration time
	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(expirationTime) {
		fmt.Println("Token has expired")
		return nil, errors.New("token_expired")
	}

	exp := 0
	err = DecodeData(claims["exp"], &exp)
	if err != nil {
		return nil, err
	}

	return &entity.Claim{
		UserID:   fmt.Sprint(claims["user_id"]),
		Username: fmt.Sprint(claims["username"]),
		Exp:      exp,
	}, nil
}
