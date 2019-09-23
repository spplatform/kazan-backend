FROM golang:1.12-alpine

RUN mkdir -p /opt/code/

WORKDIR /opt/code/

ADD ./ /opt/code/

RUN apk update && apk upgrade && \
    apk add --no-cache git
    
RUN go get
# build for alpine
RUN GOOS=linux GARCH=amd64 CGO_ENABLED=0 \
    go build  -o bin/kazan cmd/kazan-server/main.go

FROM alpine

WORKDIR /app

EXPOSE 8080

COPY --from=0 /opt/code/bin/kazan /app/
COPY --from=0 /opt/code/configs/config.yml /app/configs/config.yml

ENTRYPOINT ["./kazan"]