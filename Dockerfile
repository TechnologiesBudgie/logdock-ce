FROM golang:1.22-alpine AS build
WORKDIR /src
COPY . .
RUN go mod tidy && CGO_ENABLED=0 go build -ldflags='-s -w' -o /out/logdock ./cmd/logdock

FROM scratch
COPY --from=build /out/logdock /logdock
EXPOSE 2514 4317 4318 5140/udp 5140/tcp
ENTRYPOINT ["/logdock","serve"]
