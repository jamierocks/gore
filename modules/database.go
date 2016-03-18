package modules

import (
    "log"
    "strings"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
    DB *gorm.DB
)

func InitDatabase() {
    var err error

    switch strings.ToLower(CONFIG.Section("database").Key("DB_TYPE").String()) {
    case "sqlite":
        DB, err = gorm.Open("sqlite3", CONFIG.Section("database").Key("PATH").String())
    case "mysql":
    // TODO: mysql
    }

    if err != nil {
        log.Fatal("Failed to load database!", err)
    }
}
