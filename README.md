## Сервис динамического сегментирования пользователей
Это решение тестового задания для стажёра Backend в AvitoTech. Сценарии использования, требования и детали по заданию в [репозитории](https://github.com/avito-tech/backend-trainee-assignment-2023/).

## Задача
Требовалось реализовать сервис, хранящий **пользователя** и **сегменты**, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент).

Сервис предоставляет **HTTP API** с форматом **JSON** как при отправке запроса, так и при получении результата.

## Структура HTTP API
**Публичные endpoint'ы**, доступные всем:

```
POST /seg - создание сегмента
DELETE /seg - удаление сегмента
PUT /seg - добавление/удаление пользователя в сегмент
GET /seg - просмотр активных сегментов пользователя
```

## Схема базы данных

<p align="center">
    <img src="/assets/images/db_schema.png" width="800">
</p>

## Архитектура

<p align="center">
    <img src="/assets/images/architecture.png" width="800">
</p>

## Структура проекта
```
├── cmd
|   ├── app
|
├── configs
├── internal
|   ├── app
|   ├── controller
|   ├── entity
|   ├── repository
|   ├── usecase
|
├── migrations
├── .env
├── .gitignore
├── Makefile
├── README.md
├── go.mod
├── go.sum
```

## Запуск и отладка
Все команды, используемые в процессе разработки и тестирования, фиксировались в `Makefile`.

## Примеры curl-запросов

```bash
curl "http://localhost:8080/hello"
```

## Миграции БД

```
CREATE USER dev WITH PASSWORD 'qwerty';
CREATE DATABASE user_seg_app_dev;

migrate create -ext sql -dir migrations create_segments
migrate create -ext sql -dir migrations create_users_with_segments

migrate -path migrations -database "postgres://localhost/user_seg_app_dev?sslmode=disable&user=dev&password=qwerty" up

\c user_seg_app_dev
\d segments
\d users_with_segments

CREATE DATABASE user_seg_app_test;

migrate -path migrations -database "postgres://localhost/user_seg_app_test?sslmode=disable&user=dev&password=qwerty" up
```