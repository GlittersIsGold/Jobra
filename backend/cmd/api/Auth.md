# Команды curl для авторизации

## Регистрация

### Команда

```sh
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123"}'
```

### Ответ сервера

```json
{
    "message":"User registered successfully",
    "user":{
        "email":"test@example.com",
        "id":"user-\u0004",
        "username":"testuser"
    }
}
```

## Логин

### Команда

```sh
curl -X POST http://localhost:8080/login \     
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}'
```

### Ответ сервера

```json
{
    "message":"Login successful",
    "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoidXNlci1cdTAwMDQiLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiaXNzIjoiYmFja2VuZC1hcGkiLCJzdWIiOiJ1c2VyLVx1MDAwNCIsImV4cCI6MTc2MzU1NTIxMSwiaWF0IjoxNzYzNTU0OTExfQ.McudrRDmD1DGAXwSJcZYPNhlu9uZFqPiFgfCYoCxA1w","user":{
        "email":"test@example.com",
        "id":"user-\u0004",
        "username":"testuser"
    }
}
```

## Профиль

### Команда

```sh
curl -X GET http://localhost:8080/api/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

### Ответ сервера

```json
{
    "user": {
        "email":"test@example.com",
        "id":"user-\u0004",
        "username":"testuser"
    }
}
```
