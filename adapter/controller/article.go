package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JunNishimura/clean-architecture-with-go/adapter/gateway"
	"github.com/JunNishimura/clean-architecture-with-go/adapter/presenter"
	"github.com/JunNishimura/clean-architecture-with-go/entities"
	"github.com/JunNishimura/clean-architecture-with-go/usecase/interactor"
	"github.com/go-chi/chi/v5"
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

func (a *article) FindByID(w http.ResponseWriter, r *http.Request) {
	outputPort := presenter.NewArticle(w)
	repository := gateway.NewArticleRepository(a.db)
	inputPort := interactor.NewArticle(outputPort, repository)

	strArticleID := chi.URLParam(r, "articleID")
	articleID, err := strconv.ParseInt(strArticleID, 10, 64)
	if err != nil {
		outputPort.Render(r.Context(), &presenter.ErrResponse{
			Message: fmt.Sprintf("could not find article by '%s'", strArticleID),
		}, http.StatusNotFound)
	}
	inputPort.FindByID(r.Context(), articleID)
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
