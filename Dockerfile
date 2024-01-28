FROM golang:1.21-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /go-spa

FROM alpine
WORKDIR /
COPY --from=build /go-spa /go-spa
EXPOSE 8888
ENTRYPOINT ["/go-spa"]
