package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const cookieName = "c"

//authCookieHandler gets the JWT and the uid and the cookie. If the given uid
//is  already tied to the given cookie, it does nothing and returns success.
//If the cookie is tied to a different uid, it barfs. If there is no UID, but
//there is a cookie, it removes that row in the DB and Set-Cookie's to remove
//the cookie. If there is no cookie, it validates the JWT, and then creates a
//new cookie tyied to that uid (creating that user record if necessary), and
//Set-Cookie's it back.
func (s *Server) authCookieHandler(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		panic("This can only be called as a post.")
	}

	uid := c.PostForm("uid")
	token := c.PostForm("token")

	cookie, _ := c.Cookie(cookieName)

	log.Println("Auth Cookie Handler called", uid, token, cookie, "*")

	c.JSON(http.StatusOK, gin.H{
		"Status": "Failure",
		"Error":  "Not Yet Implemented",
	})
}
