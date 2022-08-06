FROM golang:1.17-alpine as build
WORKDIR /app
COPY ./*.go .
RUN CGO_ENABLE=0 go build -o server main.go

FROM alpine:3.12
WORKDIR /app
COPY --from=build /app/server .
#RUN chmod +x /app/server
CMD ["/app/server"]
