db.createCollection( "star",{
    "capped": false,
    "validator": {
        "$jsonSchema": {
            "bsonType": "object",
            "properties": {
                "_id": {
                    "bsonType": "objectId"
                },
                "name": {
                    "bsonType": "string"
                },
                "slug": {
                    "bsonType": "string"
                },
                "gender": {
                    "bsonType": "string"
                },
                "bio": {
                    "bsonType": "string"
                },
                "birthplace": {
                    "bsonType": "string"
                },
                "birthdate": {
                    "bsonType": "date"
                },
                "is_director": {
                    "bsonType": "bool"
                },
                "rank": {
                    "bsonType": "number"
                },
                "picture": {
                    "bsonType": "string"
                },
                "size": {
                    "bsonType": "object",
                    "properties": {
                        "weight": {
                            "bsonType": "number"
                        },
                        "waist": {
                            "bsonType": "number"
                        },
                        "bust": {
                            "bsonType": "string"
                        },
                        "height": {
                            "bsonType": "number"
                        }
                    },
                    "additionalProperties": false
                },
                "traits": {
                    "bsonType": "object",
                    "properties": {
                        "ethnicity": {
                            "bsonType": "string"
                        },
                        "haircolor": {
                            "bsonType": "string"
                        },
                        "piercings": {
                            "bsonType": "bool"
                        },
                        "tattoos": {
                            "bsonType": "bool"
                        },
                        "starsign": {
                            "bsonType": "string"
                        }
                    },
                    "additionalProperties": false
                },
                "social": {
                    "bsonType": "object",
                    "properties": {
                        "twitter": {
                            "bsonType": "string"
                        },
                        "youtube": {
                            "bsonType": "string"
                        },
                        "instagram": {
                            "bsonType": "string"
                        },
                        "snapchat": {
                            "bsonType": "string"
                        }
                    },
                    "additionalProperties": false
                },
                "_created": {
                    "bsonType": "date"
                },
                "_modified": {
                    "bsonType": "date"
                }
            },
            "required": [
                "_id",
                "name",
                "slug",
                "gender",
                "bio",
                "birthplace",
                "birthdate",
                "is_director",
                "rank",
                "picture",
                "size",
                "traits",
                "social",
                "_created",
                "_modified"
            ]
        }
    },
    "validationLevel": "off",
    "validationAction": "warn"
});
db.star.createIndex(
{
    "_id": 1
},
{
    "name": "_id_"
}
);