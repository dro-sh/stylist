FROM golang:1.19-alpine AS builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN apk update && apk add make

COPY main.go ./
COPY ["app/", "app/"]
COPY ["Makefile", "Makefile"]

RUN make server-build OUTPUT="./bin/app"

FROM scratch
WORKDIR /app

COPY --from=builder /app/bin/app ./bin/app

ENTRYPOINT [ "/app/bin/app" ]
