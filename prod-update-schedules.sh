docker-compose -f docker-compose.prod.yml exec backend go run cmd/schedule/download/main.go --url=https://crossfit1905.com/app-schedules.php
docker-compose -f docker-compose.prod.yml exec backend go run cmd/schedule/database/main.go --file=schedulesRepo.json