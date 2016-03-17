package main

import (
    "github.com/jamierocks/gore/modules"
    "github.com/gin-gonic/gin"
)

func main() {
    // Load config
    modules.InitConfig()

    // Construct Gin
    g := gin.Default()

    // Run Gore
    g.Run(":" + modules.CONFIG.Section("web").Key("PORT").String())
}
