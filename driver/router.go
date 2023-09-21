package driver

import (
	"fmt"
	"net/http"

	"github.com/JunNishimura/clean-architecture-with-go/adapter/controller"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db, err := connectDB()
	if err != nil {
		return err
	}

	articleController := controller.NewArticle(db)
	r.Get("/articles", articleController.GetArticles)
	if err := http.ListenAndServe(":8080", r); err != nil {
		return fmt.Errorf("fail to listen and serve: %v", err)
	}
	return nil
}
