package services

import (
	"time"

	"github.com/gbrlsnchs/jwt/v3"
)

// JWTPayload the struct of jwt data with claims
type JWTPayload struct {
	jwt.Payload
	ID   int
	Role int8
}

// GenerateToken generates the claims needed
func GenerateToken(secret string, ID int, role int8, duration time.Duration) (string, error) {
	now := time.Now()
	payload := JWTPayload{
		Payload: jwt.Payload{
			Issuer:         "App",
			IssuedAt:       jwt.NumericDate(now),
			ExpirationTime: jwt.NumericDate(now.Add(duration)),
		},
		ID:   ID,
		Role: role,
	}
	token, err := jwt.Sign(payload, jwt.NewHS256([]byte(secret)))
	if err != nil {
		return "", err
	}
	return string(token), nil
}

// CheckToken takes a token string and returns the ID claim
func CheckToken(secret string, token string) (int, int8, error) {
	jwtPayload := JWTPayload{}
	expValidator := jwt.ExpirationTimeValidator(time.Now())
	_, err := jwt.Verify([]byte(token), jwt.NewHS256([]byte(secret)), &jwtPayload, jwt.ValidatePayload(&jwtPayload.Payload, expValidator))
	if err != nil {
		return 0, 0, err
	}
	return jwtPayload.ID, jwtPayload.Role, err
}
