build:
	GOOS=linux go build -o app
	docker build -t contd/say .
	rm -f app