# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
# building stage
# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
FROM golang:1.25-alpine As builder 

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN go build -o app main.go

# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
# final image
# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
FROM alpine:latest

ENV TZ=Asia/Tehran
ENV ROUTEROS_ADDRESS="192.168.100.88:8728"
ENV ROUTEROS_USER="ali"
ENV ROUTEROS_PASSWORD="Ali@1384"
ENV SERVER_ADDRESS="0.0.0.0:9200"


RUN apk add --no-cache --update \
    bash \
    tzdata

COPY --from=builder /app/app /usr/local/bin/app
COPY --from=builder /app/entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

EXPOSE 9200

ENTRYPOINT ["/entrypoint.sh"] 




