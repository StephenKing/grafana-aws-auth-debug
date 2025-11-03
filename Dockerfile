FROM golang:1.25.3 AS builder
ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build

FROM scratch
COPY --from=builder /app/grafana-aws-auth-debug /grafana-aws-auth-debug
ENTRYPOINT ["/grafana-aws-auth-debug"]
