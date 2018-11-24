db.createCollection( "customer",{
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
                "email": {
                    "bsonType": "string"
                },
                "cognitoId": {
                    "bsonType": "string",
                    "minLength": 0,
                    "maxLength": 24
                },
                "purchased": {
                    "bsonType": "object",
                    "properties": {
                        "movies": {
                            "bsonType": "array",
                            "additionalItems": true,
                            "minItems": 0,
                            "maxItems": 500,
                            "uniqueItems": true,
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
                        }
                    },
                    "additionalProperties": false,
                    "patternProperties": {
                        "volumes": {
                            "bsonType": "array",
                            "additionalItems": true,
                            "uniqueItems": false,
                            "items": {
                                "bsonType": "objectId"
                            }
                        }
                    }
                },
                "wishlist": {
                    "bsonType": "object",
                    "additionalProperties": false,
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
                        }
                    },
                    "patternProperties": {
                        "volumes": {
                            "bsonType": "array",
                            "additionalItems": true,
                            "uniqueItems": false,
                            "items": {
                                "bsonType": "objectId"
                            }
                        }
                    }
                },
                "credit": {
                    "bsonType": "object",
                    "properties": {
                        "token": {
                            "bsonType": "string"
                        },
                        "is_valid": {
                            "bsonType": "bool"
                        }
                    },
                    "additionalProperties": false
                }
            },
            "required": [
                "_id",
                "cognitoId"
            ]
        }
    },
    "validationLevel": "off",
    "validationAction": "warn"
});