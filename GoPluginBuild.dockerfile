FROM golang AS build
COPY . src/ubiquitous-payment/
ENV CGO_ENABLED=0
#https://stackoverflow.com/questions/36279253/go-compiled-binary-wont-run-in-an-alpine-docker-container-on-ubuntu-host
RUN cd src/ubiquitous-payment && \
    mkdir target &&  \
    go mod download &&  \
    go mod verify &&  \
    go build -buildmode=plugin -tags netgo -a -v -o target/paypal.so psp-plugins/paypal/main.go

FROM scratch AS export-stage
COPY --from=build /go/src/ubiquitous-payment/target /