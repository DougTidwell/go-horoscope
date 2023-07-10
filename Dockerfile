FROM golang:1.20.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# SOURCE is the name of the horoscope file that should be copied
# as main.go. This lets us use one project and one Dockerfile to
# build all the different images. So use --build-args SOURCE=kafkaesque/main.go
# with a different filename each time, and change the image tag
# accordingly.

ARG SOURCE
ENV SOURCE_FILE=$SOURCE

COPY $SOURCE_FILE ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /horoscope

EXPOSE 3000

CMD ["/horoscope"]