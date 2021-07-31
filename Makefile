all:
	make main
	make module

run: all
	./helloWorld

main:
	go build -o helloWorld main.go

module:
	go build -buildmode=plugin -o module/module.so module/module.go
