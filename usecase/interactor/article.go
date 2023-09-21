package interactor

import (
	"context"

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

func (a *Article) GetArticles(ctx context.Context) {
	articles, err := a.repository.GetArticles(ctx)
	if err != nil {
		a.outputPort.RenderError(err)
		return
	}
	a.outputPort.Render(articles)
}
