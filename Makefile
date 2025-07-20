# Go tools Makefile

# 变量定义
BINARY_NAME=go-tools
MAIN_FILE=main.go

# 默认目标
.PHONY: run clean build help

# 运行程序：清理 -> 编译 -> 运行
run: clean build
	@echo "运行程序..."
	./$(BINARY_NAME)

# 编译程序
build:
	@echo "编译程序..."
	go build -o $(BINARY_NAME) $(MAIN_FILE)

# 清理编译产物
clean:
	@echo "清理编译产物..."
	@if [ -f $(BINARY_NAME) ]; then rm $(BINARY_NAME); echo "已删除 $(BINARY_NAME)"; fi

# 显示帮助信息
help:
	@echo "可用的 make 命令："
	@echo "  run    - 清理、编译并运行程序"
	@echo "  build  - 编译程序"
	@echo "  clean  - 清理编译产物"
	@echo "  help   - 显示此帮助信息"
