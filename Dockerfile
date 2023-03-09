## Build layer
FROM golang:1.19.3-buster AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./
COPY *.html ./

RUN go build -o /colormatic

## Deploy layer
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /colormatic /colormatic

EXPOSE 9090

USER nonroot:nonroot

ENTRYPOINT ["/colormatic"]