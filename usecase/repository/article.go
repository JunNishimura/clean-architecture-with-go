package repository

import (
	"context"

	"github.com/JunNishimura/clean-architecture-with-go/entities"
)

type Article interface {
	FindAll(context.Context) ([]*entities.Article, error)
	FindByID(context.Context, int64) (*entities.Article, error)
	Create(context.Context, *entities.Article) error
	Update(ctx context.Context, articleID int64, title, body *string) error
	Delete(context.Context, int64) error
}
