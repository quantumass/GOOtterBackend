migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("zem2egj8utstnhn")

  collection.createRule = ""

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("zem2egj8utstnhn")

  collection.createRule = null

  return dao.saveCollection(collection)
})
