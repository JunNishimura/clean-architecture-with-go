package gateway

import (
	"context"
	"database/sql"

	"github.com/JunNishimura/clean-architecture-with-go/entities"
	"github.com/JunNishimura/clean-architecture-with-go/usecase/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type articleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) repository.Article {
	return &articleRepository{
		db: db,
	}
}

func (r *articleRepository) FindAll(ctx context.Context) ([]*entities.Article, error) {
	return entities.Articles().All(ctx, r.db)
}

func (r *articleRepository) FindByID(ctx context.Context, articleID int64) (*entities.Article, error) {
	return entities.FindArticle(ctx, r.db, articleID)
}

func (r *articleRepository) Create(ctx context.Context, newArticle *entities.Article) error {
	return newArticle.Insert(ctx, r.db, boil.Infer())
}
