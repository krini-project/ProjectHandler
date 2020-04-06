package endpoints

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/krini-project/ProjectHandler/persistence"
)

// EndpointHandler Base struct for the API endpoints
type EndpointHandler struct {
	DatabaseHandler *persistence.DatabaseHandler
}

// CreateProject Creates a new project
// @Summary Lists all projects of a user
// @Produce  json
// @Security ApiKeyAuth
// @Param   projectname     query    string     true        "project name"
// @Param   userid     query    string     true        "user id"
// @Success 200 {string} string	"ok"
// @Router /projects/create [post]
func (hander *EndpointHandler) CreateProject(c *gin.Context) {
	userID := c.Query("userid")
	projectName := c.Query("projectname")
	if userID == "" {
		c.AbortWithError(400, errors.New("user id required"))
		return
	}

	if projectName == "" {
		c.AbortWithError(400, errors.New("project id required"))
		return
	}

	projectID, err := hander.DatabaseHandler.CreateNewProject(projectName, userID)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(400, errors.New("Error while reading user projects from database"))
		return
	}

	c.String(200, "%v", projectID)
}

// CreateUserIfNotExist Creates a new user if a user with the given id does not exist
// @Summary Creates a new user if none exists
// @Security ApiKeyAuth
// @Param   userid     query    string     true        "user id"
// @Param   usermail     query    string     true        "user email address"
// @Router /users/create [post]
func (hander *EndpointHandler) CreateUserIfNotExist(c *gin.Context) {
	userID := c.Query("userid")
	usermail := c.Query("usermail")
	if userID == "" {
		c.AbortWithError(400, errors.New("user id required"))
		return
	}

	err := hander.DatabaseHandler.CreateNewUser(userID, usermail)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(400, errors.New("Error while reading user projects from database"))
		return
	}
}
