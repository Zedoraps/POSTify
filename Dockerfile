FROM golang:1.20-alpine as build

WORKDIR /build

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
#COPY go.mod go.sum ./ re enable if we have dependencies
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -v -ldflags="-w -s" .

FROM alpine:3.18

COPY --from=build /build/POSTify /usr/local/bin/POSTify

ENTRYPOINT ["POSTify"]