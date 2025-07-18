# Go Tools

一个实用的Go命令行工具集合，提供多种常用功能。

## 功能特性

- 🔧 **JSON工具**: JSON格式化、验证和美化
- 🔐 **哈希计算**: 支持MD5、SHA1、SHA256哈希算法
- 📝 **版本信息**: 查看工具版本和构建信息
- ⚙️ **配置支持**: 支持YAML配置文件
- 📊 **详细日志**: 可开启详细输出模式

## 安装

```bash
# 克隆项目
git clone https://github.com/Yanzzp999/go-tools.git
cd go-tools

# 构建
go build -o go-tools .

# 安装到系统路径（可选）
go install .
```

## 使用方法

### 基础命令

```bash
# 显示帮助信息
./go-tools --help

# 显示版本信息
./go-tools version
```

### JSON工具

```bash
# 格式化JSON字符串
./go-tools json '{"name":"张三","age":25}'

# 从文件读取并格式化
cat data.json | ./go-tools json

# 保存格式化结果到文件
./go-tools json '{"name":"张三"}' -o output.json

# 压缩JSON（不美化）
./go-tools json '{"name": "张三", "age": 25}' --pretty=false
```

### 哈希计算

```bash
# 计算字符串的MD5哈希
./go-tools hash -s "hello world" -t md5

# 计算文件的SHA256哈希
./go-tools hash -f /path/to/file -t sha256

# 计算字符串的SHA1哈希
./go-tools hash -s "test" -t sha1
```

### 配置文件

可以在家目录创建 `.go-tools.yaml` 配置文件：

```yaml
# 默认设置
verbose: true
default_hash_type: "sha256"
```

### 命令行选项

```bash
# 开启详细输出
./go-tools --verbose [command]

# 使用自定义配置文件
./go-tools --config /path/to/config.yaml [command]
```

## 项目结构

```
go-tools/
├── cmd/                  # 命令定义
│   ├── root.go          # 根命令
│   ├── version.go       # 版本命令
│   ├── json.go          # JSON工具命令
│   └── hash.go          # 哈希计算命令
├── pkg/                 # 公共包
│   └── utils/           # 工具函数
│       └── hash.go      # 哈希工具
├── internal/            # 内部包
├── main.go              # 程序入口
├── go.mod               # Go模块定义
└── README.md            # 项目说明
```

## 依赖

- [Cobra](https://github.com/spf13/cobra) - 强大的命令行框架
- [Viper](https://github.com/spf13/viper) - 配置管理
- [Logrus](https://github.com/sirupsen/logrus) - 结构化日志

## 开发

```bash
# 安装依赖
go mod tidy

# 运行测试
go test ./...

# 格式化代码
go fmt ./...

# 静态检查
go vet ./...

# 构建
go build -o go-tools .
```

## 贡献

欢迎提交Issue和Pull Request！

## 许可证

MIT License