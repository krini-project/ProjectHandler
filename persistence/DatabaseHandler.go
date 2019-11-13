package persistence

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// DatabaseHandler Used to hold and initialize the database connection via GORM
type DatabaseHandler struct {
	Database *gorm.DB
}

// InitSQLite Initiates a database handler using SQLite
// path Path to store the database
func (handler *DatabaseHandler) InitSQLite(path string) error {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	handler.Database = db
	handler.handleMigrations()

	return nil
}

// UpdateWorkflowRunFinish Updates the status of a workflow run
func (handler *DatabaseHandler) UpdateWorkflowRunFinish(workflowRunID uint, resultBucket string, resultKey string) error {
	run := WorkflowRun{}
	if err := handler.Database.First(&run, workflowRunID).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	run.State = fmt.Sprintf("%v", Finished)
	run.ResultBucket = resultBucket
	run.ResultKey = resultKey

	handler.Database.Save(&run)
	return nil
}

// UpdateWorkflowRunWithError Updates the status of a workflow run
func (handler *DatabaseHandler) UpdateWorkflowRunWithError(workflowRunID uint, errorString string) error {
	run := WorkflowRun{}
	if err := handler.Database.First(&run, workflowRunID).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	run.State = fmt.Sprintf("%v", Error)
	run.Error = errorString

	handler.Database.Save(run)
	return nil
}

// ListProjectWorkflowRuns Lists all workflowruns associated with a project
func (handler *DatabaseHandler) ListProjectWorkflowRuns(projectID string) ([]WorkflowRun, error) {
	var workflowRuns []WorkflowRun

	project := Project{
		ProjectID: projectID,
	}

	if err := handler.Database.Model(&project).Related(workflowRuns).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return workflowRuns, nil
}

// ListProjects Lists all projects
func (handler *DatabaseHandler) ListProjects() ([]Project, error) {
	var projects []Project

	if err := handler.Database.Find(projects).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return projects, nil
}

// ListProjectsByUsers List all projects that are associated with a user
func (handler *DatabaseHandler) ListProjectsByUsers(userID string) ([]*Project, error) {
	user := User{
		UserID: userID,
	}

	if err := handler.Database.Preload("Projects").First(&user).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return user.Projects, nil
}

// CreateNewUser Creates a new User
func (handler *DatabaseHandler) CreateNewUser(userID string, userMail string) error {
	user := User{
		UserID:   userID,
		UserMail: userMail,
	}

	if err := handler.Database.Create(&user).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

// CreateNewProject Creates a new project
func (handler *DatabaseHandler) CreateNewProject(projectName string, userID string) (string, error) {
	user := User{
		UserID: userID,
	}

	if err := handler.Database.First(&user).Error; err != nil {
		log.Println(err.Error())
		return "", err
	}

	projectUUID := uuid.New().String()
	project := Project{
		ProjectID:   projectUUID,
		ProjectName: projectName,
		Users:       []*User{&user},
	}

	if err := handler.Database.Create(&project).Error; err != nil {
		log.Println(err.Error())
		return "", err
	}

	return project.ProjectID, nil
}

// CreateWorkflowRun Creates a new workflow run
func (handler *DatabaseHandler) CreateWorkflowRun(projectID uint, workflowTemplateID string) error {
	project := Project{}
	handler.Database.First(project, projectID)
	startTime := time.Now()

	workflow := WorkflowRun{
		Project:            project,
		StartTime:          startTime,
		EndTime:            startTime,
		State:              fmt.Sprintf("%v", Running),
		HasError:           false,
		HasResult:          false,
		Error:              "",
		ResultBucket:       "",
		ResultKey:          "",
		WorkflowTemplateID: workflowTemplateID,
	}

	if err := handler.Database.Create(&workflow).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

// AddUserToProject Add a user to an existing project
func (handler *DatabaseHandler) AddUserToProject(projectID string, userID string) error {
	project := Project{
		ProjectID: projectID,
	}

	user := User{
		UserID: userID,
	}

	if err := handler.Database.First(&user).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	if err := handler.Database.First(&project).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	project.Users = append(project.Users, &user)
	handler.Database.Save(project)

	return nil
}

func (handler *DatabaseHandler) handleMigrations() {
	if err := handler.Database.AutoMigrate(&User{}).Error; err != nil {
		log.Fatalln(err)
	}
	if err := handler.Database.AutoMigrate(&Project{}).Error; err != nil {
		log.Fatalln(err)
	}
	if err := handler.Database.AutoMigrate(&WorkflowRun{}).Error; err != nil {
		log.Fatalln(err)
	}
}
