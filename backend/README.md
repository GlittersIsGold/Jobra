# Бэкенд для фриланс-платформы Джобра

## Технологии

```sh
golang 1.25.4
gin 1.11.0
validator 10
sqlc
pg
```

## Конфигурации

Проект поддерживает три варианта конфигурации для запуска: флаги, переменные окружения и .env файлы

### Переменные .env

```sh
PORT=8080
ENVIRONMENT=debug

DATABASE_URL=postgresql://username:password@localhost:5432/jobra?sslmode=disable

JWT_SECRET=your-super-secret-jwt-key-change-in-production
JWT_EXPIRY_HOURS=24

TRUSTED_PROXIES=http://localhost:3000
```

### Флаги

```sh
configuration    # способ объявления конфигурации: командная строка(cli), переменные ОС(os-vars) или файл .env(env)
port             # порт на котором запускается приложение, по-умолчанию 8000
environment      # окружение в котором запускается приложение: debug, release или test
db-url           # строка подключения к БД
jwt-secret       # секрет для jwt токена, должен содержать не менее 32 символов
jwt-expire-hours # время жизни токена
trusted-proxies  # допустимые адреса для запроса на бэкенд, передаётся строкой через запятую. Cм. https://gin-gonic.com/en/docs/deployment/#dont-trust-all-proxies
```

### Переменные ОС

```sh
CONFIGURATION
PORT
ENVIRONMENT
DATABASE_URL
JWT_SECRET
JWT_EXPIRE_HOURS
TRUSTED_PROXIES
```

Для запуска нужно задать соответствующий флаг configuration c параметром на выбор: cli, os-vars или env

```go
go run main.go -configuration=cli 
```

Если параметры конфигурации не переданы или не читаются, сервер запускается с параметрами по-умолчанию

Подробнее смотри в пакете `config`

## Запуск

```go
go run backend/cmd/api/main.go
```
