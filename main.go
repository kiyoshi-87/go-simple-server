package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	fmt.Println("Server running on port 8000!")

	// Serve static files
	r.Static("/static", "./static")

	// Handling the routes
	r.GET("/", RootRoute)
	r.GET("/reservation", ReservationRoute)

	log.Fatal(r.Run(":8000"))
}

func RootRoute(c *gin.Context) {
	c.File("./static/index.html")
}

func ReservationRoute(c *gin.Context) {
	c.File("./static/reservation.html")
}
