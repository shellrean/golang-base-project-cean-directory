package service

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/shellrean/golang-base-project-cean-directory/domain"
	"github.com/shellrean/golang-base-project-cean-directory/dto"
	"github.com/shellrean/golang-base-project-cean-directory/internal/config"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

type authService struct {
	cnf            *config.Config
	userRepository domain.UserRepository
}

func NewAuth(cnf *config.Config,
	userRepository domain.UserRepository) domain.AuthService {
	return &authService{
		cnf:            cnf,
		userRepository: userRepository,
	}
}

func (a authService) Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error) {
	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthRes{}, err
	}
	if user.Id == "" {
		return dto.AuthRes{}, domain.ErrInvalidCredential
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return dto.AuthRes{}, domain.ErrInvalidCredential
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    user.Id,
			"email": user.Email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString([]byte(a.cnf.Secret.Jwt))
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthRes{}, err
	}

	return dto.AuthRes{
		AccessToken: tokenString,
	}, nil
}
