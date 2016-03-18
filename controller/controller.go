package controller

import (
    "github.com/jamierocks/gore/models"
    "gopkg.in/macaron.v1"
)

func GetExplore(ctx *macaron.Context) {
    ctx.Data["projects"] = models.GetAllProjects()
    ctx.HTML(200, "explore")
}
