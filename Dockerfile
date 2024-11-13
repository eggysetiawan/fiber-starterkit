FROM golang:alpine

RUN apk update && \
    apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    echo "Asia/Jakarta" > /etc/timezone


WORKDIR /app

# RUN go install github.com/cosmtrek/air@latest
# Install air using the updated module path
RUN go install github.com/air-verse/air@latest


COPY . .
# RUN go get

RUN go get && go mod tidy

# RUN go build -o main



# ENTRYPOINT ["/app/main"]
<<<<<<< HEAD
EXPOSE 5500
=======
EXPOSE 9000
>>>>>>> bb9c8ca9d35d4d61eaffd6012850efcf47de3920

ENV TZ=Asia/Jakarta


# CMD ["watch","-n","60","ls", "-a"]
CMD ["air", "-c", ".air.toml"]

# CMD ["go","run","main.go"]