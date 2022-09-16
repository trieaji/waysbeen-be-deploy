package routes

import (
	"golangfnl/handler"
	"golangfnl/pkg/middleware"
	"golangfnl/pkg/mysql"
	"golangfnl/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handler.HandlerProduct(productRepository)

	r.HandleFunc("/products", middleware.Auth(h.FindProducts)).Methods("GET")
	r.HandleFunc("/product/{id}", middleware.Auth(h.GetProduct)).Methods("GET")
	r.HandleFunc("/product", middleware.Auth(middleware.UploadFile(h.CreateProduct))).Methods("POST")  //memerlukan UploadFile untuk bisa mengupdate gambar
	r.HandleFunc("/product/{id}", middleware.Auth(middleware.UploadFile(h.UpdateProduct))).Methods("PATCH") //memerlukan UploadFile untuk bisa memasukkan gambar
	r.HandleFunc("/product/{id}", middleware.Auth(h.DeleteProduct)).Methods("DELETE")
}
