package api

import (
    "github.com/jamierocks/gore/models"
    "github.com/jamierocks/gore/view/api"
    "github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
    ctx.JSON(200, api.GetUserView(models.GetUser(ctx.Param("user"))))
}