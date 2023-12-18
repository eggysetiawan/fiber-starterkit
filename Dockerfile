FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .
# RUN go get

RUN go get && go mod tidy

# RUN go build -o main



# ENTRYPOINT ["/app/main"]
EXPOSE 3000

# CMD ["watch","-n","60","ls", "-a"]
CMD ["go","run","main.go"]