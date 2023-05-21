# Notes API

Простое CRUD-API, с возможностью оставлять простые заметки с некоторым текстом.
## Запуск🚀
Перед запуском приложения, создайте в корневой директории файл ```.env```, и заполните поля:
```bash
PGUSER=
PGPASSWORD=
PGHOST=localhost
PGPORT=5432
PGDATABASE=
PGSSLMODE=disable

PORT=:8080
LOG_FILE_PATH=logs.log
```
Также перед запуском нужно выполнить миграцию для таблицы с заметками.
Можно использовать эту [библиотеку ](https://github.com/golang-migrate/migrate?ysclid=lhxtjgz16c765222512)(также доступна как docker-container).
