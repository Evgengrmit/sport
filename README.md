## Работа с миграциями (локально)
Применить все неприменённые миграции:
`migrate -path=./migration -database="postgres://postgres:qwerty@127.0.0.1:5436/postgres?sslmode=disable" up`

Создать новую миграцию:
` migrate -path=./migration -database="postgres://postgres:qwerty@127.0.0.1:5436/postgres?sslmode=disable" create -dir=./migration -ext=sql some_new_table`

Откатить 1 последнюю миграцию:
` migrate -path=./migration -database="postgres://postgres:qwerty@127.0.0.1:5436/postgres?sslmode=disable" down 1`

Откатить все миграции:
` migrate -path=./migration -database="postgres://postgres:qwerty@127.0.0.1:5436/postgres?sslmode=disable" down`


## Console commands
Скачивание данных комплексов:
`go run cmd/complex/download/main.go --url=https://crossfit1905.com/index-app.php`

Запись комплексов в БД:
`go run cmd/complex/database/main.go --file=complexes.json`

Скачивание расписания:
`go run cmd/schedule/download/main.go --url=https://crossfit1905.com/app-schedules.php`

Запись расписания в БД:
`go run cmd/schedule/database/main.go --file=schedulesRepo.json`

Запуск сервера
```
$ source set_env.sh  
$ go run cmd/server/main.go
```

## Команды для продакшн окружения

Применение миграций через docker-compose:
`docker-compose -f docker-compose.prod.yml up migrate-up`

Скачивание комплексов:
`docker-compose -f docker-compose.prod.yml exec backend go run cmd/complex/download/main.go --url=https://crossfit1905.com/index-app.php`
Запись комплексов в БД:
`docker-compose -f docker-compose.prod.yml exec backend go run cmd/complex/database/main.go --file=workoutDays.json`

Скачивание расписания:
`docker-compose -f docker-compose.prod.yml exec backend go run cmd/schedule/download/main.go --url=https://crossfit1905.com/app-schedules.php`
Запись расписания:
`docker-compose -f docker-compose.prod.yml exec backend go run cmd/schedule/database/main.go --file=schedulesRepo.json`