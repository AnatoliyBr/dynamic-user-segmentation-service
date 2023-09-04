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

## Примеры запросов (Linux)
* [Создание сегмента](#создание-сегмента)
* [Удаление сегмента](#удаление-сегмента)
* [Добавление/удаление пользователя в сегмент](#добавлениеудаление-пользователя-в-сегмент)
* [Просмотр активных сегментов пользователя](#просмотр-активных-сегментов-пользователя)


### Создание сегмента
Создание нового сегмента:

```bash
curl -X POST http://localhost:8080/seg -d "{"slug":"AVITO_DISCOUNT_30"}"
```

Пример ответа:

```bash
{"seg_id":1,"slug":"AVITO_DISCOUNT_30"}
```

### Удаление сегмента
Удаление сегмента:

```bash
curl -X DELETE http://localhost:8080/seg -d "{"slug":"AVITO_DISCOUNT_30"}"
```

Пример ответа:

```bash
{"delete segment":"AVITO_DISCOUNT_30"}
```

### Добавление/удаление пользователя в сегмент
Добавление/удаление пользователя в сегмент:

```bash
curl -X PUT http://localhost:8080/seg -d "{"slug_list_add":["AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS"], "slug_list_del":["AVITO_DISCOUNT_30", "AVITO_DISCOUNT_40", "AVITO_DISCOUNT_50", "AVITO_DISCOUNT_60"], "user_id": 1}"
```

Пример ответа:

```bash
{"add segments":["AVITO_VOICE_MESSAGES","AVITO_PERFORMANCE_VAS"],"delete segments":["AVITO_DISCOUNT_30","AVITO_DISCOUNT_40","AVITO_DISCOUNT_50","AVITO_DISCOUNT_60"],"user_id":["1"]}
```

### Просмотр активных сегментов пользователя
Просмотр активных сегментов пользователя:

```bash
curl -X GET http://localhost:8080/seg -d "{"user_id": 1}"
```

Пример ответа:

```bash
[{"seg_id":1,"slug":"AVITO_VOICE_MESSAGES"}]
```

### Linux
```bash
curl "http://localhost:8080/hello"
{"test":"hello"}

curl -X POST http://localhost:8080/seg -d "{"slug":"AVITO_DISCOUNT_30"}"
{"seg_id":1,"slug":"AVITO_DISCOUNT_30"}

curl -X DELETE http://localhost:8080/seg -d "{"slug":"AVITO_DISCOUNT_30"}"
{"delete segment":"AVITO_DISCOUNT_30"}

curl -X PUT http://localhost:8080/seg -d "{"slug_list_add":["AVITO_DISCOUNT_30"], "slug_list_del":[], "user_id": 1}"
{"add segments":["AVITO_DISCOUNT_30"],"delete segments":[],"user_id":["1"]}

curl -X GET http://localhost:8080/seg -d "{"user_id": 1}"
[{"seg_id":1,"slug":"AVITO_DISCOUNT_30"}]

curl -X PUT http://localhost:8080/seg -d "{"slug_list_add":[], "slug_list_del":["AVITO_DISCOUNT_30"], "user_id": 1}"
{"add segments":[],"delete segments":["AVITO_DISCOUNT_30"],"user_id":["1"]}
```

<details>
    <summary> Более сложный пример</summary>

```bash
curl -X POST http://localhost:8080/seg -d "{"slug":"AVITO_DISCOUNT_30"}"
{"seg_id":1,"slug":"AVITO_DISCOUNT_30"}

curl -X POST http://localhost:8080/seg -d "{"slug":"AVITO_DISCOUNT_40"}"
{"seg_id":2,"slug":"AVITO_DISCOUNT_40"}

curl -X POST http://localhost:8080/seg -d "{"slug":"AVITO_DISCOUNT_50"}"
{"seg_id":3,"slug":"AVITO_DISCOUNT_50"}

curl -X POST http://localhost:8080/seg -d "{"slug":"AVITO_DISCOUNT_60"}"
{"seg_id":4,"slug":"AVITO_DISCOUNT_60"}

curl -X PUT http://localhost:8080/seg -d "{"slug_list_add":["AVITO_DISCOUNT_30", "AVITO_DISCOUNT_40", "AVITO_DISCOUNT_50", "AVITO_DISCOUNT_60"], "slug_list_del":[], "user_id": 1}"
{"add segments":["AVITO_DISCOUNT_30","AVITO_DISCOUNT_40","AVITO_DISCOUNT_50","AVITO_DISCOUNT_60"],"delete segments":[],"user_id":["1"]}

curl -X POST http://localhost:8080/seg -d "{"slug":"TEST"}"
{"seg_id":30,"slug":"TEST"}

// Одновременное добавление и удаление сегментов

curl -X PUT http://localhost:8080/seg -d "{"slug_list_add":["TEST"], "slug_list_del":["AVITO_DISCOUNT_30", "AVITO_DISCOUNT_40", "AVITO_DISCOUNT_50", "AVITO_DISCOUNT_60"], "user_id": 1}"
{"add segments":["TEST"],"delete segments":["AVITO_DISCOUNT_30","AVITO_DISCOUNT_40","AVITO_DISCOUNT_50","AVITO_DISCOUNT_60"],"user_id":["1"]}
```
</details>

### Windows
```bash
curl "http://localhost:8080/hello"
{"test":"hello"}

curl -X POST http://localhost:8080/seg -d "{\"slug\":\"AVITO_DISCOUNT_30\"}"
{"seg_id":1,"slug":"AVITO_DISCOUNT_30"}

curl -X DELETE http://localhost:8080/seg -d "{\"slug\":\"AVITO_DISCOUNT_30\"}"
{"delete segment":"AVITO_DISCOUNT_30"}

curl -X PUT http://localhost:8080/seg -d "{\"slug_list_add\":[\"AVITO_DISCOUNT_30\"], \"slug_list_del\":[], \"user_id\": 1}"
{"add segments":["AVITO_DISCOUNT_30"],"delete segments":[],"user_id":["1"]}

curl -X GET http://localhost:8080/seg -d "{\"user_id\": 1}"
[{"seg_id":1,"slug":"AVITO_DISCOUNT_30"}]

curl -X PUT http://localhost:8080/seg -d "{\"slug_list_add\":[], \"slug_list_del\":[\"AVITO_DISCOUNT_30\"], \"user_id\": 1}"
{"add segments":[],"delete segments":["AVITO_DISCOUNT_30"],"user_id":["1"]}
```

<details>
    <summary> Более сложный пример</summary>

```bash
curl -X POST http://localhost:8080/seg -d "{\"slug\":\"AVITO_DISCOUNT_30\"}"
{"seg_id":1,"slug":"AVITO_DISCOUNT_30"}

curl -X POST http://localhost:8080/seg -d "{\"slug\":\"AVITO_DISCOUNT_40\"}"
{"seg_id":2,"slug":"AVITO_DISCOUNT_40"}

curl -X POST http://localhost:8080/seg -d "{\"slug\":\"AVITO_DISCOUNT_50\"}"
{"seg_id":3,"slug":"AVITO_DISCOUNT_50"}

curl -X POST http://localhost:8080/seg -d "{\"slug\":\"AVITO_DISCOUNT_60\"}"
{"seg_id":4,"slug":"AVITO_DISCOUNT_60"}

curl -X PUT http://localhost:8080/seg -d "{\"slug_list_add\":[\"AVITO_DISCOUNT_30\", \"AVITO_DISCOUNT_40\", \"AVITO_DISCOUNT_50\", \"AVITO_DISCOUNT_60\"], \"slug_list_del\":[], \"user_id\": 1}"
{"add segments":["AVITO_DISCOUNT_30","AVITO_DISCOUNT_40","AVITO_DISCOUNT_50","AVITO_DISCOUNT_60"],"delete segments":[],"user_id":["1"]}

curl -X POST http://localhost:8080/seg -d "{\"slug\":\"TEST\"}"
{"seg_id":30,"slug":"TEST"}

// Одновременное добавление и удаление сегментов

curl -X PUT http://localhost:8080/seg -d "{\"slug_list_add\":[\"TEST\"], \"slug_list_del\":[\"AVITO_DISCOUNT_30\", \"AVITO_DISCOUNT_40\", \"AVITO_DISCOUNT_50\", \"AVITO_DISCOUNT_60\"], \"user_id\": 1}"
{"add segments":["TEST"],"delete segments":["AVITO_DISCOUNT_30","AVITO_DISCOUNT_40","AVITO_DISCOUNT_50","AVITO_DISCOUNT_60"],"user_id":["1"]}
```
</details>

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