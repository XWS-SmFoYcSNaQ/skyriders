FROM golang:alpine as build_container
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o static_server
CMD ["./mongoDBExample"]


FROM alpine
WORKDIR /app
COPY --from=build_container /app/static_server /usr/bin
EXPOSE 8080
ENTRYPOINT ["static_server"]
