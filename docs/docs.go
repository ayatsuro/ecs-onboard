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
        "/bird/onboard": {
            "post": {
                "description": "In Dell ECS, creates a IAM user and an AccessKey. In Vault, stores the secret access keys and creates a JWT auth role bound to the Brid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Brid"
                ],
                "summary": "onboard a brid to a namespace as IAM user",
                "parameters": [
                    {
                        "description": "the user to onboard",
                        "name": "brid",
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
        "/iamuser/onboard": {
            "post": {
                "description": "In Dell ECS, creates a IAM user and an AccessKey. In Vault, stores the secret access keys and creates a JWT auth role bound to the Brid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IamUser"
                ],
                "summary": "onboard a IAM user in a namespace",
                "parameters": [
                    {
                        "description": "the user to onboard",
                        "name": "brid",
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
        "/iamuser/{username}": {
            "delete": {
                "description": "In Dell ECS, deletes the IAM user. In Vault, deletes the role and the JWT auth role (if any)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IamUser"
                ],
                "summary": "delete a IAM user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the user to delete",
                        "name": "brid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/namespace/migrate": {
            "post": {
                "description": "In Dell ECS, creates a IAM user (and AccessKey) for the Native users, creates a second AccessKey for existing IAM users. In Vault, stores the Secret Access Keys",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Namespace"
                ],
                "summary": "migrates a namespace",
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
                "description": "In Dell ECS, creates a namespace, a IAM user and an AccessKey. In Vault, stores the Secret Access Key",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Namespace"
                ],
                "summary": "onboards a namespace",
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
