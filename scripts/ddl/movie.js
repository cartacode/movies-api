db.createCollection( "movie",{
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
                    "additionalProperties": false,
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
                    }
                },
                "performance": {
                    "bsonType": "object",
                    "additionalProperties": false,
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
                    }
                },
                "dynamodbid": {
                    "bsonType": "objectId"
                },
                "description": {
                    "bsonType": "string"
                },
                "trailers": {
                    "bsonType": "object",
                    "additionalProperties": false,
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
                    }
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
                        "additionalProperties": false,
                        "properties": {
                            "url": {
                                "bsonType": "string"
                            },
                            "published": {
                                "bsonType": "bool"
                            }
                        }
                    }
                },
                "thumbnails": {
                    "bsonType": "object",
                    "additionalProperties": false,
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
                    }
                },
                "information": {
                    "bsonType": "object",
                    "additionalProperties": false,
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
                    }
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
db.movie.createIndex(
{
    "_id": 1
},
{
    "name": "_id_"
}
);