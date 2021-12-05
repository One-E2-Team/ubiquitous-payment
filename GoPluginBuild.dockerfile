FROM golang AS build
COPY . src/ubiquitous-payment/
RUN cd src/ubiquitous-payment && \
    mkdir target &&  \
    go mod download &&  \
    go mod verify &&  \
    go build -buildmode=plugin -o target/paypal.so psp-plugins/paypal/main.go

FROM scratch AS export-stage
COPY --from=build /go/src/ubiquitous-payment/target /