FROM golang:1.22

ENV USER=appuser \
  UID=1000 \
  GO111MODULE=on \
  CGO_ENABLED=1

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/tmp" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid "${UID}" \
  "${USER}"

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download && \
  mkdir -p /.cache/go-build && \
  chown -R appuser:appuser /.cache

USER appuser:appuser
CMD ["air", "-c", ".air.toml"]
