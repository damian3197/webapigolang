package main

import (
	"net/http"
	"tugaswebapi/models"
	"tugaswebapi/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	// Models
	db := models.SetUpModels()
	r.Use(func(c *gin.Context){
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message" : "golang web api",
		})
	})
	r.GET("/api/v1/mahasiswa", controller.GetData)
	r.POST("/api/v1/mahasiswa", controller.CreateData)

	r.Run()
}