package port

import (
	"context"

	"github.com/JunNishimura/clean-architecture-with-go/entities"
)

type ArticleInput interface {
	FindAll(context.Context)
	Create(context.Context, *entities.Article)
}

type ArticleOutput interface {
	Render(ctx context.Context, body any, status int)
}
