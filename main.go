package main

import (
	"gin-docker-todo/domain"
	"gin-docker-todo/infrastructure"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	Env_load()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	infrastructure.DbInit()

	r.GET("/", func(ctx *gin.Context) {
		todos := infrastructure.DbRead()
		ctx.HTML(http.StatusOK, "index.html", gin.H{"todos": todos})
	})

	r.POST("/new", func(c *gin.Context) {
		text := c.PostForm("text")
		rawStatus := c.PostForm("status")
		id, err := strconv.Atoi(rawStatus)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		status := domain.Status(id)
		rawTime := c.PostForm("deadline")
		deadline, err := strconv.Atoi(rawTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		todo := domain.Todo{Text: text, Status: status, Deadline: deadline}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		infrastructure.DbCreate(todo)
		c.Redirect(302, "/")
	})

	r.POST("/new", func(c *gin.Context) {
		text := c.PostForm("text")
		rawStatus := c.PostForm("status")
		id, err := strconv.Atoi(rawStatus)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		status := domain.Status(id)
		rawTime := c.PostForm("deadline")
		deadline, err := strconv.Atoi(rawTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		todo := domain.Todo{Text: text, Status: status, Deadline: deadline}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		infrastructure.DbCreate(todo)
		c.Redirect(302, "/")
	})

	r.Run()
}
