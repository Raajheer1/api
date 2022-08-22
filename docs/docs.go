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
        "contact": {
            "name": "Daniel Hawton",
            "email": "wm@denartcc.org"
        },
        "license": {
            "name": "Apache",
            "url": "https://github.com/kzdv/api2/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/live/:facility": {
            "get": {
                "tags": [
                    "overflight"
                ],
                "summary": "Get Overflights for Facility [Legacy/Deprecated]",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Facility, defaults to ZDV if no facility id provided",
                        "name": "fac",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/internal_v1_overflight.Flightsv1"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    }
                }
            }
        },
        "/v1/overflight": {
            "get": {
                "tags": [
                    "overflight"
                ],
                "summary": "Get Overflights for Facility",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Facility, defaults to ZDV if no facility id provided",
                        "name": "fac",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/internal_v1_overflight.Flightsv1"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    }
                }
            }
        },
        "/v1/overflight/:facility": {
            "get": {
                "tags": [
                    "overflight"
                ],
                "summary": "Get Overflights for Facility",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Facility, defaults to ZDV if no facility id provided",
                        "name": "fac",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/internal_v1_overflight.Flightsv1"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    }
                }
            }
        },
        "/v1/storage": {
            "post": {
                "tags": [
                    "storage"
                ],
                "summary": "Create storage object",
                "parameters": [
                    {
                        "description": "Storage Object",
                        "name": "storage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.StorageRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.Document"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    }
                }
            }
        },
        "/v1/storage/:category": {
            "get": {
                "tags": [
                    "storage"
                ],
                "summary": "Get Storage Listing",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category, if applicable",
                        "name": "category",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.Document"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    }
                }
            }
        },
        "/v1/storage/:id": {
            "put": {
                "tags": [
                    "storage"
                ],
                "summary": "Update storage object",
                "parameters": [
                    {
                        "description": "Storage Object",
                        "name": "storage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.StorageRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Storage ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Document"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "storage"
                ],
                "summary": "Delete storage object",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Storage ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    }
                }
            }
        },
        "/v1/storage/:id/file": {
            "put": {
                "tags": [
                    "storage"
                ],
                "summary": "Upload file data to storage object",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Storage ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "File",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    }
                }
            }
        },
        "/v1/user/:cid": {
            "get": {
                "tags": [
                    "user"
                ],
                "summary": "Get User Information, if cid is not provided, defaults to logged in user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CID, defaults to logged in user",
                        "name": "cid",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Patch User Information",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponse"
                        }
                    },
                    {
                        "type": "string",
                        "description": "CID",
                        "name": "cid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    }
                }
            }
        },
        "/v1/user/:cid/roles": {
            "get": {
                "tags": [
                    "user"
                ],
                "summary": "Get User Information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CID, if not provided, defaults to logged in user",
                        "name": "cid",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    }
                }
            }
        },
        "/v1/user/:cid/roles/:role": {
            "put": {
                "tags": [
                    "user"
                ],
                "summary": "Add User Role",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Role",
                        "name": "role",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "CID",
                        "name": "cid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "309": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "user"
                ],
                "summary": "Remove User Role",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Role",
                        "name": "role",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "CID",
                        "name": "cid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    }
                }
            }
        },
        "/v1/user/login": {
            "get": {
                "tags": [
                    "user",
                    "oauth"
                ],
                "summary": "Login to account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Redirect URL",
                        "name": "redirect",
                        "in": "path"
                    }
                ],
                "responses": {
                    "307": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    }
                }
            }
        },
        "/v1/user/login/callback": {
            "get": {
                "tags": [
                    "user",
                    "oauth"
                ],
                "summary": "Login callback",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "307": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.R"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.StorageRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.UserResponse": {
            "type": "object",
            "properties": {
                "certifications": {
                    "$ref": "#/definitions/dto.UserResponseCertifications"
                },
                "cid": {
                    "type": "integer"
                },
                "controller_type": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "discord_id": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "operating_initials": {
                    "type": "string"
                },
                "rating": {
                    "type": "string"
                },
                "removal_reason": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "status": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "dto.UserResponseCertifications": {
            "type": "object",
            "properties": {
                "approach": {
                    "type": "string"
                },
                "delivery": {
                    "type": "string"
                },
                "enroute": {
                    "type": "string"
                },
                "ground": {
                    "type": "string"
                },
                "local": {
                    "type": "string"
                }
            }
        },
        "github.com_kzdv_api_internal_v1_overflight.Flightsv1": {
            "type": "object",
            "properties": {
                "alt": {
                    "type": "integer",
                    "example": 10500
                },
                "arr": {
                    "type": "string",
                    "example": "KLMO"
                },
                "callsign": {
                    "type": "string",
                    "example": "N462AW"
                },
                "cid": {
                    "type": "integer",
                    "example": 876594
                },
                "dep": {
                    "type": "string",
                    "example": "KLMO"
                },
                "facility": {
                    "type": "string",
                    "example": "ZDV"
                },
                "hdg": {
                    "type": "integer",
                    "example": 180
                },
                "lastSeen": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                },
                "lat": {
                    "type": "number",
                    "example": -33.867
                },
                "lon": {
                    "type": "number",
                    "example": 151.206
                },
                "route": {
                    "type": "string",
                    "example": "DCT"
                },
                "spd": {
                    "type": "integer",
                    "example": 150
                },
                "type": {
                    "type": "string",
                    "example": "C208"
                }
            }
        },
        "internal_v1_overflight.Flightsv1": {
            "type": "object",
            "properties": {
                "alt": {
                    "type": "integer",
                    "example": 10500
                },
                "arr": {
                    "type": "string",
                    "example": "KLMO"
                },
                "callsign": {
                    "type": "string",
                    "example": "N462AW"
                },
                "cid": {
                    "type": "integer",
                    "example": 876594
                },
                "dep": {
                    "type": "string",
                    "example": "KLMO"
                },
                "facility": {
                    "type": "string",
                    "example": "ZDV"
                },
                "hdg": {
                    "type": "integer",
                    "example": 180
                },
                "lastSeen": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                },
                "lat": {
                    "type": "number",
                    "example": -33.867
                },
                "lon": {
                    "type": "number",
                    "example": 151.206
                },
                "route": {
                    "type": "string",
                    "example": "DCT"
                },
                "spd": {
                    "type": "integer",
                    "example": 150
                },
                "type": {
                    "type": "string",
                    "example": "C208"
                }
            }
        },
        "response.R": {
            "type": "object",
            "properties": {
                "data": {},
                "status": {
                    "type": "string"
                }
            }
        },
        "types.Document": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string",
                    "example": "sops"
                },
                "created_at": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                },
                "created_by": {
                    "$ref": "#/definitions/types.User"
                },
                "created_by_id": {
                    "type": "integer",
                    "example": 1
                },
                "description": {
                    "type": "string",
                    "example": "Description of document"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "document name"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                },
                "updated_by": {
                    "$ref": "#/definitions/types.User"
                },
                "updated_by_id": {
                    "type": "integer",
                    "example": 1
                },
                "url": {
                    "type": "string",
                    "example": "https://www.example.com/document.pdf"
                }
            }
        },
        "types.Rating": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "long": {
                    "type": "string",
                    "example": "Observer"
                },
                "short": {
                    "type": "string",
                    "example": "OBS"
                }
            }
        },
        "types.Role": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "wm"
                },
                "updated_at": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.User"
                    }
                }
            }
        },
        "types.User": {
            "type": "object",
            "properties": {
                "appCertification": {
                    "description": "Must be one of : none, training, solo, certified, major, cantrain",
                    "type": "string",
                    "example": "certified"
                },
                "cid": {
                    "type": "integer",
                    "example": 876594
                },
                "controllerType": {
                    "description": "Must be one of: none, active, inactive, loa",
                    "type": "string",
                    "example": "home"
                },
                "created_at": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                },
                "ctrCertification": {
                    "description": "Must be one of : none, training, solo, certified, major, cantrain",
                    "type": "string",
                    "example": "none"
                },
                "delCertification": {
                    "description": "Must be one of : none, training, solo, certified, major, cantrain",
                    "type": "string",
                    "example": "certified"
                },
                "discord_id": {
                    "type": "string",
                    "example": "123456789012345678"
                },
                "division": {
                    "type": "string",
                    "example": "USA"
                },
                "email": {
                    "type": "string",
                    "example": "wm@denartcc.org"
                },
                "firstname": {
                    "type": "string",
                    "example": "Daniel"
                },
                "gndCertification": {
                    "description": "Must be one of : none, training, solo, certified, major, cantrain",
                    "type": "string",
                    "example": "certified"
                },
                "lastname": {
                    "type": "string",
                    "example": "Hawton"
                },
                "lclCertification": {
                    "description": "Must be one of : none, training, solo, certified, major, cantrain",
                    "type": "string",
                    "example": "certified"
                },
                "oi": {
                    "type": "string",
                    "example": "DH"
                },
                "rating": {
                    "$ref": "#/definitions/types.Rating"
                },
                "region": {
                    "type": "string",
                    "example": "AMAS"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Role"
                    }
                },
                "status": {
                    "description": "Must be one of: none, active, inactive, loa",
                    "type": "string",
                    "example": "active"
                },
                "subdivision": {
                    "description": "This may be blank",
                    "type": "string",
                    "example": "ZDV"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                },
                "updateid": {
                    "description": "Internally used identifier during scheduled updates for removals",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "network.denartcc.org",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "KZDV API",
	Description:      "Session Cookie",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
