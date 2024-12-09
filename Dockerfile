FROM golang:1.23 as builder
ARG CGO_ENABLED=0
WORKDIR /app
RUN apt-get update && apt-get install -y upx
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make release
FROM scratch
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]