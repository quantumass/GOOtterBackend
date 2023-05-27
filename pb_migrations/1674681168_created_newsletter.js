migrate((db) => {
  const collection = new Collection({
    "id": "k2uwwnkoz6uv8wt",
    "created": "2023-01-25 21:12:47.931Z",
    "updated": "2023-01-25 21:12:47.931Z",
    "name": "newsletter",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "eqhpcgpz",
        "name": "email",
        "type": "email",
        "required": true,
        "unique": false,
        "options": {
          "exceptDomains": null,
          "onlyDomains": null
        }
      }
    ],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("k2uwwnkoz6uv8wt");

  return dao.deleteCollection(collection);
})
