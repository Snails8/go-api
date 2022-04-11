DC := docker-compose exec app
a := $1
sed := $2
SED := $3
# ====================================================================
# docker command
# ====================================================================
up:
	docker-compose up --build

stop:
	docker-compose stop

restart:
	docker-compose down
	docker-compose up -d

down:
	docker-compose down

destroy:
	docker-compose down --rmi all --volumes

app:
	${DC} /bin/bash