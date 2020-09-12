package controllers

import (
	"net/http"
	"product-order-be/api/middlewares"
)

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Static Files Upload
	s.Router.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("./files"))))

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	// s.Router.HandleFunc("/register", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/email", middlewares.SetMiddlewareJSON(s.GetUserByEmail)).Methods("POST")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/orders", middlewares.SetMiddlewareJSON(s.CreateOrder)).Methods("POST")
	s.Router.HandleFunc("/upload/orders", middlewares.SetMiddlewareJSON(s.UploadOrder)).Methods("POST")
	s.Router.HandleFunc("/orders", middlewares.SetMiddlewareJSON(s.GetOrders)).Methods("GET")
	s.Router.HandleFunc("/orders/{id}", middlewares.SetMiddlewareJSON(s.GetOrder)).Methods("GET")
	s.Router.HandleFunc("/orders/user/{id}", middlewares.SetMiddlewareJSON(s.GetOrderByUserId)).Methods("GET")
	s.Router.HandleFunc("/orders/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateOrder))).Methods("PUT")
	s.Router.HandleFunc("/orders/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteOrder)).Methods("DELETE")

	//Puskesmas routes
	s.Router.HandleFunc("/puskesmas", middlewares.SetMiddlewareJSON(s.CreatePuskesmas)).Methods("POST")
	s.Router.HandleFunc("/puskesmas", middlewares.SetMiddlewareJSON(s.GetAllPuskesmas)).Methods("GET")
	s.Router.HandleFunc("/puskesmas/{id}", middlewares.SetMiddlewareJSON(s.GetPuskesmas)).Methods("GET")
	s.Router.HandleFunc("/puskesmas/user/{id}", middlewares.SetMiddlewareJSON(s.GetPuskesmasByUserId)).Methods("GET")
	s.Router.HandleFunc("/puskesmas/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePuskesmas))).Methods("PUT")
	s.Router.HandleFunc("/puskesmas/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePuskesmas)).Methods("DELETE")
}
