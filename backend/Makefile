DC := docker-compose exec app
APP_NAME := go-app
APP_TEST_NAME := gp-api-test
a := $1
sed := $2
SED := $3
# ====================================================================
# docker command
# ====================================================================
up:
	docker-compose -f docker-compose.yaml -p ${APP_NAME} up -d

down:
	docker compose -f docker-compose.yaml -p ${APP_NAME} down

restart:
	docker compose -f docker-compose.yaml -p ${APP_NAME} down
	docker-compose -f docker-compose.yaml -p ${APP_NAME} up -d --build

destroy:
	docker-compose down --rmi all --volumes

app:
	docker compose -p ${APP_NAME} exec app bash

# ====================================================================
# go command
# ====================================================================
rungo:
	docker compose -p ${APP_NAME} exec app bash -c 'go run main.go'

gomod:
	docker compose -p ${APP_NAME} exec app bash -c 'go mod init src/'

######################### seed ##############################
seed:
	docker compose -p ${APP_NAME} exec db psql \
	-h localhost \
	-U postgres \
	-d go_app \
	-v ON_ERROR_STOP=1 \
	-f /seeds/${SEED_ENV}/seeds.sql

testseed:
	docker compose -p ${APP_TEST_NAME} exec db psql \
	-h localhost \
	-U postgres \
	-d go_app \
	-v ON_ERROR_STOP=1 \
	-f /seeds/${SEED_ENV}/seeds.sql

######################### migration #########################
migratecreate:
	docker compose -p ${APP_NAME} exec app bash -c 'migrate create -ext sql -dir tools/migrations/ ${NAME}'

migrate:
	docker compose -p ${APP_NAME} exec app bash -c 'migrate -database="$$DB_DSN" -path=tools/migrations/ up'

migrateforce:
	docker compose -p ${APP_NAME} exec app bash -c 'migrate -database="$$DB_DSN" -path=tools/migrations/ force ${VERSION}'

migrateup:
	docker compose -p ${APP_NAME} exec app bash -c 'migrate -database="$$DB_DSN" -path=tools/migrations/ up 1'

migratedown:
	docker compose -p ${APP_NAME} exec app bash -c 'migrate -database="$$DB_DSN" -path=tools/migrations/ down 1'

######################### test migration ###################
testmigrate:
	docker compose -p ${APP_TEST_NAME} exec app bash -c 'migrate -database="$$DB_DSN" -path=tools/migrations/ up'

# ====================================================================
# genapi
# ====================================================================

# ====================================================================
# setup
# ====================================================================
setup:
	cp .env.exmaple .env
	docker-compose -f docker-compose.yaml -p ${APP_NAME} up -d --build
	make migrate
	make seed
	echo "########################################################"
	echo "                     Enjoy Go Life                      "
	echo "######################## finish ########################"

