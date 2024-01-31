package auth

import (
	"context"
	"log"

	"github.com/dickidarmawansaputra/go-simple-online-shop/infra/response"
	"github.com/dickidarmawansaputra/go-simple-online-shop/internal/config"
)

type Repository interface {
	GetAuthByEmail(ctx context.Context, email string) (model AuthEntity, err error)
	CreateAuth(ctx context.Context, model AuthEntity) (err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) register(ctx context.Context, req RegisterRequestPayload) (err error) {
	authEntity := NewFromRegisterRequest(req)

	if err = authEntity.Validate(); err != nil {
		return
	}

	if err = authEntity.EncryptPassword(int(config.Cfg.App.Encryption.Salt)); err != nil {
		return
	}

	model, err := s.repo.GetAuthByEmail(ctx, authEntity.Email)
	if err != nil {
		if err == response.ErrNotFound {
			return
		}
	}

	if model.IsExists() {
		return response.ErrEmailAlreadyUsed
	}

	return s.repo.CreateAuth(ctx, authEntity)
}

func (s service) login(ctx context.Context, req LoginRequestPayload) (token string, err error) {
	authEntity := NewFromLoginRequest(req)

	if err = authEntity.ValidateEmail(); err != nil {
		return
	}

	if err = authEntity.ValidatePassword(); err != nil {
		return
	}

	model, err := s.repo.GetAuthByEmail(ctx, authEntity.Email)
	if err != nil {
		log.Println("error when try to GetAuthByEmail with deail", err.Error())
		return
	}

	if err = authEntity.VerifyPlainPasswordToHash(model.Password); err != nil {
		err = response.ErrPasswordNotMatch
		return
	}

	token, err = model.GenerateToken(config.Cfg.App.Encryption.JWTSecret)
	if err != nil {
		log.Println("error when generate token", err.Error())
		return
	}

	return
}
