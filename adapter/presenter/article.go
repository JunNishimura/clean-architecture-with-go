package presenter

import (
	"net/http"

	"github.com/JunNishimura/clean-architecture-with-go/entities"
	"github.com/JunNishimura/clean-architecture-with-go/usecase/port"
)

type article struct {
	w http.ResponseWriter
}

func NewArticle(w http.ResponseWriter) port.ArticleOutput {
	return &article{
		w: w,
	}
}

func (a *article) Render([]*entities.Article) {
}

func (a *article) RenderError(error) {
}
