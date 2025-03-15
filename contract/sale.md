# Sale API Specs

## GET Single

Endpoint: GET /api/v1/auth/sales/:id

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "id": 1,
    "store_id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
    "user_id": "e24a9f20-8717-4329-99c1-8dcff54830e4",
    "sale_date": "2025-04-21T10:15:00Z",
    "sale_number": "SALE-20250421001",
    "payment_method": "cash",
    "total_amount": 250000.00,
    "created_at": "2025-04-21T10:16:00Z",
    "updated_at": "2025-04-21T10:16:00Z"
  }
}
```

Response Body (Failed):

```
{
  "message": "Sale not found",
  "errorCode": 8001,
  "statusCode": 404
}
```

## GET LIST

Endpoint: GET /api/v1/auth/sales

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
      "store_id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
      "user_id": "e24a9f20-8717-4329-99c1-8dcff54830e4",
      "sale_date": "2025-04-21T10:15:00Z",
      "sale_number": "SALE-20250421001",
      "payment_method": "cash",
      "total_amount": 250000.00,
      "created_at": "2025-04-21T10:16:00Z",
      "updated_at": "2025-04-21T10:16:00Z"
    },
    {
      "id": 2,
      "store_id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
      "user_id": "e24a9f20-8717-4329-99c1-8dcff54830e4",
      "sale_date": "2025-04-20T14:00:00Z",
      "sale_number": "SALE-20250420001",
      "payment_method": "credit_card",
      "total_amount": 120000.00,
      "created_at": "2025-04-21T10:16:00Z",
      "updated_at": "2025-04-21T10:16:00Z"
    }
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

## CREATE (Transaction With Sale Item, Single or Bulk Insert on Sale Item)

Endpoint: POST /api/v1/auth/sales

Request Body:

```
{
  "store_id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
  "user_id": "e24a9f20-8717-4329-99c1-8dcff54830e4",
  "sale_date": "2025-04-21T10:15:00Z",
  "sale_number": "SALE-20250421001",
  "payment_method": "cash",
  "total_amount": 250000.00,
  "items": [
    {
      "product_id": 101,
      "quantity": 2,
      "unit_price": 100000.00,
      "total_price": 200000.00
    },
    {
      "product_id": 102,
      "quantity": 1,
      "unit_price": 50000.00,
      "total_price": 50000.00
    }
  ]
}
```

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "sale": {
      "id": 1,
      "store_id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
      "user_id": "e24a9f20-8717-4329-99c1-8dcff54830e4",
      "sale_date": "2025-04-21T10:15:00Z",
      "sale_number": "SALE-20250421001",
      "payment_method": "cash",
      "total_amount": 250000.00,
      "created_at": "2025-04-21T10:16:00Z",
      "updated_at": "2025-04-21T10:16:00Z"
    },
    "sale_items": [
      {
        "id": 1,
        "sale_id": 1,
        "product_id": 101,
        "quantity": 2,
        "unit_price": 100000.00,
        "total_price": 200000.00,
        "created_at": "2025-04-21T10:16:00Z",
        "updated_at": "2025-04-21T10:16:00Z"
      },
      {
        "id": 2,
        "sale_id": 1,
        "product_id": 102,
        "quantity": 1,
        "unit_price": 50000.00,
        "total_price": 50000.00,
        "created_at": "2025-04-21T10:16:00Z",
        "updated_at": "2025-04-21T10:16:00Z"
      }
    ]
  }
}

```

Response Body (Failed):

```
{
  "message": "All fields must be filled",
  "errorCode": 1007,
  "statusCode": 422,
  "errors": {
    "store_id": "Store ID is required",
    "items": "At least one item must be included"
  }
}
```

## UPDATE

Endpoint: PUT /api/v1/auth/sales/:id

Request Body:

```
{
  "store_id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
  "user_id": "e24a9f20-8717-4329-99c1-8dcff54830e4",
  "sale_date": "2025-04-22T14:30:00Z",
  "sale_number": "SALE-20250422001",
  "payment_method": "transfer",
  "total_amount": 300000.00
}
```

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "id": 1,
    "store_id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
    "user_id": "e24a9f20-8717-4329-99c1-8dcff54830e4",
    "sale_date": "2025-04-22T14:30:00Z",
    "sale_number": "SALE-20250422001",
    "payment_method": "transfer",
    "total_amount": 300000.00,
    "created_at": "2025-04-21T10:16:00Z",
    "updated_at": "2025-04-22T14:35:00Z"
  }
}
```

Response Body (Failed):

```
{
  "message": "Sale not found",
  "errorCode": 8001,
  "statusCode": 404
}
```

## DELETE (Transaction With Sale Item)

Endpoint: DELETE /api/v1/auth/sales/:id

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "message": "Sale and related sale items deleted successfully"
  }
}
```

Response Body (Failed):

```
{
  "message": "Sale not found",
  "errorCode": 8001,
  "statusCode": 404
}
```
