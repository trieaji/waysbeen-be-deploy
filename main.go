package main

import (
	"fmt"
	"golangfnl/database"
	"golangfnl/pkg/mysql"
	"golangfnl/routes"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	// import this package ...
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	// initial DB here ...
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	// On http (API)
	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	//path file
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads")))) // add this code

	// Setup allowed Header, Method, and Origin for CORS here ...
	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = os.Getenv("PORT")
	fmt.Println("server running localhost:" + port)

	// Embed the setup allowed in 2 parameter on this below code ...
	http.ListenAndServe("localhost:"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))

	// fmt.Println("server running localhost:5000")
	// http.ListenAndServe("localhost:5000", r)
}
