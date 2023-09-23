package interactor

import (
	"context"
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
	Details []string `json:"details,omitempty`
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
