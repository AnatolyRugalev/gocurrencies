############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

WORKDIR /root
COPY . .

RUN go build -mod vendor -o main

############################
# STEP 2 build a small image
############################
FROM alpine

# ssl support require ca-certificates
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

# default:
ENV REMOTE_SVC_URL https://api.ratesapi.io/api/latest

EXPOSE 8080
COPY --from=builder /root/main /main

ENTRYPOINT ["/main"]
