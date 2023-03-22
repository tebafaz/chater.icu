// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Mukhamedjanov Erjan",
            "email": "tebafaz@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/guest/message": {
            "post": {
                "description": "Inserts data into database and publishes it to subscribers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Messages"
                ],
                "summary": "Publishes message as guest",
                "parameters": [
                    {
                        "description": "Inserts and publishes as guest account",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GuestMessageReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.StatusSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.StatusError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.StatusError"
                        }
                    }
                }
            }
        },
        "/last-messages": {
            "get": {
                "description": "Gets last 100 messages and latest message id for longpoll listening",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Messages"
                ],
                "summary": "Gets last messages",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MessageRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.StatusError"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Logs in by giving token. Token lasts 30 minutes and updates it at each performed action with token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Username and password",
                        "name": "tokenReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TokenReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TokenRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.StatusError"
                        }
                    }
                }
            }
        },
        "/messages": {
            "get": {
                "description": "Gets message downwards from given id with limit",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Messages"
                ],
                "summary": "Gets messages",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "message id from which gets older messages",
                        "name": "last_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "limit of messages in response",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MessageRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.StatusError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.StatusError"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Registers new user and returns token. Token lasts 30 minutes and updates it at each performed action with token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Register and get token",
                "parameters": [
                    {
                        "description": "Token",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TokenReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TokenRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.StatusError"
                        }
                    }
                }
            }
        },
        "/subscribe": {
            "get": {
                "description": "Subscribes to updates with longpoll connection",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Messages"
                ],
                "summary": "Subscribes to updates",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Client listening message id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MessageRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.StatusError"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.StatusError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.StatusError"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Logs out user by deleting his token in server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Logs out",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.StatusSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.StatusError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/user/message": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Inserts message into database and publishes it to subscribers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Messages"
                ],
                "summary": "publishes as user",
                "parameters": [
                    {
                        "description": "Message",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserMessageReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.StatusSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.StatusError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.StatusError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete message by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Messages"
                ],
                "summary": "Delete message",
                "parameters": [
                    {
                        "description": "Deletes message based on message id and provided token",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.StatusSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.StatusError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.StatusError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.GuestMessageReq": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.ID": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.Message": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "is_registered": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                },
                "sent_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.MessageRes": {
            "type": "object",
            "properties": {
                "last_id": {
                    "type": "integer"
                },
                "messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Message"
                    }
                }
            }
        },
        "models.StatusError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.StatusSuccess": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "string"
                }
            }
        },
        "models.TokenReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.TokenRes": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "models.UserMessageReq": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "tebafaz.com",
	BasePath:         "/api/v1",
	Schemes:          []string{"https"},
	Title:            "Chater.icu API",
	Description:      "Chater api made by Tebafaz using long poll as connction",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
