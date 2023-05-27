migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("k2uwwnkoz6uv8wt")

  collection.createRule = ""

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("k2uwwnkoz6uv8wt")

  collection.createRule = null

  return dao.saveCollection(collection)
})
