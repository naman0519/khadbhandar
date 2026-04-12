package controllers

import (
	"fmt"
	"strings"

	"khadbhandar/config"
	"khadbhandar/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ================== PLACE ORDER ==================

func PlaceOrder(c *gin.Context) {

	name := c.PostForm("name")
	product := strings.ToLower(c.PostForm("product"))
	phone := c.PostForm("phone")

	var quantity int
	fmt.Sscan(c.PostForm("quantity"), &quantity)

	if quantity > 30 {
		c.String(400, "Max 30 bags allowed")
		return
	}

	fmt.Println("DATA:", name, product, phone, quantity)

	// 🔥 STEP 1: UPDATE STOCK (MAIN LOGIC)
	err := UpdateStock(product, quantity)
	if err != nil {
		c.String(400, err.Error()) // STOP HERE
		return
	}

	// 🔥 STEP 2: SAVE ORDER (ONLY IF STOCK OK)
	order := models.Order{
		Name:     name,
		Product:  product,
		Phone:    phone,
		Quantity: quantity,
		Status:   "Pending",
	}

	err = config.DB.Create(&order).Error
	if err != nil {
		fmt.Println("DB ERROR:", err)
		c.String(500, "Database error")
		return
	}

	fmt.Println("ORDER SAVED")

	// SUCCESS
	c.Redirect(302, "/success?product="+product)
}

// ================== CHECK STOCK ==================

func CheckStock(product string, quantity int) (bool, error) {

	var stock int

	err := config.DB.Table("products").
		Select("stock").
		Where("LOWER(name) = ?", product).
		Scan(&stock).Error

	if err != nil {
		return false, err
	}

	return stock >= quantity, nil
}

// ================== UPDATE STOCK ==================

func UpdateStock(product string, quantity int) error {

	fmt.Println("PRODUCT:", product)
	fmt.Println("QTY:", quantity)

	if quantity <= 0 {
		return fmt.Errorf("Invalid quantity")
	}

	result := config.DB.Table("products").
		Where("LOWER(name) = ? AND stock >= ?", strings.ToLower(product), quantity).
		Update("stock", gorm.Expr("stock - ?", quantity))

	fmt.Println("ROWS:", result.RowsAffected)

	if result.RowsAffected == 0 {
		return fmt.Errorf("Out of stock or insufficient stock")
	}

	return nil
}

// ================== ADMIN - GET ORDERS ==================

func GetOrders(c *gin.Context) {

	var orders []models.Order

	config.DB.Find(&orders)

	c.HTML(200, "admin.html", gin.H{
		"orders": orders,
	})
}

// ================== DELETE ORDER ==================

func DeleteOrder(c *gin.Context) {

	id := c.Param("id")

	config.DB.Delete(&models.Order{}, id)

	c.Redirect(302, "/admin/dashboard")
}

func GetUserOrders(c *gin.Context) {

	phone := c.Query("phone")

	var orders []models.Order

	config.DB.Where("phone = ?", phone).Find(&orders)

	c.HTML(200, "my_orders.html", gin.H{
		"orders": orders,
		"phone":  phone,
	})
}
