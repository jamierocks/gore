package main

import (
    "log"
    "net/http"
    "github.com/jamierocks/gore/modules"
    "github.com/jamierocks/gore/models"
    "github.com/jamierocks/gore/controller"
    apicon "github.com/jamierocks/gore/controller/api"
    authcon"github.com/jamierocks/gore/controller/auth"
    "github.com/go-macaron/pongo2"
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
    m.Use(pongo2.Pongoer())

    // Base routes
    m.Get("/", controller.GetHomepage)
    m.Get("/explore", controller.GetExplore)

    // API routes
    m.Group("/api", func() {
        m.Get("/:user", apicon.GetUser)
        m.Get("/:user/:project", apicon.GetProject)
    })

    // Auth routes
    m.Get("/login", authcon.GetLogin)
    m.Get("/auth/callback", authcon.GetCallback)

    // Run Gore
    log.Print("Listening on 0.0.0.0:" + modules.CONFIG.Section("web").Key("PORT").String())
    http.ListenAndServe("0.0.0.0:" + modules.CONFIG.Section("web").Key("PORT").String(), m)
}
