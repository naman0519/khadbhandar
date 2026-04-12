package controllers

import (
	"fmt"
	"khadbhandar/config"
	"khadbhandar/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var StockMap = make(map[string]int)

// ================= LOGIN PAGE =================

func ShowLogin(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil)
}

// ================= LOGIN HANDLE =================

func HandleLogin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "admin" && password == "Naman@123" {
		c.Redirect(302, "/admin/dashboard")
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"error": "Invalid credentials ",
		})
	}
}

// ================= DASHBOARD =================

func GetDashboard(c *gin.Context) {

	var orders []models.Order
	config.DB.Find(&orders)

	var products []models.Product
	config.DB.Find(&products)

	totalOrders := len(orders)
	totalProducts := len(products)

	lowStock := 0
	for _, p := range products {
		if p.Stock < 50 {
			lowStock++
		}
	}

	c.HTML(200, "admin.html", gin.H{
		"orders":        orders,
		"products":      products,
		"totalOrders":   totalOrders,
		"totalProducts": totalProducts,
		"lowStock":      lowStock,
	})
}

//==============Update stock=========//

func UpdateStockAdmin(c *gin.Context) {

	product := c.PostForm("product")
	stockStr := c.PostForm("stock")

	stock, _ := strconv.Atoi(stockStr)

	fmt.Println("UPDATE HIT:", product, stock)

	// FIXED DB UPDATE
	result := config.DB.Model(&models.Product{}).
		Where("name ILIKE ?", product).
		Update("stock", stock)

	fmt.Println("Rows affected:", result.RowsAffected)

	c.JSON(200, gin.H{
		"message": "Stock updated",
	})
}

// =============DELETE PRODUCT==============//

func DeleteProduct(c *gin.Context) {

	product := c.PostForm("product")

	result := config.DB.Where("name = ?", product).Delete(&models.Product{})

	fmt.Println("DELETE HIT:", product)
	fmt.Println("Rows deleted:", result.RowsAffected)

	c.Redirect(302, "/admin/dashboard")
}

func ApproveOrder(c *gin.Context) {

	id := c.Param("id")

	config.DB.Model(&models.Order{}).
		Where("id = ?", id).
		Update("status", "Approved")

	c.Redirect(302, "/admin/dashboard")
}
