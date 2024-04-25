# Makefile
# make -B order 强制重新构建

env:
	docker-compose up -d
	docker start kibana	
	docker start elasticsearch

build:
	go build -o test
# 运行程序
run:
	./test
# 将编译和运行放在一起的命令
go:env build run

# 清理生成的目标文件
clean:
	rm -f test
