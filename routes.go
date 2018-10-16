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
	e.GET("/contact", p.Contact)
        e.GET("/team", p.Team)
        e.GET("/pm", p.Pm)
        e.GET("/ethics", p.Ethics)
        e.GET("/docs", p.Documents)
        e.GET("/blog", p.Blog)
        e.GET("/login", p.Login)
	e.Static("/css", "static/css/")

	return e
}
