# PR Reviewer Assignment Service

Сервис назначения ревьюеров на Pull Request’ы внутри команды.  
Поддерживает:
- Автоматическое назначение до двух ревьюверов на PR;
- Переназначение ревьюверов среди активных участников команды;
- Получение списка PR, назначенных конкретному пользователю;
- Управление командами и активностью пользователей;
- Идемпотентное слияние PR.

---

## Требования

- Go >= 1.24
- PostgreSQL >= 15
- Docker & Docker Compose

---

## Запуск

1. Перейти в корень проекта:
```bash
cd <repo_folder>

docker compose up --build

```


Конфигурация

Все переменные окружения находятся в .env файле
```azure
PORT_SERVER=8080
DB_PASSWORD=qwerty
DB_USERNAME=postgres
HOST=db
PORT=5432
DB_NAME=avito
SSLMODE=disable

```