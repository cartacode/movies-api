db.createCollection("scene", {
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
				"dynamodbid": {
					"bsonType": "string",
					"pattern": "[0-9a-fA-F]{8}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{12}"
				},
				"description": {
					"bsonType": "string"
				},
				"volume": {
					"bsonType": "objectId"
				},
				"price": {
					"bsonType": "double"
				},
				"slug": {
					"bsonType": "string",
					"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
				},
				"series": {
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
				}
			},
			"required": [
				"_id",
				"title",
				"slug",
				"volume",
				"price",
				"ispublished",
				"_created",
				"_modified",
			]
		}
	},
	"validationLevel": "off",
	"validationAction": "warn"
});
db.scene.createIndex(
	{
		"_id": 1
	},
	{
		"name": "_id_"
	}
);
