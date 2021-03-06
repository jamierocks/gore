package models

import (
    "time"
    "github.com/jamierocks/gore/modules"
)

type Project struct {
    ID int64 `gorm:"primary_key"`
    Name string
    SafeName string

    User User
    UserID int64

    Versions []ProjectVersion

    CreatedAt time.Time
    UpdatedAt time.Time
}

func (p Project) GetOwner() User {
    var user User
    modules.DB.Model(&p).Related(&user)
    return user
}

func (p Project) GetVersions() []ProjectVersion {
    var versions []ProjectVersion
    modules.DB.Model(&p).Related(&versions)
    return versions
}

func (p Project) GetVersion(version string) ProjectVersion {
    var projectVersion ProjectVersion
    modules.DB.First(&projectVersion, "project_id = ? AND version = ?", p.ID, version)
    return projectVersion
}

type ProjectVersion struct {
    ID int64 `gorm:"primary_key"`
    Version string
    Channel string

    Project Project
    ProjectID int64

    CreatedAt time.Time
    UpdatedAt time.Time
}

func (v ProjectVersion) getProject() Project {
    var project Project
    modules.DB.Model(&v).Related(&project)
    return project
}

func GetAllProjects() []Project {
    var projects []Project
    modules.DB.Find(&projects)
    return projects
}
