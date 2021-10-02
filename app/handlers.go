package app

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	"github.com/josemendoza/restapi/model"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var users []model.User
	// buscamos en la DB a todos los usuarios.
	if err := db.Find(&users).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)	
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&users)
	}
	
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []model.Post
	if err := db.Find(&posts).Error; err !=nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&posts)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var user model.User
	params := mux.Vars(r)

	if err := db.First(&user, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&user)
	}
}
func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post model.Post
	params := mux.Vars(r)
	if err := db.First(&post, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&post)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	loc, _ := time.LoadLocation("UTC")
	newTime := time.Now().In(loc)
	user.CreatedAt = newTime
	user.UpdatedAt = newTime

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil{
		panic(err)
	}
	user.Password = string(hashedPassword)

	if err := db.Create(&user).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}
func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post model.Post
	json.NewDecoder(r.Body).Decode(&post)

	loc, _ := time.LoadLocation("UTC")
	newTime := time.Now().In(loc)
	post.CreatedAt = newTime
	post.UpdatedAt = newTime

	if err := db.Create(&post).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	params := mux.Vars(r)
	
	response := db.First(&user, params["id"])
	if response.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&user)
	updated := db.Save(&user)
	if updated.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)

}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post model.Post
	params := mux.Vars(r)

	db.First(&post, params["id"])
	json.NewDecoder(r.Body).Decode(&post)
	db.Save(&post)
	json.NewEncoder(w).Encode(post)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {}
func DeletePost(w http.ResponseWriter, r *http.Request) {}
