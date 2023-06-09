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
        "/images": {
            "get": {
                "description": "images list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Images"
                ],
                "summary": "HandleImagesList",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Image"
                            }
                        }
                    }
                }
            }
        },
        "/images/download": {
            "get": {
                "description": "download images from links file",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Images"
                ],
                "summary": "HandleDownloadImages",
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/images/upload": {
            "post": {
                "description": "upload images",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Images"
                ],
                "summary": "HandleImagesUpload",
                "parameters": [
                    {
                        "type": "file",
                        "description": "images",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/images/{id}": {
            "get": {
                "description": "downoald an image from list",
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "Images"
                ],
                "summary": "HandleDownloadImage",
                "parameters": [
                    {
                        "type": "string",
                        "description": "image ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.Image": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "download_date": {
                    "type": "string"
                },
                "file_extension": {
                    "type": "string"
                },
                "file_size": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "local_url": {
                    "type": "string"
                },
                "original_name": {
                    "type": "string"
                },
                "original_url": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "updated_at": {
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
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
