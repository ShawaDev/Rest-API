package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func DaysUntill(c *gin.Context) {
	header := c.GetHeader("User-Role")

	if header == "admin" {
		fmt.Print("red button user detected")
	}

	currentTime := time.Now().String()
	currentYear, _ := strconv.ParseInt(currentTime[0:4], 10, 32)
	currentMonth, _ := strconv.ParseInt(currentTime[5:7], 10, 32)
	currentDay, _ := strconv.ParseInt(currentTime[8:10], 10, 32)

	currentDate := time.Date(int(currentYear), time.Month(currentMonth), int(currentDay), 0, 0, 0, 0, time.UTC)
	finalDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	days := finalDate.Sub(currentDate).Hours() / 24

	c.JSON(200, gin.H{
		"message": days,
	})
}
