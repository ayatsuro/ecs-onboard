// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/namespace/migrate": {
            "post": {
                "description": "If the safe_id is omitted in the payload, it will be derived from the namespace name (first part of a split on '-')",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Namespace"
                ],
                "summary": "Fetches all the users of an ECS namespace. For native users, creates a IAM user, an access key, stores it in a Vault role, creates a Vault policy. For IAM users, creates an access key, stores it in a Vault role, creates a Vault policy.",
                "parameters": [
                    {
                        "description": "the namespace to migrate",
                        "name": "ns",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.MigrateNamespace"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/namespace/onboard": {
            "post": {
                "description": "If the safe_id is omitted in the payload, it will be derived from the namespace (first part of a split on '-')",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Namespace"
                ],
                "summary": "Creates an IAM user, an access key, stores it in a Vault role, creates a Vault policy",
                "parameters": [
                    {
                        "description": "the namespace to onboard",
                        "name": "ns",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.OnboardNamespace"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user": {
            "post": {
                "description": "If the safe_id is omitted in the payload, it will be derived from the namespace name (first part of a split on '-')",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Creates a IAM user in an ECS namespace, creates an access key, stores it in a Vault role, creates a Vault policy. And if the username is a BRID, creates a JWT authrole.",
                "parameters": [
                    {
                        "description": "the user to onboard",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.IamUser"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/{username}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Deletes the IAM user, the Vault role and policy, and the JWT authrole if any",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the IAM user to delete",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "model.IamUser": {
            "type": "object",
            "required": [
                "namespace",
                "username"
            ],
            "properties": {
                "namespace": {
                    "type": "string"
                },
                "safe_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.MigrateNamespace": {
            "type": "object",
            "required": [
                "namespace"
            ],
            "properties": {
                "namespace": {
                    "type": "string"
                },
                "safe_id": {
                    "type": "string"
                }
            }
        },
        "model.OnboardNamespace": {
            "type": "object",
            "required": [
                "namespace",
                "username"
            ],
            "properties": {
                "namespace": {
                    "type": "string"
                },
                "safe_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
