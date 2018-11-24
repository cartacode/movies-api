db.createCollection( "series",{
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
                "coverimage": {
                    "bsonType": "string"
                },
                "information": {
                    "bsonType": "object",
                    "properties": {
                        "studio": {
                            "bsonType": "objectId"
                        },
                        "director": {
                            "bsonType": "array",
                            "additionalItems": true,
                            "uniqueItems": false,
                            "items": {
                                "bsonType": "string"
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
                "title": {
                    "bsonType": "string"
                },
                "description": {
                    "bsonType": "string"
                },
                "slug": {
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
                "category": {
                    "bsonType": "array",
                    "additionalItems": true,
                    "uniqueItems": false,
                    "items": {
                        "bsonType": "objectId"
                    }
                }
            },
            "required": [
                "_id",
                "images",
                "coverimage",
                "information",
                "performance",
                "title",
                "description",
                "slug",
                "ispublished",
                "_created",
                "_modified",
                "category"
            ]
        }
    },
    "validationLevel": "off",
    "validationAction": "warn"
});
db.series.createIndex(
{
    "_id": 1
},
{
    "name": "_id_"
}
);