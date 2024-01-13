package repository

type Auth interface {
}

type AuthRepository struct {
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}
