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

func (u User) GetProject(projectName string) Project {
    var project Project
    modules.DB.First(&project, "user_id = ? AND safe_name = ?", u.ID, projectName)
    return project
}

func GetUser(username string) User {
    var user User
    modules.DB.FirstOrInit(&user, "username = ?", username, User{Type: -1})
    return user
}

func GetAllUsers() []User {
    var users []User
    modules.DB.Find(&users)
    return users
}

func DoesUserExist(username string) bool {
    return GetUser(username).Type != -1
}
