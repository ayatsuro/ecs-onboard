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
        "/namespace/onboard": {
            "post": {
                "description": "Creates in Dell ECS: a namespace, a IAM user RW and a IAM user RO and store their secret access keys in Vault",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Namespace"
                ],
                "summary": "onboard a namespace",
                "parameters": [
                    {
                        "description": "the namespace to onboard",
                        "name": "ns",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Namespace"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "model.Namespace": {
            "type": "object",
            "required": [
                "namespace",
                "username"
            ],
            "properties": {
                "namespace": {
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