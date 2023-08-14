package routes 

import (
	"github.com/gorilla/mux" 
	"github.com/abimanyu/apimovies/pkg/controllers"
)

var RegisterMovieRoutes = func(router *mux.Router){
	router.HandleFunc("/movies", controllers.CreateMovie).Methods("POST") 
	router.HandleFunc("/movies", controllers.GetMovie).Methods("GET") 
	router.HandleFunc("/movies/{movieId}", controllers.GetMovieById).Methods("GET") 
	router.HandleFunc("/movies/{movieId}", controllers.UpdateMovie).Methods("PUT") 
	router.HandleFunc("/movies/{movieId}", controllers.DeleteMovie).Methods("DELETE")
}