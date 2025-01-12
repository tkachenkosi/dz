TARGET = dist/main
SOURCES = *.go

.PHONY: build run lint clean
.DEFAULT_GOAL = run

build:
	@echo Компиляция программы!
	go build -o ${TARGET} main.go
${TARGET}: main.go
	@echo Компиляция программы вторым способом.
	go build -o ${TARGET} main.go
run:
	go run main.go
lint:
	@echo Запускаем $@
	go vet
	golangci-lint run
clean:
ifneq (,$(wildcard $(TARGET)))
	 @echo Удаляем файл $(TARGET)
	 rm -f ${TARGET}
endif
	go clean
