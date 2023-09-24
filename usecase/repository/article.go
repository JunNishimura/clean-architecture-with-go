package repository

import (
	"context"

	"github.com/JunNishimura/clean-architecture-with-go/entities"
)

type Article interface {
	FindAll(context.Context) ([]*entities.Article, error)
	FindByID(context.Context, int64) (*entities.Article, error)
	Create(context.Context, *entities.Article) error
	Delete(context.Context, int64) error
}
