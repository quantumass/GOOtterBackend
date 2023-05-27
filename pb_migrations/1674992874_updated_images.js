migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("c83ukhcc0l9jq90")

  collection.createRule = ""

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("c83ukhcc0l9jq90")

  collection.createRule = null

  return dao.saveCollection(collection)
})
