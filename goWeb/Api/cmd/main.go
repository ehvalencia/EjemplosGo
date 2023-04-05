package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"

	handler "github.com/ehvalencia/EjemplosGo/EjemplosGo/goWeb/Api/cmd/handlers"
	"github.com/ehvalencia/EjemplosGo/EjemplosGo/goWeb/Api/internal/domain"
	"github.com/ehvalencia/EjemplosGo/EjemplosGo/goWeb/Api/internal/products"
	"github.com/ehvalencia/EjemplosGo/EjemplosGo/goWeb/Api/pkg/store"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("/Users/edvalencia/Documents/GO/EjemplosGo/goWeb/Api/token.env")
	if err != nil {
		panic(err)
	}

	//var productsList = []domain.Product{}
	//loadProducts("/Users/edvalencia/Documents/GO/EjemplosGo/goWeb/Api/products.json", &productsList)

	// Extract products data from the JSON file
	jsonStore := store.NewJsonStore("/Users/edvalencia/Documents/GO/EjemplosGo/goWeb/Api/products.json")
	productsList, err := jsonStore.GetAll()
	if err != nil {
		panic(err)
	}

	repo := products.NewRepository(productsList)
	service := products.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	products := r.Group("/products")
	{
		products.GET("", productHandler.GetAll())
		products.GET(":id", productHandler.GetByID())
		products.GET("/search", productHandler.Search())
		products.POST("", productHandler.Post())
		products.DELETE(":id", productHandler.Delete())
		products.PATCH(":id", productHandler.Patch())

	}
	r.Run(":8080")
}

// loadProducts carga los productos desde un archivo json
func loadProducts(path string, list *[]domain.Product) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}
}
