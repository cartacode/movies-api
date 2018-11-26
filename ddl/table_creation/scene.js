db.createCollection( "scene",{
    "capped": false,
    "validator": {
        "$jsonSchema": {
            "bsonType": "object",
            "properties": {
                "_id": {
                    "bsonType": "objectId"
                },
                "images": {
                    "bsonType": "object",
                    "properties": {
                        "landscape": {
                            "bsonType": "string",
                            "pattern": "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"
                        },
                        "portrait": {
                            "bsonType": "string",
                            "pattern": "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"
                        },
                        "banner": {
                            "bsonType": "string",
                            "pattern": "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"
                        },
                        "detailpage": {
                            "bsonType": "string",
                            "pattern": "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"
                        }
                    },
                    "additionalProperties": false
                },
                "thumbnails": {
                    "bsonType": "object",
                    "properties": {
                        "prefix": {
                            "bsonType": "string"
                        },
                        "count": {
                            "bsonType": "int"
                        },
                        "format": {
                            "bsonType": "string"
                        }
                    },
                    "additionalProperties": false
                },
                "trailer": {
                    "bsonType": "object",
                    "properties": {
                        "dynamodbid": {
                            "bsonType": "string",
                            "pattern": "[0-9a-fA-F]{8}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{12}"
                        },
                        "title": {
                            "bsonType": "string"
                        }
                    },
                    "additionalProperties": false
                },
                "information": {
                    "bsonType": "object",
                    "properties": {
                        "studio": {
                            "bsonType": "objectId"
                        },
                        "length": {
                            "bsonType": "int"
                        },
                        "director": {
                            "bsonType": "array",
                            "additionalItems": true,
                            "uniqueItems": false,
                            "items": {
                                "bsonType": "objectId"
                            }
                        },
                        "quality": {
                            "bsonType": "array",
                            "additionalItems": true,
                            "uniqueItems": false,
                            "items": {
                                "bsonType": "int"
                            }
                        },
                        "stars": {
                            "bsonType": "array",
                            "additionalItems": true,
                            "uniqueItems": false,
                            "items": {
                                "bsonType": "objectId"
                            }
                        }
                    },
                    "additionalProperties": false
                },
                "performance": {
                    "bsonType": "object",
                    "properties": {
                        "rank": {
                            "bsonType": "int",
                            "minimum": 0
                        },
                        "upvotes": {
                            "bsonType": "int",
                            "minimum": 0
                        },
                        "downvotes": {
                            "bsonType": "int",
                            "minimum": 0
                        },
                        "views": {
                            "bsonType": "number"
                        }
                    },
                    "additionalProperties": false
                },
                "title": {
                    "bsonType": "string"
                },
                "dynamodbid": {
                    "bsonType": "string",
                    "pattern": "[0-9a-fA-F]{8}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{12}"
                },
                "description": {
                    "bsonType": "string"
                },
                "volume": {
                    "bsonType": "objectId"
                },
                "price": {
                    "bsonType": "double"
                },
                "slug": {
                    "bsonType": "string",
                    "pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
                },
                "series": {
                    "bsonType": "string"
                },
                "ispublished": {
                    "bsonType": "bool"
                },
                "_created": {
                    "bsonType": "date"
                },
                "_modified": {
                    "bsonType": "date"
                },
                "chapters": {
                    "bsonType": "array",
                    "additionalItems": true,
                    "uniqueItems": false,
                    "items": {
                        "bsonType": "object",
                        "properties": {
                            "name": {
                                "bsonType": "string"
                            },
                            "timestamp": {
                                "bsonType": "int"
                            }
                        },
                        "additionalProperties": false
                    }
                },
                "extras": {
                    "bsonType": "array",
                    "additionalItems": true,
                    "uniqueItems": false,
                    "items": {
                        "bsonType": "object",
                        "properties": {
                            "published": {
                                "bsonType": "bool"
                            },
                            "url": {
                                "bsonType": "string",
                                "pattern": "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"
                            }
                        },
                        "additionalProperties": false,
                        "required": [
                            "published",
                            "url"
                        ]
                    }
                },
                "tags": {
                    "bsonType": "array",
                    "additionalItems": true,
                    "uniqueItems": false,
                    "items": {
                        "bsonType": "string"
                    }
                }
            },
            "required": [
                "_id",
                "images",
                "thumbnails",
                "trailer",
                "information",
                "performance",
                "title",
                "volume",
                "price",
                "series",
                "ispublished",
                "_created",
                "_modified",
                "chapters",
                "extras",
                "tags"
            ]
        }
    },
    "validationLevel": "off",
    "validationAction": "warn"
});
db.scene.createIndex(
{
    "_id": 1
},
{
    "name": "_id_"
}
);