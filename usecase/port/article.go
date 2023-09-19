package port

import (
	"context"

	"github.com/JunNishimura/clean-architecture-with-go/entities"
)

type ArticleInput interface {
	GetArticles(context.Context)
}

type ArticleOutput interface {
	Render([]*entities.Article)
	RenderError(error)
}
