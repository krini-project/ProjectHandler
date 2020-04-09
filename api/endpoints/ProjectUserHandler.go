package endpoints

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

// RemoveUserFromProject Adds a user to a project
// @Summary Removes a user from a project
// @Produce  json
// @Security ApiKeyAuth
// @Param   projectid     query    string     true        "project id"
// @Param   userid     query    string     true        "user id"
// @Success 200
// @Router /projects/removeUser [post]
func (handler *EndpointHandler) RemoveUserFromProject(c *gin.Context) {
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

	err := handler.DatabaseHandler.RemoveUserFromProject(projectid, userID)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(400, errors.New("Error while reading user projects from database"))
		return
	}

	c.Status(200)
}

// AddUserToProject Adds a user to a project
// @Summary Lists all projects of a user
// @Produce  json
// @Security ApiKeyAuth
// @Param   projectid     query    string     true        "project id"
// @Param   userid     query    string     true        "user id"
// @Success 200
// @Router /projects/adduser [post]
func (handler *EndpointHandler) AddUserToProject(c *gin.Context) {
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

	err := handler.DatabaseHandler.AddUserToProject(projectid, userID)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(400, errors.New("Error while reading user projects from database"))
		return
	}

	c.Status(200)
}

// ListUserProjects Lists all projects of a user
// @Summary Lists all projects of a user
// @Produce  json
// @Security ApiKeyAuth
// @Param   userid     query    string     true        "user id"
// @Success 200 {array} models.Project
// @Router /projects [get]
func (handler *EndpointHandler) ListUserProjects(c *gin.Context) {
	userID := c.Query("userid")
	if userID == "" {
		c.AbortWithError(400, errors.New("user id required"))
		return
	}
	projects, err := handler.DatabaseHandler.ListUserProjects(userID)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(400, errors.New("Error while reading user projects from database"))
		return
	}

	c.JSON(200, projects)
}
