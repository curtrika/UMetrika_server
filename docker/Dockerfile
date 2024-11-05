FROM golang:1.23-alpine as builder

RUN apk update && apk add --no-cache make

WORKDIR /usr/src/service
COPY go.mod .
COPY go.sum .
#COPY db/mock_data_up.sql .
#COPY db/mock_data_down.sql .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux make build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /usr/src/service/build/main .
COPY --from=builder /usr/src/service/configs/config.toml /app/configs/config.toml
#COPY --from=builder /usr/src/service/db/mock_data_up.sql /app/db/mock_data_up.sql
#COPY --from=builder /usr/src/service/db/mock_data_down.sql /app/db/mock_data_down.sql

RUN sleep 5

#CMD ["/app/main", "--config", "/app/config/config.yaml"]
#CMD ["./main", "-config-path", "configs/apiserver.toml"]
CMD ["./cmd/sso/main", "-config-path", "configs/apiserver.toml"]