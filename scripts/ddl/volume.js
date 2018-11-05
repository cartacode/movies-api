db.createCollection( "volume",{
    "storageEngine": {
        "wiredTiger": {}
    },
    "capped": false,
    "validator": {
        "$jsonSchema": {
            "bsonType": "object",
            "additionalProperties": false,
            "properties": {
                "_id": {
                    "bsonType": "objectId"
                },
                "series": {
                    "bsonType": "objectId"
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
                        "cover": {
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
                    }
                },
                "title": {
                    "bsonType": "string"
                },
                "description": {
                    "bsonType": "string"
                },
                "price": {
                    "bsonType": "double"
                },
                "reviewed": {
                    "bsonType": "bool"
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
                "_modified": {
                    "bsonType": "date"
                },
                "_created": {
                    "bsonType": "date"
                }
            },
            "required": [
                "_id",
                "images",
                "information",
                "title",
                "description",
                "price",
                "reviewed",
                "ispublished",
                "category",
                "extras",
                "trailers",
                "_modified",
                "_created"
            ]
        }
    },
    "validationLevel": "off",
    "validationAction": "warn"
});