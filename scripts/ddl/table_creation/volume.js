db.createCollection("volume", {
	"capped": false,
	"validator": {
		"$jsonSchema": {
			"bsonType": "object",
			"properties": {
				"_id": {
					"bsonType": "objectId"
				},
				"images": {
					"bsonType": "object",
					"properties": {
						"landscape": {
							"bsonType": "string",
							"pattern": "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"
						},
						"portrait": {
							"bsonType": "string",
							"pattern": "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"
						},
						"banner": {
							"bsonType": "string",
							"pattern": "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"
						},
						"detailpage": {
							"bsonType": "string",
							"pattern": "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"
						}
					},
					"additionalProperties": false
				},
				"thumbnails": {
					"bsonType": "object",
					"properties": {
						"prefix": {
							"bsonType": "string"
						},
						"count": {
							"bsonType": "int"
						},
						"format": {
							"bsonType": "string"
						}
					},
					"additionalProperties": false
				},
				"information": {
					"bsonType": "object",
					"properties": {
						"studio": {
							"bsonType": "objectId"
						},
						"length": {
							"bsonType": "int"
						},
						"director": {
							"bsonType": "array",
							"additionalItems": true,
							"uniqueItems": false,
							"items": {
								"bsonType": "objectId"
							}
						},
						"stars": {
							"bsonType": "array",
							"additionalItems": true,
							"uniqueItems": false,
							"items": {
								"bsonType": "objectId"
							}
						}
					},
					"additionalProperties": false
				},
				"performance": {
					"bsonType": "object",
					"properties": {
						"rank": {
							"bsonType": "int"
						},
						"upvotes": {
							"bsonType": "int"
						},
						"downvotes": {
							"bsonType": "int"
						},
						"views": {
							"bsonType": "int"
						}
					},
					"additionalProperties": false
				},
				"series": {
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
				"extras": {
					"bsonType": "array",
					"additionalItems": true,
					"uniqueItems": false,
					"items": {
						"bsonType": "object",
						"properties": {
							"published": {
								"bsonType": "bool"
							},
							"url": {
								"bsonType": "string",
								"pattern": "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"
							}
						},
						"additionalProperties": false,
						"required": [
							"published",
							"url"
						]
					}
				},
				"scenes": {
					"bsonType": "array",
					"additionalItems": true,
					"uniqueItems": false,
					"items": {
						"bsonType": "objectId"
					}
				},
				"tags": {
					"bsonType": "array",
					"additionalItems": true,
					"uniqueItems": false,
					"items": {
						"bsonType": "string"
					}
				},
				"trailers": {
					"bsonType": "array",
					"additionalItems": true,
					"uniqueItems": false,
					"items": {
						"bsonType": "object",
						"properties": {
							"published": {
								"bsonType": "bool"
							},
							"title": {
								"bsonType": "string"
							},
							"url": {
								"bsonType": "string",
								"pattern": "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"
							}
						},
						"additionalProperties": false,
						"required": [
							"published",
							"title",
							"url"
						]
					}
				}
			},
			"required": [
				"_id",
				"images",
				"thumbnails",
				"information",
				"performance",
				"title",
				"price",
				"ispublished",
				"_created",
				"_modified",
				"extras",
				"scenes",
				"tags",
				"trailers"
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
