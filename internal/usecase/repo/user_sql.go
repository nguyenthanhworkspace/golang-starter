package repo

import (
	"context"
	"github.com/nguyenthanhworkspace/golang-starter/internal/entity"
	"github.com/nguyenthanhworkspace/golang-starter/pkg/mysql"
)

// UserRepo =.
type UserRepo struct {
	*mysql.Mysql
}

func (u UserRepo) Store(ctx context.Context, user entity.User) error {
	//TODO implement me
	panic("implement me")
}

// NewUserRepo =.
func NewUserRepo(db *mysql.Mysql) *UserRepo {
	return &UserRepo{db}
}
