package port

import (
	"context"

	"github.com/JunNishimura/clean-architecture-with-go/entities"
)

type ArticleInput interface {
	FindAll(context.Context)
	FindByID(context.Context, int64)
	Create(context.Context, *entities.Article)
	Update(ctx context.Context, articleID int64, title, body *string)
	Delete(context.Context, int64)
}

type ArticleOutput interface {
	Render(ctx context.Context, body any, status int)
}
