package main

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published,omitempty"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var productsList = []Product{}
var idCounter = 0

func main() {
	loadProducts("products.json", &productsList)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("", GetAllProducts())
		products.GET(":id", GetProduct())
		products.GET("/search", SearchProduct())
		products.POST("", CreateProduct())
	}
	r.Run(":8080")
}

func loadProducts(path string, list *[]Product) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}

	idCounter = len(*list)
}

func GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, productsList)
	}
}

func GetProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid id"})
			return
		}
		for _, product := range productsList {
			if product.Id == id {
				ctx.JSON(200, product)
				return
			}
		}
		ctx.JSON(404, gin.H{"error": "product not found"})
	}
}

func SearchProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Query("priceGt")
		priceGt, err := strconv.ParseFloat(query, 32)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid price"})
			return
		}
		list := []Product{}
		for _, product := range productsList {
			if product.Price > priceGt {
				list = append(list, product)
			}
		}
		ctx.JSON(200, list)
	}
}

func CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newProduct Product
		if err := ctx.ShouldBindJSON(&newProduct); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if !validateProductCodeValue(newProduct.CodeValue) {
			ctx.JSON(400, gin.H{"error": "CodeValue must be unique"})
			return
		}

		if !validateProductExpiration(newProduct.Expiration) {
			ctx.JSON(400, gin.H{"error": "Invalid expiration date"})
			return
		}

		idCounter++
		newProduct.Id = idCounter
		productsList = append(productsList, newProduct)
		ctx.JSON(201, newProduct)
	}
}

func validateProductCodeValue(value string) bool {
	for _, product := range productsList {
		if product.CodeValue == value {
			return false
		}
	}
	return true
}

func validateProductExpiration(value string) bool {
	layout := "02/01/2006" // The expected date format

	expirationDate, err := time.Parse(layout, value)
	if err != nil {
		return false
	}

	if expirationDate.Month() < time.January || expirationDate.Month() > time.December {
		return false
	}

	return expirationDate.Format(layout) == value
}

//cargar un sl
