package server

import "github.com/krini-project/ProjectHandler/persistence"

// ProjectListWrapper Wrapper structs to store a list of projects
type ProjectListWrapper struct {
	UserID   string
	Projects []*persistence.Project
}
