# GoRestAPI

Простой REST API на Go с использованием фреймворка [Echo](https://echo.labstack.com/) и ORM [GORM](https://gorm.io/), подключённый к базе данных PostgreSQL.

## 📦 Функциональность

API позволяет:

- 📥 Создавать сообщения (`POST /messages`)
- 📄 Получать список всех сообщений (`GET /messages`)
- ✏️ Обновлять текст сообщения по ID (`PATCH /messages/:id`)
- ❌ Удалять сообщение по ID (`DELETE /messages/:id`)

---

## ⚙️ Используемые технологии

- Go 1.20+
- Echo v4
- GORM
- PostgreSQL (контейнер на порту 5433)

---

## 🛠 Установка и запуск

### 1. Клонируй репозиторий
```bash
git clone https://github.com/the-student-of-school-366/GoRestAPI.git
cd GoRestAPI
