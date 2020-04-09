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

// CreateUserIfNotExist Creates a new user if a user with the given id does not exist
// @Summary Creates a new user if none exists; not intended for production use: In production users are dynamically created based on external oauth2/oidc ids
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
