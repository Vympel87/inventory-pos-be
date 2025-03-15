# Store API Specs

## GET Single

Endpoint: GET /api/v1/auth/store/:id

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "id": "f1d1de08-bd2e-4b3c-9116-bf0e8a5ac25f",
    "name": "Toko Mas Amba",
    "email": "tomba@store.com",
    "address": "Jl. Ngawi-Jomokerto No. 121",
    "phone": "08123456789",
    "store_type": "CENTRAL",
    "created_at": "2025-04-17T10:11:12Z",
    "updated_at": "2025-04-17T10:11:12Z",
  }
}


```

Response Body (Failed):

```
{
  "message": "Store not found",
  "errorCode": 3001,
  "statusCode": 404
}
```

## GET Stores by Parent Store

Endpoint: GET /api/v1/auth/stores/central/:id?limit=&page=

Query Parameters:

* Page: `<page number>`
* Limit: `<limit per page>`

Response Body (Success):

```
{
  "status": "success",
  "data": [
    {
      "id": "44a8ff38-6a55-4cf4-8e3f-26166c458497",
      "name": "Toko Mas Amba",
      "email": "tomba@store.com",
      "address": "Jl. Ngawi-Jomokerto No. 121",
      "phone": "0811223344",
      "store_type": "CENTRAL",
      "created_at": "2025-04-12T08:10:00Z",
      "updated_at": "2025-04-12T08:10:00Z"
    },
    {
      "id": "44a8ff38-6a55-4cf4-8e3f-26166c458497",
      "name": "Toko Mas Amba Surabaya",
      "email": "tomba@store.com",
      "address": "Jl. Diponegoro No. 3",
      "phone": "0811223344",
      "store_type": "BRANCH",
      "created_at": "2025-04-12T08:10:00Z",
      "updated_at": "2025-04-12T08:10:00Z"
    },
    {
      "id": "44a8ff38-6a55-4cf4-8e3f-26166c458497",
      "name": "Toko Mas Amba Kediri",
      "email": "tomba@store.com",
      "address": "Jl. Sudirman No. 1",
      "phone": "0811223344",
      "store_type": "BRANCH",
      "created_at": "2025-04-12T08:10:00Z",
      "updated_at": "2025-04-12T08:10:00Z"
    },
  ],
  "pagination": {
    "total_items": 3,
    "total_pages": 1,
    "current_page": 1,
    "items_per_page": 3
  }
}

```

Response Body (Failed):

```
{
  "message": "Parent store not found",
  "errorCode": 3002,
  "statusCode": 404
}

```

## CREATE

Endpoint: POST /api/v1/auth/store

Request Body:

```
{
  "name": "Toko Mas Amba Kediri",
  "address": "Jl. Sudirman No. 1",
  "parent_store_id": "a1f63e45-4b8d-4b2c-b9d3-6c168cb6c847"
}
```

Response Body (Success) :

```
{
  "status": "success",
  "data": {
    "id": "b9a60fa2-5ad1-4d9c-b7c7-6b67fd13f880",
    "name": "Toko Mas Amba Kediri",
    "email": "cabang1@store.com",
    "address": "Jl. Sudirman No. 1",
    "phone": "081234567890",
    "store_type": "BRANCH",
    "parent_store_id": "a1f63e45-4b8d-4b2c-b9d3-6c168cb6c847",
    "created_at": "2025-04-17T10:15:00Z",
    "updated_at": "2025-04-17T10:15:00Z"
  }
}

```

Response Body (Failed):

```
{
  "message": "Validation failed",
  "errorCode": 3003,
  "statusCode": 400,
  "errors": {
    "email": "All field must be filled"
  }
}
```

## UPDATE

Endpoint: PUT /api/v1/auth/store/:id

Request Body:

```
{
  "name": "Toko Mas Amba Solo",
  "address": "Jl. Gatot Subroto No. 99"
}
```

Response Body (Success) :

```
{
  "status": "success",
  "data": {
    "id": "b9a60fa2-5ad1-4d9c-b7c7-6b67fd13f880",
    "name": "Toko Mas Amba Solo",
    "email": "cabang1@store.com",
    "address": "Jl. Gatot Subroto No. 99",
    "phone": "081234567890",
    "store_type": "BRANCH",
    "parent_store_id": "a1f63e45-4b8d-4b2c-b9d3-6c168cb6c847",
    "created_at": "2025-04-17T10:15:00Z",
    "updated_at": "2025-04-17T11:00:00Z"
  }
}

```

Response Body (Failed):

```
{
  "message": "Store not found",
  "errorCode": 3004,
  "statusCode": 404
}
```

## DELETE

Endpoint: DELETE /api/v1/auth/store/:id

Response Body (Success) :

```
{
  "status": "success",
  "message": "Store deleted successfully"
}
```

Response Body (Failed):

```
{
  "message": "Store not found",
  "errorCode": 3004,
  "statusCode": 404
}
```
