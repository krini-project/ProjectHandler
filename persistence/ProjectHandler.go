package persistence

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/krini-project/ProjectHandler/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// AddUserToProject Add a user to an existing project
func (handler *DatabaseHandler) AddUserToProject(projectID string, userID string) error {
	err := handler.insertProjectUser(projectID, userID, InternRole)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

//RemoveUserFromProject Removes a user from a project
func (handler *DatabaseHandler) RemoveUserFromProject(projectID string, userID string) error {

	return nil
}

// ListProjectWorkflowRuns Lists all workflowruns associated with a project
func (handler *DatabaseHandler) ListProjectWorkflowRuns(projectID string) error {

	return nil
}

// ListProjects Lists all projects
func (handler *DatabaseHandler) ListProjects() ([]*models.Project, error) {
	projects, err := models.Projects().All(context.TODO(), handler.DB)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var projectSlice []*models.Project
	for _, project := range projects {
		projectSlice = append(projectSlice, project)
	}

	return projectSlice, nil
}

// CreateNewProject Creates a new project
func (handler *DatabaseHandler) CreateNewProject(projectName string, userID string) error {
	err := handler.CreateNewUser(userID, "")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	projectID := uuid.New().String()

	project := models.Project{
		ProjectID:   projectID,
		ProjectName: null.NewString(projectName, true),
	}

	err = project.Insert(context.TODO(), handler.DB, boil.Infer())
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = handler.insertProjectUser(projectID, userID, AdminRole)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (handler *DatabaseHandler) ListProjectWorkflows(id int) ([]*models.Workflow, error) {
	workflowSlice, err := models.Workflows(qm.Where("Project_id=?", id)).All(context.TODO(), handler.DB)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var workflows []*models.Workflow
	for _, workflow := range workflowSlice {
		workflows = append(workflows, workflow)
	}

	return workflows, nil
}

func (handler *DatabaseHandler) ListProjectWESEndpoints(id int) ([]*models.WESEndpoint, error) {
	wesEndpointsSlice, err := models.WESEndpoints(qm.Where("Owner_id=?", id)).All(context.TODO(), handler.DB)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var endpoints []*models.WESEndpoint
	for _, endpoint := range wesEndpointsSlice {
		endpoints = append(endpoints, endpoint)
	}

	return endpoints, nil
}

func (handler *DatabaseHandler) insertProjectUser(projectID string, userID string, roleName string) error {

	role, err := models.Roles(qm.Where("RoleName=?", roleName)).One(context.TODO(), handler.DB)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	projectUser := models.ProjectUser{
		UserID:            userID,
		RoleID:            role.RoleID,
		ProjectsProjectID: projectID,
	}

	err = projectUser.Insert(context.TODO(), handler.DB, boil.Infer())
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
