# Go Commerce API

## Response Codes

| Code | Description           |
|------|-----------------------|
| 200  | OK                    |
| 201  | Created               |
| 400  | Bad Request           |
| 401  | Unauthorized          |
| 403  | Forbidden             |
| 404  | Not Found             |
| 405  | Method Not Allowed    |
| 500  | Internal Server Error |

## Response Structure

- Success Response

```json
{
  "status": "success",
  "message": "Success Message",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "",
    "phone": ""
  }
}
```

- Error Response

```json
{
  "status": "error",
  "message": "Error Message",
  "error": {
    "code": 400
  }
}
```

### Respose Message Codes (http://localhost:3000/api/error-messages)

| Code | Description                               |
|------|-------------------------------------------|
| 1    | Token is expired                          |
| 2    | Token is invalid                          |
| 3    | Invalid token type, must be refresh token |
| 4    | User is not authorized                    |
| 5    | Invalid token type, must be basic token   |
| 401  | User is not authenticated                 |

### Authentication Endpoints

| Method | Endpoint                  | Description                |
|--------|---------------------------|----------------------------|
| POST   | `/api/auth/register`      | Register a new user.       |
| POST   | `/api/auth/token/refresh` | Generate a refresh token.  |
| POST   | `/api/auth/token/access`  | Generate an access token.  |
| POST   | `/api/password/forgot`    | Send password reset email. |
| POST   | `/api/password/change`    | Change user password.      |
| POST   | `/api/profile/update`     | Update user profile.       |

### User Endpoints

| Method | Endpoint                     | Description                 |
|--------|------------------------------|-----------------------------|
| POST   | `/api/user/:id`              | Get user by ID.             |
| POST   | `/api/user/create`           | Create a new user.          |
| POST   | `/api/user/update/:id`       | Update user by ID.          |
| GET    | `/api/user/delete/:id`       | Delete user by ID.          |
| GET    | `/api/user/delete/admin/:id` | Delete user by ID as admin. |

## Core Endpoints

| Method | Endpoint                 | Description                |
|--------|--------------------------|----------------------------|
| POST   | `/api/user/:id`          | Get user by ID.            |
| GET    | `/api/error-messages`    | Get error message codes.   |
| GET    | `/api/constant`          | Get constant values.       |
| GET    | `/api/province`          | Get all provinces.         |
| GET    | `/api/city/:province_id` | Get cities by province ID. |
| GET    | `/api/district/:city_id` | Get districts by city ID.  |
| GET    | `/api/currency`          | Get all currencies.        |



