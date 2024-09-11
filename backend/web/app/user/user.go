package user

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Handler returns the user profile information.
func Handler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")
	email := session.Get("email")

	// TODO: make custom session struct for user

	if profile == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":      profile.(map[string]interface{})["sub"],
		"name":    profile.(map[string]interface{})["name"],
		"email":   email,
		"picture": profile.(map[string]interface{})["picture"],
	})
}
