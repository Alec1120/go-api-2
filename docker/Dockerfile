# docker/Dockerfile.dev
#### Ubuntu 22.04 Golang 1.22
FROM golang:1.22

#### initial stage
# set the timezone to Asia/Shanghai(optional)
# ENV GOPROXY=https://goproxy.cn,direct

# set Go module mode
ENV GO111MODULE=on

# ✅ set workdir
WORKDIR /app/go-api

# ✅ copy source code
COPY ./src /app/go-api/src
COPY ./test /app/go-api/test
COPY ./Makefile /app/go-api/Makefile
COPY ./run.sh /app/go-api/run.sh

# ✅ init go module(go-api)
RUN go mod init go-api || true

# ✅ install dependencies
RUN go mod tidy

EXPOSE 8080

#### run stage(comment out the following lines during debugging to access the container manually)
# ✅ run the app
# CMD ["make", "run"]