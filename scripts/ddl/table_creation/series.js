db.createCollection( "series",{
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
                "description": {
                    "bsonType": "string"
                },
                "is_published": {
                    "bsonType": "bool"
                },
                "tags": {
                    "bsonType": "array",
                    "additionalItems": true,
                    "uniqueItems": false,
                    "items": {
                        "bsonType": "string"
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
            "required": [
                "_id",
                "title",
                "slug",
                "description",
                "is_published",
                "tags",
                "volumes"
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