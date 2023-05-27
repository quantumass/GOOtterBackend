migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("c83ukhcc0l9jq90")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "dxoprh1d",
    "name": "convertedImage",
    "type": "file",
    "required": false,
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
      "thumbs": []
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("c83ukhcc0l9jq90")

  // remove
  collection.schema.removeField("dxoprh1d")

  return dao.saveCollection(collection)
})
