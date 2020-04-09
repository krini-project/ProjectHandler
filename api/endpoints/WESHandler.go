package endpoints

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

// AddWESEndpoint Adds a new WES endpoint
// @Summary Creates a new user if none exists; not intended for production use: In production users are dynamically created based on external oauth2/oidc ids
// @Security ApiKeyAuth
// @Param   userid      query    string     true        "user id"
// @Param   projectid   query    string     true        "project id"
// @Param   name     	query    string     true        "endpoint name"
// @Param   link     	query    string     true        "wes link"
// @Router /projects/wes/add [post]
func (hander *EndpointHandler) AddWESEndpoint(c *gin.Context) {
	userID := c.Query("userid")
	projectid := c.Query("projectid")
	name := c.Query("name")
	link := c.Query("link")
	if userID == "" {
		c.AbortWithError(400, errors.New("user id required"))
		return
	}

	if projectid == "" {
		c.AbortWithError(400, errors.New("project id required"))
		return
	}

	if name == "" {
		c.AbortWithError(400, errors.New("name required"))
		return
	}

	if projectid == "" {
		c.AbortWithError(400, errors.New("link required"))
		return
	}

	err := hander.DatabaseHandler.AddWESEndpoint(name, link, projectid)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(400, errors.New("Error while reading user projects from database"))
		return
	}
}

// ListWESEndpoints Adds a new WES endpoint
// @Summary Creates a new user if none exists; not intended for production use: In production users are dynamically created based on external oauth2/oidc ids
// @Security ApiKeyAuth
// @Param   userid      query    string     true        "user id"
// @Param   projectid   query    string     true        "project id"
// @Success 200 {array} models.WESEndpoint
// @Router /projects/wes/list [get]
func (hander *EndpointHandler) ListWESEndpoints(c *gin.Context) {
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

	endpoints, err := hander.DatabaseHandler.ListWESEndpoints(userID, projectid)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(400, errors.New("Error while reading user projects from database"))
		return
	}

	c.JSON(200, endpoints)
}
