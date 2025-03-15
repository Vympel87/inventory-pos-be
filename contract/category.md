# Category API Specs

## GET Single

Endpoint: GET /api/v1/auth/categories/:id

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "id": 1,
    "name": "Baju Blouse",
    "description": "Kategori untuk baju blouse",
    "created_at": "2025-04-17T12:34:56Z",
    "updated_at": "2025-04-17T12:34:56Z"
  }
}

```

Response Body (Failed):

```
{
  "message": "Category not found",
  "errorCode": 5001,
  "statusCode": 404
}
```

## GET List

Endpoint: GET /api/v1/auth/categories/:id?limit=&page=

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
      "name": "Beverages",
      "description": "All types of drinks",
      "created_at": "2025-04-01T10:20:30Z",
      "updated_at": "2025-04-05T14:45:00Z"
    },
    {
      "id": 2,
      "name": "Snacks",
      "description": "Packaged and fresh snacks",
      "created_at": "2025-04-02T11:15:00Z",
      "updated_at": "2025-04-06T09:00:00Z"
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

## CREATE

Endpoint: POST /api/v1/auth/categories/:id

Request Body:

```
{
  "name": "Beverages",
  "description": "All kinds of drinks including soft drinks, juices, etc."
}
```

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "id": 1,
    "name": "Beverages",
    "description": "All kinds of drinks including soft drinks, juices, etc.",
    "created_at": "2025-04-19T10:20:30Z",
    "updated_at": "2025-04-19T10:20:30Z"
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
    "name": "Name is required"
  }
}
```

## UPDATE

Endpoint: PUT /api/v1/auth/categories/:id

Request Body:

```
{
  "name": "Minuman Ringan",
  "description": "Kategori untuk semua jenis minuman ringan"
}
```

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "id": 1,
    "name": "Minuman Ringan",
    "description": "Kategori untuk semua jenis minuman ringan",
    "created_at": "2024-08-17T10:00:00Z",
    "updated_at": "2025-04-19T14:22:11Z"
  }
}
```

Response Body (Failed):

```
{
  "message": "Store not found",
  "errorCode": 5001,
  "statusCode": 404
}
```

## DELETE

Endpoint: DELETE /api/v1/auth/categories/:id

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "message": "Category deleted successfully"
  }
}
```

Response Body (Failed):

```
{
  "message": "Category not found",
  "errorCode": 5001,
  "statusCode": 404
}
```
