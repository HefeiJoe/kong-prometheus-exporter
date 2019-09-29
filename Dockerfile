FROM golang:latest as build-env
WORKDIR /go/src/kongExporter
ADD . /go/src/kongExporter
RUN go build -o /go/app

FROM harbor.wise-paas.io/distroless/base:latest as prod-env
WORKDIR /go/
COPY --from=build-env /go/app .
EXPOSE 8080
CMD ["./app"]