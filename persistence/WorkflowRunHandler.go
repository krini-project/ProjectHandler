package persistence

// UpdateWorkflowRunFinish Updates the status of a workflow run
func (handler *DatabaseHandler) UpdateWorkflowRunFinish(workflowRunID uint, resultBucket string, resultKey string) error {

	return nil
}

// UpdateWorkflowRunWithError Updates the status of a workflow run
func (handler *DatabaseHandler) UpdateWorkflowRunWithError(workflowRunID uint, errorString string) error {

	return nil
}

// CreateWorkflowRun Creates a new workflow run
func (handler *DatabaseHandler) CreateWorkflowRun(projectID uint, workflowTemplateID string) error {

	return nil
}
