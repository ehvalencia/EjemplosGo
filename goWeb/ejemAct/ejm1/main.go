package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type users struct {
	Name     string `json:"name"`
	Apellido string `json:"apellido"`
}

func main() {
	var err error
	var user users
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "Pong")
	})

	router.POST("/saludo", func(a *gin.Context) {
		a.BindJSON(&user)

		a.String(200, "Hola "+user.Name+" "+user.Apellido)
	})

	if err = http.ListenAndServe(":8989", router); err != nil {
		panic(err)
	}

}
