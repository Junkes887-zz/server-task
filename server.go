package main

import (
	"encoding/json"
	"net/http"
)

// User ...
type User struct {
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/", helloServer)
	http.HandleFunc("/users", listemUsers)
	http.ListenAndServe(":8080", nil)
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal("Server run port 8080...")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func listemUsers(w http.ResponseWriter, r *http.Request) {
	var user2 = User{"Henrique", 18}
	var user1 = User{"Eder", 43}
	listUser := [2]User{user1, user2}
	js, err := json.Marshal(listUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
