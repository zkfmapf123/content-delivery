up:
	docker compose up --build -d

down:
	docker compose down

dup: down
	make up