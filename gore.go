package main

import (
    "github.com/jamierocks/gore/modules"
    "github.com/jamierocks/gore/models"
    apicon "github.com/jamierocks/gore/controller/api"
    "github.com/jamierocks/gore/controller/auth"
    "github.com/gin-gonic/gin"
)

func main() {
    // Load config
    modules.InitConfig()

    // Load database
    modules.InitDatabase()

    // Run migrations
    models.AutoMigrate()

    // Construct Gin
    g := gin.Default()
    g.Static("/maven", "storage")

    // API routes
    g.GET("/api/:user", apicon.GetUser)
    g.GET("/api/:user/:project", apicon.GetProject)

    // Auth routes
    g.GET("/login", auth.GetLogin)
    g.GET("/auth/callback", auth.GetCallback)

    // Run Gore
    g.Run(":" + modules.CONFIG.Section("web").Key("PORT").String())
}
