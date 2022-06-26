DC := docker-compose exec front-app
ARG := $1

up:
	docker-compose up -d --build
	make yarn


	
reinstall:
	@make destroy
	@make install

stop:
	docker-compose stop

restart:
	docker-compose down
	docker-compose up -d

down:
	docker-compose down

destroy:
	docker-compose down --rmi all --volumes

ps:
	docker-compose ps

front-app:
	docker-compose exec front-app /bin/ash

# =======================================================
# setup
# =======================================================
yarn:
	docker-compose exec front-app yarn start

build:
	docker-compose exec front-app yarn build

install:
	docker-compose exec front-app yarn
	docker-compose exec front-app npx create-react-app frontend --template typescript
	cp -R frontend/. ./
	rm -rf frontend
	docker-compose exec front-app yarn build
	docker-compose exec front-app yarn start	