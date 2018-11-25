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
				"birthplace": {
					"bsonType": "string"
				},
				"slug": {
					"bsonType": "string"
				},
				"social": {
					"bsonType": "object",
					"properties": {
						"twitter": {
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
				"favorites": {
					"bsonType": "int"
				},
				"traits": {
					"bsonType": "object",
					"properties": {
						"haircolor": {
							"bsonType": "string"
						},
						"ethnicity": {
							"bsonType": "string"
						},
						"piercings": {
							"bsonType": "bool"
						},
						"tattoos": {
							"bsonType": "bool"
						}
					},
					"additionalProperties": false
				},
				"gender": {
					"bsonType": "string"
				},
				"size": {
					"bsonType": "object",
					"properties": {
						"weight": {
							"bsonType": "int"
						},
						"waist": {
							"bsonType": "int"
						},
						"bust": {
							"bsonType": "string"
						},
						"height": {
							"bsonType": "int"
						}
					},
					"additionalProperties": false
				},
				"images": {
					"bsonType": "object",
					"properties": {
						"portrait": {
							"bsonType": "string",
							"pattern": "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"
						},
						"landscape": {
							"bsonType": "string",
							"pattern": "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"
						}
					},
					"additionalProperties": false
				},
				"nationality": {
					"bsonType": "string"
				},
				"orientation": {
					"bsonType": "string"
				},
				"bio": {
					"bsonType": "string"
				},
				"birthdate": {
					"bsonType": "date"
				}
			},
			"required": [
				"_id",
				"name",
				"director",
				"birthplace",
				"slug",
				"social",
				"favorites",
				"traits",
				"gender",
				"size",
				"images",
				"nationality",
				"orientation",
				"birthdate"
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
