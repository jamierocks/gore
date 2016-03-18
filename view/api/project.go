package api

import (
    "github.com/jamierocks/gore/models"
)

type ProjectView struct {
    ID int64 `json:"id"`
    Name string `json:"name"`
    SafeName string `json:"safeName"`
    Versions []ProjectVersionView `json:"versions"`
}

type ProjectVersionView struct {
    Version string `json:"version"`
    Channel string `json:"channel"`
}

func GetProjectView(project models.Project) ProjectView {
    var versions []ProjectVersionView

    for _, version := range project.GetVersions() {
        versions = append(versions, getProjectVersionView(version))
    }

    return ProjectView{
        ID: project.ID,
        Name: project.Name,
        SafeName: project.SafeName,
        Versions: versions,
    }
}

func getProjectVersionView(version models.ProjectVersion) ProjectVersionView {
    return ProjectVersionView{
        Version: version.Version,
        Channel: version.Channel,
    }
}
