package router

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"bucks/platform/authenticator"
	"bucks/platform/invitation"
	"bucks/platform/middleware"
	"bucks/web/app/callback"
	"bucks/web/app/login"
	"bucks/web/app/logout"
	"bucks/web/app/user"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	// router.LoadHTMLGlob("web/template/*")

	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/logout", logout.Handler)
	router.GET("/user", middleware.UserMiddleware, user.Handler)

	router.POST("/invitation-code", middleware.UserMiddleware, invitation.Handler)

	// router.GET("/test", middleware.IsAuthenticated, func(c *gin.Context) {
	// 	c.Status(200)
	// })

	return router
}
