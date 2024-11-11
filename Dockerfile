FROM golang:alpine

RUN apk update && \
    apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    echo "Asia/Jakarta" > /etc/timezone


WORKDIR /app

# RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/air-verse/air@latest


COPY . .
# RUN go get

RUN go get && go mod tidy

# RUN go build -o main



# ENTRYPOINT ["/app/main"]
EXPOSE 9000

ENV TZ=Asia/Jakarta


# CMD ["watch","-n","60","ls", "-a"]
CMD ["air", "-c", ".air.toml"]

# CMD ["go","run","main.go"]