package router

import (
	"LinkShortener/internal/jwtToken"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)

	GetLink(w http.ResponseWriter, r *http.Request)

	AllLinks(w http.ResponseWriter, r *http.Request)
	CreateLink(w http.ResponseWriter, r *http.Request)
	DeleteLink(w http.ResponseWriter, r *http.Request)
}

func NewRouter(router Router) http.Handler {
	r := chi.NewRouter()

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", router.Register)
		r.Post("/login", router.Login)
	})

	//route used for short link
	r.Get("/*", router.GetLink)

	r.Route("/user", func(r chi.Router) {
		r.Use(jwtToken.CheckLogin)
		
		r.Get("/allLinks", router.AllLinks)
		r.Post("/createLink", router.CreateLink)
		r.Post("/deleteLink", router.DeleteLink)
	})

	return r
}
