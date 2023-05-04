// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/topics": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "topics"
                ],
                "summary": "Create Topic",
                "operationId": "Topic-Create",
                "parameters": [
                    {
                        "description": "New Topic Information",
                        "name": "options",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/topic.CreateInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/topics/owner/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "topics"
                ],
                "summary": "Find Topics By Owner ID",
                "operationId": "Topic-Find-Owned",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Owner ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/topic.Topic"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/topics/pair/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "topics"
                ],
                "summary": "Find Paired Topics By User ID",
                "operationId": "Topic-Find-Paired",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Involved ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/topic.Topic"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/topics/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "topics"
                ],
                "summary": "Find Topic By ID",
                "operationId": "Topic-Find",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Topic ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/topic.Topic"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "topics"
                ],
                "summary": "Delete Topic",
                "operationId": "Topic-Delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Topic ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login With Credentials",
                "operationId": "User-Login",
                "parameters": [
                    {
                        "description": "User Credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register New User",
                "operationId": "User-Register",
                "parameters": [
                    {
                        "description": "New User Info",
                        "name": "options",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/update-password": {
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Upate Password",
                "operationId": "User-Update-Password",
                "parameters": [
                    {
                        "description": "User Credentials and New Password",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UpdatePasswordInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Find User By ID",
                "operationId": "User-FindByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "topic.Category": {
            "type": "string",
            "enum": [
                "software",
                "social_sciences",
                "other"
            ],
            "x-enum-varnames": [
                "Software",
                "SocialSciences",
                "Other"
            ]
        },
        "topic.CreateInput": {
            "type": "object",
            "required": [
                "capacity",
                "category",
                "title"
            ],
            "properties": {
                "capacity": {
                    "type": "integer",
                    "minimum": 1
                },
                "category": {
                    "$ref": "#/definitions/topic.Category"
                },
                "title": {
                    "type": "string",
                    "minLength": 3
                }
            }
        },
        "topic.Topic": {
            "type": "object",
            "properties": {
                "capacity": {
                    "description": "Max expected parties",
                    "type": "integer"
                },
                "category": {
                    "$ref": "#/definitions/topic.Category"
                },
                "created_at": {
                    "type": "string"
                },
                "finished_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "owner": {
                    "description": "ID of the owner User",
                    "type": "string"
                },
                "parties": {
                    "description": "IDs of invloved Users",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "user.LoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "maxLength": 24,
                    "minLength": 3
                }
            }
        },
        "user.RegisterInput": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "maxLength": 24,
                    "minLength": 3
                }
            }
        },
        "user.UpdatePasswordInput": {
            "type": "object",
            "required": [
                "new_password",
                "password"
            ],
            "properties": {
                "new_password": {
                    "type": "string",
                    "minLength": 6
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.User": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Paiwr Server",
	Description:      "Paiwr Server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
