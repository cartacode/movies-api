db.createCollection( "studio",{
    "capped": false,
    "validator": {
        "$jsonSchema": {
            "bsonType": "object",
            "properties": {
                "_id": {
                    "bsonType": "objectId"
                },
                "description": {
                    "bsonType": "string"
                },
                "name": {
                    "bsonType": "string"
                },
                "slug": {
                    "bsonType": "string"
                },
                "images": {
                    "bsonType": "object",
                    "properties": {
                        "thumbnail": {
                            "bsonType": "string",
                            "pattern": "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"
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
                "thumbnail": {
                    "bsonType": "string"
                }
            },
            "required": [
                "_id",
                "name",
                "slug"
            ]
        }
    },
    "validationLevel": "off",
    "validationAction": "warn"
});
db.studio.createIndex(
{
    "_id": 1
},
{
    "name": "_id_"
}
);