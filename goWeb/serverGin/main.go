package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	router := gin.Default()

	router.POST("/edwin", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{"message": "pong", "txt": "txt"})
	})
	if err = http.ListenAndServe(":4568", router); err != nil {
		panic(err)
	}
}
