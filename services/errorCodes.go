package services

const (
	// BadFields error code when request has errors
	BadFields = "BAD_FIELDS"
	// EmailExists to indicate email already exists
	EmailExists = "EMAIL_EXISTS"
	// EmailNotExists to indicate email not exists
	EmailNotExists = "EMAIL_NOT_EXISTS"
	// ExpiredToken when JWT is expired
	ExpiredToken = "EXPIRED_TOKEN"
	// InvalidToken when token is invalid
	InvalidToken = "INVALID_TOKEN"
)
