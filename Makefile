APP=./cmd/app/main.go
ATTACK_URL="http://127.0.0.1:8080/api/orders/b563feb7b2b84b6test"


app:
	docker compose up -d                                                                                                                                                                    ✔  at 18:10:14 
	go run $(APP)

down:
	docker compose down


stress_test:
	echo "GET $(ATTACK_URL)" | vegeta attack -rate=50 -duration=30s | vegeta plot > plot.html