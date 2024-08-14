TARGET = main

default: $(TARGET)
$(TARGET):
	@go build -o bin/$(TARGET) cmd/main.go

build: clean 
	@templ generate
	@go build -o bin/$(TARGET) cmd/main.go

run: build
	@./bin/$(TARGET)

test: clean default
	@go test ./...

sec: clean default
	@gosec ./...

fmt: clean default
	@go fmt ./...

clean:
	@rm -f ./main

templ:
	@templ generate

dep:
	@go mod init github.com/SisyphianLiger/dream_mail
	@go get github.com/securego/gosec/v2/cmd/gosec
	@go get github.com/a-h/templ
	@go get github.com/labstack/echo/v4
	@go get github.com/mailgun/mailgun-go/v4
	@go get github.com/SparkPost/gosparkpost

startup: 
	@rm -f go.mod
	@rm -f go.sum
	@make dep
	@make run
	


