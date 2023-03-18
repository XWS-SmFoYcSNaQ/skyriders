FROM golang:alpine as build_container
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o api


FROM alpine
WORKDIR /app
COPY --from=build_container /app/api /usr/bin
EXPOSE 8080
ENTRYPOINT ["api"]
