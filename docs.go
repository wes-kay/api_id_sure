// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "io.labs.development@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/signin": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Sign Into Account",
                "parameters": [
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "name": "email_validated",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "name": "temp",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SignInResponse"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Creates an account on the service",
                "parameters": [
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "passwordconfirm",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "name": "temp",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/v1/account": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Get's account info using the session token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Auth",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CourseCertificate"
                        }
                    }
                }
            }
        },
        "/v1/account/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Gets an account by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Auth",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "account id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Account"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Updates an account by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Auth",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "account id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "name": "email_validated",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "image",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "role",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "status",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "name": "temp",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Deletes an account by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Auth",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "account id",
                        "name": "id",
                        "in": "query",
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
        "/v1/admin/accounts": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Gets all accounts for Admin",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Auth",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Account"
                            }
                        }
                    }
                }
            }
        },
        "/v1/certificate": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "certificate"
                ],
                "summary": "Creates a certificate by session ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Auth",
                        "in": "header",
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
        "/v1/certificate/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "certificate"
                ],
                "summary": "Gets certificate by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Auth",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Certificate ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Certificate"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "certificate"
                ],
                "summary": "Updates a certificate by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Auth",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Certificate ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "accessed",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "accountID",
                        "in": "formData"
                    },
                    {
                        "type": "boolean",
                        "name": "activated",
                        "in": "formData"
                    },
                    {
                        "type": "boolean",
                        "name": "completed",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "courseID",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "created",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "dateOfBirth",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "experience",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "issued",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "kin",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "licence",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "nationality",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "passport",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "rank",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "shipExpereince",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "type",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "uid",
                        "in": "formData"
                    },
                    {
                        "type": "boolean",
                        "name": "verified",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "certificate"
                ],
                "summary": "Deletes a certificate by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Auth",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Certificate ID",
                        "name": "id",
                        "in": "query",
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
        "/v1/certificates/course/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "certificate"
                ],
                "summary": "Gets all certificate by course ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Auth",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Certificate"
                            }
                        }
                    }
                }
            }
        },
        "/v1/certificates/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "certificate"
                ],
                "summary": "Gets all certificate by account ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Auth",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Certificate"
                            }
                        }
                    }
                }
            }
        },
        "/v1/course/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "Gets a course by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Auth",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Course ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Course"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "Update a course by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Auth",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Course ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "additional_description",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "certificate_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "country",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "course_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "course_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "created",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "expire",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "image",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "Creates a new course using session ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Auth",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "Delete a course by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Auth",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Course ID",
                        "name": "id",
                        "in": "query",
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
        "/v1/courses/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "Returns all courses under account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "Auth",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Course ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Course"
                            }
                        }
                    }
                }
            }
        },
        "/v1/signup/welcome/{uid}": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "verify"
                ],
                "summary": "Verifies an email with a UID sent to the account email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uid from email",
                        "name": "uid",
                        "in": "query",
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
        "/v1/view/certificate/{uid}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "certificate"
                ],
                "summary": "Loads certificate from uid",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uid from url",
                        "name": "uid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CourseCertificate"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Account": {
            "type": "object",
            "required": [
                "email",
                "email_validated",
                "id",
                "name",
                "phone",
                "role",
                "status"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "email_validated": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "integer"
                },
                "role": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "temp": {
                    "type": "boolean"
                }
            }
        },
        "model.Certificate": {
            "type": "object",
            "properties": {
                "accessed": {
                    "type": "integer"
                },
                "accountID": {
                    "type": "integer"
                },
                "activated": {
                    "type": "boolean"
                },
                "completed": {
                    "type": "boolean"
                },
                "courseID": {
                    "type": "integer"
                },
                "created": {
                    "type": "string"
                },
                "dateOfBirth": {
                    "type": "string"
                },
                "experience": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "issued": {
                    "type": "string"
                },
                "kin": {
                    "type": "string"
                },
                "licence": {
                    "type": "string"
                },
                "nationality": {
                    "type": "string"
                },
                "passport": {
                    "type": "string"
                },
                "rank": {
                    "type": "string"
                },
                "shipExpereince": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                },
                "uid": {
                    "type": "string"
                },
                "verified": {
                    "type": "boolean"
                }
            }
        },
        "model.Course": {
            "type": "object",
            "required": [
                "address",
                "certificate_name",
                "country",
                "course_id",
                "course_name",
                "description",
                "expire"
            ],
            "properties": {
                "additional_description": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "certificate_name": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "course_id": {
                    "type": "integer"
                },
                "course_name": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "expire": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                }
            }
        },
        "model.CourseCertificate": {
            "type": "object",
            "properties": {
                "account": {
                    "$ref": "#/definitions/model.Account"
                },
                "certificate": {
                    "$ref": "#/definitions/model.Certificate"
                },
                "course": {
                    "$ref": "#/definitions/model.Course"
                }
            }
        },
        "model.SignInResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
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
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "Digital ID",
	Description: "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
