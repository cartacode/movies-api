db.createCollection("volume", {
	"capped": false,
	"validator": {
		"$jsonSchema": {
			"bsonType": "object",
			"properties": {
				"_id": {
					"bsonType": "objectId"
				},
				"slug": {
					"bsonType": "string",
					"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
				},
				"description": {
					"bsonType": "string"
				},
				"title": {
					"bsonType": "string"
				},
				"price": {
					"bsonType": "double"
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
				"scenes": {
					"bsonType": "array",
					"additionalItems": true,
					"uniqueItems": true,
					"items": {
						"bsonType": "objectId"
					}
				},
				"tags": {
					"bsonType": "array",
					"additionalItems": true,
					"uniqueItems": true,
					"items": {
						"bsonType": "string"
					}
				}
			},
			"required": [
				"_id",
				"title",
				"price",
				"slug",
				"ispublished",
				"_created",
				"_modified"
			]
		}
	},
	"validationLevel": "off",
	"validationAction": "warn"
});
db.volume.createIndex(
	{
		"_id": 1
	},
	{
		"name": "_id_"
	}
);
