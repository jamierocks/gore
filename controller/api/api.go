package api

import (
    "github.com/jamierocks/gore/models"
    "github.com/jamierocks/gore/view/api"
    "gopkg.in/macaron.v1"
)

func GetUser(ctx *macaron.Context) {
    ctx.JSON(200, api.GetUserView(models.GetUser(ctx.Params("user"))))
}

func GetProject(ctx *macaron.Context) {
    ctx.JSON(200, api.GetProjectView(models.GetUser(ctx.Params("user")).GetProject(ctx.Params("project"))))
}
