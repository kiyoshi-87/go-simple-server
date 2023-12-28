package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("Server running on port 8000!")

	//Handling the routes
	r.HandleFunc("/", RootRoute).Methods("GET")
	r.HandleFunc("/reservation", ReservationRoute).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Server running on port 8000!")
}

func RootRoute(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func ReservationRoute(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/reservation.html")
}
