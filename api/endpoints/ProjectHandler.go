package endpoints

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

// CreateProject Creates a new project
// @Summary Lists all projects of a user
// @Description Create project
// @Produce  json
// @Security ApiKeyAuth
// @Param   projectname     query    string     true        "project name"
// @Param   userid     query    string     true        "user id"
// @Success 200
// @Router /projects/create [post]
func (hander *EndpointHandler) CreateProject(c *gin.Context) {
	userID := c.Query("userid")
	projectName := c.Query("projectname")
	if userID == "" {
		c.AbortWithError(400, errors.New("user id required"))
		return
	}

	if projectName == "" {
		c.AbortWithError(400, errors.New("project name required"))
		return
	}

	err := hander.DatabaseHandler.CreateNewProject(projectName, userID)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(400, errors.New("Error while reading user projects from database"))
		return
	}

	c.Status(200)
}
