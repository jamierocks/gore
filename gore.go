package main

import (
    "github.com/jamierocks/gore/modules"
    "github.com/jamierocks/gore/models"
    apicon "github.com/jamierocks/gore/controller/api"
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

    // API routes
    g.GET("/api/:user", apicon.GetUser)
    g.GET("/api/:user/:project", apicon.GetProject)

    // Run Gore
    g.Run(":" + modules.CONFIG.Section("web").Key("PORT").String())
}
