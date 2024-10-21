package interfaces

import "github.com/deecodeid/api_nowted/domain/entities"

type TokenVerificationRepository interface {
	GenerateToken(userId string, tokeType string) string
	FindToken(token, email string) (*entities.TokenVerification, error)
	UpdateToken(*entities.TokenVerification) error
	FindLatestToken(userId string) (*entities.TokenVerification, error)
}
