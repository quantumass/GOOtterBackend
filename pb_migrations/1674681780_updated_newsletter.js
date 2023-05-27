migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("k2uwwnkoz6uv8wt")

  collection.name = "newsletters"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("k2uwwnkoz6uv8wt")

  collection.name = "newsletter"

  return dao.saveCollection(collection)
})
