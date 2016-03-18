package api

import (
    "github.com/jamierocks/gore/models"
)

type UserView struct {
    ID int64 `json:"id"`
    Username string `json:"username"`
    FullName string `json:"fullName"`
    Type int64 `json:"type"`
    Projects []UserProjectView `json:"projects"`
}

type UserProjectView struct {
    Name string `json:"name"`
    SafeName string `json:"safeName"`
}

func GetUserView(user models.User) UserView {
    var projects []UserProjectView

    for _, project := range user.GetProjects() {
        projects = append(projects, getUserProjectView(project))
    }

    return UserView{
        ID: user.ID,
        Username: user.Username,
        FullName: user.FullName,
        Type: user.Type,
        Projects: projects,
    }
}

func getUserProjectView(project models.Project) UserProjectView {
    return UserProjectView{
        Name: project.Name,
        SafeName: project.SafeName,
    }
}
