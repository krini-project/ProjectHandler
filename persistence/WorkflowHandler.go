package persistence

import (
	"fmt"
	"log"
	"time"
)

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

// CreateWorkflowRun Creates a new workflow run
func (handler *DatabaseHandler) CreateWorkflowRun(projectID uint, workflowTemplateID string) error {
	project := Project{}
	handler.Database.First(project, projectID)
	startTime := time.Now()

	workflow := WorkflowRun{
		Project:      project,
		StartTime:    startTime,
		EndTime:      startTime,
		State:        fmt.Sprintf("%v", Running),
		HasError:     false,
		HasResult:    false,
		Error:        "",
		ResultBucket: "",
		ResultKey:    "",
	}

	if err := handler.Database.Create(&workflow).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
