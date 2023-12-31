// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/todos": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get all items in the todo list",
                "operationId": "get-all-todos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Todo"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "add a new item to the todo list",
                "operationId": "create-todo",
                "parameters": [
                    {
                        "description": "Todo Request",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TodoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Todo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get a todo item by ID",
                "operationId": "get-todo-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Todo"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "update a todo item by ID",
                "operationId": "update-todo-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Todo Request",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TodoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Todo"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "delete a todo item by ID",
                "operationId": "delete-todo-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Todo"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Todo not found"
                }
            }
        },
        "model.Todo": {
            "type": "object",
            "properties": {
                "deadline": {
                    "type": "string",
                    "example": "2022-09-15T14:30:45.0000001+07:00"
                },
                "description": {
                    "type": "string",
                    "example": "Something to do in the weekend"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Assigntment 1"
                },
                "status": {
                    "type": "string",
                    "example": "New"
                }
            }
        },
        "model.TodoRequest": {
            "type": "object",
            "properties": {
                "deadline": {
                    "type": "string",
                    "example": "2022-09-15T14:30:45.0000001+07:00"
                },
                "description": {
                    "type": "string",
                    "example": "Something to do in the weekend"
                },
                "name": {
                    "type": "string",
                    "example": "Assigntment 1"
                },
                "status": {
                    "type": "string",
                    "example": "New"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Final Project 1 [Go + Gin Todo API]",
	Description:      "This is a simple server todo server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
