# 🧠 Go Todo API Project - 2
<div align="center">

![Go Test](https://github.com/Alec1120/go-api-2/actions/workflows/test.yml/badge.svg)

</div>

<div align="center">

[中文](./README.md) | English

</div>

## 🧩 Project Overview

This is a modular backend API built with Go, integrating:

- Gin: High-performance HTTP routing framework  
- GORM + SQLite: Lightweight persistent storage  
- Testify: Unit testing and assertion framework  
- Docker: Containerized development environment  
- GitHub Actions: Automated testing and coverage reports  
- Codecov: Coverage analysis platform

---

## 🚀 Quick Start

```bash
# Build Docker image
docker build -f docker/Dockerfile -t go-api .

# Run service (with volume mounts)
docker run -it --rm \
  -v "$(pwd)/src:/app/go-api/src" \
  -v "$(pwd)/database:/app/go-api/database" \
  -v "$(pwd)/test:/app/go-api/test" \
  -p 8080:8080 go-api

# Run inside Docker container
make run

# Run tests inside Docker container
make test
```

## 📂 项目结构
```
go-api/
├── .github/workflows/   # GitHub Actions CI workflows
├── build/               # Build & test scripts
│   ├── docker-build.sh  # Docker build script
│   └── docker-run.sh    # Docker run/enter script
├── docker/
│   └── Dockerfile       # Docker configuration
├── database/            # SQLite database
├── src/                 # Source code
│   └── handlers/        # Route logic
│       └── todos.go
├── tests/               # Unit test code
├── main.go              # Application entry point
├── Makefile             # Common CLI commands
└── run.sh               # Internal docker run script
```