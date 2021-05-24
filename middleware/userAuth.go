package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authz validates token and authorizes users
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, err := c.Request.Cookie("token")
		if err == http.ErrNoCookie {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.AbortWithStatus(http.StatusTemporaryRedirect)
			c.Abort()
		}

		//claims, err := helper.ValidateToken(clientToken)
		//if err != "" {
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		//	c.Abort()
		//	return
		//}

		//c.Set("email", claims.Email)
		//c.Set("first_name", claims.First_name)
		//c.Set("last_name", claims.Last_name)
		//c.Set("uid", claims.Uid)
		//c.Set("user_type", claims.User_type)

		c.Next()

	}
}
