# User API Specs

## GET Single

Endpoint: GET /api/v1/auth/user/:id

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "id": "a3f01db2-7e1a-42ad-bb6a-728c67db7bfe",
    "name": "John Doe",
    "email": "johndoe@example.com",
    "role": "ADMIN",
    "account_status": "ACTIVE",
    "created_at": "2024-08-17T10:20:30Z",
    "updated_at": "2024-08-17T10:20:30Z"
  }
}
```

Response Body (Failed):

```
{
  "message": "User not found",
  "errorCode": 4001,
  "statusCode": 404
}
```

## GET Users by Store

Endpoint: GET /api/v1/auth/users/store/:id?limit=&page=

Query Parameters:

* Page: `<page number>`
* Limit: `<limit per page>`

Response Body (Success):

```
{
  "status": "success",
  "data": [
    {
      "id": "cd6b7cde-45f2-4b6a-97f6-fb49a2b9e7b7",
      "name": "John Doe",
      "email": "john@example.com",
      "role": "Supervisor",
      "account_status": "ACTIVE",
      "created_at": "2024-08-01T10:00:00Z",
      "updated_at": "2024-08-10T08:15:00Z"
    },
    {
      "id": "e2a1de47-ef28-4a93-97e2-d4d4341e8f88",
      "name": "Jane Smith",
      "email": "jane@example.com",
      "role": "ADMIN",
      "account_status": "INACTIVE",
      "created_at": "2024-08-02T14:20:00Z",
      "updated_at": "2024-08-09T12:30:00Z"
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
  "message": "Store not found",
  "errorCode": 4001,
  "statusCode": 404
}
```

## GET All Users in All Stores

GET /api/v1/auth/users/by-store-email?email=store@example.com&limit=&page=

Query Parameters:

* Email: `<store email>`
* Page: `<page number>`
* Limit: `<limit per page>`

Response Body (Success):

```
{
  "status": "success",
  "data": [
    {
      "id": "cd6b7cde-45f2-4b6a-97f6-fb49a2b9e7b7",
      "name": "John Doe",
      "email": "john@branch.com",
      "role": "OWNER",
      "account_status": "ACTIVE",
      "store": {
        "id": "uuid",
        "name": "Branch A",
        "store_type": "CENTRAL"
      },
      "created_at": "2025-04-17T10:00:00Z"
      "updated_at": "2024-08-09T12:30:00Z"
    },
    {
      "id": "cd6b7cde-45f2-4b6a-97f6-fb49a2b9e7b7",
      "name": "John Doe",
      "email": "john@branch.com",
      "role": "SUPERVISOR",
      "account_status": "ACTIVE",
      "store": {
        "id": "uuid",
        "name": "Branch A",
        "store_type": "BRANCH"
      },
      "created_at": "2025-04-17T10:00:00Z"
      "updated_at": "2024-08-09T12:30:00Z"
    },
    // ... more users
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
  "message": "Store not found",
  "errorCode": 3001,
  "statusCode": 404
}
```

## CREATE

Endpoint: POST /api/v1/auth/user

Request Body:

```
{
  "store_id": "cd6b7cde-45f2-4b6a-97f6-fb49a2b9e7b7",
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securePassword123",
  "role": "ADMIN"
}
```

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "id": "cd6b7cde-45f2-4b6a-97f6-fb49a2b9e7b7",
    "name": "John Doe",
    "email": "john@example.com",
    "role": "ADMIN",
    "account_status": "ACTIVE",
    "created_at": "2025-04-17T12:34:56Z",
    "updated_at": "2024-08-17T11:02:16.345Z"
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
    "email": "Email is required",
    "name": "Name is required"
  }
}
```

## UPDATE

Endpoint: PUT /api/v1/auth/user/:id

Request Body:

```
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "role": "ADMIN",        
  "account_status": "ACTIVE",   
  "store_id": "store-uuid-here"
}
```

Response Body (Success):

```
{
  "status": "success",
  "data": {
    "id": "cd6b7cde-45f2-4b6a-97f6-fb49a2b9e7b7",
    "name": "John Doe",
    "email": "john.doe@example.com",
    "role": "ADMIN",
    "account_status": "ACTIVE",
    "created_at": "2024-08-01T14:25:43.511Z",
    "updated_at": "2024-08-17T11:02:16.345Z"
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
    "email": "Email must be valid"
  }
}
```

## DELETE

Endpoint: DELETE /api/v1/auth/user/:id

Response Body (Success):

```
{
  "status": "success",
  "message": "User deleted successfully"
}
```

Response Body (Failed):

```
{
  "message": "User not found",
  "errorCode": 4001,
  "statusCode": 404
}
```
