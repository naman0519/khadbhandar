package main

import (
	"fmt"

	"khadbhandar/config"
	database "khadbhandar/config"
	"khadbhandar/models"
	"khadbhandar/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println(" SERVER STARTED")

	// DB connect
	database.Connect()

	// migrate tables
	config.DB.AutoMigrate(&models.Product{}, &models.Order{})

	// gin init
	r := gin.Default()

	//  SESSION SETUP (IMPORTANT)
	store := cookie.NewStore([]byte("secret"))

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600, // 1 hour
		HttpOnly: true,
	})

	r.Use(sessions.Sessions("mysession", store))

	// static files
	r.Static("/static", "./static")

	// templates load
	r.LoadHTMLGlob("./templates/*")

	// routes load
	fmt.Println(" ROUTES LOADING...")
	routes.SetupRoutes(r)

	// run server
	r.Run(":8080")
}
