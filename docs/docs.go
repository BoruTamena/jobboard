// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Boru Tamene Yadeta",
            "url": "https://github.com/BoruTamena"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/signin": {
            "post": {
                "description": "user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "login to the system",
                "parameters": [
                    {
                        "description": "user login request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "create new user on boardapi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "create new user",
                "parameters": [
                    {
                        "description": "user request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/user/{id}/profile": {
            "get": {
                "description": "create profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "get users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "create profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "get users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user_id",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserProfie"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "create profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "create new user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user profile request body",
                        "name": "profile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserProfie"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.UserDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "dto.UserLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.UserProfie": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "skills": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "JobBoard",
	Description:      "This is a  Swagger API documentation for JobBoard Open source Project.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
