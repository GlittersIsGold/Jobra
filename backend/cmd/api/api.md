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

## Обновление профиля

### Команда

```sh
curl -X PUT http://localhost:8080/api/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"username":"newusername","email":"new@example.com"}'
```

### Ответ сервера

```json
{
  "status": "success",
  "message": "Profile updated successfully",
  "data": {
    "id": 123,
    "name": "Новое Имя",
    "email": "new@example.com"
  }
}
```

## Обновление пароля

### Команда

```sh
curl -X POST http://localhost:8080/api/profile/password \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_token_here" \
  -d '{
    "currentPassword": "oldPassword123",
    "newPassword": "newPassword456"
  }'
```

### Ответ сервера

```json
{
  "status": "success",
  "message": "Password changed successfully"
}
```

## Создание заказа

### Команда

```sh
curl -X POST http://localhost:8080/api/orders \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_token_here" \
  -d '{
    "title": "Разработка веб-сайта",
    "description": "Нужно разработать корпоративный сайт с адаптивным дизайном",
    "budget": 50000,
    "category": "Разработка",
    "section": "Веб-разработка"
  }'
```

### Ответ сервера

```json
{
  "message": "Order created successfully",
  "order": {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "customer_id": "user-12345",
    "title": "Разработка веб-сайта",
    "description": "Нужно разработать корпоративный сайт с адаптивным дизайном",
    "budget": 50000,
    "category": "Разработка",
    "section": "Веб-разработка",
    "status": "active",
    "created_at": "2024-01-24T10:30:00Z",
    "updated_at": "2024-01-24T10:30:00Z"
  }
}
```

## Получение заказов пользователя

### Команда

```sh
curl -X GET http://localhost:8080/api/orders \
  -H "Authorization: Bearer your_token_here"
```

### Ответ сервера

```json
{
  "orders": [
    {
      "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
      "customer_id": "user-12345",
      "title": "Разработка веб-сайта",
      "description": "Нужно разработать корпоративный сайт с адаптивным дизайном",
      "budget": 50000,
      "category": "Разработка",
      "section": "Веб-разработка",
      "status": "active",
      "created_at": "2024-01-24T10:30:00Z",
      "updated_at": "2024-01-24T10:30:00Z"
    }
  ],
  "count": 1
}
```

## Получение всех заказов (для админа)

### Команда

```sh
curl -X GET http://localhost:8080/api/orders/all \
  -H "Authorization: Bearer your_token_here"
```

### Ответ сервера

```json
{
  "orders": [
    {
      "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
      "customer_id": "user-12345",
      "title": "Разработка веб-сайта",
      "description": "Нужно разработать корпоративный сайт с адаптивным дизайном",
      "budget": 50000,
      "category": "Разработка",
      "section": "Веб-разработка",
      "status": "active",
      "created_at": "2024-01-24T10:30:00Z",
      "updated_at": "2024-01-24T10:30:00Z"
    }
  ],
  "count": 1
}
```

## Получение конкретного заказа

### Команда

```sh
curl -X GET http://localhost:8080/api/orders/a1b2c3d4-e5f6-7890-abcd-ef1234567890 \
  -H "Authorization: Bearer your_token_here"
```

### Ответ сервера

```json
{
  "order": {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "customer_id": "user-12345",
    "title": "Разработка веб-сайта",
    "description": "Нужно разработать корпоративный сайт с адаптивным дизайном",
    "budget": 50000,
    "category": "Разработка",
    "section": "Веб-разработка",
    "status": "active",
    "created_at": "2024-01-24T10:30:00Z",
    "updated_at": "2024-01-24T10:30:00Z"
  }
}
```

## Обновление заказа

### Команда

```sh
curl -X PUT http://localhost:8080/api/orders/a1b2c3d4-e5f6-7890-abcd-ef1234567890 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_token_here" \
  -d '{
    "title": "Разработка веб-сайта (обновлено)",
    "description": "Нужно разработать корпоративный сайт с адаптивным дизайном и CMS",
    "budget": 60000,
    "category": "Разработка",
    "section": "Веб-разработка",
    "status": "active"
  }'
```

### Ответ сервера

```json
{
  "message": "Order updated successfully",
  "order": {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "customer_id": "user-12345",
    "title": "Разработка веб-сайта (обновлено)",
    "description": "Нужно разработать корпоративный сайт с адаптивным дизайном и CMS",
    "budget": 60000,
    "category": "Разработка",
    "section": "Веб-разработка",
    "status": "active",
    "created_at": "2024-01-24T10:30:00Z",
    "updated_at": "2024-01-24T11:45:00Z"
  }
}
```

## Удаление заказа

### Команда

```sh
curl -X DELETE http://localhost:8080/api/orders/a1b2c3d4-e5f6-7890-abcd-ef1234567890 \
  -H "Authorization: Bearer your_token_here"
```

### Ответ сервера

```json
{
  "message": "Order deleted successfully"
}
```

## Получение профиля пользователя

### Команда

```sh
curl -X GET http://localhost:8080/api/profile \
  -H "Authorization: Bearer your_token_here"
```

## Ответ сервера

```json
{
  "user": {
    "id": "user-12345",
    "username": "john_doe",
    "email": "john@example.com"
  }
}
```

## Обновление профиля пользователя

### Команда

```sh
curl -X PUT http://localhost:8080/api/profile \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_token_here" \
  -d '{
    "username": "john_doe_updated",
    "email": "john.updated@example.com"
  }'
```

### Ответ сервера

```json
{
  "message": "Profile updated successfully",
  "user": {
    "id": "user-12345",
    "username": "john_doe_updated",
    "email": "john.updated@example.com"
  }
}
```

## Получение всех пользователей

### Команда

```sh
curl -X GET http://localhost:8080/api/users \
  -H "Authorization: Bearer your_token_here"
```

### Ответ сервера

```json
{
  "users": [
    {
      "id": "user-12345",
      "username": "john_doe",
      "email": "john@example.com"
    }
  ],
  "count": 1
}
```
