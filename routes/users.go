package routes

import (
	"golangfnl/handler"
	"golangfnl/pkg/middleware"
	"golangfnl/pkg/mysql"
	"golangfnl/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handler.HandlerUser(userRepository)
	r.HandleFunc("/users", middleware.Auth(h.FindUsers)).Methods("GET")
	r.HandleFunc("/user/{id}", middleware.Auth(h.GetUser)).Methods("GET")
	r.HandleFunc("/user", middleware.Auth(h.CreateUser)).Methods("POST")
	r.HandleFunc("/user/{id}", middleware.Auth(h.UpdateUser)).Methods("PATCH")
	r.HandleFunc("/user/{id}", middleware.Auth(h.DeleteteUser)).Methods("DELETE")
}
