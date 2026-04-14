package main

import (
	"fmt"

	"khadbhandar/config"
	"khadbhandar/models"
	"khadbhandar/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("SERVER STARTED")

	config.Connect()

	config.DB.AutoMigrate(&models.Product{}, &models.Order{})

	// gin init
	r := gin.Default()

	// SESSION SETUP
	store := cookie.NewStore([]byte("secret"))

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	})

	r.Use(sessions.Sessions("mysession", store))

	// static files
	r.Static("/static", "./static")

	// templates
	r.LoadHTMLGlob("./templates/*")

	// routes
	fmt.Println("ROUTES LOADING...")
	routes.SetupRoutes(r)

	// run server
	r.Run(":8080")
}
