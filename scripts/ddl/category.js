db.createCollection( "category",{
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
                "description": {
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
                "name",
                "slug",
                "description",
                "_created",
                "_modified"
            ]
        }
    },
    "validationLevel": "off",
    "validationAction": "warn"
});
db.category.createIndex(
{
    "_id": 1
},
{
    "name": "_id_"
}
);