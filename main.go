package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/AfterShip/email-verifier"
)

var (
	verifier = emailverifier.
		NewVerifier().
		EnableAutoUpdateDisposable()
)

func main() {
	r := gin.Default()

	r.GET("/verify-email", func(c *gin.Context) {
		email := c.Query("email")
		if email == "" {
			c.JSON(400, gin.H{"error": "Email parameter is missing"})
			return
		}

		ret, err := verifier.Verify(email)
		if err != nil {
			c.JSON(500, gin.H{"error": "Email verification failed", "message": err.Error()})
			return
		}

		if !ret.Syntax.Valid {
			c.JSON(400, gin.H{"error": "Email address syntax is invalid"})
			return
		}

		c.JSON(200, gin.H{
			"email":         ret.Email,
			"disposable":    ret.Disposable,
			"role_account":  ret.RoleAccount,
			"free":          ret.Free,
			"username":      ret.Syntax.Username,
			"domain":        ret.Syntax.Domain,
			"valid":         ret.Syntax.Valid,
			"has_mx_records": ret.HasMxRecords,
			"smtp":          ret.SMTP,
			"gravatar":      ret.Gravatar,
		})
	})

	r.Run(":8080")
}
