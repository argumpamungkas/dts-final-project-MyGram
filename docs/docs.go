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
        "/comment/delete/{commentID}": {
            "delete": {
                "description": "Delete data comment by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delete Comment"
                ],
                "summary": "Delete data comment with commentID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "commentID of the data comment to be deleted",
                        "name": "commentID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/comment/getAll": {
            "get": {
                "description": "Get details of all comment or add query parameter photo_id for all comment from photo_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get All Comment"
                ],
                "summary": "Get details of All comment",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/comment/getOne/{commentID}": {
            "get": {
                "description": "Get details of comment corresponding to the input commentID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Comment by ID"
                ],
                "summary": "Get details for a given commentID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the photo",
                        "name": "commentID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/comment/post": {
            "post": {
                "description": "Post a new Comment, NOTE : id auto increment, so in body id is deleted. and add query parameter photo_id for comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create Comment"
                ],
                "summary": "Post Comment",
                "parameters": [
                    {
                        "description": "create comment",
                        "name": "models.Comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/comment/update/{commentID}": {
            "put": {
                "description": "Update data comment by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Update Comment"
                ],
                "summary": "Updated data comment with commentID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "commentID of the data comment to be updated",
                        "name": "commentID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "updated comment",
                        "name": "models.Comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/photo/delete/{photoID}": {
            "delete": {
                "description": "Delete data photo by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delete Photo"
                ],
                "summary": "Delete data photo with photoID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "photoID of the data photo to be deleted",
                        "name": "photoID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            }
        },
        "/photo/getAll": {
            "get": {
                "description": "Get details of all photo or add query parameter user_id for all photo from user_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get All Photo"
                ],
                "summary": "Get details of All photo",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            }
        },
        "/photo/getOne/{photoID}": {
            "get": {
                "description": "Get details of photo corresponding to the input photoID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Photo by ID"
                ],
                "summary": "Get details for a given photoID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the photo",
                        "name": "photoID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            }
        },
        "/photo/post": {
            "post": {
                "description": "Post a new Photo, NOTE : id auto increment, so in body id is deleted",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create Photo"
                ],
                "summary": "Post Photo",
                "parameters": [
                    {
                        "description": "create photo",
                        "name": "models.Photo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            }
        },
        "/photo/update/{photoID}": {
            "put": {
                "description": "Update data social media by id, NOTE: photo is not updated, just title and caption can be updated, so in the body photo_url doesn't use",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Update Photo"
                ],
                "summary": "Updated data photo with socialMediaID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "photoID of the data photo to be updated",
                        "name": "photoID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "updated photo",
                        "name": "models.Photo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            }
        },
        "/social-media/create": {
            "post": {
                "description": "Post create a new social media, NOTE : id auto increment, so in body id is deleted",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create Social Media"
                ],
                "summary": "Post Social media",
                "parameters": [
                    {
                        "description": "create social media",
                        "name": "models.SocialMedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            }
        },
        "/social-media/delete/{socialMediaID}": {
            "delete": {
                "description": "Delete data social media by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delete Social Media"
                ],
                "summary": "Delete data social media with socialMediaID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "socialMediaID of the data social media to be deleted",
                        "name": "socialMediaID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            }
        },
        "/social-media/getAll": {
            "get": {
                "description": "Get details of all social media or add query parameter user_id for all social media from it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get All Social Media"
                ],
                "summary": "Get details of All social media",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            }
        },
        "/social-media/getOne/{socialMediaID}": {
            "get": {
                "description": "Get details of social media corresponding to the input socialMediaID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Social Media by ID"
                ],
                "summary": "Get details for a given socialMediaID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the social media",
                        "name": "socialMediaID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            }
        },
        "/social-media/update/{socialMediaID}": {
            "put": {
                "description": "Update data social media by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Update Social Media"
                ],
                "summary": "Updated data social media with socialMediaID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "socialMediaID of the data social media to be updated",
                        "name": "socialMediaID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "updated social media",
                        "name": "models.SocialMedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login user needed for crud of the photo, social media and comment because if you login you have token for that",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login user"
                ],
                "summary": "Login User",
                "parameters": [
                    {
                        "description": "login user",
                        "name": "models.User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Comment": {
            "type": "object",
            "properties": {
                "comment_message": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "photo_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Photo": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "comment_message": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Comment"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "photo_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.SocialMedia": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Comment"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Photo"
                    }
                },
                "socials_media": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SocialMedia"
                    }
                },
                "updated_at": {
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
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "My Gram API",
	Description:      "This is a final project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
