package common

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken creates a JSON Web Token (JWT) with the specified TTL, payload, and key.
func CreateToken(ttl time.Duration, payload interface{}, key string) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(ttl).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = payload

	// Sign the token with the key
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}

// ValidateToken validates a JSON Web Token (JWT) using the specified key.
func ValidateToken(tokenString string, key string) (interface{}, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the key for validation
		return []byte(key), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("error extracting claims")
	}

	// You can access the individual claims here
	subject := claims["sub"]

	return subject, nil
}
