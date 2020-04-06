package persistence

import (
	"log"

	"github.com/google/uuid"
)

// AddUserToProject Add a user to an existing project
func (handler *DatabaseHandler) AddUserToProject(projectID string, userID string) error {
	project := Project{
		ProjectID: projectID,
	}

	user := User{
		UserID: userID,
	}

	role := Role{
		Name: InternRole,
	}

	if err := handler.Database.First(&user).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	if err := handler.Database.First(&project).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	if err := handler.Database.First(&role).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	userProjectRole := ProjectUser{
		User: &user,
		Role: &role,
	}

	project.Users = append(project.Users, &userProjectRole)
	handler.Database.Save(project)

	return nil
}

//RemoveUserFromProject Removes a user from a project
func (handler *DatabaseHandler) RemoveUserFromProject(projectID string, userID string) error {
	project := Project{
		ProjectID: projectID,
	}

	user := User{
		UserID: userID,
	}

	projectuser := ProjectUser{
		User: &user,
	}

	if err := handler.Database.First(&project).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	if err := handler.Database.First(&user).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	if err := handler.Database.First(&projectuser).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	if err := handler.Database.Model(&project).Association("Users").Delete(&projectuser).Error; err != nil {
		log.Println(err.Error())
		return err
	}
	handler.Database.Save(project)

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
		Users:       []*ProjectUser{},
	}

	if err := handler.Database.Create(&project).Error; err != nil {
		log.Println(err.Error())
		return "", err
	}

	return project.ProjectID, nil
}
