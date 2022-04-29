# Sport Clubs API
## Task 1
Необходимо подготовить консольную команду, цель которой - получать JSON с данными спортивных комплексов.

Суть работы команды:

1) Используя метод GET получить JSON [данные](https://crossfit1905.com/index-app.php)

Метод возвращает список спортивных комплексов. Каждый комплекс состоит из названия (title), даты проведения (scheduledAt) и описания (description)

2) Вывести в консоль название комплекса (title) и его дату (scheduledAt).

## Task 2
Используя данные из JSON, необходимо записать данные спортивных комплексов в БД

Для этого необходимо:

1) Создать в БД (PostgreSQL) таблицу workout_day

Поля таблицы:

- ID (GUID)
- title
- description
- scheduled_at (timestamp, дата проведения комплекса)
- created_at (дата создания комплекса)

2) Создать ORM модель для сущности

3) В консольной команде записывать комплексы в БД, используя созданную ORM модель

При повторном запуске команды комплекс не должен записываться в БД, если такой комплекс уже существует (проверка по title + scheduledAt)
## Task 3
Реализовать API метод GET /schedules, который должен возвращать информацию в формате:
Поля trainer, duration - пока харкодим( будем дорабатывать в рамках других задач)

## Console commands
Скачивание данных комплексов:
`go run cmd/complex/download/main.go --url=https://crossfit1905.com/index-app.php`

Пример запуска через docker-compose:
`docker-compose -f docker-compose.prod.yml exec backend go run cmd/complex/download/main.go --url=https://crossfit1905.com/index-app.php`

Запись комплексов в БД:
`go run cmd/complex/database/main.go --file=complexes.json`

Запуск сервера
```
$ source set_env.sh  
$ go run cmd/server/main.go
```

Применение миграций через docker-compose:
`docker-compose -f docker-compose.prod.yml up migrate-up`