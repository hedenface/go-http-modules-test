.PHONY: all clean server module run

all:
	make server
	make module

run: clean all
	./server/server -config config/server.cfg

server:
	go build -o server/server server/server.go

module:
	go build -buildmode=plugin -o module/module.so module/module.go

clean:
	rm server/server
	rm module/module.so