db.createCollection("movie", {
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
							"bsonType": "number"
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
							"bsonType": "number"
						},
						"director": {
							"bsonType": "array",
							"additionalItems": true,
							"uniqueItems": false,
							"items": {
								"bsonType": "objectId"
							}
						},
						"quality": {
							"bsonType": "array",
							"additionalItems": true,
							"uniqueItems": false,
							"items": {
								"bsonType": "int"
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
							"bsonType": "number"
						},
						"upvotes": {
							"bsonType": "number"
						},
						"downvotes": {
							"bsonType": "number"
						},
						"views": {
							"bsonType": "number"
						}
					},
					"additionalProperties": false
				},
				"title": {
					"bsonType": "string"
				},
				"dynamodbid": {
					"bsonType": "string"
				},
				"description": {
					"bsonType": "string"
				},
				"price": {
					"bsonType": "double"
				},
				"slug": {
					"bsonType": "string",
					"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
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
				"chapters": {
					"bsonType": "array",
					"additionalItems": true,
					"uniqueItems": false
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
				"dynamodbid",
				"description",
				"price",
				"slug",
				"ispublished",
				"_created",
				"_modified",
				"chapters",
				"extras",
				"tags",
				"trailers"
			]
		}
	},
	"validationLevel": "off",
	"validationAction": "warn"
});
db.movie.createIndex(
	{
		"_id": 1
	},
	{
		"name": "_id_"
	}
);
