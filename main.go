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

	// Serve static files for each route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("static/images"))))
	r.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("static/fonts"))))
	r.PathPrefix("/videos/").Handler(http.StripPrefix("/videos/", http.FileServer(http.Dir("static/videos"))))

	// Handling the routes
	r.HandleFunc("/", RootRoute).Methods("GET")
	r.HandleFunc("/reservation", ReservationRoute).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func RootRoute(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func ReservationRoute(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/reservation.html")
}
