.PHONY: all clean server module run

all:
	make server
	make module

run: clean all
	./main -config config/server.cfg

server:
	go build -o server/server server/server.go
	go build -o main main.go

module:
	go build -buildmode=plugin -o module/module.so module/module.go

clean:
	rm -f server/server
	rm -f server/main
	rm -f module/module.so