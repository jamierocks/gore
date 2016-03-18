package main

import (
    "log"
    "net/http"
    "github.com/jamierocks/gore/modules"
    "github.com/jamierocks/gore/models"
    apicon "github.com/jamierocks/gore/controller/api"
    "github.com/jamierocks/gore/controller/auth"
    "gopkg.in/macaron.v1"
)

func main() {
    // Load config
    modules.InitConfig()

    // Load database
    modules.InitDatabase()

    // Run migrations
    models.AutoMigrate()

    // Construct Macaron
    m := macaron.Classic()
    m.Use(macaron.Renderer())

    // API routes
    m.Group("/api", func() {
        m.Get("/:user", apicon.GetUser)
        m.Get("/:user/:project", apicon.GetProject)
    })

    // Auth routes
    m.Get("/login", auth.GetLogin)
    m.Get("/auth/callback", auth.GetCallback)

    // Run Gore
    log.Print("Listening on 0.0.0.0:" + modules.CONFIG.Section("web").Key("PORT").String())
    http.ListenAndServe("0.0.0.0:" + modules.CONFIG.Section("web").Key("PORT").String(), m)
}
