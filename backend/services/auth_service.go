package services

import (
	"os"
	"time"

	"gin-fleamarket/models"
	"gin-fleamarket/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SignUp(email string, password string) error
	Login(email string, password string) (*string, error)
}

type AuthService struct {
	repository repositories.IAuthRepository
}

func NewAuthService(repository repositories.IAuthRepository) IAuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) SignUp(email string, password string) error {
	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	return s.repository.CreateUser(user)
}

func (s *AuthService) Login(email string, password string) (*string, error) {
	foundUser, err := s.repository.FindUser(email)
	if err != nil {
		return nil, err
	}

	// パスワードの比較
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	// トークンの作成
	token, err := CreateToken(foundUser.ID, foundUser.Email)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func CreateToken(userId uint, email string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// subはsubjectの略で、JWTの主題を表す
		// emailはユーザーのメールアドレス
		// expはJWTの有効期限を表す
		"sub":   userId,
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	// 署名
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
