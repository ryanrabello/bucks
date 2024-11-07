package invitation

import (
	"net/http"

	"bucks/database"

	"github.com/gin-gonic/gin"
)

type InvitationCodePayload struct {
	Code        string `json:"code"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

func Handler(c *gin.Context) {
	user := c.MustGet("user").(database.User)

	if !user.IsRyan {
		c.String(http.StatusForbidden, "Only Ryan can create invitation codes")
		return
	}

	var payload InvitationCodePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if payload.Amount <= 0 {
		c.String(http.StatusBadRequest, "Amount must be greater than 0")
		return
	}

	code := database.InvitationCode{
		Code:        payload.Code,
		Amount:      payload.Amount,
		Description: payload.Description,
	}

	err := database.DB.Create(&code).Error
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invitation code created successfully"})
}
