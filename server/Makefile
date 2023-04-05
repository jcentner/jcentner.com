build:
	go vet
	go build -o stats-tracker

run:
	rm -f stats-tracker
	go vet
	go build -o stats-tracker
	./kill-stats-tracker
	./stats-tracker > server.log &

tidy:
	go mod tidy
