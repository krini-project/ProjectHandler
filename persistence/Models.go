package persistence

import (
	"time"

	"github.com/jinzhu/gorm"
)

// State of a workflow
type State int

const (
	Running  State = 0
	Error    State = 1
	Finished State = 2
)

//Right model
type Right struct {
	gorm.Model
	Name string `gorm:"UNIQUE_INDEX"`
}

//Role model
type Role struct {
	gorm.Model
	Name   string   `gorm:"UNIQUE_INDEX"`
	Rights []*Right `gorm:"many2many:role_rights;"`
}

//ProjectUser model
type ProjectUser struct {
	gorm.Model
	User *User
	Role *Role
}

// Workflow model to store a workflow
// Workflowdescription: The base64 encoded workflow yaml
type Workflow struct {
	gorm.Model
	WorkflowName        string `gorm:"index"`
	WorkflowDescription string
	Version             string
}

// User Model for a normal user
type User struct {
	UserID   string `gorm:"primary_key; unique_index"`
	Username string
	UserMail string
	Projects []*Project `gorm:"many2many:user_projects;"`
}

// Project Model for projects
type Project struct {
	ProjectID   string         `gorm:"primary_key;type:varchar(100)"`
	ProjectName string         `gorm:"index"`
	Users       []*ProjectUser `gorm:"many2many:user_projects;"`
	Workflows   []*Workflow    `gorm:"many2many:project_worfklows;"`
}

// WorkflowRun Model for storing issued worfklows
type WorkflowRun struct {
	gorm.Model
	Project      Project `gorm:"index"`
	StartTime    time.Time
	EndTime      time.Time
	State        string
	HasError     bool
	HasResult    bool
	Error        string
	Workflow     Workflow
	ResultBucket string
	ResultKey    string
}
