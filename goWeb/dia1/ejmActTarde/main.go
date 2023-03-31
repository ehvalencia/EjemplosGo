/*
Problema
Un supermercado necesita un sistema para gestionar los productos frescos que tienen publicados en su página web. Para poder hacer esto, necesitan un servidor que ejecute una API que les permita manipular los productos cargados desde distintos clientes. Los campos que conforman un producto son:
Nombre
Tipo de dato
JSON
Tipo de dato
GO
Descripción | Ejemplo
id
number
int
Identificador en conjunto de datos | 15
name
string
string
Nombre caracterizado | Cheese - St. Andre
quantity
number
int
Cantidad almacenada | 60
code_value
string
string
Código alfanumérico característico | S73191A
is_published
boolean
bool
El producto se encuentra publicado o no |  True
expiration
string
string
Fecha de vencimiento | 12/04/2022
price
number
float64
Precio del producto | 50.15

Ejercicio 1 : Iniciando el proyecto
Debemos crear un repositorio en github.com para poder subir nuestros avances. Este repositorio es el que vamos a utilizar para llevar lo que realicemos durante las distintas prácticas de Go Web.
Primero debemos clonar el repositorio creado, luego iniciar nuestro proyecto de go con con el comando go mod init.
El siguiente paso será crear un archivo main.go donde deberán cargar en una slice, desde un archivo JSON, los datos de productos. Esta slice se debe cargar cada vez que se inicie la API para realizar las distintas consultas.
El archivo para trabajar es el siguiente:

Ejercicio 2 : Creando un servidor
Vamos a levantar un servidor utilizando el paquete gin en el puerto 8080. Para probar nuestros endpoints haremos uso de postman.
Crear una ruta /ping que debe respondernos con un string que contenga pong con el status 200 OK.
Crear una ruta /products que nos devuelva la lista de todos los productos en la slice.
Crear una ruta /products/:id que nos devuelva un producto por su id.
Crear una ruta /products/search que nos permita buscar por parámetro los productos cuyo precio sean mayor a un valor priceGt.
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var products []Product

func ReadText(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		panic("El archivo indicado no fue encontrado o esta daniado")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&products); err != nil { // el archivo se incenta a un slice
		panic(err) // Retornar errores
	}

	fmt.Println(products)
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, "Pong")
}

func ProductList(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func ProductById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	exist := false
	var product Product

	for _, current_product := range products {
		if current_product.Id == id {
			exist = true
			product = current_product
			break
		}
	}

	if exist == false {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Product does not exist",
		})
	} else {
		c.JSON(http.StatusOK, product)
	}

}

func SearchProductByParam(c *gin.Context) {
	param, err := strconv.ParseFloat(c.Query("priceGt"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Param",
		})
	}

	var products_response []Product

	for _, current_product := range products {
		if current_product.Price > param {
			products_response = append(products_response, current_product)
		}
	}

	if len(products_response) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No existen datos",
		})
	}

	c.JSON(http.StatusOK, products_response)

}

func main() {
	ReadText("products.json")

	server := gin.Default()
	server.GET("/ping", Pong)
	server.GET("/products", ProductList)
	server.GET("/products/:id", ProductById)
	server.GET("/products/search", SearchProductByParam)

	if err := http.ListenAndServe(":8080", server); err != nil {
		panic(err)
	}
}
