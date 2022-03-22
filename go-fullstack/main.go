package main

import (
	"net/http"

	"go-fullstack/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		stories, err := models.LoadTopStories()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"stories": stories,
		})
	})
	r.GET("/loaderio-345f0e73092256270286441af8d5e2fe", func(c *gin.Context){
		c.Header("Content-Type", "text/plain; charset=utf-8")
		c.String("loaderio-345f0e73092256270286441af8d5e2fe")	
	})
	r.Run("0.0.0.0:8000")
}
