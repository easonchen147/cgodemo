BINARY="cgodemo"

build:
	rm -rf cgodemo.o
	rm -rf libcgodemo.a
	gcc -Wall -c cgodemo.c
	ar -rv libcgodemo.a cgodemo.o
	go build -o cgodemo

run:
	./${BINARY}

help:
	@echo "make build 编译静态链接库"
	@echo "make run 运行cgodemo"