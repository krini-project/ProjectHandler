package persistence

import (
	"context"
	"log"

	"github.com/krini-project/ProjectHandler/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

//AddWorkflow Adds a workflow to a project
func (handler *DatabaseHandler) AddWorkflow(projectid string, userid string, name string, link string, version string) error {
	workflow := models.Workflow{
		Link:         link,
		Workflowname: name,
		ProjectID:    projectid,
		Version:      null.NewString(version, true),
	}

	if err := workflow.Insert(context.TODO(), handler.DB, boil.Infer()); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

//GetWorkflow Retrives the informations about a workflow of a project
func (handler *DatabaseHandler) GetWorkflow(id int, projectid string, userid string) (*models.Workflow, error) {
	workflow, err := models.FindWorkflow(context.TODO(), handler.DB, id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return workflow, nil
}

//ListWorkflows Lists all workflows of a given projects
func (handler *DatabaseHandler) ListWorkflows(projectid string, userid string) ([]*models.Workflow, error) {
	project, err := models.FindProject(context.TODO(), handler.DB, projectid)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	workflowSlice, err := project.Workflows().All(context.TODO(), handler.DB)
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
