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
	cd backend && \
	make up 

down:
	cd backend && \
	make down

restart:
	cd backend && \
	make restart

destroy:
	cd backend && \
	make destroy

app:
	cd backend && \
	make app

# ====================================================================
# genapi
# ====================================================================

# ====================================================================
# setup
# ====================================================================
setup:
	cd backend && \
	cp .env.exmaple .env
	docker-compose -f docker-compose.yaml -p ${APP_NAME} up -d --build
	make migrate
	make seed
	echo "########################################################"
	echo "                     Enjoy Go Life                      "
	echo "######################## finish ########################"

