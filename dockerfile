FROM golang:1.21.0-alpine3.18 AS BuildStage

ENV DISCORD_TOKEN ${DISCORD_TOKEN}

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN echo ${DISCORD_TOKEN}

RUN CGO_ENABLED=0 GOOS=linux go build -o /zenbot

CMD ["/zenbot"]

#Deploy stage

FROM alpine:latest

WORKDIR /

COPY --from=BuildStage /zenbot /zenbot

ENTRYPOINT ["/zenbot"]

