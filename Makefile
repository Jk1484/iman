proto:
	./scripts/gen_proto.sh
up:
	docker compose up --build -d
stop:
	docker compose stop
start:
	docker compose start
down:
	docker compose down