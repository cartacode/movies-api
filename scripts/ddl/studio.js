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
                "_created": {
                    "bsonType": "date"
                },
                "_modified": {
                    "bsonType": "date"
                }
            },
            "required": [
                "_id",
                "description",
                "name",
                "slug",
                "_created",
                "_modified"
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