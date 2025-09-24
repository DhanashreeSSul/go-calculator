FROM golang:1.25-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o calculator

FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/calculator .
COPY index.html .
EXPOSE 8080
CMD ["./calculator"]
