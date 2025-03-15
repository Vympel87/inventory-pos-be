# Sale Item API Specs

## GET Single

Endpoint: GET /api/v1/auth/sale_items/:id

Response Body (Success):

```
{
  "id": 12,
  "sale_id": 34,
  "product_id": 56,
  "quantity": 2,
  "unit_price": 10000,
  "total_price": 20000,
  "created_at": "2025-04-22T12:00:00Z",
  "updated_at": "2025-04-22T12:10:00Z",
  "product": {
    "id": 56,
    "name": "Pensil 2B",
    "size": "Standard",
    "color": "Hitam",
    "category": {
      "id": 3,
      "name": "Alat Tulis"
    }
  },
  "sale": {
    "id": 34,
    "store_id": "21d9e598-3a4e-4e0e-9a79-f113a2a6898e",
    "user_id": "79b8456c-5826-465d-84a9-184fd7e98721",
    "sale_date": "2025-04-22T10:00:00Z",
    "sale_number": "SL-20250422-001",
    "payment_method": "Cash",
    "total_amount": 20000
  }
}

```

Response Body (Failed):

```
{
  "message": "Sale item not found",
  "errorCode": 1008,
  "statusCode": 404
}
```

## GET LIST

Endpoint: GET /api/v1/auth/sale_items

Query Parameters:

* Page: `<page number>`
* Limit: `<limit per page>`

Response Body (Success):

```
{
  "status": "success",
  "data": [
    {
      "id": 12,
      "sale_id": 34,
      "product_id": 56,
      "quantity": 2,
      "unit_price": 10000,
      "total_price": 20000,
      "created_at": "2025-04-22T12:00:00Z",
      "updated_at": "2025-04-22T12:10:00Z",
      "product": {
        "id": 56,
        "name": "Pensil 2B",
        "size": "Standard",
        "color": "Hitam",
        "category": {
          "id": 3,
          "name": "Alat Tulis"
        }
      },
      "sale": {
        "id": 34,
        "store_id": "21d9e598-3a4e-4e0e-9a79-f113a2a6898e",
        "user_id": "79b8456c-5826-465d-84a9-184fd7e98721",
        "sale_date": "2025-04-22T10:00:00Z",
        "sale_number": "SL-20250422-001",
        "payment_method": "Cash",
        "total_amount": 20000
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
  "message": "No sale items found",
  "errorCode": 1008,
  "statusCode": 404
}
```

## CREATE (Single Or Bulk Insert)

Endpoint: POST /api/v1/auth/sale_items

Request Body:

```
[
  {
    "sale_id": 34,
    "product_id": 56,
    "quantity": 2,
    "unit_price": 10000,
    "total_price": 20000
  },
  {
    "sale_id": 34,
    "product_id": 57,
    "quantity": 1,
    "unit_price": 15000,
    "total_price": 15000
  }
]
```

Response Body (Success):

```
{
  "status": "success",
  "data": [
    {
      "id": 101,
      "sale_id": 34,
      "product_id": 56,
      "quantity": 2,
      "unit_price": 10000,
      "total_price": 20000,
      "created_at": "2025-04-22T12:00:00Z",
      "updated_at": "2025-04-22T12:00:00Z"
    },
    {
      "id": 102,
      "sale_id": 34,
      "product_id": 57,
      "quantity": 1,
      "unit_price": 15000,
      "total_price": 15000,
      "created_at": "2025-04-22T12:00:00Z",
      "updated_at": "2025-04-22T12:00:00Z"
    }
  ]
}

```

Response Body (Failed):

```
{
  "message": "All fields must be filled",
  "errorCode": 1007,
  "statusCode": 422,
  "errors": {
    "store_id": "Product ID is required",
    "items": "At least one item must be included"
  }
}

```

## UPDATE

Endpoint: PUT /api/v1/auth/sale_items/:id

Request Body:

```
{
  "sale_id": 34,
  "product_id": 56,
  "quantity": 3,
  "unit_price": 11000,
  "total_price": 33000
}
```

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "id": 101,
    "sale_id": 34,
    "product_id": 56,
    "quantity": 3,
    "unit_price": 11000,
    "total_price": 33000,
    "created_at": "2025-04-22T12:00:00Z",
    "updated_at": "2025-04-22T12:00:00Z"
  }
}
```

Response Body (Failed):

```
{
  "message": "Sale item not found",
  "errorCode": 1008,
  "statusCode": 404
}
```

## DELETE

Endpoint: DELETE /api/v1/auth/sale_items/:id

Response Body (Success):

```
{
  "status": "success",
  "message": "Sale item deleted successfully"
}
```

Response Body (Failed):

```
{
  "message": "Sale item not found",
  "errorCode": 1008,
  "statusCode": 404
}
```
