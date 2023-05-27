migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("c83ukhcc0l9jq90")

  collection.listRule = "@request.auth.id=user.id"
  collection.viewRule = "@request.auth.id=user.id"
  collection.updateRule = "@request.auth.id=user.id"
  collection.deleteRule = "@request.auth.id=user.id"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("c83ukhcc0l9jq90")

  collection.listRule = null
  collection.viewRule = null
  collection.updateRule = null
  collection.deleteRule = null

  return dao.saveCollection(collection)
})
