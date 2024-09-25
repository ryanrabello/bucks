package middleware

import (
	"net/http"

	"bucks/database"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserMiddleware makes the user available to the context.
func UserMiddleware(ctx *gin.Context) {
	userId := sessions.Default(ctx).Get("userId")

	if userId == nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user database.User
	tx := database.DB.First(&user, userId)

	if tx.Error != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("user", user)
	ctx.Next()
}
