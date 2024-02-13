# My Home Libray REST API Docs

## Join - `/api/auth/join`

### Request

### Response

## Login - `/api/auth/login`

### Request
```json
{
    "email": "me@hyunsang.dev",
    "password": "q1w2e3r4"
}
```

### Response
```json
{
    "status": {
        "code": 200,
        "message": "어서와요! 오늘도 열심히 책을 읽어봐요 :)"
    },
    "data": {
        "access_uuid": "1853acd3-e16a-412a-9717-15331ac2c1c3",
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjE4NTNhY2QzLWUxNmEtNDEyYS05NzE3LTE1MzMxYWMyYzFjMyIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTcwNzgwMTYyMiwidXNlcl9pZCI6ImFmNGUzOWUxLWRjMmQtNGU4NS1iZmUwLWNlY2JlOGI3OTM3YiJ9.NhvE4MGx4Kp-042Ray3sPSrkZeRDtMDdiRzJK_6Y_gI",
        "refresh_uuid": "4b941c10-26ef-4510-96a5-1bc17631d901",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDg0MDU1MjIsInJlZnJlc2hfdXVpZCI6IjRiOTQxYzEwLTI2ZWYtNDUxMC05NmE1LTFiYzE3NjMxZDkwMSIsInVzZXJfaWQiOiJhZjRlMzllMS1kYzJkLTRlODUtYmZlMC1jZWNiZThiNzkzN2IifQ.6ig5hDtHqLAesoR1ffcDv_1NI0YemNskVw6DFLnu-C0"
    },
    "repsonsed_at": "2024-02-13T14:05:22.049309+09:00"
}
```

### in Redis
![in Redis](./jwt-redis.png)

- Refresh (Token) UUID + UserUUID
- Access (Token) UUID + UserUUID