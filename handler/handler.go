package handler

import (
	"net/http"
	"database/sql"
	"github.com/gin-gonic/gin"
	"Piyaporn_agnos_backend/model"
)

func AddDataHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData struct {
			PasswordInput string `json:"password_input"`
		}
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !model.IsValidInput(requestData.PasswordInput) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := model.AddData(db, requestData.PasswordInput); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add data"})
			return
		}

		numStep := model.CalculateNumStep(requestData.PasswordInput)
		c.JSON(http.StatusOK, gin.H{"num_of_steps": numStep})
	}
}
