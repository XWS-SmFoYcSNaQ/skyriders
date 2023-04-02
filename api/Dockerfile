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
COPY --from=build_container /app/app.env /app/app.env 
COPY --from=build_container /app/config/rbac_model.conf /app/config/rbac_model.conf
EXPOSE 8080
ENTRYPOINT ["api"]
