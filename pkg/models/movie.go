package models

import (
	"github.com/abimanyu/apimovies/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB 

type Movie struct{
	gorm.Model 
	Title string `json:"title"`
	Director string `json:"director"`
	Year string `json:"year"`

}

func init(){
	config.Connect() 
	db = config.GetDB() 
	db.AutoMigrate(&Movie{})
}

func (m *Movie) CreateMovie() *Movie{
	db.Create(&m)
	return m 
}

func GetAllMovies() []Movie{
	var Movies []Movie 
	db.Find(&Movies) 
	return Movies 
}

func GetMovieById(id int64) (*Movie, *gorm.DB){
	var movie Movie 
	db := db.Where("ID=?", id).Find(&movie) 
	return &movie, db 
}

func DeleteMovie(ID int64) Movie{
	var movie Movie 
	db.Where("ID=?", ID).Delete(movie) 
	return movie 
}