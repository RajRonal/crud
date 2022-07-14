package server

import (
	"awesomeProject1/handlers"
	"awesomeProject1/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	chi.Router
}

func SetupRoutes() *Server {
	router := chi.NewRouter()
	router.Route("/api", func(api chi.Router) {
		api.Post("/login", handlers.Login)

		api.Route("/ronal", func(r chi.Router) {
			r.Use(middleware.Middleware)
			r.Post("/addUser", handlers.AddRow)
			r.Put("/updateUser", handlers.UpdateRow)
			r.Delete("/deleteUser", handlers.DeleteRow)
		})
	})

	return &Server{router}
}

func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}
