build:
	@echo "Compilando..."
	go build -o bin/netpulse main.go

run:
	go run main.go check --file websites.csv

test:
	go test ./... -v

clean:
	rm -rf bin/