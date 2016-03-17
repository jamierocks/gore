package modules

import (
    "log"
    "gopkg.in/ini.v1"
)

var (
    CONFIG *ini.File
)

func InitConfig() {
    var err error

    CONFIG, err = ini.Load([]byte(""), "config.ini")
    if err != nil {
        log.Fatal("Failed to load config", err)
    }
}
