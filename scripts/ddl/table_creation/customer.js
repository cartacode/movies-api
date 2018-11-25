db.createCollection( "customer",{
    "capped": false,
    "validator": {
        "$jsonSchema": {
            "bsonType": "object",
            "properties": {
                "_id": {
                    "bsonType": "objectId"
                },
                "email": {
                    "bsonType": "string"
                },
                "password": {
                    "bsonType": "string"
                },
                "active": {
                    "bsonType": "bool"
                },
                "admin": {
                    "bsonType": "bool"
                },
                "liked": {
                    "bsonType": "object",
                    "properties": {
                        "movies": {
                            "bsonType": "array",
                            "additionalItems": true,
                            "uniqueItems": false,
                            "items": {
                                "bsonType": "objectId"
                            }
                        },
                        "scenes": {
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
                        },
                        "volumes": {
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
                "disliked": {
                    "bsonType": "object",
                    "properties": {
                        "movies": {
                            "bsonType": "array",
                            "additionalItems": true,
                            "uniqueItems": false,
                            "items": {
                                "bsonType": "objectId"
                            }
                        },
                        "scenes": {
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
                        },
                        "volumes": {
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
                "credit": {
                    "bsonType": "object",
                    "properties": {
                        "infostored": {
                            "bsonType": "bool"
                        },
                        "key": {
                            "bsonType": "string"
                        }
                    },
                    "additionalProperties": false
                },
                "purchased": {
                    "bsonType": "object",
                    "properties": {
                        "movies": {
                            "bsonType": "array",
                            "additionalItems": true,
                            "uniqueItems": false,
                            "items": {
                                "bsonType": "objectId"
                            }
                        },
                        "scenes": {
                            "bsonType": "array",
                            "additionalItems": true,
                            "uniqueItems": false,
                            "items": {
                                "bsonType": "objectId"
                            }
                        },
                        "volumes": {
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
                "wishlist": {
                    "bsonType": "object",
                    "properties": {
                        "movies": {
                            "bsonType": "array",
                            "additionalItems": true,
                            "uniqueItems": false,
                            "items": {
                                "bsonType": "objectId"
                            }
                        },
                        "scenes": {
                            "bsonType": "array",
                            "additionalItems": true,
                            "uniqueItems": false,
                            "items": {
                                "bsonType": "objectId"
                            }
                        },
                        "volumes": {
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
                "_created": {
                    "bsonType": "date"
                },
                "_modified": {
                    "bsonType": "date"
                },
                "preferences": {
                    "bsonType": "array",
                    "additionalItems": true,
                    "uniqueItems": false,
                    "items": {
                        "bsonType": "object",
                        "properties": {
                            "tag": {
                                "bsonType": "string"
                            },
                            "weight": {
                                "bsonType": "number"
                            }
                        },
                        "additionalProperties": false,
                        "required": [
                            "weight"
                        ]
                    }
                }
            },
            "required": [
                "_id",
                "email",
                "password",
                "active",
                "admin",
                "liked",
                "disliked",
                "credit",
                "purchased",
                "wishlist",
                "_created",
                "_modified",
                "preferences"
            ]
        }
    },
    "validationLevel": "off",
    "validationAction": "warn"
});
db.customer.createIndex(
{
    "_id": 1
},
{
    "name": "_id_"
}
);