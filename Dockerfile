############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
# Create appuser.
ENV USER=appuser
ENV UID=10001
# See https://stackoverflow.com/a/55757473/12429735RUN
RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid "${UID}" \
  "${USER}"
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Using go mod.
# RUN go mod download
# RUN go mod verify

# This container exposes port 8080 to the outside world
EXPOSE 8080
# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/simples-go-api
############################
# STEP 2 build a small image
############################
FROM scratch
# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
# Copy our static executable.
COPY --from=builder /go/bin/simples-go-api /go/bin/simples-go-api
# Use an unprivileged user.
USER appuser:appuser
# Run the simples-go-api binary.
ENTRYPOINT ["/go/bin/simples-go-api"]
