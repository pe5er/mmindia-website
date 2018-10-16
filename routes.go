//routes.go

package main

import (
	"github.com/gin-gonic/gin"
	p "github.com/pe5er/mmindia-website/pages"
)

func routes() *gin.Engine {

	e := gin.Default()

	e.NoRoute(func(c *gin.Context) {
		c.Redirect(301, "/")
	})

	// Static Pages
	e.GET("/", p.Home)
	e.GET("/gallery", p.Gallery)
	e.GET("/news", p.News)
	e.GET("/robots.txt", p.Robots)
	e.GET("/privacy", p.Privacy)
	e.GET("/history", p.History)
	e.GET("/contact", p.Contact)
	e.GET("/team", p.Team)
	e.Static("/css", "static/css/")

	return e
}
