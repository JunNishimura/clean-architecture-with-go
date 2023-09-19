package repository

import (
	"context"

	"github.com/JunNishimura/clean-architecture-with-go/entities"
)

type Article interface {
	GetArticles(context.Context) ([]*entities.Article, error)
}
