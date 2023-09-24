package interactor

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/JunNishimura/clean-architecture-with-go/entities"
	"github.com/JunNishimura/clean-architecture-with-go/usecase/port"
	"github.com/JunNishimura/clean-architecture-with-go/usecase/repository"
)

type Article struct {
	outputPort port.ArticleOutput
	repository repository.Article
}

func NewArticle(outputPort port.ArticleOutput, repository repository.Article) port.ArticleInput {
	return &Article{
		outputPort: outputPort,
		repository: repository,
	}
}

type ErrResponse struct {
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

func (a *Article) FindAll(ctx context.Context) {
	articles, err := a.repository.FindAll(ctx)
	if err != nil {
		a.outputPort.Render(ctx, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	a.outputPort.Render(ctx, articles, http.StatusOK)
}

func (a *Article) FindByID(ctx context.Context, articleID int64) {
	article, err := a.repository.FindByID(ctx, articleID)
	if err != nil {
		var status int
		if errors.Is(err, sql.ErrNoRows) {
			status = http.StatusNotFound
		} else {
			status = http.StatusInternalServerError
		}
		a.outputPort.Render(ctx, &ErrResponse{
			Message: fmt.Sprintf("could not find article by '%d'", articleID),
		}, status)
		return
	}
	a.outputPort.Render(ctx, article, http.StatusOK)
}

func (a *Article) Create(ctx context.Context, newArticle *entities.Article) {
	if err := a.repository.Create(ctx, newArticle); err != nil {
		a.outputPort.Render(ctx, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct {
		ID int64 `json:"id"`
	}{ID: newArticle.ID}
	a.outputPort.Render(ctx, rsp, http.StatusOK)
}

func (a *Article) Update(ctx context.Context, articleID int64, title, body *string) {
	if err := a.repository.Update(ctx, articleID, title, body); err != nil {
		a.outputPort.Render(ctx, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	a.outputPort.Render(ctx, struct{}{}, http.StatusOK)
}

func (a *Article) Delete(ctx context.Context, articleID int64) {
	if err := a.repository.Delete(ctx, articleID); err != nil {
		a.outputPort.Render(ctx, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	a.outputPort.Render(ctx, struct{}{}, http.StatusOK)
}
