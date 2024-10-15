package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default() // in this version two default mid are attached, logger, recovery
	//r := gin.New() // this router have no middlewares attached
	// Simple "Hello, World!" endpoint
	r.GET("/", home)

	// JSON Response
	r.GET("/json", func(c *gin.Context) {

		c.JSON(200, struct {
			Message string `json:"message"`
		}{
			Message: "Hello, World!",
		})

	})

	// Route Parameters
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello, %s! (Gin)", name)
	})

	// Query Parameters
	r.GET("/welcome", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstName", "Guest")
		lastName := c.Query("lastName")
		c.String(200, "Hello, %s %s! (Gin)", firstName, lastName)
	})

	// Grouping Routes
	v1 := r.Group("/v1")
	{
		v1.GET("/users", func(c *gin.Context) {
			c.String(200, "Users v1 (Gin)")
		})
		v1.POST("/posts", func(c *gin.Context) {
			c.String(200, "Posts v1 (Gin)")
		})
	}

	v2 := r.Group("/v2")
	{
		r.Use(gin.Logger())
		v2.GET("/users", func(c *gin.Context) {
			c.String(200, "Users v2 (Gin)")
		})
		v2.GET("/posts", func(c *gin.Context) {
			c.String(200, "Posts v2 (Gin)")
		})
	}

	r.GET("/custom-error", func(c *gin.Context) {
		err := errors.New("custom error message")
		//AbortWithStatusJSON would not quit the request, we need to return manually
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	})
	r.Run(":8081")
}

func home(c *gin.Context) {
	//go panic("boom!") // if panic happens in seperate goroutine,
	//then we need to do manual panic recovery, otherwise http panic recovery can handle it
	c.String(200, "Hello, World! (Gin)")
}

/*
# GET /
curl http://localhost:8081/

# GET /json
curl http://localhost:8081/json

# GET /user/:name (example with 'name' as 'John')
curl http://localhost:8081/user/John

# GET /welcome with query parameters
curl "http://localhost:8081/welcome?firstName=John&lastName=Doe"

# GET /v1/users
curl http://localhost:8081/v1/users

# GET /v1/posts
curl http://localhost:8081/v1/posts

# GET /v2/users
curl http://localhost:8081/v2/users

# GET /v2/posts
curl http://localhost:8081/v2/posts

# GET /custom-error
curl http://localhost:8081/custom-error
*/
