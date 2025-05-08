APP_NAME=renkdegistir
GOFILES=main.go

build:
	go build -o $(APP_NAME) $(GOFILES)

clean:
	rm -f $(APP_NAME)

run: build
	./$(APP_NAME)

ms:
	GOOS=windows GOARCH=amd64 go build -o $(APP_NAME).exe $(GOFILES)