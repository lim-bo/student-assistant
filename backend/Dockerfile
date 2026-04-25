FROM golang:1.25.6-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /main cmd/main.go

FROM alpine:latest
COPY --from=build /main /bin/main

CMD [ "/bin/main" ]