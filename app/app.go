package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var router *mux.Router
var db *gorm.DB

func Initilize() {
	dsn := os.Getenv("DB_CONNECTION")

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil{
		panic("Failed to connect to DB")
	}

	fmt.Println("DB connected")

	router = mux.NewRouter().StrictSlash(true)
	setRoutes()

	log.Fatal(http.ListenAndServe(":3000", router))
}

func setRoutes() {
	router.HandleFunc("/getAllUsers", GetAllUsers).Methods("GET")
	router.HandleFunc("/getAllPosts", GetAllPosts).Methods("GET")

	router.HandleFunc("/getUser/{id}", GetUser).Methods("GET")
	router.HandleFunc("/getPost/{id}", GetPost).Methods("GET")
	
	router.HandleFunc("/createUser", CreateUser).Methods("POST")
	router.HandleFunc("/createPost", CreatePost).Methods("POST")
	
	router.HandleFunc("/updateUser/{id}", UpdateUser).Methods("PATCH")
	router.HandleFunc("/updatePost/{id}", UpdatePost).Methods("PATCH")
	
	router.HandleFunc("/deleteUser/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/deletePost/{id}", DeletePost).Methods("DELETE")

}
