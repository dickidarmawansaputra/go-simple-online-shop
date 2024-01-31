package auth

import (
	"strings"
	"time"

	"github.com/dickidarmawansaputra/go-simple-online-shop/infra/response"
	"github.com/dickidarmawansaputra/go-simple-online-shop/utility"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	ROLE_ADMIN = "admin"
	ROLE_USER  = "user"
)

type AuthEntity struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	PublicId  uuid.UUID `json:"public_id"`
	Password  string    `json:"password"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewFromRegisterRequest(req RegisterRequestPayload) AuthEntity {
	return AuthEntity{
		Email:     req.Email,
		Password:  req.Password,
		Role:      ROLE_USER,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		PublicId:  uuid.New(),
	}
}

func NewFromLoginRequest(req LoginRequestPayload) AuthEntity {
	return AuthEntity{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (a AuthEntity) Validate() (err error) {
	if err = a.ValidateEmail(); err != nil {
		return
	}

	if err = a.ValidatePassword(); err != nil {
		return
	}

	return
}

func (a AuthEntity) ValidateEmail() (err error) {
	if a.Email == "" {
		return response.ErrEmailRequired
	}

	emails := strings.Split(a.Email, "@")
	if len(emails) != 2 {
		return response.ErrEmailInvalid
	}

	return
}

func (a AuthEntity) ValidatePassword() (err error) {
	if a.Password == "" {
		return response.ErrPasswordRequired
	}

	if len(a.Password) < 6 {
		return response.ErrPasswordInvalid
	}

	return
}

func (a AuthEntity) IsExists() bool {
	return a.Id != 0
}

func (a *AuthEntity) EncryptPassword(salt int) (err error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	a.Password = string(encryptedPassword)
	return nil
}

func (a AuthEntity) VerifyPlainPasswordToHash(hashedPassword string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(a.Password))
}

func (a AuthEntity) GenerateToken(secret string) (tokenString string, err error) {
	return utility.GenerateToken(a.PublicId.String(), string(a.Role), secret)
}
