// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package swagger_docs

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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Daryl",
            "url": "https://www.daryl.top/",
            "email": "susecjh@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/config/property-system-class-enum": {
            "get": {
                "description": "查询系统属性类型",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "场景管理"
                ],
                "summary": "查询系统属性类型",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/config/property-type-enum": {
            "get": {
                "description": "查询属性分类",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "场景管理"
                ],
                "summary": "查询属性分类",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object",
                                            "additionalProperties": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/namespace": {
            "get": {
                "description": "查询命名空间",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "命名空间"
                ],
                "summary": "查询命名空间",
                "parameters": [
                    {
                        "type": "string",
                        "description": "描述(模糊查询)",
                        "name": "description",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "主键id",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "名称(模糊查询)",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页：查询第某页",
                        "name": "pageIndex",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页：每页数量",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序字段，默认是以主键倒序，例如：-name,这个例子是以name为倒序排列",
                        "name": "sortedField",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResultPaged"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/namespace.Namespace"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "创建命名空间",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "命名空间"
                ],
                "summary": "创建命名空间",
                "parameters": [
                    {
                        "description": "命名空间属性",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/namespace.CreateNamespaceReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/namespace.Namespace"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/namespace-set": {
            "get": {
                "description": "查询属性集",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "场景管理"
                ],
                "summary": "查询属性集",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据集CODE",
                        "name": "code",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "描述",
                        "name": "description",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "主键ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "命名空间ID",
                        "name": "namespaceID",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页：查询第某页",
                        "name": "pageIndex",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页：每页数量",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序字段，默认是以主键倒序，例如：-name,这个例子是以name为倒序排列",
                        "name": "sortedField",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResultPaged"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/property.PropertySet"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/namespace/{id}": {
            "delete": {
                "description": "删除命名空间",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "命名空间"
                ],
                "summary": "删除命名空间",
                "parameters": [
                    {
                        "minLength": 1,
                        "type": "string",
                        "description": "命名空间id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpx.JSONResult"
                        }
                    }
                }
            },
            "patch": {
                "description": "修改命名空间",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "命名空间"
                ],
                "summary": "修改命名空间",
                "parameters": [
                    {
                        "minLength": 1,
                        "type": "string",
                        "description": "命名空间id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "命名空间更新信息",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/namespace.UpdateOneReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/namespace.Namespace"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/property": {
            "get": {
                "description": "查询属性",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "场景管理"
                ],
                "summary": "查询属性",
                "parameters": [
                    {
                        "type": "string",
                        "description": "描述(模糊查询)",
                        "name": "description",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "主键id",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "名称(模糊查询)",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "命名空间ID",
                        "name": "namespaceID",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页：查询第某页",
                        "name": "pageIndex",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页：每页数量",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "属性集code",
                        "name": "propertySetCode",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序字段，默认是以主键倒序，例如：-name,这个例子是以name为倒序排列",
                        "name": "sortedField",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResultPaged"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/property.Property"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "创建属性",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "场景管理"
                ],
                "summary": "创建属性",
                "parameters": [
                    {
                        "description": "属性基本信息",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/property.CreatePropertyReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/property.Property"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/property-set": {
            "post": {
                "description": "创建属性集",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "场景管理"
                ],
                "summary": "创建属性集",
                "parameters": [
                    {
                        "description": "属性集基本信息",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/property.CreatePropertySetReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/property.PropertySet"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/property-set-map": {
            "get": {
                "description": "查询属性地图",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "场景管理"
                ],
                "summary": "查询属性地图",
                "parameters": [
                    {
                        "minLength": 1,
                        "type": "string",
                        "description": "属性集id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResultPaged"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/property.PropertyMap"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/property-set/{id}": {
            "delete": {
                "description": "删除属性集",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "场景管理"
                ],
                "summary": "删除属性集",
                "parameters": [
                    {
                        "minLength": 1,
                        "type": "string",
                        "description": "属性集id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpx.JSONResult"
                        }
                    }
                }
            },
            "patch": {
                "description": "修改属性集",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "场景管理"
                ],
                "summary": "修改属性集",
                "parameters": [
                    {
                        "minLength": 1,
                        "type": "string",
                        "description": "属性集id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "属性集更新信息",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/property.UpdatePropertySetReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/property.PropertySet"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/property/{id}": {
            "delete": {
                "description": "删除属性",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "场景管理"
                ],
                "summary": "删除属性",
                "parameters": [
                    {
                        "minLength": 1,
                        "type": "string",
                        "description": "属性id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpx.JSONResult"
                        }
                    }
                }
            },
            "patch": {
                "description": "修改属性",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "场景管理"
                ],
                "summary": "修改属性",
                "parameters": [
                    {
                        "minLength": 1,
                        "type": "string",
                        "description": "属性id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "属性更新信息",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/property.UpdatePropertyReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/property.Property"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httpx.JSONResult": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "回包code，表明是否正确，在code == 0时，表明服务正常",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "object"
                },
                "message": {
                    "description": "回报message，在code != 0时，展示给前端",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "httpx.JSONResultPaged": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "回包code，表明是否正确，在code == 0时，表明服务正常",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "object"
                },
                "message": {
                    "description": "回报message，在code != 0时，展示给前端",
                    "type": "string",
                    "example": "success"
                },
                "total": {
                    "description": "总数量",
                    "type": "integer"
                }
            }
        },
        "namespace.CreateNamespaceReq": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "rule-engine"
                },
                "name": {
                    "type": "string",
                    "example": "longquan-sword"
                }
            }
        },
        "namespace.Namespace": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "string"
                },
                "name": {
                    "description": "名称",
                    "type": "string"
                }
            }
        },
        "namespace.UpdateOneReq": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "rule-engine"
                },
                "name": {
                    "type": "string",
                    "example": "longquan-sword"
                }
            }
        },
        "property.CreatePropertyReq": {
            "type": "object",
            "properties": {
                "class": {
                    "description": "字段类型 int, string, float, ${property_set_code}",
                    "type": "string"
                },
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "name": {
                    "description": "字段名",
                    "type": "string"
                },
                "namespace_id": {
                    "description": "关联命名空间",
                    "type": "string"
                },
                "property_set_code": {
                    "description": "关联属性集",
                    "type": "string"
                },
                "type": {
                    "description": "字段分类 system or custome",
                    "type": "string"
                }
            }
        },
        "property.CreatePropertySetReq": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "属性集code",
                    "type": "string"
                },
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "name": {
                    "description": "字段名",
                    "type": "string"
                },
                "namespace_id": {
                    "description": "关联命名空间",
                    "type": "string"
                }
            }
        },
        "property.Property": {
            "type": "object",
            "properties": {
                "class": {
                    "description": "字段类型 int, string, float, ${property_set_code}",
                    "type": "string"
                },
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "string"
                },
                "name": {
                    "description": "字段名",
                    "type": "string"
                },
                "namespace_id": {
                    "description": "关联命名空间",
                    "type": "string"
                },
                "property_set_code": {
                    "description": "关联属性集",
                    "type": "string"
                },
                "type": {
                    "description": "字段分类 system or custome",
                    "type": "string"
                }
            }
        },
        "property.PropertyMap": {
            "type": "object",
            "properties": {
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/property.PropertyMap"
                    }
                },
                "class": {
                    "description": "字段类型 int, string, float, ${property_set_code}",
                    "type": "string"
                },
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "string"
                },
                "name": {
                    "description": "字段名",
                    "type": "string"
                },
                "namespace_id": {
                    "description": "关联命名空间",
                    "type": "string"
                },
                "property_set_code": {
                    "description": "关联属性集",
                    "type": "string"
                },
                "type": {
                    "description": "字段分类 system or custome",
                    "type": "string"
                }
            }
        },
        "property.PropertySet": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "数据集CODE",
                    "type": "string"
                },
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "string"
                },
                "namespace_id": {
                    "description": "命名空间ID",
                    "type": "string"
                }
            }
        },
        "property.UpdatePropertyReq": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "描述",
                    "type": "string"
                }
            }
        },
        "property.UpdatePropertySetReq": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "描述",
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
	BasePath:    "/engine/api",
	Schemes:     []string{},
	Title:       "龙泉决策引擎",
	Description: "龙泉决策引擎",
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
