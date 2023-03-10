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

# Embed the GitHub Actions Run URL so we can link to it in the application.
ARG GITHUB_ACTIONS_RUN_URL
ENV GITHUB_ACTIONS_RUN_URL=$GITHUB_ACTIONS_RUN_URL

ENTRYPOINT ["/colormatic"]