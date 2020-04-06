package endpoints

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/krini-project/ProjectHandler/models"
)

// RemoveUserFromProject Adds a user to a project
// @Summary Removes a user from a project
// @Produce  json
// @Security ApiKeyAuth
// @Param   projectid     query    string     true        "project id"
// @Param   userid     query    string     true        "user id"
// @Success 200
// @Router /projects/removeUser [post]
func (hander *EndpointHandler) RemoveUserFromProject(c *gin.Context) {
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

	err := hander.DatabaseHandler.RemoveUserFromProject(projectid, userID)
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
func (hander *EndpointHandler) AddUserToProject(c *gin.Context) {
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

	err := hander.DatabaseHandler.AddUserToProject(projectid, userID)
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
// @Success 200 {object} models.ProjectListWrapper
// @Router /projects [get]
func (hander *EndpointHandler) ListUserProjects(c *gin.Context) {
	userID := c.Query("userid")
	if userID == "" {
		c.AbortWithError(400, errors.New("user id required"))
		return
	}
	projects, err := hander.DatabaseHandler.ListUserProjects(userID)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(400, errors.New("Error while reading user projects from database"))
		return
	}

	wrapper := models.ProjectListWrapper{
		Projects: projects,
		UserID:   userID,
	}

	c.JSON(200, wrapper)
}
