package main 

import (
	"log" 
	"net/http" 
	"github.com/gorilla/mux" 
	_"gorm.io/driver/mysql" 
	
	"github.com/abimanyu/apimovies/pkg/routes"
)

func main(){
	r := mux.NewRouter() 
	routes.RegisterMovieRoutes(r) 
	http.Handle("/", r) 
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}