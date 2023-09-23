package repository

import (
	"context"

	"github.com/JunNishimura/clean-architecture-with-go/entities"
)

type Article interface {
	FindAll(context.Context) ([]*entities.Article, error)
}
