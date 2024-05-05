# Example SMS Restful API

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

-   Success ResponseP

```json
{
    "status": true,
    "data": {
        "id": 1,
        "name": "John Doe",
        "email": "",
        "phone": ""
    }
}
```

-   Error Response

```json
{
    "status": false,
    "error": {
        "code": 400
    }
}
```

### Respose Message Codes

| Code | Description                               |
|------|-------------------------------------------|
| 1    | Token is expired                          |
| 2    | Token is invalid                          |
| 3    | Invalid token type, must be refresh token |
| 4    | User is not authorized                    |
| 5    | Invalid token type, must be basic token   |
| 401  | User is not authenticated                 |

### Authentication Endpoints

| Method | Endpoint                                    | Description               |
|--------|---------------------------------------------|---------------------------|
| POST   | `/api/v1/auth/token/generate/refresh-token` | Generate a refresh token. |
| POST   | `/api/v1/auth/token/generate/token`         | Generate an access token. |
| POST   | `/api/v1/auth/register`                     | Register a new user.      |

### SMS Endpoints

| Method | Endpoint                   | Description                             |
|--------|----------------------------|-----------------------------------------|
| POST   | `/api/v1/sms/send`         | Send an SMS.                            |
| GET    | `/api/v1/sms/report`       | Get SMS delivery reports.               |
| GET    | `/api/v1/sms/report/{sms}` | Get detailed report for a specific SMS. |

## Dependencies

| Package Name | Description      |
|--------------|------------------|
| JWT Auth     | firebase/php-jwt |