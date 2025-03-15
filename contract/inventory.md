# Inventory API Specs

## GET Single Inventory With Product and Product Category

Endpoint: GET /api/v1/auth/inventories/:id

## GET List Inventory With Product and Product Category

Endpoint: GET /api/v1/auth/inventories?limit=&page=

Query Parameters:

* Page: `<page number>`
* Limit: `<limit per page>`

## CREATE Inventory - (Branch Store)

Endpoint: POST /api/v1/auth/inventories

Request Body:

Response Body (Success):

Response Body (Failed):

## UPDATE Inventory

Endpoint: PUT /api/v1/auth/inventories/:id

Request Body:

Response Body (Success):

Response Body (Failed):

## DELETE Inventory - (Branch Store)

Endpoint: DELETE /api/v1/auth/inventories/:id

Request Body:

Response Body (Success):

Response Body (Failed):
