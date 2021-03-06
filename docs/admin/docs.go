// Package admin GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package admin

import "github.com/swaggo/swag"

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
                "security": [
                    {
                        "ApikeyAuth": []
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
                "summary": "ALl",
                "operationId": "role-all",
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
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "DeleteOne",
                "operationId": "role-delete-one",
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
                        "description": ""
                    }
                }
            }
        },
        "/staff": {
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
                    "Staff"
                ],
                "summary": "All",
                "operationId": "staff-all",
                "parameters": [
                    {
                        "type": "string",
                        "name": "active",
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
            }
        },
        "/staff/{id}/change-status": {
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
                    "Staff"
                ],
                "summary": "ChangeStatus",
                "operationId": "change-status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "staff id",
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
        },
        "/staffs": {
            "post": {
                "security": [
                    {
                        "ApikeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Staff"
                ],
                "summary": "Create",
                "operationId": "staff-create",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apimodel.StaffCreate"
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
        "/staffs/login-with-email": {
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
                    "Staff"
                ],
                "summary": "LoginWithEmail",
                "operationId": "login-with-email",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apimodel.StaffLoginWithEmail"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.StaffLogin"
                        }
                    }
                }
            }
        },
        "/staffs/me": {
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
                    "Staff"
                ],
                "summary": "GetMe",
                "operationId": "get-me",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.Staff"
                        }
                    }
                }
            }
        },
        "/staffs/{id}": {
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
                    "Staff"
                ],
                "summary": "Update",
                "operationId": "update-staff",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apimodel.StaffUpdate"
                        }
                    },
                    {
                        "type": "string",
                        "description": "staff id",
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
        "apimodel.StaffCreate": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "apimodel.StaffLoginWithEmail": {
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
        "apimodel.StaffUpdate": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
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
        "responsemodel.RoleShort": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "isAdmin": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "responsemodel.Staff": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "active": {
                    "type": "boolean"
                },
                "avatar": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "isRoot": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/responsemodel.RoleShort"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "responsemodel.StaffLogin": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "token": {
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
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Admin API",
	Description:      "API for admin management.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
