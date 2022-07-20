// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/nguyenthanhworkspace/golang-starter/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Translation -.
	Translation interface {
		Translate(context.Context, entity.Translation) (entity.Translation, error)
		History(context.Context) ([]entity.Translation, error)
	}

	// TranslationRepo -.
	TranslationRepo interface {
		Store(context.Context, entity.Translation) error
		GetHistory(context.Context) ([]entity.Translation, error)
	}

	// TranslationWebAPI -.
	TranslationWebAPI interface {
		Translate(entity.Translation) (entity.Translation, error)
	}
)

type (
	// UserRepo -.
	UserRepo interface {
		Store(context.Context, entity.User) error
		//GetUsers(context.Context) ([]entity.User, error)
	}

	User interface {
		usersIndex(context.Context) ([]entity.User, error)
	}
)
