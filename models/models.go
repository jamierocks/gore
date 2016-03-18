package models

import (
    "github.com/jamierocks/gore/modules"
)

func AutoMigrate() {
    modules.DB.AutoMigrate(
        &User{},
        &Project{},
        &ProjectVersion{},
    )
}
