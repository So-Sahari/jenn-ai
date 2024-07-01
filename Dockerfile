FROM golang:1.22 as build

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

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build \
  -o /bin/jenn-ai \
  -ldflags "-w -s -linkmode external -extldflags "-static"" \ 
  . && \
  chown appuser:appuser /bin/jenn-ai

FROM scratch

# Needed for ssl requirements 
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Used for the unprivileged user
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group 

WORKDIR /app
COPY --from=build /bin/jenn-ai /bin/jenn-ai
COPY --from=build /src/templates /app/templates
USER appuser:appuser
ENTRYPOINT ["/bin/jenn-ai", "server"]
