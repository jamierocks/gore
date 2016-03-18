package models

import (
    "time"
    "github.com/jamierocks/gore/modules"
)

type User struct {
    ID int64 `gorm:"primary_key"`
    Username string
    FullName string
    Type int64

    Projects []Project

    CreatedAt time.Time
    UpdatedAt time.Time
}

func (u User) IsOrganisation() bool {
    return u.Type == 1
}

func (u User) GetProjects() []Project {
    var projects []Project
    modules.DB.Model(&u).Related(&projects)
    return projects
}

func GetUser(username string) User {
    var user User
    modules.DB.First(&user, "username = ?", username)
    return user
}

func GetAllUsers() []User {
    var users []User
    modules.DB.Find(&users)
    return users
}