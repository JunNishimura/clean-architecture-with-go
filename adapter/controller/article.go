package controller

import (
	"database/sql"
	"net/http"

	"github.com/JunNishimura/clean-architecture-with-go/usecase/interactor"
	"github.com/JunNishimura/clean-architecture-with-go/usecase/port"
	"github.com/JunNishimura/clean-architecture-with-go/usecase/repository"
)

type article struct {
	outputPortFactory func(http.ResponseWriter) port.ArticleOutput
	inputPortFactory  func(port.ArticleInput, repository.Article) port.ArticleInput
	repositoryFactory func(*sql.DB) repository.Article
	db                *sql.DB
}

func NewArticle(db *sql.DB) *article {
	return &article{
		db: db,
	}
}

func (a *article) GetArticles(w http.ResponseWriter, r *http.Request) {
	outputPort := a.outputPortFactory(w)
	repository := a.repositoryFactory(a.db)
	inputPort := interactor.NewArticle(outputPort, repository)
	inputPort.GetArticles(r.Context())
}
