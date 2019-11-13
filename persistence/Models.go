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

// User Model for a normal user
type User struct {
	UserID   string `gorm:"primary_key; unique_index"`
	UserMail string
	Projects []*Project `gorm:"many2many:user_projects;"`
}

// Project Model for projects
type Project struct {
	ProjectID   string  `gorm:"primary_key;type:varchar(100)"`
	ProjectName string  `gorm:"index"`
	Users       []*User `gorm:"many2many:user_projects;"`
}

// WorkflowRun Model for storing issued worfklows
type WorkflowRun struct {
	gorm.Model
	Project            Project `gorm:"index"`
	StartTime          time.Time
	EndTime            time.Time
	State              string
	HasError           bool
	HasResult          bool
	Error              string
	WorkflowTemplateID string
	ResultBucket       string
	ResultKey          string
}
