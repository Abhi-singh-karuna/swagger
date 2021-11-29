package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Abhishek Singh",
            "email": "Abhishek@gmail.com"
        },
        
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/product": {
            "get": {
                "description": "Get all products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get all products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handlers.ResponseHTTP"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.product"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/handlers.ResponseHTTP"
                        }
                    }
                }
            },
            "post": {
                "description": "Register product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Register a new product",
                "parameters": [
                    {
                        "description": "Register product",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handlers.ResponseHTTP"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.product"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ResponseHTTP"
                        }
                    }
                }
            }
        },
        "/product/{id}": {
            "get": {
                "description": "Get product by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get product by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handlers.ResponseHTTP"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.product"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/handlers.ResponseHTTP"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove product by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Remove product by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.ResponseHTTP"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.ResponseHTTP"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/handlers.ResponseHTTP"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.ResponseHTTP": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.product": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string",
                    "example": "Golang"
                },
                "description": {
                    "type": "string",
                    "example": "Go is an open source programming language .."
                },
                "amount": {
                    "type": "integer",
                    "example": "999"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Product App [ CRUD ]",
	Description: "This is testing api for Product App ",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
