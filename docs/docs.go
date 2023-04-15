// Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/videos/download/{name}": {
            "get": {
                "description": "下载视频",
                "produces": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Videos"
                ],
                "summary": "下载视频",
                "parameters": [
                    {
                        "type": "string",
                        "description": "视频文件名",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "{\"code\":\"404\",\"msg\":\"视频文件不存在\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/videos/upload": {
            "post": {
                "description": "上传视频",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Videos"
                ],
                "summary": "上传视频",
                "parameters": [
                    {
                        "type": "file",
                        "description": "上传文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "00:00:00",
                        "name": "start_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "00:00:00",
                        "name": "end_time",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"上传成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"code\":\"400\",\"msg\":\"上传失败\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/videos-list": {
            "get": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "视频输入",
                "parameters": [
                    {
                        "type": "file",
                        "description": "上传文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"上传成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"code\":\"400\",\"msg\":\"上传失败\"}",
                        "schema": {
                            "type": "string"
                        }
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
