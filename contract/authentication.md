# Authentication API Specs

## REGISTER (Transaction)

Endpoint: POST /api/v1/auth/register

Request Body:

```
{
  "store_name": "Toko Mas Amba",
  "store_address": "Jl. Ngawi-Jomokerto No.121",
  "store_phone": "08123456789",
  "store_email": "tomba771@gmail.com",
  "user_name": "Mas Amba",
  "user_email": "Ambasuke232@gmail.com",
  "user_password": "rahasiBGT"
}
```

Response Body (Success) :

```
{
  "status": "success",
  "data": {
    "store": {
       "id": "2cf3a1bc-91ad-4ef4-9505-123456789abc",
       "name": "Toko Mas Amba",
       "email": "tomba771@gmail.com",
       "address": "Jl. Ngawi-Jomokerto No.121",
       "phone": "08123456789",
       "store_type": "CENTRAL",
       "created_at": "2025-03-18T11:01:10Z",
       "updated_at": "2025-03-18T11:01:10Z"
    },
    "user": {
      "id": "dc3f9ac0-34e2-4c02-8c4f-123456789abc",
      "name": "Mas Amba",
      "email": "Ambasuke232@gmail.com",
      "role": "OWNER",
      "account_status": "ACTIVE",
      "created_at": "2025-03-18T11:01:10Z",
      "updated_at": "2025-03-18T11:01:10Z"
    }
  }
}

```

Response Body (Failed):

```
{
  "message": "Email already registered",
  "errorCode": 2001,
  "statusCode": 400,
  "errors": {
    "admin_email": "email already taken"
  }
}

```

## LOGIN

Endpoint: POST /api/v1/auth/login

Request Body:

```
{
  "email": "admin1@gmail.com",
  "password": "12345678"
}

```

Response Body (Success) :

```
{
  "status": "success",
  "data": {
    "id": "dc3f9ac0-34e2-4c02-8c4f-123456789abc",
    "name": "Mas Amba",
    "email": "ambasuke232@gmail.com",
    "role": "OWNER",
    "account_status": "ACTIVE",
    "last_login": "2025-04-17T12:15:30Z"
  }
}
```

Response Body (Failed):

```
{
  "message": "Invalid email or password",
  "errorCode": 2002,
  "statusCode": 401
}
```

## LOGOUT

Endpoint: POST /api/v1/auth/logout

Response Body (Success):

```
{
  "status": "success",
  "message": "Logout successful"
}

```

Response Body (Failed):

```
{
  "message": "Unauthorized",
  "errorCode": 1003,
  "statusCode": 401
}
```
