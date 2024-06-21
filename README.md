# Task Manager API

## Запуск🚀
Перед запуском приложения, создайте в корневой директории файл ```.env```, и заполните поля:
```bash
PGUSER=user
PGPASSWORD=password
PGHOST=localhost
PGPORT=5432
PGDATABASE=db
PGSSLMODE=disable

ENVIRONMENT=debug
HTTP_PORT=8080
LOG_FILE_PATH=logs.log
GIN_MODE=release
```

Дальше для запуска запустить команду:
```bash
docker compose up --build
```

