//routes.go

package main

import (
	"github.com/gin-gonic/gin"
	p "github.com/pe5er/mmindia-website/pages"
	"fmt"
	"net/http"
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
	e.GET("/userarea", p.Userarea)
	e.Static("/css", "static/css/")
  	e.GET("/auth/github", redirectHandler)
  	e.GET("/auth/github/callback", callbackHandler)



	return e
}


func redirectHandler(c *gin.Context) {
 // Define our settings
    appSettings := map[string]string{
        "clientID":     "c5d5ccfa293ba72d15c9",
        "clientSecret": "dc96f2fa2f9d41b3eb95ed57b9b57dc7394b1855",
        "redirectURL":  "http://www.mmindia.uk/auth/github/callback",
    }

    // Retrieve the URL
    authURL, err := gocial.New().
        Driver("github").
        Scopes([]string{"public_repo"}). // Specify custom scopes
        Redirect(
            appSettings["clientID"],
            appSettings["clientSecret"],
            appSettings["redirectURL"],
                )

        // Check for errors
        if err != nil {
                c.Writer.Write([]byte("Error: " + err.Error()))
                return
        }

        // Redirect to authURL with status 302
        c.Redirect(http.StatusFound, authURL)
}

// Collect user info
func callbackHandler(c *gin.Context) {
        // Retrieve query params for state and code
        state := c.Query("state")
        code := c.Query("code")

        // Handle callback and check for errors
        user, token, err := gocial.Handle(state, code)
        if err != nil {
                c.Writer.Write([]byte("Error: " + err.Error()))
                return
        }

        // Print token information
    fmt.Printf("%#v", token)
        // Print in terminal user information
    fmt.Printf("%#v", user)

        // If no errors, show provider name
        c.Writer.Write([]byte("Hi, " + user.FullName))
}

