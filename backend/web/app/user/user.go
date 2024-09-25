package user

import (
	"net/http"

	"bucks/database"

	"github.com/gin-gonic/gin"
)

// Handler returns the user profile information.
func Handler(ctx *gin.Context) {
	user := ctx.MustGet("user").(database.User)

	ctx.JSON(http.StatusOK, gin.H{
		"id":      user.ID,
		"name":    user.Name,
		"email":   user.Email,
		"picture": user.ProfilePicture,
	})
}
