FROM golang:alpine
RUN apk update && apk add --no-cache git
WORKDIR /app/crowdfunding/backend
COPY . .
RUN go mod tidy
RUN go build -o server
ENTRYPOINT ["/app/crowdfunding/backend/server"]

