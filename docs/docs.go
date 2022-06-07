// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"github.com/swaggo/swag"
)

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://selly.vn",
        "contact": {
            "name": "Dev team",
            "url": "https://selly.vn",
            "email": "dev@cashbag.vn"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/roles": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.ResponseList"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "Create",
                "operationId": "role-create",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apimodel.RoleCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.ResponseCreate"
                        }
                    }
                }
            }
        },
        "/roles/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "Detail",
                "operationId": "role-detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.RoleDetail"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "Update",
                "operationId": "role-update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apimodel.RoleUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.ResponseUpdate"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "All",
                "operationId": "user-all",
                "parameters": [
                    {
                        "type": "string",
                        "name": "banned",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "fromAt",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "toAt",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.User"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Detail",
                "operationId": "user-detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.UserDetail"
                        }
                    }
                }
            }
        },
        "/users/{id}/banned": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Banned",
                "operationId": "user-banned",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.ResponseUpdate"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apimodel.RoleCreate": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "permission": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "apimodel.RoleUpdate": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "permission": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "constant.Permission": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.UserStatistic": {
            "type": "object",
            "properties": {
                "totalCustomer": {
                    "type": "integer"
                },
                "totalOrder": {
                    "type": "integer"
                },
                "totalOrderCanceled": {
                    "type": "integer"
                },
                "totalOrderCompleted": {
                    "type": "integer"
                },
                "totalOrderDelivered": {
                    "type": "integer"
                },
                "totalOrderDelivering": {
                    "type": "integer"
                },
                "totalOrderPending": {
                    "type": "integer"
                },
                "totalOrderWaitingApproved": {
                    "type": "integer"
                }
            }
        },
        "ptime.TimeResponse": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                }
            }
        },
        "responsemodel.ResponseCreate": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "responsemodel.ResponseList": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "list": {},
                "total": {
                    "type": "integer"
                }
            }
        },
        "responsemodel.ResponseUpdate": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "responsemodel.Role": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "createdAt": {
                    "$ref": "#/definitions/ptime.TimeResponse"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "permissions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/constant.Permission"
                    }
                },
                "updatedAt": {
                    "$ref": "#/definitions/ptime.TimeResponse"
                }
            }
        },
        "responsemodel.RoleDetail": {
            "type": "object",
            "properties": {
                "role": {
                    "$ref": "#/definitions/responsemodel.Role"
                }
            }
        },
        "responsemodel.User": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "avatar": {
                    "type": "string"
                },
                "banned": {
                    "type": "boolean"
                },
                "createdAt": {
                    "$ref": "#/definitions/ptime.TimeResponse"
                },
                "email": {
                    "description": "unique",
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "description": "unique",
                    "type": "string"
                },
                "statistic": {
                    "$ref": "#/definitions/model.UserStatistic"
                },
                "updatedAt": {
                    "$ref": "#/definitions/ptime.TimeResponse"
                }
            }
        },
        "responsemodel.UserDetail": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/responsemodel.User"
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
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "GoGo - Admin API",
	Description:      "API for admin management.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}