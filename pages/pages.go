//pages.go

package pages

import (
	"github.com/gin-gonic/gin"
)

func sHTML(c *gin.Context, code int, name string, obj gin.H) {
	c.HTML(code, name, obj)
}

func Contact(c *gin.Context) {
	sHTML(c, 200, "contact.html", nil)
}
func Gallery(c *gin.Context) {
	sHTML(c, 200, "gallery.html", nil)
}
func Robots(c *gin.Context) {
	sHTML(c, 200, "robots.txt", nil)
}
func Privacy(c *gin.Context) {
	sHTML(c, 200, "privacy.html", nil)
}
func News(c *gin.Context) {
	sHTML(c, 200, "news.html", nil)
}
func Team(c *gin.Context) {
	sHTML(c, 200, "team.html", nil)
}
func Pm(c *gin.Context) {
        sHTML(c, 200, "pm.html", nil)
}
func Ethics(c *gin.Context) {
        sHTML(c, 200, "ethics.html", nil)
}
func Documents(c *gin.Context) {
        sHTML(c, 200, "documents.html", nil)
}
func Blog(c *gin.Context) {
        sHTML(c, 200, "blog.html", nil)
}
func Login(c *gin.Context) {
        sHTML(c, 200, "login.html", nil)
}

