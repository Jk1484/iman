proto:
	./scripts/gen_proto.sh
run:
	docker compose up --build -d
stop:
	docker compose down