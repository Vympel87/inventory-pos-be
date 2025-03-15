# Product API Specs

## GET Single Product With Category

Endpoint: GET /api/v1/auth/products/:id

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "id": 1,
    "name": "Celana Jeans",
    "description": "Celana panjang berbahan denim",
    "sku": "PRD001",
    "image_url": "https://example.com/images/jeans.png",
    "size": "L",
    "color": "Biru",
    "price": 250000.00,
    "category": {
      "id": 2,
      "name": "Pakaian"
    },
    "created_at": "2025-04-18T10:12:00Z",
    "updated_at": "2025-04-19T09:44:33Z"
  }
}
```

Response Body (Failed):

```
{
  "message": "Product not found",
  "errorCode": 6001,
  "statusCode": 404
}
```

## CREATE Product (Transaction With Inventory) - (Central Store)

Endpoint: POST /api/v1/auth/products

Request Body:

```
{
  "category_id": 1,
  "name": "Air Jordan 1",
  "description": "High-top sneakers",
  "sku": "AJ1-001",
  "image_url": "https://example.com/images/aj1.jpg",
  "size": "42",
  "color": "Red",
  "price": 249.99,
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
    "product": {
      "id": 1,
      "name": "Air Jordan 1",
      "description": "High-top sneakers",
      "sku": "AJ1-001",
      "image_url": "https://example.com/images/aj1.jpg",
      "size": "42",
      "color": "Red",
      "price": 249.99,
      "category": 1
      },
      "created_at": "2025-04-19T12:00:00Z",
      "updated_at": "2025-04-19T12:00:00Z"
    },
    "inventory": {
      "id": 7,
      "store_id": "52e7c2d8-235f-4673-b285-02df547bb4ff",
      "quantity": 100,
      "min_stock_level": 10,
      "created_at": "2025-04-19T12:00:00Z",
      "updated_at": "2025-04-19T12:00:00Z"
    }
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
    "sku": "SKU already exists",
    "price": "Price must be greater than 0"
  }
}
```

## UPDATE Product

Endpoint: PUT /api/v1/auth/products/:id

Request Body:

```
{
  "category_id": 1,
  "name": "Air Jordan 1 Retro",
  "description": "High-top sneakers, updated edition",
  "sku": "AJ1-001-R",
  "image_url": "https://example.com/images/aj1-retro.jpg",
  "size": "43",
  "color": "Black/Red",
  "price": 279.99
}
```

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "id": 123,
    "category": 1
    "name": "Air Jordan 1 Retro",
    "description": "High-top sneakers, updated edition",
    "sku": "AJ1-001-R",
    "image_url": "https://example.com/images/aj1-retro.jpg",
    "size": "43",
    "color": "Black/Red",
    "price": 279.99,
    "created_at": "2025-04-10T10:30:00Z",
    "updated_at": "2025-04-19T17:45:00Z"
  }
}
```

Response Body (Failed):

```
{
  "message": "Product not found",
  "errorCode": 6001,
  "statusCode": 404
}
```

## DELETE Product (Transaction With Inventory) - (Central Store)

Endpoint: DELETE /api/v1/auth/products/:id

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "message": "Product deleted successfully"
  }
}

```

Response Body (Failed):

```
{
  "message": "Product not found",
  "errorCode": 6001,
  "statusCode": 404
}
```
