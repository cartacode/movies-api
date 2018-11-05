db.createCollection( "scene",{
    "capped": false,
    "validator": {
        "$jsonSchema": {
            "bsonType": "object",
            "properties": {
                "_id": {
                    "bsonType": "objectId"
                },
                "title": {
                    "bsonType": "string"
                },
                "slug": {
                    "bsonType": "string"
                },
                "coverimage": {
                    "bsonType": "string"
                },
                "images": {
                    "bsonType": "object",
                    "properties": {
                        "detailpage": {
                            "bsonType": "string"
                        },
                        "traythumbnail": {
                            "bsonType": "string"
                        },
                        "trayfeaturedthumbnail": {
                            "bsonType": "string"
                        },
                        "mobilethumbnail": {
                            "bsonType": "string"
                        },
                        "coverimages": {
                            "bsonType": "array",
                            "additionalItems": true,
                            "uniqueItems": false,
                            "items": {
                                "bsonType": "string"
                            }
                        }
                    },
                    "additionalProperties": false
                },
                "performance": {
                    "bsonType": "object",
                    "properties": {
                        "rank": {
                            "bsonType": "number"
                        },
                        "upvotes": {
                            "bsonType": "number"
                        },
                        "downvotes": {
                            "bsonType": "number"
                        },
                        "views": {
                            "bsonType": "number"
                        }
                    },
                    "additionalProperties": false
                },
                "dynamodbid": {
                    "bsonType": "objectId"
                },
                "description": {
                    "bsonType": "string"
                },
                "trailers": {
                    "bsonType": "object",
                    "properties": {
                        "selected": {
                            "bsonType": "string"
                        },
                        "available": {
                            "bsonType": "array",
                            "additionalItems": true,
                            "uniqueItems": false,
                            "items": {
                                "bsonType": "string"
                            }
                        }
                    },
                    "additionalProperties": false
                },
                "price": {
                    "bsonType": "double"
                },
                "reviewed": {
                    "bsonType": "bool"
                },
                "volume": {
                    "bsonType": "objectId"
                },
                "ispublished": {
                    "bsonType": "bool"
                },
                "category": {
                    "bsonType": "array",
                    "additionalItems": true,
                    "uniqueItems": false,
                    "items": {
                        "bsonType": "objectId"
                    }
                },
                "extras": {
                    "bsonType": "array",
                    "additionalItems": true,
                    "uniqueItems": false,
                    "items": {
                        "bsonType": "object",
                        "properties": {
                            "url": {
                                "bsonType": "string"
                            },
                            "published": {
                                "bsonType": "bool"
                            }
                        },
                        "additionalProperties": false
                    }
                },
                "playlist": {
                    "bsonType": "array",
                    "additionalItems": true,
                    "uniqueItems": false
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
                "information": {
                    "bsonType": "object",
                    "properties": {
                        "studio": {
                            "bsonType": "objectId"
                        },
                        "length": {
                            "bsonType": "int",
                            "description": "Length of the scene in seconds"
                        },
                        "director": {
                            "bsonType": "array",
                            "additionalItems": true,
                            "uniqueItems": false,
                            "items": {
                                "bsonType": "objectId"
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
                "_modified": {
                    "bsonType": "date"
                },
                "_created": {
                    "bsonType": "date"
                }
            },
            "required": [
                "_id",
                "title",
                "slug",
                "coverimage",
                "images",
                "performance",
                "dynamodbid",
                "description",
                "trailers",
                "price",
                "reviewed",
                "volume",
                "ispublished",
                "category",
                "extras",
                "playlist",
                "thumbnails",
                "information",
                "_modified",
                "_created"
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