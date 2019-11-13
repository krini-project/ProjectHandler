package server

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

// ListUserProjects Lists all projects of a user
// @Summary Lists all projects of a user
// @Produce  json
// @Param   userid     query    string     true        "user id"
// @Success 200 {object} server.ProjectListWrapper
// @Router /projects [get]
func (server *EndpointHandler) ListUserProjects(c *gin.Context) {
	userID := c.Query("userid")
	if userID == "" {
		c.AbortWithError(400, errors.New("user id required"))
		return
	}
	projects, err := server.DatabaseHandler.ListProjectsByUsers(userID)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(400, errors.New("Error while reading user projects from database"))
		return
	}

	wrapper := ProjectListWrapper{
		Projects: projects,
		UserID:   userID,
	}

	c.JSON(200, wrapper)
}

// AddUserToProject Adds a user to a project
// @Summary Lists all projects of a user
// @Produce  json
// @Param   projectid     query    string     true        "project id"
// @Param   userid     query    string     true        "user id"
// @Success 200
// @Router /projects/adduser [post]
func (server *EndpointHandler) AddUserToProject(c *gin.Context) {
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

	err := server.DatabaseHandler.AddUserToProject(projectid, userID)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(400, errors.New("Error while reading user projects from database"))
		return
	}

	c.Status(200)
}

// CreateProject Creates a new project
// @Summary Lists all projects of a user
// @Produce  json
// @Param   projectname     query    string     true        "project name"
// @Param   userid     query    string     true        "user id"
// @Success 200 {string} string	"ok"
// @Router /projects/create [post]
func (server *EndpointHandler) CreateProject(c *gin.Context) {
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

	projectID, err := server.DatabaseHandler.CreateNewProject(projectName, userID)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(400, errors.New("Error while reading user projects from database"))
		return
	}

	c.String(200, "%v", projectID)
}

// CreateUserIfNotExist Creates a new user if a user with the given id does not exist
// @Summary Creates a new user if none exists
// @Param   userid     query    string     true        "user id"
// @Param   usermail     query    string     true        "user email address"
// @Router /users/create [post]
func (server *EndpointHandler) CreateUserIfNotExist(c *gin.Context) {
	userID := c.Query("userid")
	usermail := c.Query("usermail")
	if userID == "" {
		c.AbortWithError(400, errors.New("user id required"))
		return
	}

	err := server.DatabaseHandler.CreateNewUser(userID, usermail)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(400, errors.New("Error while reading user projects from database"))
		return
	}
}
