package usecase

type Auth interface {
}

type AuthUsecase struct{}

func NewAuthUsecase() *AuthUsecase {
	return &AuthUsecase{}
}
