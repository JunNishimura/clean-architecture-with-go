package controller

import (
	"database/sql"
	"net/http"

	"github.com/JunNishimura/clean-architecture-with-go/adapter/gateway"
	"github.com/JunNishimura/clean-architecture-with-go/adapter/presenter"
	"github.com/JunNishimura/clean-architecture-with-go/usecase/interactor"
)

type article struct {
	db *sql.DB
}

func NewArticle(db *sql.DB) *article {
	return &article{
		db: db,
	}
}

func (a *article) GetArticles(w http.ResponseWriter, r *http.Request) {
	outputPort := presenter.NewArticle(w)
	repository := gateway.NewArticleRepository(a.db)
	inputPort := interactor.NewArticle(outputPort, repository)
	inputPort.GetArticles(r.Context())
}
