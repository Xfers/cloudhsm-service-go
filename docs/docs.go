// Package docs GENERATED BY SWAG; DO NOT EDIT
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
        "/digest": {
            "post": {
                "description": "digest the data currently using sha256",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "digest"
                ],
                "summary": "Digest the data.",
                "parameters": [
                    {
                        "description": "Data to be digested",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.DigestRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.DigestResponse"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "health check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.HealthResponse"
                        }
                    }
                }
            }
        },
        "/pure-sign": {
            "post": {
                "description": "sign the data using openssl or cloudhsm",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pure-sign"
                ],
                "summary": "Sign the data.",
                "parameters": [
                    {
                        "description": "Data to be signed",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.PureSignRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.PureSignResponse"
                        }
                    }
                }
            }
        },
        "/sign": {
            "post": {
                "description": "sign the digest using openssl or cloudhsm",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sign"
                ],
                "summary": "Sign the digest.",
                "parameters": [
                    {
                        "description": "Digest to be signed",
                        "name": "digest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.SignRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SignResponse"
                        }
                    }
                }
            }
        },
        "/verify": {
            "post": {
                "description": "verify the data using provided signature and public key, using openssl or cloudhsm",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "verify"
                ],
                "summary": "Verify the Data.",
                "parameters": [
                    {
                        "description": "Data to be verified",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.VerifyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.VerifyResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.DigestRequest": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                }
            }
        },
        "controllers.DigestResponse": {
            "type": "object",
            "properties": {
                "digest": {
                    "type": "string"
                }
            }
        },
        "controllers.HealthResponse": {
            "description": "health check response",
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "controllers.PureSignRequest": {
            "type": "object",
            "properties": {
                "digest": {
                    "type": "string"
                }
            }
        },
        "controllers.PureSignResponse": {
            "type": "object",
            "properties": {
                "signature": {
                    "type": "string"
                }
            }
        },
        "controllers.SignRequest": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                }
            }
        },
        "controllers.SignResponse": {
            "type": "object",
            "properties": {
                "signature": {
                    "type": "string"
                }
            }
        },
        "controllers.VerifyRequest": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "signature": {
                    "type": "string"
                }
            }
        },
        "controllers.VerifyResponse": {
            "type": "object",
            "properties": {
                "valid": {
                    "type": "boolean"
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
