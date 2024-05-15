package domain

import (
	"context"
	"github.com/shellrean/golang-base-project-cean-directory/dto"
)

type AuthService interface {
	Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error)
}
