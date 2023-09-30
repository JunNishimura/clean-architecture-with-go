.PHONY: init up down build gen

init:
	cp compose.override.yml.sample compose.override.yml

up:
	docker compose up --build -d

down:
	docker compose down

build: 
	docker compose build

gen:
	sqlboiler mysql -c sqlboiler.toml -o entities -p entities --no-tests --wipe