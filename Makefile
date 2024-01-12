build:
	docker build -t cowrvalera/go-qr-generator:1.0.0 --platform linux/amd64 .

run:
	docker run -d -p 8080:8080 --platform linux/amd64 cowrvalera/go-qr-generator:1.0.0

	
run_dev:
	docker-compose up -d

stop_dev:
	docker-compose down