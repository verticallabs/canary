FROM golang:1.10.1-alpine as build
ARG GOOS
ARG GOARCH
ARG GOARM

WORKDIR /go/src/github.com/verticallabs/canary
COPY . .
RUN GOOS=$GOOS GOARCH=$GOARCH GOARM=$GOARM go build

FROM armhf/alpine
RUN apk add --update curl bind-tools
WORKDIR /app
COPY --from=build /go/src/github.com/verticallabs/canary/canary canary

ENTRYPOINT ["/app/canary"]