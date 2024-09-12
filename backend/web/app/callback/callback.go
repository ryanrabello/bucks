package callback

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"bucks/database"
	"bucks/platform/authenticator"
)

// Handler for our callback.
func Handler(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if ctx.Query("state") != session.Get("state") {
			ctx.String(http.StatusBadRequest, "Invalid state parameter.")
			return
		}

		// Exchange an authorization code for a token.
		token, err := auth.Exchange(ctx.Request.Context(), ctx.Query("code"))
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Failed to convert an authorization code into a token.")
			return
		}

		idToken, err := auth.VerifyIDToken(ctx.Request.Context(), token)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to verify ID Token.")
			return
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Extract email from the profile
		email, ok := profile["email"].(string)
		if !ok {
			ctx.String(http.StatusInternalServerError, "Failed to get email from profile")
			return
		}

		// Check if the user already exists
		var user database.User
		database.DB.Where("idp_type = ? AND idp_user_id = ?", "auth0", idToken.Subject).First(&user)

		if user.ID == 0 {
			// User doesn't exist, create a new user
			user = database.User{
				Name:           profile["name"].(string),
				Email:          email,
				IDPUserId:      idToken.Subject,
				IDPType:        "auth0",
				ProfilePicture: profile["picture"].(string),
			}

			if err := database.DB.Create(&user).Error; err != nil {
				ctx.String(http.StatusInternalServerError, "Failed to create new user")
				return
			}
		}

		session.Set("access_token", token.AccessToken)
		session.Set("userId", user.ID)

		if err := session.Save(); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Redirect to logged in page.
		ctx.Redirect(http.StatusTemporaryRedirect, "/user")
	}
}
