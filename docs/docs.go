// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "Developpement",
        "contact": {
            "name": "AD0791",
            "email": "alexandrodisla@hotmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/file/todos": {
            "get": {
                "description": "Retrieve all todos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Get all todos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/filemodel.Todo"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new todo to the list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Create a new todo",
                "parameters": [
                    {
                        "description": "Todo to create",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.TodoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/filemodel.Todo"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/file/todos/{id}": {
            "get": {
                "description": "Retrieve a single todo by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Get todo by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/filemodel.Todo"
                        }
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Modify an existing todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Update a todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated todo data",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.TodoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.TodoResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove a todo from the json list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Delete a todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/schema.MessageResponse"
                        }
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/service/todos": {
            "get": {
                "description": "Retrieve all todos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "service"
                ],
                "summary": "Get all todos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.TodoResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new todo to the list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "service"
                ],
                "summary": "Create a new todo",
                "parameters": [
                    {
                        "description": "Todo to create",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.TodoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/schema.TodoResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/service/todos/{id}": {
            "get": {
                "description": "Retrieve a single todo by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "service"
                ],
                "summary": "Get todo by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.TodoResponse"
                        }
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Modify an existing todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "service"
                ],
                "summary": "Update a todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated todo data",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.TodoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.TodoResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove a todo from the list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "service"
                ],
                "summary": "Delete a todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/schema.MessageResponse"
                        }
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sql/todos": {
            "get": {
                "description": "Retrieve all todos from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sql"
                ],
                "summary": "Get all todos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.TodoSQLResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Database error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sql"
                ],
                "summary": "Create todo",
                "parameters": [
                    {
                        "description": "Todo request body",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.TodoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/schema.TodoSQLResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Database error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sql/todos/{id}": {
            "get": {
                "description": "Retrieve a single todo by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sql"
                ],
                "summary": "Get todo by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.TodoSQLResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sql"
                ],
                "summary": "Update todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated todo data",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.TodoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.TodoSQLResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Database error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Soft delete a todo by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sql"
                ],
                "summary": "Delete todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/schema.MessageSQLResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Database error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "filemodel.Todo": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string",
                    "minLength": 3
                }
            }
        },
        "schema.MessageResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "schema.MessageSQLResponse": {
            "type": "object",
            "properties": {
                "deleted_at": {
                    "type": "string",
                    "example": "2023-01-01"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "schema.TodoRequest": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "completed": {
                    "description": "example:\"false\"",
                    "type": "boolean",
                    "example": false
                },
                "title": {
                    "description": "example:\"Sample Todo\"",
                    "type": "string",
                    "minLength": 3,
                    "example": "Sample Todo"
                }
            }
        },
        "schema.TodoResponse": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string",
                    "minLength": 3
                }
            }
        },
        "schema.TodoSQLResponse": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": false
                },
                "created_at": {
                    "description": "Formatted date",
                    "type": "string",
                    "example": "2023-01-01"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "title": {
                    "type": "string",
                    "example": "Sample Todo"
                },
                "updated_at": {
                    "description": "Formatted date",
                    "type": "string",
                    "example": "2023-01-01"
                }
            }
        }
    },
    "tags": [
        {
            "description": "File-based todo operations",
            "name": "file"
        },
        {
            "description": "Service-based todo operations",
            "name": "service"
        },
        {
            "description": "SQL-based todo operations",
            "name": "sql"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.1",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Todo API Documentation",
	Description:      "API documentation for Todo service with persistence",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
