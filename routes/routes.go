package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	AuthRoutes(r)
	ProfileRoutes(r)
	Transaction(r)
	CartRoutes(r)
	ProductRoutes(r)
}
