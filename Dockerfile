#FROM golang:latest as build-env
#WORKDIR $GOPATH/src/kong-prometheus-exporter
#ADD . $GOPATH/src/kong-prometheus-exporter
#RUN go build -o .
#
#FROM harbor.wise-paas.io/distroless/base:latest as prod-env
#WORKDIR $GOPATH/
#COPY --from=build-env $GOPATH/src/kong-prometheus-exporter .
#EXPOSE 8080
#CMD ["./kong-prometheus-exporter"]

FROM golang:latest as build-env
WORKDIR /go/src/kong-prometheus-exporter
ADD . /go/src/kong-prometheus-exporter
RUN go build -o /go/app

FROM harbor.wise-paas.io/distroless/base:latest as prod-env
WORKDIR /go/
COPY --from=build-env /go/app .
EXPOSE 8080
CMD ["./kong-prometheus-exporter"]