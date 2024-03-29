// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/transactions": {
            "get": {
                "description": "Returns top level transactions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Top level transactions",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Items per page limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.TransactionsResponse"
                        }
                    }
                }
            }
        },
        "/transactions/{id}/sublist": {
            "get": {
                "description": "Returns related transactions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions_sub_list"
                ],
                "summary": "Top level transactions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Transaction id",
                        "name": "id",
                        "in": "path"
                    },
                    {
                        "type": "integer",
                        "description": "Items per page limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.TransactionsResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.Transaction": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "category": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "source": {
                    "type": "string"
                }
            }
        },
        "types.TransactionsResponse": {
            "type": "object",
            "properties": {
                "total": {
                    "type": "integer"
                },
                "transactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Transaction"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
