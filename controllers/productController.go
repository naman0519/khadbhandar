package controllers

import (
	"fmt"
	"khadbhandar/config"
	"khadbhandar/models"

	"github.com/gin-gonic/gin"
)

// ================== GET ALL PRODUCTS (HOME PAGE) ==================

func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)

	fmt.Println("Products:", products)

	//  stock map create
	stockMap := make(map[string]int)

	for _, p := range products {
		stockMap[p.Name] = p.Stock
	}

	c.HTML(200, "index.html", gin.H{
		"products": products,
		"stockMap": stockMap,
	})
}

// ================== ADMIN PRODUCTS PAGE ==================

func AdminProducts(c *gin.Context) {

	var products []models.Product

	config.DB.Find(&products)

	c.HTML(200, "admin_products.html", gin.H{
		"products": products,
	})
}

func AddProduct(c *gin.Context) {

	name := c.PostForm("name")
	category := c.PostForm("category")
	stockStr := c.PostForm("stock")
	price := c.PostForm("price")
	image := c.PostForm("image")

	var stock int
	fmt.Sscan(stockStr, &stock)

	product := models.Product{
		Name:     name,
		Category: category,
		Stock:    stock,
		Price:    price,
		Image:    image,
	}

	result := config.DB.Create(&product)

	if result.Error != nil {
		fmt.Println(" DB ERROR:", result.Error)
		c.String(500, "DB error")
		return
	}

	fmt.Println("PRODUCT SAVED:", product.Name)

	c.Redirect(302, "/admin/dashboard")
}
