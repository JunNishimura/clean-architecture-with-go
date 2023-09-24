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

func (r *articleRepository) Update(ctx context.Context, articleID int64, title, body *string) error {
	article, err := entities.FindArticle(ctx, r.db, articleID)
	if err != nil {
		return err
	}
	if title != nil {
		article.Title = *title
	}
	if body != nil {
		article.Body = *body
	}
	if _, err := article.Update(ctx, r.db, boil.Infer()); err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) Delete(ctx context.Context, articleID int64) error {
	article, err := entities.FindArticle(ctx, r.db, articleID)
	if err != nil {
		return err
	}
	if _, err := article.Delete(ctx, r.db); err != nil {
		return err
	}
	return nil
}
