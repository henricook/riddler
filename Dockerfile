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
    && go build -o /go/bin/riddler

##############################
# STEP 2 build a small image #
##############################
FROM alpine

# Copy our static executable and certs
COPY --from=builder /go/bin/* /go/bin/

WORKDIR /go/bin

# Run the riddler binary.
ENTRYPOINT ["./riddler"]
