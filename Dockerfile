# Build
FROM golang:1.14-alpine3.11 AS build
RUN apk add --no-cache make git protobuf protobuf-dev curl && \
    rm -rf /var/cache/apk/*
ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /build
COPY . .
RUN make

# Production
FROM alpine:3.11
RUN apk add --no-cache ca-certificates su-exec && \
    rm -rf /var/cache/apk/*
RUN addgroup -S qvspot && adduser -S qvspot -G qvspot
RUN mkdir -p /opt/qvspot
WORKDIR /opt/qvspot
EXPOSE 8080
COPY --from=build /build/api .
CMD [ "su-exec", "qvspot:qvspot", "/opt/qvspot/api", "api" ]
