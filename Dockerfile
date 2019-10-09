FROM golang:latest as build-env
WORKDIR $GOPATH/src/kong-prometheus-exporter
ADD . $GOPATH/src/kong-prometheus-exporter
RUN go build -o $GOPATH/app

#FROM harbor.wise-paas.io/distroless/base:latest as prod-env
#WORKDIR $GOPATH/
#COPY --from=build-env $GOPATH/app .
EXPOSE 8080
CMD ["./kong-prometheus-exporter"]