package endpoints

import (
	"errors"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddWorkflow Creates a new user if a user with the given id does not exist
// @Summary Creates a new user if none exists; not intended for production use: In production users are dynamically created based on external oauth2/oidc ids
// @Security ApiKeyAuth
// @Param   userid     query    string     true        "user id"
// @Param   projectid     query    string     true        "project id"
// @Param   name     query    string     true        "workflow name"
// @Param   link     query    string     true        "workflow link"
// @Param   version     query    string     true        "workflow version"
// @Success 200
// @Router /projects/workflows/add [post]
func (handler *EndpointHandler) AddWorkflow(c *gin.Context) {
	userID := c.Query("userid")
	projectid := c.Query("projectid")
	name := c.Query("name")
	link := c.Query("link")
	version := c.Query("version")
	if userID == "" {
		c.AbortWithError(400, errors.New("user id required"))
		return
	}

	if projectid == "" {
		c.AbortWithError(400, errors.New("project id required"))
		return
	}

	if name == "" {
		c.AbortWithError(400, errors.New("workflow name required"))
		return
	}

	if link == "" {
		c.AbortWithError(400, errors.New("link required"))
		return
	}

	err := handler.DatabaseHandler.AddWorkflow(projectid, userID, name, link, version)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(400, errors.New("Error while reading user projects from database"))
		return
	}
}

// GetWorkflow Returns all workflows of a given project
// @Summary Returns all workflows of a given project
// @Security ApiKeyAuth
// @Param   userid      query    string     true        "user id"
// @Param   projectid   query    string     true        "project id"
// @Param	workflowID	query    int		true		"workflow id"
// @Success 200 {object} models.Workflow
// @Router /projects/workflows/get [get]
func (handler *EndpointHandler) GetWorkflow(c *gin.Context) {
	userID := c.Query("userid")
	projectid := c.Query("projectid")
	workflowID := c.Query("workflowID")

	if userID == "" {
		c.AbortWithError(400, errors.New("user id required"))
		return
	}

	if projectid == "" {
		c.AbortWithError(400, errors.New("project id required"))
		return
	}

	if workflowID == "" {
		c.AbortWithError(400, errors.New("workflow id required"))
		return
	}

	workflowIDInt, err := strconv.Atoi(workflowID)
	if err != nil {
		c.AbortWithError(400, errors.New("workflow id needs to be integer"))
		return
	}

	workflow, err := handler.DatabaseHandler.GetWorkflow(workflowIDInt, projectid, userID)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	c.JSON(200, workflow)
}

// ListWorkflows Returns all workflows of a given project
// @Summary Returns all workflows of a given project
// @Security ApiKeyAuth
// @Param   userid      query    string     true        "user id"
// @Param   projectid   query    string     true        "project id"
// @Success 200 {array} models.Workflow
// @Router /projects/workflows/list [get]
func (handler *EndpointHandler) ListWorkflows(c *gin.Context) {
	userID := c.Query("userid")
	projectid := c.Query("projectid")

	if userID == "" {
		c.AbortWithError(400, errors.New("user id required"))
		return
	}

	if projectid == "" {
		c.AbortWithError(400, errors.New("project id required"))
		return
	}

	workflow, err := handler.DatabaseHandler.ListWorkflows(projectid, userID)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	c.JSON(200, workflow)
}
