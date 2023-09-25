.PHONY: up down build gen

up:
	docker compose up --build -d

down:
	docker compose down

build: 
	docker compose build

gen:
	sqlboiler mysql -c sqlboiler.toml -o entities -p entities --no-tests --wipe