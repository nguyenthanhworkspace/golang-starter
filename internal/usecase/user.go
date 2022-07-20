package usecase

import (
	"context"
	"github.com/nguyenthanhworkspace/golang-starter/internal/entity"
)

type UserUseCase struct {
	repo UserRepo
}

// NewUserUseCase -.
func NewUserUseCase(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (u UserUseCase) usersIndex(ctx context.Context) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}
