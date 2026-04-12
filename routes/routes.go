package routes

import (
	"fmt"
	"khadbhandar/config"
	"khadbhandar/controllers"
	"khadbhandar/models"
	"strings"

	"github.com/gin-gonic/gin"
)

// ================= ADMIN AUTH =================

// func AdminAuth(c *gin.Context) {
// 	c.Next()
// }

// ================= ROUTES =================

func SetupRoutes(r *gin.Engine) {

	fmt.Println("---- ROUTES START ----")

	// HOME
	// r.GET("/", controllers.GetProducts)

	// PRODUCT
	r.GET("/product", showProduct)

	// ORDER
	r.GET("/order", func(c *gin.Context) {

		productName := c.Query("product")

		var productData models.Product

		config.DB.Where("LOWER(name) = ?", strings.ToLower(productName)).First(&productData)

		c.HTML(200, "order.html", gin.H{
			"product": productName,
			"stock":   productData.Stock, // 🔥 THIS LINE IMPORTANT
		})
	})

	r.GET("/success", func(c *gin.Context) {
		product := c.Query("product")
		c.HTML(200, "success.html", gin.H{
			"product": product,
		})
	})

	// USER
	r.POST("/place-order", controllers.PlaceOrder)
	r.POST("/admin/add-product", controllers.AddProduct)
	r.POST("/admin/delete-product", controllers.DeleteProduct)

	// ADMIN LOGIN
	r.GET("/admin/login", controllers.ShowLogin)
	r.POST("/admin/login", controllers.HandleLogin)

	fmt.Println("ADMIN ROUTES REGISTERED")
	//  DASHBOARD (IMPORTANT)
	r.GET("/admin/dashboard", controllers.GetDashboard)
	r.GET("/", controllers.GetProducts)
	r.GET("/my-orders", controllers.GetUserOrders)
	r.GET("/admin/approve/:id", controllers.ApproveOrder)

	// ADMIN ACTIONS
	r.GET("/admin/delete/:id", controllers.DeleteOrder)
	r.POST("/admin/update-stock", controllers.UpdateStockAdmin)

	r.GET("/test", func(c *gin.Context) {
		c.String(200, "TEST OK")
	})
}

// ================= PRODUCT DATA =================
func showProduct(c *gin.Context) {

	product := c.Query("name")

	var productData = map[string]map[string]string{

		"urea": {
			"price": "₹266 / Bag",
			"desc":  "Best nitrogen fertilizer",
			"usage": "Apply after 20-25 days",
			"crop":  "Wheat, Rice",
			"image": "/static/images/urea1.jpg",
		},

		"dap": {
			"price": "₹1350 / Bag",
			"desc":  "High phosphorus fertilizer",
			"usage": "Use at sowing time",
			"crop":  "All crops",
			"image": "/static/images/dap1.jpg",
		},

		"npk": {
			"price": "₹1470 / Bag",
			"desc":  "Balanced fertilizer",
			"usage": "Early stage",
			"crop":  "Vegetables, Mustard",
			"image": "/static/images/npk1.jpg",
		},

		"potash": {
			"price": "₹1700 / Bag",
			"desc":  "Improves crop quality",
			"usage": "Flowering stage",
			"crop":  "All crops",
			"image": "/static/images/potash1.jpg",
		},

		"nano": {
			"price": "₹240 / Bottle",
			"desc":  "Advanced liquid nitrogen fertilizer",
			"usage": "Spray on crops",
			"crop":  "All crops",
			"image": "/static/images/nano1.jpg",
		},

		"ssp": {
			"price": "₹450 / Bag",
			"desc":  "Single super phosphate fertilizer",
			"usage": "Use at sowing time",
			"crop":  "All crops",
			"image": "/static/images/ssp1.jpg",
		},
		"zinc": {
			"price": "₹500 / Bag",
			"desc":  "Improves crop growth and yield",
			"usage": "Apply during early growth stage",
			"crop":  "Wheat, Rice, Vegetables, Mustard",
			"image": "/static/images/zinc1.jpg",
		},

		"sulphur": {
			"price": "₹280 / Bag",
			"desc":  "Essential nutrient for oilseed crops",
			"usage": "Use before flowering",
			"crop":  "Mustard, Pulses",
			"image": "/static/images/sulphur1.jpg",
		},

		"sagarika": {
			"price": "₹450",
			"desc":  "Seaweed based plant growth promoter",
			"usage": "Spray during growth stage",
			"crop":  "Vegetables, Fruits",
			"image": "/static/images/sagarika1.jpg",
		},

		"atrazine": {
			"price": "₹650",
			"desc":  "Weed control herbicide",
			"usage": "Apply after sowing",
			"crop":  "Bajra, Maize",
			"image": "/static/images/atrazine1.jpg",
		},

		"glyphosate": {
			"price": "₹700",
			"desc":  "Non selective weed killer",
			"usage": "Spray before planting",
			"crop":  "All crops",
			"image": "/static/images/glyphosate1.jpg",
		},

		"seeds": {
			"price": "1700 / 1100 / 3000",
			"desc":  "High quality crop seeds",
			"usage": "Sowing time",
			"crop":  "Wheat, Mustard, Bajra",
			"image": "/static/images/seeds1.jpg",
		},

		"sardar nutri": {
			"price": "1000",
			"desc":  "High quality crop seeds",
			"usage": "Sowing time",
			"crop":  "Wheat, Mustard, Bajra, Vegitables",
			"image": "/static/images/nutri1.jpg",
		},

		"fortex": {
			"price": "150 to 850",
			"desc":  "It protects crops like rice, wheat, and vegetables by being applied to the soil, providing long-term systemic protection",
			"usage": "at the time of planting or early crop stages ",
			"crop":  "wheat, paddy (rice), and various vegetables.",
			"image": "/static/images/fortex1.jpg",
		},

		"imidacloprid": {
			"price": "Various",
			"desc":  " control sucking and chewing insects in agriculture, horticulture, and pet care",
			"usage": "early in the morning or late afternoon",
			"crop":  "cotton, paddy, sugarcane, vegetables (tomato, okra), and fruits",
			"image": "/static/images/imidacloprid1.jpg",
		},

		"cotton seeds": {
			"price": "750 to 900",
			"desc":  " Cottonseed is the oil-rich, ovoid seed of the cotton plant, produced inside the cotton boll alongside fiber, (e.g., Rasi RCH 947, RCH 999, MRC 7373)",
			"usage": "primarily used to produce cottonseed oil, high-protein livestock feed (meal), and linters for industrial materials",
			"crop":  "cotton",
			"image": "/static/images/cotton1.jpg",
		},

		"palak seeds": {
			"price": "120",
			"desc":  " small, light brown, and slightly irregular in shape, designed for easy, fast-growing cultivation of nutritious green leafy vegetables",
			"usage": " seeds are used to grow nutrient-dense leafy greens",
			"crop":  "palak",
			"image": "/static/images/palak1.jpg",
		},

		"tomato seeds": {
			"price": "150 / Packet",
			"desc":  "High yield hybrid tomato seeds",
			"usage": "Sowing",
			"crop":  "Vegetable",
			"image": "/static/images/tomato1.jpg",
		},
	}

	// ================= PRODUCT FUNCTION =================

	data, exists := productData[product]

	if !exists {
		c.String(404, "Product not found")
		return
	}

	c.HTML(200, "product.html", gin.H{
		"product": product,
		"price":   data["price"],
		"desc":    data["desc"],
		"usage":   data["usage"],
		"crop":    data["crop"],
		"image":   data["image"],
	})
}
