FROM golang:1.25-alpine AS build
RUN adduser -D -g '' appuser
WORKDIR /src
COPY . .
RUN go mod tidy && CGO_ENABLED=0 go build -ldflags='-s -w' -o /out/logdock ./cmd/logdock

FROM scratch
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
USER appuser
COPY --from=build /out/logdock /logdock
EXPOSE 2514 4317 4318 5140/udp 5140/tcp
ENTRYPOINT ["/logdock","serve"]
