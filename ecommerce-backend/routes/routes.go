package routes

import (
	"backend/controllers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", controllers.LoginUser).Methods("POST")

	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", controllers.GetProductByID).Methods("GET")
	r.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")
	return r
}
