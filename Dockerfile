FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build cmd/main.go

FROM alpine:3.15
WORKDIR /app
COPY --from=build-env /src/main /app/
ENTRYPOINT ./main

EXPOSE 2113