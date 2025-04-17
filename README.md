# 🧠 Go Todo API 项目 - 2
<div align="center">

![Go Test](https://github.com/Alec1120/go-api-2/actions/workflows/test.yml/badge.svg)

</div>

<div align="center">

中文 | [English](./README_en.md)

</div>

## 🧩 项目简介

这是一个使用 Go 构建的模块化后端服务，集成了：

- Gin：高性能 HTTP 路由框架
- GORM + SQLite：轻量级持久化
- Testify：单元测试与断言
- Docker：容器化开发环境
- GitHub Actions：自动化测试与覆盖率报告
- Codecov：覆盖率分析平台

---

## 🚀 快速开始

```bash
# 构建 docker
docker build -f docker/Dockerfile -t go-api .

# 启动服务
docker run -it --rm \
  -v "$(pwd)/src:/app/go-api/src" \
  -v "$(pwd)/database:/app/go-api/database" \
  -v "$(pwd)/test:/app/go-api/test" \
  -p 8080:8080 go-api

# 在 Docker 内运行
make run

# 在 Docker 内测试
make test

```
## 📂 项目结构
```
go-api/
├── .github/workflows/   # GitHub Actions、自动化流水线
├── build/               # 测试
    ├── docker-build.sh  # 构造 docker 镜像脚本
    └── docker-run.sh    # 运行/进入docker
├── docker/
    └── Dockerfile       # docker配置
├── database/            # 数据库
├── src/                 # 源代码
    └── handlers/        # 路由逻辑
        └── todos.go
├── tests/               # 测试代码
├── main.go              # 程序入口
├── Makefile             # 快捷命令
└── run.sh               # docker 内执行脚本
```