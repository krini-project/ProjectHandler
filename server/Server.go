package server

import (
	"fmt"
	"log"
	"path"

	"github.com/krini-project/ProjectHandler/persistence"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

// Server base struct for the endpoint server
type Server struct {
	Engine  *gin.Engine
	Handler *EndpointHandler
}

// Run Starts the endpoint server
// tlsPath Path to the tls cert and key file
// port Port for the server to listen on
// TODO: Currently only for self-signed certs (testing)
func (server *Server) Run(tlsPath string, port int) {
	certFile := path.Join(tlsPath, "certificate.crt")
	keyFile := path.Join(tlsPath, "certificate.key")

	err := server.Engine.RunTLS(fmt.Sprintf(":%v", port), certFile, keyFile)
	if err != nil {
		log.Println(err.Error())
	}
}

// Init Init method for the API endpoint to create all required handlers and the swagger endpoint
// @title Project handler API for the Krini project
// @version 0.1
// @basePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
func (server *Server) Init(databaseHandler *persistence.DatabaseHandler, port int) {
	server.Engine = gin.Default()
	server.Handler = &EndpointHandler{
		DatabaseHandler: databaseHandler,
	}

	server.handleSwagger()
	server.addHandlers()

	if !(port > 1 && port < 65535) {
		log.Fatalln(fmt.Sprintf("Could not start server, %v is not a valid port!", port))
	}
	server.Engine.Run(fmt.Sprintf(":%v", port))

}

func (server *Server) addHandlers() {
	server.Engine.GET("/projects", server.Handler.ListUserProjects)
	server.Engine.POST("/projects/adduser", server.Handler.AddUserToProject)
	server.Engine.POST("/projects/create", server.Handler.CreateProject)
	server.Engine.POST("/users/create", server.Handler.CreateUserIfNotExist)
}

func (server *Server) handleSwagger() {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	server.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
