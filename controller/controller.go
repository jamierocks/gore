package controller

import (
    "github.com/jamierocks/gore/models"
    "gopkg.in/macaron.v1"
)

func GetHomepage(ctx *macaron.Context) {
    ctx.HTML(200, "index")
}

func GetExplore(ctx *macaron.Context) {
    ctx.Data["projects"] = models.GetAllProjects()
    ctx.HTML(200, "explore")
}
