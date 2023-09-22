package port

import (
	"context"
)

type ArticleInput interface {
	GetArticles(context.Context)
}

type ArticleOutput interface {
	Render(ctx context.Context, body any, status int)
}
