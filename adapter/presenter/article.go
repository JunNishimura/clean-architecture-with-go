package presenter

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

type ErrResponse struct {
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

func (a *article) Render(ctx context.Context, body any, status int) {
	a.w.Header().Set("Content-Type", "applications/json; charset=utf-8")
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		a.w.WriteHeader(http.StatusInternalServerError)
		rsp := ErrResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		}
		if err := json.NewEncoder(a.w).Encode(rsp); err != nil {
			log.Printf("fail to write response error: %v", err)
		}
		return
	}
	a.w.WriteHeader(status)
	if _, err := fmt.Fprintf(a.w, "%s", bodyBytes); err != nil {
		log.Printf("fail to write response: %v", err)
	}
}
