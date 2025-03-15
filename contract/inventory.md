# Inventory API Specs

## GET Single Inventory With Product and Product Category

Endpoint: GET /api/v1/auth/inventories/:id

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "id": 101,
    "store_id": "52e7c2d8-235f-4673-b285-02df547bb4ff",
    "quantity": 100,
    "min_stock_level": 10,
    "created_at": "2025-04-20T14:30:00Z",
    "updated_at": "2025-04-20T14:30:00Z",
    "product": {
      "id": 123,
      "name": "Air Jordan 1",
      "image_url": "https://example.com/images/aj1.jpg",
      "size": "42",
      "color": "Red",
      "price": 249.99,
      "category": {
        "id": 1,
        "name": "Sneakers"
      }
    }
  }
}
```

Response Body (Failed):

```
{
  "message": "Inventory not found",
  "errorCode": 7001,
  "statusCode": 404
}
```

## GET List Inventory With Product and Product Category

Endpoint: GET /api/v1/auth/inventories?limit=&page=

Query Parameters:

* Page: `<page number>`
* Limit: `<limit per page>`

Response Body (Success):

```
{
  "status": "success",
  "data": [
    {
      "id": 1,
      "store_id": "52e7c2d8-235f-4673-b285-02df547bb4ff",
      "quantity": 100,
      "min_stock_level": 10,
      "created_at": "2025-04-20T09:30:00Z",
      "updated_at": "2025-04-20T09:30:00Z",
      "product": {
        "id": 12,
        "name": "Air Jordan 1",
        "image_url": "https://example.com/images/aj1.jpg",
        "size": "42",
        "color": "Red",
        "price": 249.99,
        "category": {
          "id": 3,
          "name": "Shoes"
        }
      }
    },
    ...
  ],
  "pagination": {
    "total_items": 2,
    "total_pages": 1,
    "current_page": 1,
    "items_per_page": 2
  }
}
```

Response Body (Failed):

```
{
  "message": "Invalid query parameters",
  "errorCode": 1004,
  "statusCode": 400,
  "errors": {
    "limit": "Limit must be a positive number",
    "page": "Page must be a positive number"
  }
}
```

## CREATE Inventory - (Branch Store)

Endpoint: POST /api/v1/auth/inventories

Request Body:

```
{
  "product_id": 1,
  "store_id": "52e7c2d8-235f-4673-b285-02df547bb4ff",
  "quantity": 100,
  "min_stock_level": 10
}
```

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "id": 25,
    "product_id": 1,
    "store_id": "52e7c2d8-235f-4673-b285-02df547bb4ff",
    "quantity": 100,
    "min_stock_level": 10,
    "created_at": "2025-04-20T10:20:00Z",
    "updated_at": "2025-04-20T10:20:00Z"
  }
}

```

Response Body (Failed):

```
{
  "message": "Validation failed",
  "errorCode": 1006,
  "statusCode": 400,
  "errors": {
    "store_id": "Need to be filled",
    "min_stock_level": "Stock must be greater than 0"
  }
}
```

## UPDATE Inventory

Endpoint: PUT /api/v1/auth/inventories/:id

Request Body:

```
{
  "store_id": "52e7c2d8-235f-4673-b285-02df547bb4ff",
  "product_id": 1,
  "quantity": 80,
  "min_stock_level": 5
}
```

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "id": 3,
    "store_id": "52e7c2d8-235f-4673-b285-02df547bb4ff",
    "product_id": 1,
    "quantity": 80,
    "min_stock_level": 5,
    "created_at": "2025-04-20T10:20:30Z",
    "updated_at": "2025-04-20T11:00:00Z"
  }
}
```

Response Body (Failed):

```
{
  "message": "Inventory not found",
  "errorCode": 7001,
  "statusCode": 404
}
```

## DELETE Inventory - (Branch Store)

Endpoint: DELETE /api/v1/auth/inventories/:id

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "message": "Inventory deleted successfully"
  }
}

```

Response Body (Failed):

```
{
  "message": "Inventory not found",
  "errorCode": 7001,
  "statusCode": 404
}
```
