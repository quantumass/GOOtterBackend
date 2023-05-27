migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("c83ukhcc0l9jq90")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "pcijogbu",
    "name": "image",
    "type": "file",
    "required": true,
    "unique": false,
    "options": {
      "maxSelect": 1,
      "maxSize": 5242880,
      "mimeTypes": [
        "image/jpg",
        "image/jpeg",
        "image/png",
        "image/gif",
        "image/webp"
      ],
      "thumbs": [
        "50x50"
      ]
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("c83ukhcc0l9jq90")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "pcijogbu",
    "name": "image",
    "type": "file",
    "required": true,
    "unique": false,
    "options": {
      "maxSelect": 1,
      "maxSize": 5242880,
      "mimeTypes": [
        "image/jpg",
        "image/jpeg",
        "image/png",
        "image/svg+xml",
        "image/gif",
        "image/webp"
      ],
      "thumbs": [
        "50x50"
      ]
    }
  }))

  return dao.saveCollection(collection)
})
