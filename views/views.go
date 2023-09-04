// views/views.go
package views

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var loggedInUsers = make(map[string]bool)

func Welcome(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("static/login.html"))
	tmpl.Execute(c.Writer, nil)
}

func Home(c *gin.Context) {
	if loggedInUsers[c.ClientIP()] {
		tmpl := template.Must(template.ParseFiles("static/home.html"))
		tmpl.Execute(c.Writer, nil)
	} else {
		c.Redirect(http.StatusSeeOther, "/")
	}
}

func Logout(c *gin.Context) {
	loggedInUsers[c.ClientIP()] = false
	c.Redirect(http.StatusSeeOther, "/")
}

func HandleLogin(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "12" && password == "13" {
			loggedInUsers[c.ClientIP()] = true
			http.Redirect(c.Writer, c.Request, "/Home", http.StatusSeeOther)
		} else {
			tmpl := template.Must(template.ParseFiles("static/login.html"))
			tmpl.Execute(c.Writer, gin.H{"Error": "Invalid credentials"})
		}
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func WD(router *gin.Engine) {
	router.GET("/", Welcome)
	router.GET("/Home", Home)
	router.GET("/Logout", Logout)
	router.POST("/login", HandleLogin)
	router.Static("/static", "./static")
}
