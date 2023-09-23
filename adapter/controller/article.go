package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/JunNishimura/clean-architecture-with-go/adapter/gateway"
	"github.com/JunNishimura/clean-architecture-with-go/adapter/presenter"
	"github.com/JunNishimura/clean-architecture-with-go/entities"
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

func (a *article) FindAll(w http.ResponseWriter, r *http.Request) {
	outputPort := presenter.NewArticle(w)
	repository := gateway.NewArticleRepository(a.db)
	inputPort := interactor.NewArticle(outputPort, repository)
	inputPort.FindAll(r.Context())
}

func (a *article) Create(w http.ResponseWriter, r *http.Request) {
	outputPort := presenter.NewArticle(w)
	repository := gateway.NewArticleRepository(a.db)
	inputPort := interactor.NewArticle(outputPort, repository)

	var b struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		outputPort.Render(r.Context(), &presenter.ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
	}

	newArticle := &entities.Article{
		Title: b.Title,
		Body:  b.Body,
	}
	inputPort.Create(r.Context(), newArticle)
}
