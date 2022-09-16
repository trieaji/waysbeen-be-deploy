package routes

import (
	"golangfnl/handler"
	"golangfnl/pkg/middleware"
	"golangfnl/pkg/mysql"
	"golangfnl/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handler.HandlerCart(cartRepository)

	r.HandleFunc("/carts", h.FindCart).Methods("GET")
	r.HandleFunc("/cart/{id}", h.GetCart).Methods("GET")
	r.HandleFunc("/cart", middleware.Auth(h.CreateCart)).Methods("POST")
	r.HandleFunc("/cart/{id}", h.DeleteCart).Methods("DELETE")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.UpdateCart)).Methods("PATCH")
}
