// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/access-control": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "access-control"
                ],
                "summary": "manipulate message from Kafka",
                "parameters": [
                    {
                        "description": "model_type",
                        "name": "model_type",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/request.Request"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "HTTPError"
                        }
                    }
                }
            }
        },
        "/casbin/{role}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "casbin"
                ],
                "summary": "return of api that role can access to",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role",
                        "name": "role",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "HTTPError"
                        }
                    }
                }
            }
        },
        "/general/{type}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "general"
                ],
                "summary": "Show the list of objects by limit input.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "type",
                        "name": "type",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "HTTPError"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "general"
                ],
                "summary": "Create Object with Specified Type",
                "parameters": [
                    {
                        "type": "string",
                        "description": "type",
                        "name": "type",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "model-value",
                        "name": "model-value",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "HTTPError"
                        }
                    }
                }
            }
        },
        "/general/{type}/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "general"
                ],
                "summary": "Get active object by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "type",
                        "name": "type",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "HTTPError"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "general"
                ],
                "summary": "Update specified object with id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "type",
                        "name": "type",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "model_value",
                        "name": "model_value",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "general"
                ],
                "summary": "Deactive object by user_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "type",
                        "name": "type",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.Request": {
            "type": "object",
            "properties": {
                "method": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
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
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Access Control Service",
	Description: "manipulate data from Kafka",
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
