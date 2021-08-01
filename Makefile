all:
	make main
	make module

run: clean all
	./helloWorld

main:
	go build -o helloWorld main.go

module:
	go build -buildmode=plugin -o module/module.so module/module.go

clean:
	rm helloWorld
	rm module/module.so