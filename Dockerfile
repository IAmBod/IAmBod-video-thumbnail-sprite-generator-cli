FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download \
     && go mod verify

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/video-storybook-cli

FROM builder AS test

RUN go test -v ./...

FROM jrottenberg/ffmpeg:7-alpine

COPY --from=builder /usr/local/bin/video-storybook-cli /usr/local/bin/video-storybook-cli

ENTRYPOINT ["video-storybook-cli"]