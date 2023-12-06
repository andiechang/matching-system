// utils.go

package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

// RespondWithError 用於處理錯誤響應。
func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}

// LogError 用於記錄錯誤。
func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}
