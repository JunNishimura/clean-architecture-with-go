package port

import (
	"context"
)

type ArticleInput interface {
	FindAll(context.Context)
}

type ArticleOutput interface {
	Render(ctx context.Context, body any, status int)
}
