db.createCollection("star", {
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
				"director": {
					"bsonType": "bool"
				},

				"slug": {
					"bsonType": "string",
					"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
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
				"director",
				"slug",
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
