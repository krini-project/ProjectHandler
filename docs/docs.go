// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-11-13 15:14:10.296930131 +0100 CET m=+0.030155985

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/projects": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Lists all projects of a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "userid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.ProjectListWrapper"
                        }
                    }
                }
            }
        },
        "/projects/adduser": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Lists all projects of a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "projectid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "userid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/projects/create": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Lists all projects of a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project name",
                        "name": "projectname",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "userid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/create": {
            "post": {
                "summary": "Creates a new user if none exists",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "userid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user email address",
                        "name": "usermail",
                        "in": "query",
                        "required": true
                    }
                ]
            }
        }
    },
    "definitions": {
        "persistence.Project": {
            "type": "object",
            "properties": {
                "projectID": {
                    "type": "string"
                },
                "projectName": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/persistence.User"
                    }
                }
            }
        },
        "persistence.User": {
            "type": "object",
            "properties": {
                "projects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/persistence.Project"
                    }
                },
                "userID": {
                    "type": "string"
                },
                "userMail": {
                    "type": "string"
                }
            }
        },
        "server.ProjectListWrapper": {
            "type": "object",
            "properties": {
                "projects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/persistence.Project"
                    }
                },
                "userID": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "X-API-Key",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Project handler API for the Krini project",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
