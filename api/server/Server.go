package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/krini-project/ProjectHandler/api/endpoints"
	middleware "github.com/krini-project/ProjectHandler/api/middlware"
	"github.com/krini-project/ProjectHandler/persistence"

	"github.com/gin-gonic/gin"
)

// Server base struct for the endpoint server
type Server struct {
	Engine  *gin.Engine
	Handler *endpoints.EndpointHandler
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
// @basePath /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
func (server *Server) Init(databaseHandler *persistence.DatabaseHandler, port int) {
	server.Engine = gin.Default()
	server.Handler = &endpoints.EndpointHandler{
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
	apiKey := os.Getenv("APIKey")
	if apiKey == "" {
		apiKey = "test123"
	}

	v1Group := server.Engine.Group("/v1")
	v1Group.Use(middleware.ApiKeyMiddleware(apiKey))

	v1Group.Handle("GET", "/projects", server.Handler.ListUserProjects)
	v1Group.Handle("POST", "/projects/adduser", server.Handler.AddUserToProject)
	v1Group.Handle("POST", "/projects/create", server.Handler.CreateProject)
	v1Group.Handle("POST", "/users/create", server.Handler.CreateUserIfNotExist)
}

func (server *Server) handleSwagger() {
	swaggerStaticPath := "/home/marius/Code/GA4GH/ProjectHandler/www/swagger-ui"
	if os.Getenv("swaggerStaticPath") != "" {
		swaggerStaticPath = os.Getenv("swaggerStaticPath")
	}

	fs := http.FileSystem(http.Dir(swaggerStaticPath))
	server.Engine.StaticFS("/swagger-ui/", fs)
}
