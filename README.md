# 🚀 URL Shortener (на Go + PostgreSQL)

Простой и быстрый сервис сокращения ссылок, написанный на Go с использованием `chi`, `slog`, `PostgreSQL`, покрытием юнит- и интеграционными тестами, и удобным запуском через Docker.

---

## 📦 Возможности

- Сокращение ссылок (с кастомным alias)
- Редирект по alias
- Удаление сохранённой ссылки
- Поддержка Basic Auth
- Хранение ссылок в PostgreSQL
- Запуск и миграции через Docker
- Юнит- и интеграционные тесты

---

## 🛠 Стек

- Go 1.22+
- PostgreSQL
- Chi (router)
- Docker
- slog (structured logging)
- Testify + Httpexpect

---

## ⚙️ Установка и запуск

### 🔧 1. Клонируй репозиторий

```bash
git clone https://github.com/gerneram/url-shortener.git
```

### 🧪 2. Создай `.env` файл

Создай файл `local.env` в корне проекта:

```env
CONFIG_PATH=./config/local.yaml
HTTP_SERVER_PASSWORD=admin

POSTGRES_USER=admin
POSTGRES_PASSWORD=admin
POSTGRES_DB=db_shortner_url
```

### 🐳 3. Запусти через Docker

```bash
docker run --name postgres-db -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=db_shortner_url -p 5432:5432 -d postgres:latest
```

Это поднимет:

- PostgreSQL на `localhost:5432`

---

### 🚀Запусти сервер
```bash
go mod tidy  
go run cmd/url-shortener/main.go 
```
---
## 📮 Примеры запросов

### 🔗 Сократить ссылку

```bash
curl -X POST http://localhost:8082/url -u admin:admin -H "Content-Type: application/json" -d '{"url": "https://google.com", "alias": "ggl"}'
```

### 🔁 Редирект

Открой в браузере:

```
http://localhost:8082/ggl
```

### ❌ Удалить ссылку

```bash
curl -X DELETE http://localhost:8082/url/ggl -u admin:admin
```
