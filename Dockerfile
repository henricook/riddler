FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/henricook/riddler/
COPY . .

# Fetch dependencies.
# Using go get.
RUN go get -d -v

RUN mv server.crt /go/bin/server.crt \
    && mv server.key /go/bin/server.key \
    && mv *.txt /go/bin/ \
    && GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/riddler

##############################
# STEP 2 build a small image #
##############################
FROM alpine

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

# Copy our static executable and certs
COPY --from=builder /go/bin/* /go/bin/

RUN chown -R appuser:appuser /go/bin

# Use an unprivileged user.
USER appuser:appuser

WORKDIR /go/bin

# Run the riddler binary.
ENTRYPOINT ["./riddler"]

## TODO Secure