package controllers 


import (
	"encoding/json" 
	"fmt" 
	"github.com/gorilla/mux" 
	"net/http" 
	"strconv" 
	"github.com/abimanyu/apimovies/pkg/utils" 
	"github.com/abimanyu/apimovies/pkg/models"
)

var NewMovie models.Movie 

func GetMovie(w http.ResponseWriter, r *http.Request){
	newMovies := models.GetAllMovies() 
	res, _ := json.Marshal(newMovies) 
	w.Header().Set("Content-Type", "application/json") 
	w.WriteHeader(http.StatusOK) 
	w.Write(res) 
}

func GetMovieById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r) 
	movieId := vars["movieId"] 
	ID, err := strconv.ParseInt(movieId, 0, 0) 
	if err != nil{
		fmt.Println("error while parsing")
	}

	movieDetails, _ := models.GetMovieById(ID) 
	res, _ := json.Marshal(movieDetails) 
	w.Header().Set("Content-Type", "application/json") 
	w.WriteHeader(http.StatusOK) 
	w.Write(res)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	movie := &models.Movie{} 
	utils.ParseBody(r, movie) 
	payload := movie.CreateMovie() 
	res, _ := json.Marshal(payload) 
	w.WriteHeader(http.StatusOK) 
	w.Write(res)

}

func DeleteMovie(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r) 
	movieId := vars["movieId"] 
	ID, err := strconv.ParseInt(movieId, 0, 0) 
	if err != nil{
		fmt.Println("error while parsing parameter")
	}
	movie := models.DeleteMovie(ID) 
	res, _ := json.Marshal(movie) 
	w.Header().Set("Content-Type", "application/json") 
	w.WriteHeader(http.StatusOK) 
	w.Write(res)

}

func UpdateMovie(w http.ResponseWriter, r *http.Request){

	var movieToUpdate = &models.Movie{} 
	utils.ParseBody(r, movieToUpdate) 
	vars := mux.Vars(r) 
	movieId := vars["movieId"] 
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil{
		fmt.Println("error while parsing request")
	}
	movieDetails, db := models.GetMovieById(ID) 
	if movieToUpdate.Title != ""{
		movieDetails.Title = movieToUpdate.Title
	}
	if movieToUpdate.Director != ""{
		movieDetails.Director = movieToUpdate.Director
	}
	if movieToUpdate.Year != ""{
		movieDetails.Year = movieToUpdate.Year 
	}

	db.Save(&movieDetails) 
	res, _ := json.Marshal(movieDetails) 
	w.Header().Set("Content-Type", "application/json") 
	w.WriteHeader(http.StatusOK) 
	w.Write(res)

}