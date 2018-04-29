ARG GOOS
ARG GOARCH
ARG GOARM

FROM golang:1.10.1-alpine as build
WORKDIR /go/src/github.com/verticallabs/canary
COPY . .
RUN GOOS=$GOOS GOARCH=$GOARM GOARM=$GOARM go build

FROM alpine
RUN apk add --update curl bind-tools
WORKDIR /app
COPY --from=build /go/src/github.com/verticallabs/canary/canary canary

ENTRYPOINT ["/app/canary"]