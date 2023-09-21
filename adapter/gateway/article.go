package gateway

import (
	"context"
	"database/sql"

	"github.com/JunNishimura/clean-architecture-with-go/entities"
	"github.com/JunNishimura/clean-architecture-with-go/usecase/repository"
)

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) repository.Article {
	return &ArticleRepository{
		db: db,
	}
}

func (r *ArticleRepository) GetArticles(ctx context.Context) ([]*entities.Article, error) {
	return entities.Articles().All(ctx, r.db)
}
