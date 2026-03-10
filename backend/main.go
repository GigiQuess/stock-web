package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/k3243/stock-web/backend/ent"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Database connection
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run migrations
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

    // A simple endpoint to create a user
	r.POST("/api/users", func(c *gin.Context) {
        var u struct {
            Name string `json:"name"`
            Email string `json:"email"`
        }
        if err := c.ShouldBindJSON(&u); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
		user, err := client.User.
			Create().
			SetName(u.Name).
			SetEmail(u.Email).
			Save(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
		}
		c.JSON(http.StatusOK, user)
	})

	// Endpoint to get all users
	r.GET("/api/users", func(c *gin.Context) {
		users, err := client.User.Query().All(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})


	r.Run(":8080")
}
